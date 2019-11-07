/*
Copyright The KubeVault Authors.

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

package s3

import (
	"fmt"
	"testing"

	api "kubevault.dev/operator/apis/kubevault/v1alpha1"

	"github.com/stretchr/testify/assert"
)

func TestOptions_GetStorageConfig(t *testing.T) {
	opts, err := NewOptions(api.S3Spec{
		Bucket:             "test",
		EndPoint:           "endpoint",
		Region:             "test-region",
		CredentialSecret:   "credential",
		SessionTokenSecret: "session",
		MaxParallel:        128,
		S3ForcePathStyle:   true,
		DisableSSL:         true,
	})
	assert.Nil(t, err)

	out := `
storage "s3" {
bucket = "test"
endpoint = "endpoint"
region = "test-region"
s3_force_path_style = "true"
disable_ssl = "true"
max_parallel = 128
}
`
	t.Run("S3 storage config", func(t *testing.T) {
		got, err := opts.GetStorageConfig()
		assert.Nil(t, err)
		if !assert.Equal(t, out, got) {
			fmt.Println("expected:", out)
			fmt.Println("got:", got)
		}
	})
}
