package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCProvisioningPolicy{}

type CloudPCProvisioningPolicy struct {
	// The URL of the alternate resource that links to this provisioning policy. Read-only.
	AlternateResourceUrl nullable.Type[string] `json:"alternateResourceUrl,omitempty"`

	// A defined collection of provisioning policy assignments. Represents the set of Microsoft 365 groups and security
	// groups in Microsoft Entra ID that have provisioning policy assigned. Returned only on $expand. For an example about
	// how to get the assignments relationship, see Get cloudPcProvisioningPolicy.
	Assignments *[]CloudPCProvisioningPolicyAssignment `json:"assignments,omitempty"`

	// The display name of the Cloud PC group that the Cloud PCs reside in. Read-only.
	CloudPCGroupDisplayName nullable.Type[string] `json:"cloudPcGroupDisplayName,omitempty"`

	// The template used to name Cloud PCs provisioned using this policy. The template can contain custom text and
	// replacement tokens, including %USERNAME:x% and %RAND:x%, which represent the user's name and a randomly generated
	// number, respectively. For example, CPC-%USERNAME:4%-%RAND:5% means that the name of the Cloud PC starts with CPC-,
	// followed by a four-character username, a - character, and then five random characters. The total length of the text
	// generated by the template can't exceed 15 characters. Supports $filter, $select, and $orderby.
	CloudPCNamingTemplate nullable.Type[string] `json:"cloudPcNamingTemplate,omitempty"`

	// The provisioning policy description. Supports $filter, $select, and $orderBy.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name for the provisioning policy.
	DisplayName *string `json:"displayName,omitempty"`

	// Specifies a list ordered by priority on how Cloud PCs join Microsoft Entra ID (Azure AD). Supports $select.
	DomainJoinConfigurations *[]CloudPCDomainJoinConfiguration `json:"domainJoinConfigurations,omitempty"`

	// True if the provisioned Cloud PC can be accessed by single sign-on. False indicates that the provisioned Cloud PC
	// doesn't support this feature. The default value is false. Windows 365 users can use single sign-on to authenticate to
	// Microsoft Entra ID with passwordless options (for example, FIDO keys) to access their Cloud PC. Optional.
	EnableSingleSignOn nullable.Type[bool] `json:"enableSingleSignOn,omitempty"`

	// The number of hours to wait before reprovisioning/deprovisioning happens. Read-only.
	GracePeriodInHours nullable.Type[int64] `json:"gracePeriodInHours,omitempty"`

	// The display name of the operating system image that is used for provisioning. For example, Windows 11 Preview +
	// Microsoft 365 Apps 23H2 23H2. Supports $filter, $select, and $orderBy.
	ImageDisplayName *string `json:"imageDisplayName,omitempty"`

	// The unique identifier that represents an operating system image that is used for provisioning new Cloud PCs. The
	// format for a gallery type image is: {publisherNameofferNameskuName}. Supported values for each of the parameters
	// are:publisher: Microsoftwindowsdesktop offer: windows-ent-cpc sku: 21h1-ent-cpc-m365, 21h1-ent-cpc-os,
	// 20h2-ent-cpc-m365, 20h2-ent-cpc-os, 20h1-ent-cpc-m365, 20h1-ent-cpc-os, 19h2-ent-cpc-m365, and 19h2-ent-cpc-os
	// Supports $filter, $select, and $orderBy.
	ImageId *string `json:"imageId,omitempty"`

	ImageType *CloudPCProvisioningPolicyImageType `json:"imageType,omitempty"`

	// When true, the local admin is enabled for Cloud PCs; false indicates that the local admin isn't enabled for Cloud
	// PCs. The default value is false. Supports $filter, $select, and $orderBy.
	LocalAdminEnabled nullable.Type[bool] `json:"localAdminEnabled,omitempty"`

	// The specific settings to microsoftManagedDesktop that enables Microsoft Managed Desktop customers to get device
	// managed experience for Cloud PC. To enable microsoftManagedDesktop to provide more value, an admin needs to specify
	// certain settings in it. Supports $filter, $select, and $orderBy.
	MicrosoftManagedDesktop *MicrosoftManagedDesktop `json:"microsoftManagedDesktop,omitempty"`

	// Specifies the type of license used when provisioning Cloud PCs using this policy. By default, the license type is
	// dedicated if the provisioningType isn't specified when you create the cloudPcProvisioningPolicy. You can't change
	// this property after the cloudPcProvisioningPolicy was created. Possible values are: dedicated, shared,
	// unknownFutureValue.
	ProvisioningType *CloudPCProvisioningType `json:"provisioningType,omitempty"`

	// Indicates a specific Windows setting to configure during the creation of Cloud PCs for this provisioning policy.
	// Supports $select.
	WindowsSetting *CloudPCWindowsSetting `json:"windowsSetting,omitempty"`

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

func (s CloudPCProvisioningPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCProvisioningPolicy{}

func (s CloudPCProvisioningPolicy) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCProvisioningPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCProvisioningPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCProvisioningPolicy: %+v", err)
	}

	delete(decoded, "alternateResourceUrl")
	delete(decoded, "cloudPcGroupDisplayName")
	delete(decoded, "gracePeriodInHours")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcProvisioningPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCProvisioningPolicy: %+v", err)
	}

	return encoded, nil
}
