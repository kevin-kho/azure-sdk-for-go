//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListAuthorizationProviders.json
func ExampleAuthorizationProviderClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAuthorizationProviderClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.AuthorizationProviderClientListByServiceOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AuthorizationProviderCollection = armapimanagement.AuthorizationProviderCollection{
		// 	Value: []*armapimanagement.AuthorizationProviderContract{
		// 		{
		// 			Name: to.Ptr("aadwithauthcode"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithauthcode"),
		// 			Properties: &armapimanagement.AuthorizationProviderContractProperties{
		// 				DisplayName: to.Ptr("aadwithauthcode"),
		// 				IdentityProvider: to.Ptr("aad"),
		// 				Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
		// 					GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
		// 						AuthorizationCode: map[string]*string{
		// 							"clientId": to.Ptr("53790825-fdd3-4b80-bc7a-4c3aaf25801d"),
		// 							"loginUri": to.Ptr("https://login.windows.net"),
		// 							"resourceUri": to.Ptr("https://graph.microsoft.com"),
		// 							"scopes": to.Ptr("User.Read.All Group.Read.All"),
		// 							"tenantId": to.Ptr("common"),
		// 						},
		// 					},
		// 					RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("aadwithclientcred"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithclientcred"),
		// 			Properties: &armapimanagement.AuthorizationProviderContractProperties{
		// 				DisplayName: to.Ptr("aadwithclientcred"),
		// 				IdentityProvider: to.Ptr("aad"),
		// 				Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
		// 					GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
		// 						ClientCredentials: map[string]*string{
		// 							"loginUri": to.Ptr("https://login.windows.net"),
		// 							"resourceUri": to.Ptr("https://graph.microsoft.com"),
		// 							"scopes": to.Ptr("User.Read.All Group.Read.All"),
		// 							"tenantId": to.Ptr("common"),
		// 						},
		// 					},
		// 					RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("google"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/google"),
		// 			Properties: &armapimanagement.AuthorizationProviderContractProperties{
		// 				DisplayName: to.Ptr("google"),
		// 				IdentityProvider: to.Ptr("google"),
		// 				Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
		// 					GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
		// 						AuthorizationCode: map[string]*string{
		// 							"clientId": to.Ptr("99999999-xxxxxxxxxxxxxxxxxxxxxxxx.apps.googleusercontent.com"),
		// 							"scopes": to.Ptr("openid https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email"),
		// 						},
		// 					},
		// 					RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("eventbrite"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/eventbrite"),
		// 			Properties: &armapimanagement.AuthorizationProviderContractProperties{
		// 				DisplayName: to.Ptr("eventbrite"),
		// 				IdentityProvider: to.Ptr("oauth2"),
		// 				Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
		// 					GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
		// 						AuthorizationCode: map[string]*string{
		// 							"authorizationUrl": to.Ptr("https://www.eventbrite.com/oauth/authorize"),
		// 							"clientId": to.Ptr("clientid"),
		// 							"refreshUrl": to.Ptr("https://www.eventbrite.com/oauth/token"),
		// 							"scopes": nil,
		// 							"tokenUrl": to.Ptr("https://www.eventbrite.com/oauth/token"),
		// 						},
		// 					},
		// 					RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetAuthorizationProvider.json
func ExampleAuthorizationProviderClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationProviderClient().Get(ctx, "rg1", "apimService1", "aadwithauthcode", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationProviderContract = armapimanagement.AuthorizationProviderContract{
	// 	Name: to.Ptr("aadwithauthcode"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithauthcode"),
	// 	Properties: &armapimanagement.AuthorizationProviderContractProperties{
	// 		DisplayName: to.Ptr("aadwithauthcode"),
	// 		IdentityProvider: to.Ptr("aad"),
	// 		Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
	// 			GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
	// 				AuthorizationCode: map[string]*string{
	// 					"clientId": to.Ptr("53790825-fdd3-4b80-bc7a-4c3aaf25801d"),
	// 					"loginUri": to.Ptr("https://login.windows.net"),
	// 					"resourceUri": to.Ptr("https://graph.microsoft.com"),
	// 					"scopes": to.Ptr("User.Read.All Group.Read.All"),
	// 					"tenantId": to.Ptr("common"),
	// 				},
	// 			},
	// 			RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateAuthorizationProviderAADAuthCode.json
func ExampleAuthorizationProviderClient_CreateOrUpdate_apiManagementCreateAuthorizationProviderAadAuthCode() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationProviderClient().CreateOrUpdate(ctx, "rg1", "apimService1", "aadwithauthcode", armapimanagement.AuthorizationProviderContract{
		Properties: &armapimanagement.AuthorizationProviderContractProperties{
			DisplayName:      to.Ptr("aadwithauthcode"),
			IdentityProvider: to.Ptr("aad"),
			Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
				GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
					AuthorizationCode: map[string]*string{
						"clientId":     to.Ptr("clientsecretid"),
						"clientSecret": to.Ptr("clientsecretvalue"),
						"resourceUri":  to.Ptr("https://graph.microsoft.com"),
						"scopes":       to.Ptr("User.Read.All Group.Read.All"),
					},
				},
				RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
			},
		},
	}, &armapimanagement.AuthorizationProviderClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationProviderContract = armapimanagement.AuthorizationProviderContract{
	// 	Name: to.Ptr("aadwithauthcode"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithauthcode"),
	// 	Properties: &armapimanagement.AuthorizationProviderContractProperties{
	// 		DisplayName: to.Ptr("aadwithauthcode"),
	// 		IdentityProvider: to.Ptr("aad"),
	// 		Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
	// 			GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
	// 				AuthorizationCode: map[string]*string{
	// 					"clientId": to.Ptr("53790825-fdd3-4b80-bc7a-4c3aaf25801d"),
	// 					"loginUri": to.Ptr("https://login.windows.net"),
	// 					"resourceUri": to.Ptr("https://graph.microsoft.com"),
	// 					"scopes": to.Ptr("User.Read.All Group.Read.All"),
	// 					"tenantId": to.Ptr("common"),
	// 				},
	// 			},
	// 			RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateAuthorizationProviderAADClientCred.json
func ExampleAuthorizationProviderClient_CreateOrUpdate_apiManagementCreateAuthorizationProviderAadClientCred() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationProviderClient().CreateOrUpdate(ctx, "rg1", "apimService1", "aadwithclientcred", armapimanagement.AuthorizationProviderContract{
		Properties: &armapimanagement.AuthorizationProviderContractProperties{
			DisplayName:      to.Ptr("aadwithclientcred"),
			IdentityProvider: to.Ptr("aad"),
			Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
				GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
					AuthorizationCode: map[string]*string{
						"resourceUri": to.Ptr("https://graph.microsoft.com"),
						"scopes":      to.Ptr("User.Read.All Group.Read.All"),
					},
				},
				RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
			},
		},
	}, &armapimanagement.AuthorizationProviderClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationProviderContract = armapimanagement.AuthorizationProviderContract{
	// 	Name: to.Ptr("aadwithclientcred"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/aadwithclientcred"),
	// 	Properties: &armapimanagement.AuthorizationProviderContractProperties{
	// 		DisplayName: to.Ptr("aadwithclientcred"),
	// 		IdentityProvider: to.Ptr("aad"),
	// 		Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
	// 			GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
	// 				ClientCredentials: map[string]*string{
	// 					"loginUri": to.Ptr("https://login.windows.net"),
	// 					"resourceUri": to.Ptr("https://graph.microsoft.com"),
	// 					"scopes": to.Ptr("User.Read.All Group.Read.All"),
	// 					"tenantId": to.Ptr("common"),
	// 				},
	// 			},
	// 			RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateAuthorizationProviderGenericOAuth2.json
func ExampleAuthorizationProviderClient_CreateOrUpdate_apiManagementCreateAuthorizationProviderGenericOAuth2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationProviderClient().CreateOrUpdate(ctx, "rg1", "apimService1", "eventbrite", armapimanagement.AuthorizationProviderContract{
		Properties: &armapimanagement.AuthorizationProviderContractProperties{
			DisplayName:      to.Ptr("eventbrite"),
			IdentityProvider: to.Ptr("oauth2"),
			Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
				GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
					AuthorizationCode: map[string]*string{
						"authorizationUrl": to.Ptr("https://www.eventbrite.com/oauth/authorize"),
						"clientId":         to.Ptr("clientid"),
						"clientSecret":     to.Ptr("clientsecretvalue"),
						"refreshUrl":       to.Ptr("https://www.eventbrite.com/oauth/token"),
						"scopes":           nil,
						"tokenUrl":         to.Ptr("https://www.eventbrite.com/oauth/token"),
					},
				},
				RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
			},
		},
	}, &armapimanagement.AuthorizationProviderClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationProviderContract = armapimanagement.AuthorizationProviderContract{
	// 	Name: to.Ptr("eventbrite"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/eventbrite"),
	// 	Properties: &armapimanagement.AuthorizationProviderContractProperties{
	// 		DisplayName: to.Ptr("eventbrite"),
	// 		IdentityProvider: to.Ptr("oauth2"),
	// 		Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
	// 			GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
	// 				AuthorizationCode: map[string]*string{
	// 					"authorizationUrl": to.Ptr("https://www.eventbrite.com/oauth/authorize"),
	// 					"clientId": to.Ptr("clientid"),
	// 					"refreshUrl": to.Ptr("https://www.eventbrite.com/oauth/token"),
	// 					"scopes": nil,
	// 					"tokenUrl": to.Ptr("https://www.eventbrite.com/oauth/token"),
	// 				},
	// 			},
	// 			RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateAuthorizationProviderOOBGoogle.json
func ExampleAuthorizationProviderClient_CreateOrUpdate_apiManagementCreateAuthorizationProviderOobGoogle() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizationProviderClient().CreateOrUpdate(ctx, "rg1", "apimService1", "google", armapimanagement.AuthorizationProviderContract{
		Properties: &armapimanagement.AuthorizationProviderContractProperties{
			DisplayName:      to.Ptr("google"),
			IdentityProvider: to.Ptr("google"),
			Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
				GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
					AuthorizationCode: map[string]*string{
						"clientId":     to.Ptr("99999999-xxxxxxxxxxxxxxxxxxxxxxxx.apps.googleusercontent.com"),
						"clientSecret": to.Ptr("clientsecretvalue"),
						"scopes":       to.Ptr("openid https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email"),
					},
				},
				RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
			},
		},
	}, &armapimanagement.AuthorizationProviderClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationProviderContract = armapimanagement.AuthorizationProviderContract{
	// 	Name: to.Ptr("google"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/authorizationProviders"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/authorizationProviders/google"),
	// 	Properties: &armapimanagement.AuthorizationProviderContractProperties{
	// 		DisplayName: to.Ptr("google"),
	// 		IdentityProvider: to.Ptr("google"),
	// 		Oauth2: &armapimanagement.AuthorizationProviderOAuth2Settings{
	// 			GrantTypes: &armapimanagement.AuthorizationProviderOAuth2GrantTypes{
	// 				AuthorizationCode: map[string]*string{
	// 					"clientId": to.Ptr("99999999-xxxxxxxxxxxxxxxxxxxxxxxx.apps.googleusercontent.com"),
	// 					"scopes": to.Ptr("openid https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email"),
	// 				},
	// 			},
	// 			RedirectURL: to.Ptr("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementDeleteAuthorizationProvider.json
func ExampleAuthorizationProviderClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAuthorizationProviderClient().Delete(ctx, "rg1", "apimService1", "aadwithauthcode", "*", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
