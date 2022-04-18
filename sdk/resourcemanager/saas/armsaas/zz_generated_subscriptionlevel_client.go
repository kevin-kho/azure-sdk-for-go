//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsaas

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

// SubscriptionLevelClient contains the methods for the SaasSubscriptionLevel group.
// Don't use this type directly, use NewSubscriptionLevelClient() instead.
type SubscriptionLevelClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSubscriptionLevelClient creates a new instance of SubscriptionLevelClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSubscriptionLevelClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SubscriptionLevelClient, error) {
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
	client := &SubscriptionLevelClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a SaaS resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// resourceName - The name of the resource.
// parameters - Parameters supplied to the create or update subscription level saas operation.
// options - SubscriptionLevelClientBeginCreateOrUpdateOptions contains the optional parameters for the SubscriptionLevelClient.BeginCreateOrUpdate
// method.
func (client *SubscriptionLevelClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters ResourceCreation, options *SubscriptionLevelClientBeginCreateOrUpdateOptions) (*armruntime.Poller[SubscriptionLevelClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[SubscriptionLevelClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[SubscriptionLevelClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a SaaS resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SubscriptionLevelClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters ResourceCreation, options *SubscriptionLevelClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SubscriptionLevelClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters ResourceCreation, options *SubscriptionLevelClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SaaS/resources/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified SaaS.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// resourceName - The name of the resource.
// options - SubscriptionLevelClientBeginDeleteOptions contains the optional parameters for the SubscriptionLevelClient.BeginDelete
// method.
func (client *SubscriptionLevelClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *SubscriptionLevelClientBeginDeleteOptions) (*armruntime.Poller[SubscriptionLevelClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[SubscriptionLevelClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[SubscriptionLevelClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified SaaS.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SubscriptionLevelClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *SubscriptionLevelClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SubscriptionLevelClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *SubscriptionLevelClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SaaS/resources/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets information about the specified Subscription Level SaaS.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// resourceName - The name of the resource.
// options - SubscriptionLevelClientGetOptions contains the optional parameters for the SubscriptionLevelClient.Get method.
func (client *SubscriptionLevelClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *SubscriptionLevelClientGetOptions) (SubscriptionLevelClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return SubscriptionLevelClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionLevelClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionLevelClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SubscriptionLevelClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *SubscriptionLevelClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SaaS/resources/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubscriptionLevelClient) getHandleResponse(resp *http.Response) (SubscriptionLevelClientGetResponse, error) {
	result := SubscriptionLevelClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Resource); err != nil {
		return SubscriptionLevelClientGetResponse{}, err
	}
	return result, nil
}

// ListAccessToken - Gets the ISV access token for a specified Subscription Level SaaS.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// resourceName - The name of the resource.
// options - SubscriptionLevelClientListAccessTokenOptions contains the optional parameters for the SubscriptionLevelClient.ListAccessToken
// method.
func (client *SubscriptionLevelClient) ListAccessToken(ctx context.Context, resourceGroupName string, resourceName string, options *SubscriptionLevelClientListAccessTokenOptions) (SubscriptionLevelClientListAccessTokenResponse, error) {
	req, err := client.listAccessTokenCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return SubscriptionLevelClientListAccessTokenResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionLevelClientListAccessTokenResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionLevelClientListAccessTokenResponse{}, runtime.NewResponseError(resp)
	}
	return client.listAccessTokenHandleResponse(resp)
}

// listAccessTokenCreateRequest creates the ListAccessToken request.
func (client *SubscriptionLevelClient) listAccessTokenCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *SubscriptionLevelClientListAccessTokenOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SaaS/resources/{resourceName}/listAccessToken"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAccessTokenHandleResponse handles the ListAccessToken response.
func (client *SubscriptionLevelClient) listAccessTokenHandleResponse(resp *http.Response) (SubscriptionLevelClientListAccessTokenResponse, error) {
	result := SubscriptionLevelClientListAccessTokenResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessTokenResult); err != nil {
		return SubscriptionLevelClientListAccessTokenResponse{}, err
	}
	return result, nil
}

// NewListByAzureSubscriptionPager - Gets information about all the Subscription Level SaaS in a certain Azure subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - SubscriptionLevelClientListByAzureSubscriptionOptions contains the optional parameters for the SubscriptionLevelClient.ListByAzureSubscription
// method.
func (client *SubscriptionLevelClient) NewListByAzureSubscriptionPager(options *SubscriptionLevelClientListByAzureSubscriptionOptions) *runtime.Pager[SubscriptionLevelClientListByAzureSubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[SubscriptionLevelClientListByAzureSubscriptionResponse]{
		More: func(page SubscriptionLevelClientListByAzureSubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubscriptionLevelClientListByAzureSubscriptionResponse) (SubscriptionLevelClientListByAzureSubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAzureSubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SubscriptionLevelClientListByAzureSubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SubscriptionLevelClientListByAzureSubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SubscriptionLevelClientListByAzureSubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAzureSubscriptionHandleResponse(resp)
		},
	})
}

// listByAzureSubscriptionCreateRequest creates the ListByAzureSubscription request.
func (client *SubscriptionLevelClient) listByAzureSubscriptionCreateRequest(ctx context.Context, options *SubscriptionLevelClientListByAzureSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.SaaS/resources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAzureSubscriptionHandleResponse handles the ListByAzureSubscription response.
func (client *SubscriptionLevelClient) listByAzureSubscriptionHandleResponse(resp *http.Response) (SubscriptionLevelClientListByAzureSubscriptionResponse, error) {
	result := SubscriptionLevelClientListByAzureSubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceResponseWithContinuation); err != nil {
		return SubscriptionLevelClientListByAzureSubscriptionResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets information about all the Subscription Level SaaS in a certain resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - SubscriptionLevelClientListByResourceGroupOptions contains the optional parameters for the SubscriptionLevelClient.ListByResourceGroup
// method.
func (client *SubscriptionLevelClient) NewListByResourceGroupPager(resourceGroupName string, options *SubscriptionLevelClientListByResourceGroupOptions) *runtime.Pager[SubscriptionLevelClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[SubscriptionLevelClientListByResourceGroupResponse]{
		More: func(page SubscriptionLevelClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubscriptionLevelClientListByResourceGroupResponse) (SubscriptionLevelClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SubscriptionLevelClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SubscriptionLevelClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SubscriptionLevelClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SubscriptionLevelClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SubscriptionLevelClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SaaS/resources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SubscriptionLevelClient) listByResourceGroupHandleResponse(resp *http.Response) (SubscriptionLevelClientListByResourceGroupResponse, error) {
	result := SubscriptionLevelClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceResponseWithContinuation); err != nil {
		return SubscriptionLevelClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginMoveResources - Move a specified Subscription Level SaaS.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// moveResourceParameter - Object that represents the resources to move.
// options - SubscriptionLevelClientBeginMoveResourcesOptions contains the optional parameters for the SubscriptionLevelClient.BeginMoveResources
// method.
func (client *SubscriptionLevelClient) BeginMoveResources(ctx context.Context, resourceGroupName string, moveResourceParameter MoveResource, options *SubscriptionLevelClientBeginMoveResourcesOptions) (*armruntime.Poller[SubscriptionLevelClientMoveResourcesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.moveResources(ctx, resourceGroupName, moveResourceParameter, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[SubscriptionLevelClientMoveResourcesResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[SubscriptionLevelClientMoveResourcesResponse](options.ResumeToken, client.pl, nil)
	}
}

// MoveResources - Move a specified Subscription Level SaaS.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SubscriptionLevelClient) moveResources(ctx context.Context, resourceGroupName string, moveResourceParameter MoveResource, options *SubscriptionLevelClientBeginMoveResourcesOptions) (*http.Response, error) {
	req, err := client.moveResourcesCreateRequest(ctx, resourceGroupName, moveResourceParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// moveResourcesCreateRequest creates the MoveResources request.
func (client *SubscriptionLevelClient) moveResourcesCreateRequest(ctx context.Context, resourceGroupName string, moveResourceParameter MoveResource, options *SubscriptionLevelClientBeginMoveResourcesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/moveResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, moveResourceParameter)
}

// BeginUpdate - Updates a SaaS Subscription Level resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// resourceName - The name of the resource.
// parameters - Parameters supplied to the update saas operation.
// options - SubscriptionLevelClientBeginUpdateOptions contains the optional parameters for the SubscriptionLevelClient.BeginUpdate
// method.
func (client *SubscriptionLevelClient) BeginUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters ResourceCreation, options *SubscriptionLevelClientBeginUpdateOptions) (*armruntime.Poller[SubscriptionLevelClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, resourceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[SubscriptionLevelClientUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[SubscriptionLevelClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates a SaaS Subscription Level resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SubscriptionLevelClient) update(ctx context.Context, resourceGroupName string, resourceName string, parameters ResourceCreation, options *SubscriptionLevelClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *SubscriptionLevelClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters ResourceCreation, options *SubscriptionLevelClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SaaS/resources/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginUpdateToUnsubscribed - Unsubscribe from a specified Subscription Level SaaS.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// resourceName - The name of the resource.
// parameters - Parameters supplied to unsubscribe saas operation.
// options - SubscriptionLevelClientBeginUpdateToUnsubscribedOptions contains the optional parameters for the SubscriptionLevelClient.BeginUpdateToUnsubscribed
// method.
func (client *SubscriptionLevelClient) BeginUpdateToUnsubscribed(ctx context.Context, resourceGroupName string, resourceName string, parameters DeleteOptions, options *SubscriptionLevelClientBeginUpdateToUnsubscribedOptions) (*armruntime.Poller[SubscriptionLevelClientUpdateToUnsubscribedResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateToUnsubscribed(ctx, resourceGroupName, resourceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[SubscriptionLevelClientUpdateToUnsubscribedResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[SubscriptionLevelClientUpdateToUnsubscribedResponse](options.ResumeToken, client.pl, nil)
	}
}

// UpdateToUnsubscribed - Unsubscribe from a specified Subscription Level SaaS.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SubscriptionLevelClient) updateToUnsubscribed(ctx context.Context, resourceGroupName string, resourceName string, parameters DeleteOptions, options *SubscriptionLevelClientBeginUpdateToUnsubscribedOptions) (*http.Response, error) {
	req, err := client.updateToUnsubscribedCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateToUnsubscribedCreateRequest creates the UpdateToUnsubscribed request.
func (client *SubscriptionLevelClient) updateToUnsubscribedCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters DeleteOptions, options *SubscriptionLevelClientBeginUpdateToUnsubscribedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SaaS/resources/{resourceName}/unsubscribe"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// ValidateMoveResources - Validate whether a specified Subscription Level SaaS can be moved.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// moveResourceParameter - Object that represents the resources to move.
// options - SubscriptionLevelClientValidateMoveResourcesOptions contains the optional parameters for the SubscriptionLevelClient.ValidateMoveResources
// method.
func (client *SubscriptionLevelClient) ValidateMoveResources(ctx context.Context, resourceGroupName string, moveResourceParameter MoveResource, options *SubscriptionLevelClientValidateMoveResourcesOptions) (SubscriptionLevelClientValidateMoveResourcesResponse, error) {
	req, err := client.validateMoveResourcesCreateRequest(ctx, resourceGroupName, moveResourceParameter, options)
	if err != nil {
		return SubscriptionLevelClientValidateMoveResourcesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionLevelClientValidateMoveResourcesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionLevelClientValidateMoveResourcesResponse{}, runtime.NewResponseError(resp)
	}
	return SubscriptionLevelClientValidateMoveResourcesResponse{}, nil
}

// validateMoveResourcesCreateRequest creates the ValidateMoveResources request.
func (client *SubscriptionLevelClient) validateMoveResourcesCreateRequest(ctx context.Context, resourceGroupName string, moveResourceParameter MoveResource, options *SubscriptionLevelClientValidateMoveResourcesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/validateMoveResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-beta")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, moveResourceParameter)
}
