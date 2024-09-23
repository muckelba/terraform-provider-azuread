package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHuntingQueryResults struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The results of the hunting query.
	Results *[]SecurityHuntingRowResult `json:"results,omitempty"`

	// The schema for the response.
	Schema *[]SecuritySinglePropertySchema `json:"schema,omitempty"`
}
