//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentscripts

const (
	moduleName    = "armdeploymentscripts"
	moduleVersion = "v0.5.0"
)

// CleanupOptions - The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
type CleanupOptions string

const (
	CleanupOptionsAlways       CleanupOptions = "Always"
	CleanupOptionsOnExpiration CleanupOptions = "OnExpiration"
	CleanupOptionsOnSuccess    CleanupOptions = "OnSuccess"
)

// PossibleCleanupOptionsValues returns the possible values for the CleanupOptions const type.
func PossibleCleanupOptionsValues() []CleanupOptions {
	return []CleanupOptions{
		CleanupOptionsAlways,
		CleanupOptionsOnExpiration,
		CleanupOptionsOnSuccess,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
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

// ManagedServiceIdentityType - Type of the managed identity.
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeUserAssigned ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// ScriptProvisioningState - State of the script execution. This only appears in the response.
type ScriptProvisioningState string

const (
	ScriptProvisioningStateCanceled              ScriptProvisioningState = "Canceled"
	ScriptProvisioningStateCreating              ScriptProvisioningState = "Creating"
	ScriptProvisioningStateFailed                ScriptProvisioningState = "Failed"
	ScriptProvisioningStateProvisioningResources ScriptProvisioningState = "ProvisioningResources"
	ScriptProvisioningStateRunning               ScriptProvisioningState = "Running"
	ScriptProvisioningStateSucceeded             ScriptProvisioningState = "Succeeded"
)

// PossibleScriptProvisioningStateValues returns the possible values for the ScriptProvisioningState const type.
func PossibleScriptProvisioningStateValues() []ScriptProvisioningState {
	return []ScriptProvisioningState{
		ScriptProvisioningStateCanceled,
		ScriptProvisioningStateCreating,
		ScriptProvisioningStateFailed,
		ScriptProvisioningStateProvisioningResources,
		ScriptProvisioningStateRunning,
		ScriptProvisioningStateSucceeded,
	}
}

// ScriptType - Type of the script.
type ScriptType string

const (
	ScriptTypeAzureCLI        ScriptType = "AzureCLI"
	ScriptTypeAzurePowerShell ScriptType = "AzurePowerShell"
)

// PossibleScriptTypeValues returns the possible values for the ScriptType const type.
func PossibleScriptTypeValues() []ScriptType {
	return []ScriptType{
		ScriptTypeAzureCLI,
		ScriptTypeAzurePowerShell,
	}
}
