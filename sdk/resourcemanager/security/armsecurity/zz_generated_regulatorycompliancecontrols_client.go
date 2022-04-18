//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RegulatoryComplianceControlsClient contains the methods for the RegulatoryComplianceControls group.
// Don't use this type directly, use NewRegulatoryComplianceControlsClient() instead.
type RegulatoryComplianceControlsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRegulatoryComplianceControlsClient creates a new instance of RegulatoryComplianceControlsClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRegulatoryComplianceControlsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RegulatoryComplianceControlsClient, error) {
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
	client := &RegulatoryComplianceControlsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Selected regulatory compliance control details and state
// If the operation fails it returns an *azcore.ResponseError type.
// regulatoryComplianceStandardName - Name of the regulatory compliance standard object
// regulatoryComplianceControlName - Name of the regulatory compliance control object
// options - RegulatoryComplianceControlsClientGetOptions contains the optional parameters for the RegulatoryComplianceControlsClient.Get
// method.
func (client *RegulatoryComplianceControlsClient) Get(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, options *RegulatoryComplianceControlsClientGetOptions) (RegulatoryComplianceControlsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, regulatoryComplianceStandardName, regulatoryComplianceControlName, options)
	if err != nil {
		return RegulatoryComplianceControlsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RegulatoryComplianceControlsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RegulatoryComplianceControlsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RegulatoryComplianceControlsClient) getCreateRequest(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, options *RegulatoryComplianceControlsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls/{regulatoryComplianceControlName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regulatoryComplianceStandardName == "" {
		return nil, errors.New("parameter regulatoryComplianceStandardName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regulatoryComplianceStandardName}", url.PathEscape(regulatoryComplianceStandardName))
	if regulatoryComplianceControlName == "" {
		return nil, errors.New("parameter regulatoryComplianceControlName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regulatoryComplianceControlName}", url.PathEscape(regulatoryComplianceControlName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RegulatoryComplianceControlsClient) getHandleResponse(resp *http.Response) (RegulatoryComplianceControlsClientGetResponse, error) {
	result := RegulatoryComplianceControlsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegulatoryComplianceControl); err != nil {
		return RegulatoryComplianceControlsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - All supported regulatory compliance controls details and state for selected standard
// If the operation fails it returns an *azcore.ResponseError type.
// regulatoryComplianceStandardName - Name of the regulatory compliance standard object
// options - RegulatoryComplianceControlsClientListOptions contains the optional parameters for the RegulatoryComplianceControlsClient.List
// method.
func (client *RegulatoryComplianceControlsClient) NewListPager(regulatoryComplianceStandardName string, options *RegulatoryComplianceControlsClientListOptions) *runtime.Pager[RegulatoryComplianceControlsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[RegulatoryComplianceControlsClientListResponse]{
		More: func(page RegulatoryComplianceControlsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegulatoryComplianceControlsClientListResponse) (RegulatoryComplianceControlsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, regulatoryComplianceStandardName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RegulatoryComplianceControlsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RegulatoryComplianceControlsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RegulatoryComplianceControlsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RegulatoryComplianceControlsClient) listCreateRequest(ctx context.Context, regulatoryComplianceStandardName string, options *RegulatoryComplianceControlsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regulatoryComplianceStandardName == "" {
		return nil, errors.New("parameter regulatoryComplianceStandardName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regulatoryComplianceStandardName}", url.PathEscape(regulatoryComplianceStandardName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RegulatoryComplianceControlsClient) listHandleResponse(resp *http.Response) (RegulatoryComplianceControlsClientListResponse, error) {
	result := RegulatoryComplianceControlsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegulatoryComplianceControlList); err != nil {
		return RegulatoryComplianceControlsClientListResponse{}, err
	}
	return result, nil
}
