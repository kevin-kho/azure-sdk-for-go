// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WorkspaceAPIExportClient contains the methods for the WorkspaceAPIExport group.
// Don't use this type directly, use NewWorkspaceAPIExportClient() instead.
type WorkspaceAPIExportClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceAPIExportClient creates a new instance of WorkspaceAPIExportClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceAPIExportClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceAPIExportClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceAPIExportClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the details of the API specified by its identifier in the format specified to the Storage Blob with SAS Key
// valid for 5 minutes.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - formatParam - Format in which to export the Api Details to the Storage Blob with Sas Key valid for 5 minutes.
//   - export - Query parameter required to export the API details.
//   - options - WorkspaceAPIExportClientGetOptions contains the optional parameters for the WorkspaceAPIExportClient.Get method.
func (client *WorkspaceAPIExportClient) Get(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, formatParam ExportFormat, export ExportAPI, options *WorkspaceAPIExportClientGetOptions) (WorkspaceAPIExportClientGetResponse, error) {
	var err error
	const operationName = "WorkspaceAPIExportClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, apiID, formatParam, export, options)
	if err != nil {
		return WorkspaceAPIExportClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceAPIExportClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceAPIExportClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkspaceAPIExportClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, formatParam ExportFormat, export ExportAPI, _ *WorkspaceAPIExportClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/apis/{apiId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	reqQP.Set("export", string(export))
	reqQP.Set("format", string(formatParam))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceAPIExportClient) getHandleResponse(resp *http.Response) (WorkspaceAPIExportClientGetResponse, error) {
	result := WorkspaceAPIExportClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIExportResult); err != nil {
		return WorkspaceAPIExportClientGetResponse{}, err
	}
	return result, nil
}
