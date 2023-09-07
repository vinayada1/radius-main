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

package converter

import (
	"encoding/json"
	"fmt"

	v1 "github.com/radius-project/radius/pkg/armrpc/api/v1"
	"github.com/radius-project/radius/pkg/ucp/api/v20220901privatepreview"
	"github.com/radius-project/radius/pkg/ucp/datamodel"
)

type AWSResource struct {
}

func (aws *AWSResource) ConvertFrom(model v1.DataModelInterface) error {
	return nil
}

func (aws *AWSResource) ConvertTo() (v1.DataModelInterface, error) {
	return nil, nil
}

// AWSDataModelToVersioned converts version agnostic plane datamodel to versioned model.
func AWSDataModelToVersioned(model *datamodel.AWSResource, version string) (v1.VersionedModelInterface, error) {
	switch version {
	case v20220901privatepreview.Version:
		versioned := &AWSResource{}
		return versioned, nil

	default:
		return nil, v1.ErrUnsupportedAPIVersion
	}
}

// PlaneDataModelFromVersioned converts versioned plane model to datamodel.
func AWSDataModelFromVersioned(content []byte, version string) (*datamodel.AWSResource, error) {
	switch version {
	case v20220901privatepreview.Version:
		// return &datamodel.AWSResource{
		// 	Properties: map[string]any{
		// 		"Name":            "test",
		// 		"RetentionPeriod": 1,
		// 	},
		// }, nil
		awsResource := datamodel.AWSResource{}
		if err := json.Unmarshal(content, &awsResource); err != nil {
			return nil, err
		}
		fmt.Println("@@@@ content: ", string(content))
		return &awsResource, nil

	default:
		return nil, v1.ErrUnsupportedAPIVersion
	}
}
