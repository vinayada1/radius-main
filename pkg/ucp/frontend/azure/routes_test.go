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

package azure

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"

	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	"github.com/project-radius/radius/pkg/armrpc/rpctest"
	"github.com/project-radius/radius/pkg/ucp/api/v20220901privatepreview"
	"github.com/project-radius/radius/pkg/ucp/dataprovider"
	"github.com/project-radius/radius/pkg/ucp/frontend/modules"
	"github.com/project-radius/radius/pkg/ucp/hostoptions"
	"github.com/project-radius/radius/pkg/ucp/secret"
	secretprovider "github.com/project-radius/radius/pkg/ucp/secret/provider"
)

const pathBase = "/some-path-base"

func Test_Routes(t *testing.T) {
	tests := []rpctest.HandlerTestSpec{
		{
			OperationType: v1.OperationType{Type: v20220901privatepreview.AzureCredentialType, Method: v1.OperationList},
			Method:        http.MethodGet,
			Path:          "/planes/azure/azurecloud/providers/System.Azure/credentials",
		}, {
			OperationType: v1.OperationType{Type: v20220901privatepreview.AzureCredentialType, Method: v1.OperationGet},
			Method:        http.MethodGet,
			Path:          "/planes/azure/azurecloud/providers/System.Azure/credentials/default",
		}, {
			OperationType: v1.OperationType{Type: v20220901privatepreview.AzureCredentialType, Method: v1.OperationPut},
			Method:        http.MethodPut,
			Path:          "/planes/azure/azurecloud/providers/System.Azure/credentials/default",
		}, {
			OperationType: v1.OperationType{Type: v20220901privatepreview.AzureCredentialType, Method: v1.OperationDelete},
			Method:        http.MethodDelete,
			Path:          "/planes/azure/azurecloud/providers/System.Azure/credentials/default",
		}, {
			OperationType:               v1.OperationType{Type: OperationTypeUCPAzureProxy, Method: v1.OperationProxy},
			Method:                      http.MethodGet,
			Path:                        "/planes/azure/azurecloud/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/some-group/providers/Microsoft.Storage/storageAccounts/test-account",
			SkipOperationTypeValidation: true,
		}, {
			OperationType:               v1.OperationType{Type: OperationTypeUCPAzureProxy, Method: v1.OperationProxy},
			Method:                      http.MethodPut,
			Path:                        "/planes/azure/azurecloud/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/some-group/providers/Microsoft.Storage/storageAccounts/test-account",
			SkipOperationTypeValidation: true,
		},
	}

	ctrl := gomock.NewController(t)
	dataProvider := dataprovider.NewMockDataStorageProvider(ctrl)
	dataProvider.EXPECT().GetStorageClient(gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()

	secretClient := secret.NewMockClient(ctrl)
	secretProvider := secretprovider.NewSecretProvider(secretprovider.SecretProviderOptions{})
	secretProvider.SetClient(secretClient)

	options := modules.Options{
		Address:        "localhost",
		PathBase:       pathBase,
		Config:         &hostoptions.UCPConfig{},
		DataProvider:   dataProvider,
		SecretProvider: secretProvider,
	}

	rpctest.AssertRouters(t, tests, pathBase, "", func(ctx context.Context) (chi.Router, error) {
		module := NewModule(options)
		router, err := module.Initialize(ctx)
		return router.(chi.Router), err
	})
}