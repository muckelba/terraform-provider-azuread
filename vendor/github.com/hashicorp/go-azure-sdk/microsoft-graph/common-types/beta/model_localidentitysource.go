package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PermissionsDefinitionIdentitySource = LocalIdentitySource{}

type LocalIdentitySource struct {

	// Fields inherited from PermissionsDefinitionIdentitySource

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s LocalIdentitySource) PermissionsDefinitionIdentitySource() BasePermissionsDefinitionIdentitySourceImpl {
	return BasePermissionsDefinitionIdentitySourceImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = LocalIdentitySource{}

func (s LocalIdentitySource) MarshalJSON() ([]byte, error) {
	type wrapper LocalIdentitySource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LocalIdentitySource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LocalIdentitySource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.localIdentitySource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LocalIdentitySource: %+v", err)
	}

	return encoded, nil
}
