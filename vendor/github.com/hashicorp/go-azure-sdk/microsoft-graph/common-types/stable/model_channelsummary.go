package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelSummary struct {
	// Count of guests in a channel.
	GuestsCount nullable.Type[int64] `json:"guestsCount,omitempty"`

	// Indicates whether external members are included on the channel.
	HasMembersFromOtherTenants nullable.Type[bool] `json:"hasMembersFromOtherTenants,omitempty"`

	// Count of members in a channel.
	MembersCount nullable.Type[int64] `json:"membersCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Count of owners in a channel.
	OwnersCount nullable.Type[int64] `json:"ownersCount,omitempty"`
}
