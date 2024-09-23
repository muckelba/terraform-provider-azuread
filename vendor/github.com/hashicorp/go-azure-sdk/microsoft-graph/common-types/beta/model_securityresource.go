package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityResource struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Name of the resource that is related to current alert. Required.
	Resource nullable.Type[string] `json:"resource,omitempty"`

	// Represents type of security resources related to an alert. Possible values are: attacked, related.
	ResourceType *SecurityResourceType `json:"resourceType,omitempty"`
}
