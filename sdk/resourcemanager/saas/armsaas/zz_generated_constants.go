//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsaas

const (
	moduleName    = "armsaas"
	moduleVersion = "v0.4.0"
)

// PaymentChannelType - The Payment channel for the SaasSubscription.
type PaymentChannelType string

const (
	PaymentChannelTypeCustomerDelegated     PaymentChannelType = "CustomerDelegated"
	PaymentChannelTypeSubscriptionDelegated PaymentChannelType = "SubscriptionDelegated"
)

// PossiblePaymentChannelTypeValues returns the possible values for the PaymentChannelType const type.
func PossiblePaymentChannelTypeValues() []PaymentChannelType {
	return []PaymentChannelType{
		PaymentChannelTypeCustomerDelegated,
		PaymentChannelTypeSubscriptionDelegated,
	}
}

// SaasAppStatus - the Saas resource status.
type SaasAppStatus string

const (
	SaasAppStatusDeactivated  SaasAppStatus = "Deactivated"
	SaasAppStatusPending      SaasAppStatus = "Pending"
	SaasAppStatusSubscribed   SaasAppStatus = "Subscribed"
	SaasAppStatusSuspended    SaasAppStatus = "Suspended"
	SaasAppStatusUnsubscribed SaasAppStatus = "Unsubscribed"
)

// PossibleSaasAppStatusValues returns the possible values for the SaasAppStatus const type.
func PossibleSaasAppStatusValues() []SaasAppStatus {
	return []SaasAppStatus{
		SaasAppStatusDeactivated,
		SaasAppStatusPending,
		SaasAppStatusSubscribed,
		SaasAppStatusSuspended,
		SaasAppStatusUnsubscribed,
	}
}

// SaasResourceStatus - The SaaS Subscription Status.
type SaasResourceStatus string

const (
	SaasResourceStatusNotStarted              SaasResourceStatus = "NotStarted"
	SaasResourceStatusPendingFulfillmentStart SaasResourceStatus = "PendingFulfillmentStart"
	SaasResourceStatusSubscribed              SaasResourceStatus = "Subscribed"
	SaasResourceStatusSuspended               SaasResourceStatus = "Suspended"
	SaasResourceStatusUnsubscribed            SaasResourceStatus = "Unsubscribed"
)

// PossibleSaasResourceStatusValues returns the possible values for the SaasResourceStatus const type.
func PossibleSaasResourceStatusValues() []SaasResourceStatus {
	return []SaasResourceStatus{
		SaasResourceStatusNotStarted,
		SaasResourceStatusPendingFulfillmentStart,
		SaasResourceStatusSubscribed,
		SaasResourceStatusSuspended,
		SaasResourceStatusUnsubscribed,
	}
}
