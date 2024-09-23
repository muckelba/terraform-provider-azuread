package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CallRecordsParticipantBase = CallRecordsOrganizer{}

type CallRecordsOrganizer struct {

	// Fields inherited from CallRecordsParticipantBase

	// List of administrativeUnitInfo of the call participant.
	AdministrativeUnitInfos *[]CallRecordsAdministrativeUnitInfo `json:"administrativeUnitInfos,omitempty"`

	// The identity of the call participant.
	Identity *CommunicationsIdentitySet `json:"identity,omitempty"`

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

func (s CallRecordsOrganizer) CallRecordsParticipantBase() BaseCallRecordsParticipantBaseImpl {
	return BaseCallRecordsParticipantBaseImpl{
		AdministrativeUnitInfos: s.AdministrativeUnitInfos,
		Identity:                s.Identity,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s CallRecordsOrganizer) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CallRecordsOrganizer{}

func (s CallRecordsOrganizer) MarshalJSON() ([]byte, error) {
	type wrapper CallRecordsOrganizer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallRecordsOrganizer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsOrganizer: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callRecords.organizer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallRecordsOrganizer: %+v", err)
	}

	return encoded, nil
}
