// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dependencymap/armdependencymap"
	"net/http"
	"net/url"
	"regexp"
)

// MapsServer is a fake server for instances of the armdependencymap.MapsClient type.
type MapsServer struct {
	// BeginCreateOrUpdate is the fake for method MapsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, mapName string, resource armdependencymap.MapsResource, options *armdependencymap.MapsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdependencymap.MapsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method MapsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, mapName string, options *armdependencymap.MapsClientBeginDeleteOptions) (resp azfake.PollerResponder[armdependencymap.MapsClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginExportDependencies is the fake for method MapsClient.BeginExportDependencies
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginExportDependencies func(ctx context.Context, resourceGroupName string, mapName string, body armdependencymap.ExportDependenciesRequest, options *armdependencymap.MapsClientBeginExportDependenciesOptions) (resp azfake.PollerResponder[armdependencymap.MapsClientExportDependenciesResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method MapsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, mapName string, options *armdependencymap.MapsClientGetOptions) (resp azfake.Responder[armdependencymap.MapsClientGetResponse], errResp azfake.ErrorResponder)

	// BeginGetConnectionsForProcessOnFocusedMachine is the fake for method MapsClient.BeginGetConnectionsForProcessOnFocusedMachine
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginGetConnectionsForProcessOnFocusedMachine func(ctx context.Context, resourceGroupName string, mapName string, body armdependencymap.GetConnectionsForProcessOnFocusedMachineRequest, options *armdependencymap.MapsClientBeginGetConnectionsForProcessOnFocusedMachineOptions) (resp azfake.PollerResponder[armdependencymap.MapsClientGetConnectionsForProcessOnFocusedMachineResponse], errResp azfake.ErrorResponder)

	// BeginGetConnectionsWithConnectedMachineForFocusedMachine is the fake for method MapsClient.BeginGetConnectionsWithConnectedMachineForFocusedMachine
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginGetConnectionsWithConnectedMachineForFocusedMachine func(ctx context.Context, resourceGroupName string, mapName string, body armdependencymap.GetConnectionsWithConnectedMachineForFocusedMachineRequest, options *armdependencymap.MapsClientBeginGetConnectionsWithConnectedMachineForFocusedMachineOptions) (resp azfake.PollerResponder[armdependencymap.MapsClientGetConnectionsWithConnectedMachineForFocusedMachineResponse], errResp azfake.ErrorResponder)

	// BeginGetDependencyViewForFocusedMachine is the fake for method MapsClient.BeginGetDependencyViewForFocusedMachine
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginGetDependencyViewForFocusedMachine func(ctx context.Context, resourceGroupName string, mapName string, body armdependencymap.GetDependencyViewForFocusedMachineRequest, options *armdependencymap.MapsClientBeginGetDependencyViewForFocusedMachineOptions) (resp azfake.PollerResponder[armdependencymap.MapsClientGetDependencyViewForFocusedMachineResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method MapsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armdependencymap.MapsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armdependencymap.MapsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method MapsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armdependencymap.MapsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armdependencymap.MapsClientListBySubscriptionResponse])

	// BeginUpdate is the fake for method MapsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, mapName string, properties armdependencymap.MapsResourceTagsUpdate, options *armdependencymap.MapsClientBeginUpdateOptions) (resp azfake.PollerResponder[armdependencymap.MapsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewMapsServerTransport creates a new instance of MapsServerTransport with the provided implementation.
// The returned MapsServerTransport instance is connected to an instance of armdependencymap.MapsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMapsServerTransport(srv *MapsServer) *MapsServerTransport {
	return &MapsServerTransport{
		srv:                     srv,
		beginCreateOrUpdate:     newTracker[azfake.PollerResponder[armdependencymap.MapsClientCreateOrUpdateResponse]](),
		beginDelete:             newTracker[azfake.PollerResponder[armdependencymap.MapsClientDeleteResponse]](),
		beginExportDependencies: newTracker[azfake.PollerResponder[armdependencymap.MapsClientExportDependenciesResponse]](),
		beginGetConnectionsForProcessOnFocusedMachine:            newTracker[azfake.PollerResponder[armdependencymap.MapsClientGetConnectionsForProcessOnFocusedMachineResponse]](),
		beginGetConnectionsWithConnectedMachineForFocusedMachine: newTracker[azfake.PollerResponder[armdependencymap.MapsClientGetConnectionsWithConnectedMachineForFocusedMachineResponse]](),
		beginGetDependencyViewForFocusedMachine:                  newTracker[azfake.PollerResponder[armdependencymap.MapsClientGetDependencyViewForFocusedMachineResponse]](),
		newListByResourceGroupPager:                              newTracker[azfake.PagerResponder[armdependencymap.MapsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:                               newTracker[azfake.PagerResponder[armdependencymap.MapsClientListBySubscriptionResponse]](),
		beginUpdate:                                              newTracker[azfake.PollerResponder[armdependencymap.MapsClientUpdateResponse]](),
	}
}

// MapsServerTransport connects instances of armdependencymap.MapsClient to instances of MapsServer.
// Don't use this type directly, use NewMapsServerTransport instead.
type MapsServerTransport struct {
	srv                                                      *MapsServer
	beginCreateOrUpdate                                      *tracker[azfake.PollerResponder[armdependencymap.MapsClientCreateOrUpdateResponse]]
	beginDelete                                              *tracker[azfake.PollerResponder[armdependencymap.MapsClientDeleteResponse]]
	beginExportDependencies                                  *tracker[azfake.PollerResponder[armdependencymap.MapsClientExportDependenciesResponse]]
	beginGetConnectionsForProcessOnFocusedMachine            *tracker[azfake.PollerResponder[armdependencymap.MapsClientGetConnectionsForProcessOnFocusedMachineResponse]]
	beginGetConnectionsWithConnectedMachineForFocusedMachine *tracker[azfake.PollerResponder[armdependencymap.MapsClientGetConnectionsWithConnectedMachineForFocusedMachineResponse]]
	beginGetDependencyViewForFocusedMachine                  *tracker[azfake.PollerResponder[armdependencymap.MapsClientGetDependencyViewForFocusedMachineResponse]]
	newListByResourceGroupPager                              *tracker[azfake.PagerResponder[armdependencymap.MapsClientListByResourceGroupResponse]]
	newListBySubscriptionPager                               *tracker[azfake.PagerResponder[armdependencymap.MapsClientListBySubscriptionResponse]]
	beginUpdate                                              *tracker[azfake.PollerResponder[armdependencymap.MapsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for MapsServerTransport.
func (m *MapsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *MapsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if mapsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = mapsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "MapsClient.BeginCreateOrUpdate":
				res.resp, res.err = m.dispatchBeginCreateOrUpdate(req)
			case "MapsClient.BeginDelete":
				res.resp, res.err = m.dispatchBeginDelete(req)
			case "MapsClient.BeginExportDependencies":
				res.resp, res.err = m.dispatchBeginExportDependencies(req)
			case "MapsClient.Get":
				res.resp, res.err = m.dispatchGet(req)
			case "MapsClient.BeginGetConnectionsForProcessOnFocusedMachine":
				res.resp, res.err = m.dispatchBeginGetConnectionsForProcessOnFocusedMachine(req)
			case "MapsClient.BeginGetConnectionsWithConnectedMachineForFocusedMachine":
				res.resp, res.err = m.dispatchBeginGetConnectionsWithConnectedMachineForFocusedMachine(req)
			case "MapsClient.BeginGetDependencyViewForFocusedMachine":
				res.resp, res.err = m.dispatchBeginGetDependencyViewForFocusedMachine(req)
			case "MapsClient.NewListByResourceGroupPager":
				res.resp, res.err = m.dispatchNewListByResourceGroupPager(req)
			case "MapsClient.NewListBySubscriptionPager":
				res.resp, res.err = m.dispatchNewListBySubscriptionPager(req)
			case "MapsClient.BeginUpdate":
				res.resp, res.err = m.dispatchBeginUpdate(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (m *MapsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := m.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DependencyMap/maps/(?P<mapName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdependencymap.MapsResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mapNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mapName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, mapNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		m.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		m.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		m.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (m *MapsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if m.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := m.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DependencyMap/maps/(?P<mapName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mapNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mapName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginDelete(req.Context(), resourceGroupNameParam, mapNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		m.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		m.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		m.beginDelete.remove(req)
	}

	return resp, nil
}

func (m *MapsServerTransport) dispatchBeginExportDependencies(req *http.Request) (*http.Response, error) {
	if m.srv.BeginExportDependencies == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportDependencies not implemented")}
	}
	beginExportDependencies := m.beginExportDependencies.get(req)
	if beginExportDependencies == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DependencyMap/maps/(?P<mapName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportDependencies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdependencymap.ExportDependenciesRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mapNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mapName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginExportDependencies(req.Context(), resourceGroupNameParam, mapNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginExportDependencies = &respr
		m.beginExportDependencies.add(req, beginExportDependencies)
	}

	resp, err := server.PollerResponderNext(beginExportDependencies, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		m.beginExportDependencies.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExportDependencies) {
		m.beginExportDependencies.remove(req)
	}

	return resp, nil
}

func (m *MapsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DependencyMap/maps/(?P<mapName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	mapNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mapName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, mapNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).MapsResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MapsServerTransport) dispatchBeginGetConnectionsForProcessOnFocusedMachine(req *http.Request) (*http.Response, error) {
	if m.srv.BeginGetConnectionsForProcessOnFocusedMachine == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetConnectionsForProcessOnFocusedMachine not implemented")}
	}
	beginGetConnectionsForProcessOnFocusedMachine := m.beginGetConnectionsForProcessOnFocusedMachine.get(req)
	if beginGetConnectionsForProcessOnFocusedMachine == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DependencyMap/maps/(?P<mapName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getConnectionsForProcessOnFocusedMachine`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdependencymap.GetConnectionsForProcessOnFocusedMachineRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mapNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mapName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginGetConnectionsForProcessOnFocusedMachine(req.Context(), resourceGroupNameParam, mapNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGetConnectionsForProcessOnFocusedMachine = &respr
		m.beginGetConnectionsForProcessOnFocusedMachine.add(req, beginGetConnectionsForProcessOnFocusedMachine)
	}

	resp, err := server.PollerResponderNext(beginGetConnectionsForProcessOnFocusedMachine, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		m.beginGetConnectionsForProcessOnFocusedMachine.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGetConnectionsForProcessOnFocusedMachine) {
		m.beginGetConnectionsForProcessOnFocusedMachine.remove(req)
	}

	return resp, nil
}

func (m *MapsServerTransport) dispatchBeginGetConnectionsWithConnectedMachineForFocusedMachine(req *http.Request) (*http.Response, error) {
	if m.srv.BeginGetConnectionsWithConnectedMachineForFocusedMachine == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetConnectionsWithConnectedMachineForFocusedMachine not implemented")}
	}
	beginGetConnectionsWithConnectedMachineForFocusedMachine := m.beginGetConnectionsWithConnectedMachineForFocusedMachine.get(req)
	if beginGetConnectionsWithConnectedMachineForFocusedMachine == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DependencyMap/maps/(?P<mapName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getConnectionsWithConnectedMachineForFocusedMachine`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdependencymap.GetConnectionsWithConnectedMachineForFocusedMachineRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mapNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mapName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginGetConnectionsWithConnectedMachineForFocusedMachine(req.Context(), resourceGroupNameParam, mapNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGetConnectionsWithConnectedMachineForFocusedMachine = &respr
		m.beginGetConnectionsWithConnectedMachineForFocusedMachine.add(req, beginGetConnectionsWithConnectedMachineForFocusedMachine)
	}

	resp, err := server.PollerResponderNext(beginGetConnectionsWithConnectedMachineForFocusedMachine, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		m.beginGetConnectionsWithConnectedMachineForFocusedMachine.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGetConnectionsWithConnectedMachineForFocusedMachine) {
		m.beginGetConnectionsWithConnectedMachineForFocusedMachine.remove(req)
	}

	return resp, nil
}

func (m *MapsServerTransport) dispatchBeginGetDependencyViewForFocusedMachine(req *http.Request) (*http.Response, error) {
	if m.srv.BeginGetDependencyViewForFocusedMachine == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetDependencyViewForFocusedMachine not implemented")}
	}
	beginGetDependencyViewForFocusedMachine := m.beginGetDependencyViewForFocusedMachine.get(req)
	if beginGetDependencyViewForFocusedMachine == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DependencyMap/maps/(?P<mapName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getDependencyViewForFocusedMachine`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdependencymap.GetDependencyViewForFocusedMachineRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mapNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mapName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginGetDependencyViewForFocusedMachine(req.Context(), resourceGroupNameParam, mapNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGetDependencyViewForFocusedMachine = &respr
		m.beginGetDependencyViewForFocusedMachine.add(req, beginGetDependencyViewForFocusedMachine)
	}

	resp, err := server.PollerResponderNext(beginGetDependencyViewForFocusedMachine, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		m.beginGetDependencyViewForFocusedMachine.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGetDependencyViewForFocusedMachine) {
		m.beginGetDependencyViewForFocusedMachine.remove(req)
	}

	return resp, nil
}

func (m *MapsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := m.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DependencyMap/maps`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		m.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armdependencymap.MapsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		m.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (m *MapsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := m.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DependencyMap/maps`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := m.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		m.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armdependencymap.MapsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		m.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (m *MapsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := m.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DependencyMap/maps/(?P<mapName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdependencymap.MapsResourceTagsUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mapNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mapName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginUpdate(req.Context(), resourceGroupNameParam, mapNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		m.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		m.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to MapsServerTransport
var mapsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
