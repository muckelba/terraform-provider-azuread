package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthorizationSystemTypeAction = GcpAuthorizationSystemTypeAction{}

type GcpAuthorizationSystemTypeAction struct {
	Service *AuthorizationSystemTypeService `json:"service,omitempty"`

	// Fields inherited from AuthorizationSystemTypeAction

	// The type of action allowed in the authorization system's service. The possible values are: delete, read,
	// unknownFutureValue. Supports $filter and (eq).
	ActionType *AuthorizationSystemActionType `json:"actionType,omitempty"`

	// The display name of an action. Read-only. Supports $filter and (eq).
	ExternalId *string `json:"externalId,omitempty"`

	// The resource types in the authorization system's service where the action can be performed. Supports $filter and
	// (eq).
	ResourceTypes *[]string `json:"resourceTypes,omitempty"`

	Severity *AuthorizationSystemActionSeverity `json:"severity,omitempty"`

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

func (s GcpAuthorizationSystemTypeAction) AuthorizationSystemTypeAction() BaseAuthorizationSystemTypeActionImpl {
	return BaseAuthorizationSystemTypeActionImpl{
		ActionType:    s.ActionType,
		ExternalId:    s.ExternalId,
		ResourceTypes: s.ResourceTypes,
		Severity:      s.Severity,
		Id:            s.Id,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
	}
}

func (s GcpAuthorizationSystemTypeAction) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GcpAuthorizationSystemTypeAction{}

func (s GcpAuthorizationSystemTypeAction) MarshalJSON() ([]byte, error) {
	type wrapper GcpAuthorizationSystemTypeAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GcpAuthorizationSystemTypeAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GcpAuthorizationSystemTypeAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.gcpAuthorizationSystemTypeAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GcpAuthorizationSystemTypeAction: %+v", err)
	}

	return encoded, nil
}
