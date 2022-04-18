//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlocks

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// AuthorizationOperationsClient contains the methods for the AuthorizationOperations group.
// Don't use this type directly, use NewAuthorizationOperationsClient() instead.
type AuthorizationOperationsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewAuthorizationOperationsClient creates a new instance of AuthorizationOperationsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAuthorizationOperationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AuthorizationOperationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AuthorizationOperationsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// NewListPager - Lists all of the available Microsoft.Authorization REST API operations.
// If the operation fails it returns an *azcore.ResponseError type.
// options - AuthorizationOperationsClientListOptions contains the optional parameters for the AuthorizationOperationsClient.List
// method.
func (client *AuthorizationOperationsClient) NewListPager(options *AuthorizationOperationsClientListOptions) *runtime.Pager[AuthorizationOperationsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[AuthorizationOperationsClientListResponse]{
		More: func(page AuthorizationOperationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AuthorizationOperationsClientListResponse) (AuthorizationOperationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AuthorizationOperationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AuthorizationOperationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AuthorizationOperationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AuthorizationOperationsClient) listCreateRequest(ctx context.Context, options *AuthorizationOperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AuthorizationOperationsClient) listHandleResponse(resp *http.Response) (AuthorizationOperationsClientListResponse, error) {
	result := AuthorizationOperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationListResult); err != nil {
		return AuthorizationOperationsClientListResponse{}, err
	}
	return result, nil
}
