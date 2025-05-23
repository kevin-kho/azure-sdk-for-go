// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

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

// ProtectionContainerOperationResultsClient contains the methods for the ProtectionContainerOperationResults group.
// Don't use this type directly, use NewProtectionContainerOperationResultsClient() instead.
type ProtectionContainerOperationResultsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProtectionContainerOperationResultsClient creates a new instance of ProtectionContainerOperationResultsClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProtectionContainerOperationResultsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProtectionContainerOperationResultsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProtectionContainerOperationResultsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Fetches the result of any operation on the container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name associated with the container.
//   - containerName - Container name whose information should be fetched.
//   - operationID - Operation ID which represents the operation whose result needs to be fetched.
//   - options - ProtectionContainerOperationResultsClientGetOptions contains the optional parameters for the ProtectionContainerOperationResultsClient.Get
//     method.
func (client *ProtectionContainerOperationResultsClient) Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, operationID string, options *ProtectionContainerOperationResultsClientGetOptions) (ProtectionContainerOperationResultsClientGetResponse, error) {
	var err error
	const operationName = "ProtectionContainerOperationResultsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, operationID, options)
	if err != nil {
		return ProtectionContainerOperationResultsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProtectionContainerOperationResultsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ProtectionContainerOperationResultsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ProtectionContainerOperationResultsClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, operationID string, _ *ProtectionContainerOperationResultsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/operationResults/{operationId}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProtectionContainerOperationResultsClient) getHandleResponse(resp *http.Response) (ProtectionContainerOperationResultsClientGetResponse, error) {
	result := ProtectionContainerOperationResultsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionContainerResource); err != nil {
		return ProtectionContainerOperationResultsClientGetResponse{}, err
	}
	return result, nil
}
