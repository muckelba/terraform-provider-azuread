package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessApplications struct {
	ApplicationFilter *ConditionalAccessFilter `json:"applicationFilter,omitempty"`

	// Can be one of the following: The list of client IDs (appId) explicitly excluded from the policy. Office365 - For the
	// list of apps included in Office365, see Apps included in Conditional Access Office 365 app suite
	// MicrosoftAdminPortals - For more information, see Conditional Access Target resources: Microsoft Admin Portals
	ExcludeApplications *[]string `json:"excludeApplications,omitempty"`

	// Can be one of the following: The list of client IDs (appId) the policy applies to, unless explicitly excluded (in
	// excludeApplications) All Office365 - For the list of apps included in Office365, see Apps included in Conditional
	// Access Office 365 app suite MicrosoftAdminPortals - For more information, see Conditional Access Target resources:
	// Microsoft Admin Portals
	IncludeApplications *[]string `json:"includeApplications,omitempty"`

	IncludeAuthenticationContextClassReferences *[]string `json:"includeAuthenticationContextClassReferences,omitempty"`

	// User actions to include. Supported values are urn:user:registersecurityinfo and urn:user:registerdevice
	IncludeUserActions *[]string `json:"includeUserActions,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
