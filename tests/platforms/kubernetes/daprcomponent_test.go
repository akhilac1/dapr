/*
Copyright 2022 The Dapr Authors
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

package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDaprComponentSpec(t *testing.T) {
	t.Run("should set name when specified", func(t *testing.T) {
		const testName = "fake-name"
		daprComponent := DaprComponent{component: ComponentDescription{
			Name: testName,
		}}
		assert.Equal(t, daprComponent.toComponentSpec().Name, testName)
	})

	t.Run("should set typename when specified", func(t *testing.T) {
		const testTypeName = "state.redis"
		daprComponent := DaprComponent{component: ComponentDescription{
			TypeName: testTypeName,
		}}
		assert.Equal(t, daprComponent.toComponentSpec().Spec.Type, testTypeName)
	})

	t.Run("should add metadata when specified", func(t *testing.T) {
		const testKey, testValue = "key", `"value"`
		daprComponent := DaprComponent{component: ComponentDescription{
			MetaData: map[string]MetadataValue{
				testKey: {Raw: testValue},
			},
		}}
		metadata := daprComponent.toComponentSpec().Spec.Metadata
		assert.Len(t, metadata, 1)
		assert.Equal(t, metadata[0].Name, testKey)
		assert.Equal(t, metadata[0].Value.Raw, []byte(testValue))
	})

	t.Run("should add secretkeyref as metadata value when specified", func(t *testing.T) {
		const testSecretKey, fromSecretName, fromSecretKey = "secretKey", "secretName", "secretKey"
		daprComponent := DaprComponent{component: ComponentDescription{
			MetaData: map[string]MetadataValue{
				testSecretKey: {FromSecretRef: &SecretRef{
					Name: fromSecretName,
					Key:  fromSecretKey,
				}},
			},
		}}
		metadata := daprComponent.toComponentSpec().Spec.Metadata
		assert.Len(t, metadata, 1)
		assert.Equal(t, metadata[0].Name, testSecretKey)
		assert.Equal(t, metadata[0].SecretKeyRef.Name, fromSecretName)
		assert.Equal(t, metadata[0].SecretKeyRef.Key, fromSecretKey)
	})
}
