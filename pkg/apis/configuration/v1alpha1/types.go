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

package v1alpha1

import (
	"strconv"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +kubebuilder:object:root=true

// Configuration describes an Dapr configuration setting.
type Configuration struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec ConfigurationSpec `json:"spec,omitempty"`
}

// ConfigurationSpec is the spec for an configuration.
type ConfigurationSpec struct {
	// +optional
	AppHTTPPipelineSpec PipelineSpec `json:"appHttpPipeline,omitempty"`
	// +optional
	HTTPPipelineSpec PipelineSpec `json:"httpPipeline,omitempty"`
	// +optional
	TracingSpec TracingSpec `json:"tracing,omitempty"`
	// +kubebuilder:default={enabled:true}
	MetricSpec MetricSpec `json:"metric,omitempty"`
	// +kubebuilder:default={enabled:true}
	MetricsSpec MetricSpec `json:"metrics,omitempty"`
	// +optional
	MTLSSpec MTLSSpec `json:"mtls,omitempty"`
	// +optional
	Secrets SecretsSpec `json:"secrets,omitempty"`
	// +optional
	AccessControlSpec AccessControlSpec `json:"accessControl,omitempty"`
	// +optional
	NameResolutionSpec NameResolutionSpec `json:"nameResolution,omitempty"`
	// +optional
	Features []FeatureSpec `json:"features,omitempty"`
	// +optional
	APISpec APISpec `json:"api,omitempty"`
	// +optional
	ComponentsSpec ComponentsSpec `json:"components,omitempty"`
	// +optional
	LoggingSpec LoggingSpec `json:"logging,omitempty"`
}

// APISpec describes the configuration for Dapr APIs.
type APISpec struct {
	Allowed []APIAccessRule `json:"allowed,omitempty"`
}

// APIAccessRule describes an access rule for allowing a Dapr API to be enabled and accessible by an app.
type APIAccessRule struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	// +optional
	Protocol string `json:"protocol"`
}

// NameResolutionSpec is the spec for name resolution configuration.
type NameResolutionSpec struct {
	Component     string       `json:"component"`
	Version       string       `json:"version"`
	Configuration DynamicValue `json:"configuration"`
}

// SecretsSpec is the spec for secrets configuration.
type SecretsSpec struct {
	Scopes []SecretsScope `json:"scopes"`
}

// SecretsScope defines the scope for secrets.
type SecretsScope struct {
	// +optional
	DefaultAccess string `json:"defaultAccess,omitempty"`
	StoreName     string `json:"storeName"`
	// +optional
	AllowedSecrets []string `json:"allowedSecrets,omitempty"`
	// +optional
	DeniedSecrets []string `json:"deniedSecrets,omitempty"`
}

// PipelineSpec defines the middleware pipeline.
type PipelineSpec struct {
	Handlers []HandlerSpec `json:"handlers"`
}

// HandlerSpec defines a request handlers.
type HandlerSpec struct {
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	SelectorSpec SelectorSpec `json:"selector,omitempty"`
}

// MTLSSpec defines mTLS configuration.
type MTLSSpec struct {
	Enabled bool `json:"enabled"`
	// +optional
	WorkloadCertTTL string `json:"workloadCertTTL"`
	// +optional
	AllowedClockSkew string `json:"allowedClockSkew"`
}

// SelectorSpec selects target services to which the handler is to be applied.
type SelectorSpec struct {
	Fields []SelectorField `json:"fields"`
}

// SelectorField defines a selector fields.
type SelectorField struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

// TracingSpec defines distributed tracing configuration.
type TracingSpec struct {
	SamplingRate string `json:"samplingRate"`
	// +optional
	Stdout bool `json:"stdout"`
	// +optional
	Zipkin ZipkinSpec `json:"zipkin"`
	// +optional
	Otel OtelSpec `json:"otel"`
}

// OtelSpec defines Otel exporter configurations.
type OtelSpec struct {
	Protocol        string `json:"protocol" yaml:"protocol"`
	EndpointAddress string `json:"endpointAddress" yaml:"endpointAddress"`
	IsSecure        bool   `json:"isSecure" yaml:"isSecure"`
}

// ZipkinSpec defines Zipkin trace configurations.
type ZipkinSpec struct {
	EndpointAddresss string `json:"endpointAddress"`
}

// MetricSpec defines metrics configuration.
type MetricSpec struct {
	Enabled bool `json:"enabled"`
	// +optional
	Rules []MetricsRule `json:"rules"`
}

// MetricsRule defines configuration options for a metric.
type MetricsRule struct {
	Name   string        `json:"name"`
	Labels []MetricLabel `json:"labels"`
}

// MetricsLabel defines an object that allows to set regex expressions for a label.
type MetricLabel struct {
	Name  string            `json:"name"`
	Regex map[string]string `json:"regex"`
}

// AppPolicySpec defines the policy data structure for each app.
type AppPolicySpec struct {
	AppName string `json:"appId" yaml:"appId"`
	// +optional
	DefaultAction string `json:"defaultAction" yaml:"defaultAction"`
	// +optional
	TrustDomain string `json:"trustDomain" yaml:"trustDomain"`
	// +optional
	Namespace string `json:"namespace" yaml:"namespace"`
	// +optional
	AppOperationActions []AppOperationAction `json:"operations" yaml:"operations"`
}

// AppOperationAction defines the data structure for each app operation.
type AppOperationAction struct {
	Operation string `json:"name" yaml:"name"`
	// +optional
	HTTPVerb []string `json:"httpVerb" yaml:"httpVerb"`
	Action   string   `json:"action" yaml:"action"`
}

// AccessControlSpec is the spec object in ConfigurationSpec.
type AccessControlSpec struct {
	// +optional
	DefaultAction string `json:"defaultAction" yaml:"defaultAction"`
	// +optional
	TrustDomain string `json:"trustDomain" yaml:"trustDomain"`
	// +optional
	AppPolicies []AppPolicySpec `json:"policies" yaml:"policies"`
}

// FeatureSpec defines the features that are enabled/disabled.
type FeatureSpec struct {
	Name    string `json:"name" yaml:"name"`
	Enabled bool   `json:"enabled" yaml:"enabled"`
}

// ComponentsSpec describes the configuration for Dapr components
type ComponentsSpec struct {
	// Denylist of component types that cannot be instantiated
	// +optional
	Deny []string `json:"deny,omitempty" yaml:"deny,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationList is a list of Dapr event sources.
type ConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Configuration `json:"items"`
}

// LoggingSpec defines the configuration for logging.
type LoggingSpec struct {
	// Configure API logging.
	// +optional
	APILogging APILoggingSpec `json:"apiLogging" yaml:"apiLogging"`
}

// APILoggingSpec defines the configuration for API logging.
type APILoggingSpec struct {
	// Default value for enabling API logging. Sidecars can always override this by setting `--enable-api-logging` to true or false explicitly.
	// The default value is false.
	// +optional
	Enabled bool `json:"enabled" yaml:"enabled"`
	// If true, health checks are not reported in API logs. Default: false.
	// This option has no effect if API logging is disabled.
	// +optional
	OmitHealthChecks bool `json:"omitHealthChecks" yaml:"omitHealthChecks"`
}

// DynamicValue is a dynamic value struct for the component.metadata pair value.
type DynamicValue struct {
	v1.JSON `json:",inline"`
}

// String returns the string representation of the raw value.
// If the value is a string, it will be unquoted as the string is guaranteed to be a JSON serialized string.
func (d *DynamicValue) String() string {
	s := string(d.Raw)
	c, err := strconv.Unquote(s)
	if err == nil {
		s = c
	}
	return s
}
