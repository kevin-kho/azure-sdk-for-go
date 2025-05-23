// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armservicenetworking

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
	moduleVersion = "v1.2.0-beta.1"
)

// ActionType - Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	// ActionTypeInternal - Actions are for internal-only APIs.
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// AssociationType - Association Type Enum
type AssociationType string

const (
	// AssociationTypeSubnets - Association of Type Subnet
	AssociationTypeSubnets AssociationType = "subnets"
)

// PossibleAssociationTypeValues returns the possible values for the AssociationType const type.
func PossibleAssociationTypeValues() []AssociationType {
	return []AssociationType{
		AssociationTypeSubnets,
	}
}

// CreatedByType - The kind of entity that created the resource.
type CreatedByType string

const (
	// CreatedByTypeApplication - The entity was created by an application.
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey - The entity was created by a key.
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity - The entity was created by a managed identity.
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser - The entity was created by a user.
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// IPAccessRuleAction - Action of Ip Access Rule
type IPAccessRuleAction string

const (
	// IPAccessRuleActionAllow - Allow Source Ip Prefixes
	IPAccessRuleActionAllow IPAccessRuleAction = "allow"
	// IPAccessRuleActionDeny - Deny Source Ip Prefixes
	IPAccessRuleActionDeny IPAccessRuleAction = "deny"
)

// PossibleIPAccessRuleActionValues returns the possible values for the IPAccessRuleAction const type.
func PossibleIPAccessRuleActionValues() []IPAccessRuleAction {
	return []IPAccessRuleAction{
		IPAccessRuleActionAllow,
		IPAccessRuleActionDeny,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	// OriginSystem - Indicates the operation is initiated by a system.
	OriginSystem Origin = "system"
	// OriginUser - Indicates the operation is initiated by a user.
	OriginUser Origin = "user"
	// OriginUserSystem - Indicates the operation is initiated by a user or system.
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// PolicyType - Policy Type of the Security Policy
type PolicyType string

const (
	// PolicyTypeIPAccessRules - Policy of Type IpAccessRules
	PolicyTypeIPAccessRules PolicyType = "ipAccessRules"
	// PolicyTypeWAF - Policy of Type WAF
	PolicyTypeWAF PolicyType = "waf"
)

// PossiblePolicyTypeValues returns the possible values for the PolicyType const type.
func PossiblePolicyTypeValues() []PolicyType {
	return []PolicyType{
		PolicyTypeIPAccessRules,
		PolicyTypeWAF,
	}
}

// ProvisioningState - Resource Provisioning State Enum
type ProvisioningState string

const (
	// ProvisioningStateAccepted - Resource in Accepted State
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Resource in Canceled State
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleting - Resource in Deleting State
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource in Failed State
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateProvisioning - Resource in Provisioning State
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	// ProvisioningStateSucceeded - Resource in Succeeded State
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - Resource in Updating State
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}
