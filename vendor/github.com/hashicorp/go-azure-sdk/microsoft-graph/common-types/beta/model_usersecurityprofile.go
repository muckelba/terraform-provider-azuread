package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserSecurityProfile{}

type UserSecurityProfile struct {
	Accounts             *[]UserAccount             `json:"accounts,omitempty"`
	AzureSubscriptionId  nullable.Type[string]      `json:"azureSubscriptionId,omitempty"`
	AzureTenantId        *string                    `json:"azureTenantId,omitempty"`
	CreatedDateTime      nullable.Type[string]      `json:"createdDateTime,omitempty"`
	DisplayName          nullable.Type[string]      `json:"displayName,omitempty"`
	LastModifiedDateTime nullable.Type[string]      `json:"lastModifiedDateTime,omitempty"`
	RiskScore            nullable.Type[string]      `json:"riskScore,omitempty"`
	Tags                 *[]string                  `json:"tags,omitempty"`
	UserPrincipalName    nullable.Type[string]      `json:"userPrincipalName,omitempty"`
	VendorInformation    *SecurityVendorInformation `json:"vendorInformation,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s UserSecurityProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserSecurityProfile{}

func (s UserSecurityProfile) MarshalJSON() ([]byte, error) {
	type wrapper UserSecurityProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserSecurityProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserSecurityProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userSecurityProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserSecurityProfile: %+v", err)
	}

	return encoded, nil
}
