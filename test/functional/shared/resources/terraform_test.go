/*
Copyright 2023 The Radius Authors.

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

package resource_test

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/project-radius/radius/test/functional"
	"github.com/project-radius/radius/test/functional/shared"
	"github.com/project-radius/radius/test/step"
	"github.com/project-radius/radius/test/validation"
)

// Test_TerraformRecipe_Redis covers the following terraform recipe scenario:
//
// - Create an extender resource using a Terraform recipe that deploys Redis on Kubernetes.
// - The recipe deployment creates a Kubernetes deployment and a Kubernetes service.
func Test_TerraformRecipe_KubernetesRedis(t *testing.T) {
	template := "testdata/corerp-resources-terraform-redis.bicep"
	name := "corerp-resources-terraform-redis"
	appName := "corerp-resources-terraform-redis-app"
	redisCacheName := "tf-redis-cache"

	test := shared.NewRPTest(t, name, []shared.TestStep{
		{
			Executor: step.NewDeployExecutor(template, functional.GetTerraformRecipeModuleServerURL(), "appName="+appName, "redisCacheName="+redisCacheName),
			RPResources: &validation.RPResourceSet{
				Resources: []validation.RPResource{
					{
						Name: "corerp-resources-terraform-redis-env",
						Type: validation.EnvironmentsResource,
					},
					{
						Name: appName,
						Type: validation.ApplicationsResource,
					},
					{
						Name:            "corerp-resources-terraform-redis",
						Type:            validation.ExtendersResource,
						App:             appName,
						OutputResources: []validation.OutputResourceResponse{}, // No output resources because Terraform Recipe outputs aren't integreted yet.
					},
				},
			},
			K8sObjects: &validation.K8sObjectSet{
				Namespaces: map[string][]validation.K8sObject{
					appName: {
						validation.NewK8sServiceForResource(appName, redisCacheName).ValidateLabels(false),
					},
				},
			},
			SkipResourceDeletion: true, // Skip deletion because Terraform Recipe deletion isn't supported yet.
		},
	})
	test.Test(t)
}

func Test_TerraformRecipe_Context(t *testing.T) {
	template := "testdata/corerp-resources-terraform-context.bicep"
	name := "corerp-resources-terraform-context"

	appNamespace := "corerp-resources-terraform-context-app"

	test := shared.NewRPTest(t, name, []shared.TestStep{
		{
			Executor: step.NewDeployExecutor(template, functional.GetTerraformRecipeModuleServerURL()),
			RPResources: &validation.RPResourceSet{
				Resources: []validation.RPResource{
					{
						Name: name,
						Type: validation.ApplicationsResource,
					},
					{
						Name: name,
						Type: validation.ExtendersResource,
					},
				},
			},
			K8sObjects: &validation.K8sObjectSet{
				Namespaces: map[string][]validation.K8sObject{
					appNamespace: {
						validation.NewK8sSecretForResource(name, name),
					},
				},
			},
			PostStepVerify: func(ctx context.Context, t *testing.T, test shared.RPTest) {
				// `k8ssecret-context` recipe should have created a secret with the populated recipe context.
				s, err := test.Options.K8sClient.CoreV1().Secrets(appNamespace).Get(ctx, name, metav1.GetOptions{})
				require.NoError(t, err)

				tests := []struct {
					key      string
					expected string
				}{
					{
						key:      "resource.id",
						expected: "/planes/radius/local/resourcegroups/kind-radius/providers/Applications.Link/extenders/corerp-resources-terraform-context",
					},
					{
						key:      "resource.type",
						expected: "Applications.Link/extenders",
					},
					{
						key:      "azure.subscription_id",
						expected: "00000000-0000-0000-0000-100000000000",
					},
					{
						key:      "recipe_context",
						expected: "{\"application\":{\"id\":\"/planes/radius/local/resourcegroups/kind-radius/providers/Applications.Core/applications/corerp-resources-terraform-context\",\"name\":\"corerp-resources-terraform-context\"},\"aws\":null,\"azure\":{\"resourceGroup\":{\"id\":\"/subscriptions/00000000-0000-0000-0000-100000000000/resourceGroups/rg-terraform-context\",\"name\":\"rg-terraform-context\"},\"subscription\":{\"id\":\"/subscriptions/00000000-0000-0000-0000-100000000000\",\"subscriptionId\":\"00000000-0000-0000-0000-100000000000\"}},\"environment\":{\"id\":\"/planes/radius/local/resourcegroups/kind-radius/providers/Applications.Core/environments/corerp-resources-terraform-context\",\"name\":\"corerp-resources-terraform-context\"},\"resource\":{\"id\":\"/planes/radius/local/resourcegroups/kind-radius/providers/Applications.Link/extenders/corerp-resources-terraform-context\",\"name\":\"corerp-resources-terraform-context\",\"type\":\"Applications.Link/extenders\"},\"runtime\":{\"kubernetes\":{\"environmentNamespace\":\"corerp-resources-terraform-context-env\",\"namespace\":\"corerp-resources-terraform-context-app\"}}}",
					},
				}

				for _, tc := range tests {
					decoded, err := base64.StdEncoding.DecodeString(string(s.Data[tc.key]))
					require.NoErrorf(t, err, "failed to decode secret data, key: %s", tc.key)
					require.Equalf(t, tc.expected, string(decoded), "secret data mismatch, key: %s", tc.key)
				}
			},
			SkipResourceDeletion: true,
		},
	})
	test.Test(t)
}