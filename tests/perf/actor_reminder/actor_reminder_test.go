//go:build perf
// +build perf

/*
Copyright 2021 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package actor_reminder_perf

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/dapr/dapr/tests/perf"
	"github.com/dapr/dapr/tests/perf/utils"
	kube "github.com/dapr/dapr/tests/platforms/kubernetes"
	"github.com/dapr/dapr/tests/runner"
	"github.com/stretchr/testify/require"
)

const (
	numHealthChecks        = 60 // Number of times to check for endpoint health per app.
	actorType              = "PerfTestActorReminder"
	serviceApplicationName = "perf-actor-reminder-service"
	clientApplicationName  = "perf-actor-reminder-client"
)

var tr *runner.TestRunner

func TestMain(m *testing.M) {
	utils.SetupLogs("actor_reminder")

	testApps := []kube.AppDescription{
		{
			AppName:           serviceApplicationName,
			DaprEnabled:       true,
			ImageName:         "perf-actorfeatures",
			Replicas:          1,
			IngressEnabled:    true,
			AppPort:           3000,
			DaprCPULimit:      "4.0",
			DaprCPURequest:    "0.1",
			DaprMemoryLimit:   "512Mi",
			DaprMemoryRequest: "250Mi",
			AppCPULimit:       "4.0",
			AppCPURequest:     "0.1",
			AppMemoryLimit:    "800Mi",
			AppMemoryRequest:  "2500Mi",
			AppEnv: map[string]string{
				"TEST_APP_ACTOR_TYPE": actorType,
			},
			Labels: map[string]string{
				"daprtest": serviceApplicationName,
			},
		},
		{
			AppName:           clientApplicationName,
			DaprEnabled:       true,
			ImageName:         "perf-tester",
			Replicas:          1,
			IngressEnabled:    true,
			AppPort:           3001,
			DaprCPULimit:      "4.0",
			DaprCPURequest:    "0.1",
			DaprMemoryLimit:   "512Mi",
			DaprMemoryRequest: "250Mi",
			AppCPULimit:       "4.0",
			AppCPURequest:     "0.1",
			AppMemoryLimit:    "800Mi",
			AppMemoryRequest:  "2500Mi",
			Labels: map[string]string{
				"daprtest": clientApplicationName,
			},
			PodAffinityLabels: map[string]string{
				"daprtest": serviceApplicationName,
			},
		},
	}

	tr = runner.NewTestRunner("actorreminder", testApps, nil, nil)
	os.Exit(tr.Start(m))
}

func TestActorReminderRegistrationPerformance(t *testing.T) {
	p := perf.Params(
		perf.WithQPS(500),
		perf.WithConnections(8),
		perf.WithDuration("1m"),
		perf.WithPayload("{}"),
	)

	// Get the ingress external url of test app
	testAppURL := tr.Platform.AcquireAppExternalURL(serviceApplicationName)
	require.NotEmpty(t, testAppURL, "test app external URL must not be empty")

	// Check if test app endpoint is available
	t.Logf("test app url: %s", testAppURL+"/health")
	_, err := utils.HTTPGetNTimes(testAppURL+"/health", numHealthChecks)
	require.NoError(t, err)

	// Get the ingress external url of tester app
	testerAppURL := tr.Platform.AcquireAppExternalURL(clientApplicationName)
	require.NotEmpty(t, testerAppURL, "tester app external URL must not be empty")

	// Check if tester app endpoint is available
	t.Logf("tester app url: %s", testerAppURL)
	_, err = utils.HTTPGetNTimes(testerAppURL, numHealthChecks)
	require.NoError(t, err)

	// Perform dapr test
	endpoint := fmt.Sprintf("http://127.0.0.1:3500/v1.0/actors/%v/{uuid}/reminders/myreminder", actorType)
	p.TargetEndpoint = endpoint
	p.Payload = "{\"dueTime\":\"24h\",\"period\":\"24h\"}"
	body, err := json.Marshal(&p)
	require.NoError(t, err)

	t.Logf("running dapr test with params: %s", body)
	daprResp, err := utils.HTTPPost(fmt.Sprintf("%s/test", testerAppURL), body)
	t.Log("checking err...")
	require.NoError(t, err)
	require.NotEmpty(t, daprResp)
	// fast fail if daprResp starts with error
	require.False(t, strings.HasPrefix(string(daprResp), "error"))

	// Let test run for 10 minutes triggering the timers and collect metrics.
	time.Sleep(10 * time.Minute)

	appUsage, err := tr.Platform.GetAppUsage(serviceApplicationName)
	require.NoError(t, err)

	sidecarUsage, err := tr.Platform.GetSidecarUsage(serviceApplicationName)
	require.NoError(t, err)

	restarts, err := tr.Platform.GetTotalRestarts(serviceApplicationName)
	require.NoError(t, err)

	t.Logf("dapr test results: %s", string(daprResp))
	t.Logf("target dapr app consumed %vm CPU and %vMb of Memory", appUsage.CPUm, appUsage.MemoryMb)
	t.Logf("target dapr sidecar consumed %vm CPU and %vMb of Memory", sidecarUsage.CPUm, sidecarUsage.MemoryMb)
	t.Logf("target dapr app or sidecar restarted %v times", restarts)

	var daprResult perf.TestResult
	err = json.Unmarshal(daprResp, &daprResult)
	require.NoError(t, err)

	percentiles := map[int]string{2: "90th", 3: "99th"}

	for k, v := range percentiles {
		daprValue := daprResult.DurationHistogram.Percentiles[k].Value
		t.Logf("%s percentile: %sms", v, fmt.Sprintf("%.2f", daprValue*1000))
	}
	t.Logf("Actual QPS: %.2f, expected QPS: %d", daprResult.ActualQPS, p.QPS)

	report := perf.NewTestReport(
		[]perf.TestResult{daprResult},
		"Actor Reminder",
		sidecarUsage,
		appUsage)
	report.SetTotalRestartCount(restarts)
	err = utils.UploadAzureBlob(report)

	if err != nil {
		t.Error(err)
	}

	require.Equal(t, 0, daprResult.RetCodes.Num400)
	require.Equal(t, 0, daprResult.RetCodes.Num500)
	require.Equal(t, 0, restarts)
	require.True(t, daprResult.ActualQPS > 55)
}
