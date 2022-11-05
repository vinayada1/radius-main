// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package daprinvokehttproutes

import (
	"context"
	"errors"
	"net/http"

	v1 "github.com/project-radius/radius/pkg/armrpc/api/v1"
	ctrl "github.com/project-radius/radius/pkg/armrpc/frontend/controller"
	"github.com/project-radius/radius/pkg/armrpc/rest"
	"github.com/project-radius/radius/pkg/linkrp/datamodel"
	"github.com/project-radius/radius/pkg/linkrp/datamodel/converter"
	"github.com/project-radius/radius/pkg/ucp/store"
)

var _ ctrl.Controller = (*GetDaprInvokeHttpRoute)(nil)

// GetDaprInvokeHttpRoute is the controller implementation to get the daprInvokeHttpRoute conenctor resource.
type GetDaprInvokeHttpRoute struct {
	ctrl.BaseController
}

// NewGetDaprInvokeHttpRoute creates a new instance of GetDaprInvokeHttpRoute.
func NewGetDaprInvokeHttpRoute(opts ctrl.Options) (ctrl.Controller, error) {
	return &GetDaprInvokeHttpRoute{ctrl.NewBaseController(opts)}, nil
}

func (daprHttpRoute *GetDaprInvokeHttpRoute) Run(ctx context.Context, w http.ResponseWriter, req *http.Request) (rest.Response, error) {
	serviceCtx := v1.ARMRequestContextFromContext(ctx)

	existingResource := &datamodel.DaprInvokeHttpRoute{}
	_, err := daprHttpRoute.GetResource(ctx, serviceCtx.ResourceID.String(), existingResource)
	if err != nil {
		if errors.Is(&store.ErrNotFound{}, err) {
			return rest.NewNotFoundResponse(serviceCtx.ResourceID), nil
		}
		return nil, err
	}

	versioned, _ := converter.DaprInvokeHttpRouteDataModelToVersioned(existingResource, serviceCtx.APIVersion)
	return rest.NewOKResponse(versioned), nil
}