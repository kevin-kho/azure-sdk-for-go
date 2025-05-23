//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationFabrics_List.json
func ExampleReplicationFabricsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationFabricsClient().NewListPager("resourceGroupPS1", "vault1", nil)
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
		// page.FabricCollection = armrecoveryservicessiterecovery.FabricCollection{
		// 	Value: []*armrecoveryservicessiterecovery.Fabric{
		// 		{
		// 			Name: to.Ptr("cloud1"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics"),
		// 			ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
		// 			Properties: &armrecoveryservicessiterecovery.FabricProperties{
		// 				BcdrState: to.Ptr("Valid"),
		// 				CustomDetails: &armrecoveryservicessiterecovery.HyperVSiteDetails{
		// 					InstanceType: to.Ptr("HyperVSite"),
		// 				},
		// 				EncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
		// 					KekState: to.Ptr("None"),
		// 				},
		// 				FriendlyName: to.Ptr("cloud1"),
		// 				Health: to.Ptr("Normal"),
		// 				HealthErrorDetails: []*armrecoveryservicessiterecovery.HealthError{
		// 				},
		// 				InternalIdentifier: to.Ptr("6d224fc6-f326-5d35-96de-fbf51efb3179"),
		// 				RolloverEncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
		// 					KekState: to.Ptr("None"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationFabrics_Get.json
func ExampleReplicationFabricsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationFabricsClient().Get(ctx, "resourceGroupPS1", "vault1", "cloud1", &armrecoveryservicessiterecovery.ReplicationFabricsClientGetOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Fabric = armrecoveryservicessiterecovery.Fabric{
	// 	Name: to.Ptr("cloud1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
	// 	Properties: &armrecoveryservicessiterecovery.FabricProperties{
	// 		BcdrState: to.Ptr("Valid"),
	// 		CustomDetails: &armrecoveryservicessiterecovery.HyperVSiteDetails{
	// 			InstanceType: to.Ptr("HyperVSite"),
	// 		},
	// 		EncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
	// 			KekState: to.Ptr("None"),
	// 		},
	// 		FriendlyName: to.Ptr("cloud1"),
	// 		Health: to.Ptr("Normal"),
	// 		HealthErrorDetails: []*armrecoveryservicessiterecovery.HealthError{
	// 		},
	// 		InternalIdentifier: to.Ptr("6d224fc6-f326-5d35-96de-fbf51efb3179"),
	// 		RolloverEncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
	// 			KekState: to.Ptr("None"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationFabrics_Create.json
func ExampleReplicationFabricsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationFabricsClient().BeginCreate(ctx, "resourceGroupPS1", "vault1", "cloud1", armrecoveryservicessiterecovery.FabricCreationInput{
		Properties: &armrecoveryservicessiterecovery.FabricCreationInputProperties{
			CustomDetails: nil,
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Fabric = armrecoveryservicessiterecovery.Fabric{
	// 	Name: to.Ptr("cloud1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
	// 	Properties: &armrecoveryservicessiterecovery.FabricProperties{
	// 		BcdrState: to.Ptr("Valid"),
	// 		CustomDetails: &armrecoveryservicessiterecovery.HyperVSiteDetails{
	// 			InstanceType: to.Ptr("HyperVSite"),
	// 		},
	// 		EncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
	// 			KekState: to.Ptr("None"),
	// 		},
	// 		FriendlyName: to.Ptr("cloud1"),
	// 		Health: to.Ptr("Normal"),
	// 		HealthErrorDetails: []*armrecoveryservicessiterecovery.HealthError{
	// 		},
	// 		InternalIdentifier: to.Ptr("6d224fc6-f326-5d35-96de-fbf51efb3179"),
	// 		RolloverEncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
	// 			KekState: to.Ptr("None"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationFabrics_Purge.json
func ExampleReplicationFabricsClient_BeginPurge() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationFabricsClient().BeginPurge(ctx, "resourceGroupPS1", "vault1", "cloud1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationFabrics_CheckConsistency.json
func ExampleReplicationFabricsClient_BeginCheckConsistency() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationFabricsClient().BeginCheckConsistency(ctx, "resourceGroupPS1", "vault1", "cloud1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Fabric = armrecoveryservicessiterecovery.Fabric{
	// 	Name: to.Ptr("cloud1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
	// 	Properties: &armrecoveryservicessiterecovery.FabricProperties{
	// 		BcdrState: to.Ptr("Valid"),
	// 		CustomDetails: &armrecoveryservicessiterecovery.HyperVSiteDetails{
	// 			InstanceType: to.Ptr("HyperVSite"),
	// 		},
	// 		EncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
	// 			KekState: to.Ptr("None"),
	// 		},
	// 		FriendlyName: to.Ptr("cloud1"),
	// 		Health: to.Ptr("Normal"),
	// 		HealthErrorDetails: []*armrecoveryservicessiterecovery.HealthError{
	// 		},
	// 		InternalIdentifier: to.Ptr("6d224fc6-f326-5d35-96de-fbf51efb3179"),
	// 		RolloverEncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
	// 			KekState: to.Ptr("None"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationFabrics_MigrateToAad.json
func ExampleReplicationFabricsClient_BeginMigrateToAAD() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationFabricsClient().BeginMigrateToAAD(ctx, "resourceGroupPS1", "vault1", "cloud1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationFabrics_ReassociateGateway.json
func ExampleReplicationFabricsClient_BeginReassociateGateway() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationFabricsClient().BeginReassociateGateway(ctx, "MadhaviVRG", "MadhaviVault", "GRACE-V2A-1", armrecoveryservicessiterecovery.FailoverProcessServerRequest{
		Properties: &armrecoveryservicessiterecovery.FailoverProcessServerRequestProperties{
			ContainerName:         to.Ptr("cloud_1f3c15af-2256-4568-9e06-e1ef4f728f75"),
			SourceProcessServerID: to.Ptr("AFA0EC54-1894-4E44-9CAB02DB8854B117"),
			TargetProcessServerID: to.Ptr("5D3ED340-85AE-C646-B338641E015DA405"),
			UpdateType:            to.Ptr("ServerLevel"),
			VMsToMigrate: []*string{
				to.Ptr("Vm1"),
				to.Ptr("Vm2")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Fabric = armrecoveryservicessiterecovery.Fabric{
	// 	Name: to.Ptr("bc15edf300344660d9c2965f5d9225593d99734f6580613c997033abc3512a56"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics"),
	// 	ID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/MadhaviVRG/providers/Microsoft.RecoveryServices/vaults/MadhaviVault/replicationFabrics/bc15edf300344660d9c2965f5d9225593d99734f6580613c997033abc3512a56"),
	// 	Properties: &armrecoveryservicessiterecovery.FabricProperties{
	// 		BcdrState: to.Ptr("Valid"),
	// 		CustomDetails: &armrecoveryservicessiterecovery.VMwareDetails{
	// 			InstanceType: to.Ptr("VMware"),
	// 		},
	// 		EncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
	// 			KekState: to.Ptr("None"),
	// 		},
	// 		FriendlyName: to.Ptr("GRACE-V2A-1"),
	// 		Health: to.Ptr("Normal"),
	// 		HealthErrorDetails: []*armrecoveryservicessiterecovery.HealthError{
	// 		},
	// 		InternalIdentifier: to.Ptr("1f3c15af-2256-4568-9e06-e1ef4f728f75"),
	// 		RolloverEncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
	// 			KekState: to.Ptr("None"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationFabrics_Delete.json
func ExampleReplicationFabricsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationFabricsClient().BeginDelete(ctx, "resourceGroupPS1", "vault1", "cloud1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationFabrics_RenewCertificate.json
func ExampleReplicationFabricsClient_BeginRenewCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationFabricsClient().BeginRenewCertificate(ctx, "resourceGroupPS1", "vault1", "cloud1", armrecoveryservicessiterecovery.RenewCertificateInput{
		Properties: &armrecoveryservicessiterecovery.RenewCertificateInputProperties{
			RenewCertificateType: to.Ptr("Cloud"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Fabric = armrecoveryservicessiterecovery.Fabric{
	// 	Name: to.Ptr("cloud1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
	// 	Properties: &armrecoveryservicessiterecovery.FabricProperties{
	// 		BcdrState: to.Ptr("Valid"),
	// 		CustomDetails: &armrecoveryservicessiterecovery.HyperVSiteDetails{
	// 			InstanceType: to.Ptr("HyperVSite"),
	// 		},
	// 		EncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
	// 			KekState: to.Ptr("None"),
	// 		},
	// 		FriendlyName: to.Ptr("cloud1"),
	// 		Health: to.Ptr("Normal"),
	// 		HealthErrorDetails: []*armrecoveryservicessiterecovery.HealthError{
	// 		},
	// 		InternalIdentifier: to.Ptr("6d224fc6-f326-5d35-96de-fbf51efb3179"),
	// 		RolloverEncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
	// 			KekState: to.Ptr("None"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationInfrastructure_Delete.json
func ExampleReplicationFabricsClient_BeginRemoveInfra() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationFabricsClient().BeginRemoveInfra(ctx, "resourceGroupPS1", "vault1", "cloud1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
