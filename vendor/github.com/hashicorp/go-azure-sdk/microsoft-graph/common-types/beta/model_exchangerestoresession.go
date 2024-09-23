package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RestoreSessionBase = ExchangeRestoreSession{}

type ExchangeRestoreSession struct {
	// A collection of restore points and destination details that can be used to restore Exchange mailboxes.
	MailboxRestoreArtifacts *[]MailboxRestoreArtifact `json:"mailboxRestoreArtifacts,omitempty"`

	// Fields inherited from RestoreSessionBase

	// The time of completion of the restore session.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The identity of person who created the restore session.
	CreatedBy IdentitySet `json:"createdBy"`

	// The time of creation of the restore session.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Contains error details if the restore session fails or completes with an error.
	Error *PublicError `json:"error,omitempty"`

	// Identity of the person who last modified the restore session.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// Timestamp of the last modification of the restore session.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Status of the restore session. The value is an aggregated status of the restored artifacts. The possible values are:
	// draft, activating, active, completedWithError, completed, unknownFutureValue, failed. You must use the Prefer:
	// include-unknown-enum-members request header to get the following value in this evolvable enum: failed.
	Status *RestoreSessionStatus `json:"status,omitempty"`

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

func (s ExchangeRestoreSession) RestoreSessionBase() BaseRestoreSessionBaseImpl {
	return BaseRestoreSessionBaseImpl{
		CompletedDateTime:    s.CompletedDateTime,
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		Error:                s.Error,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Status:               s.Status,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s ExchangeRestoreSession) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExchangeRestoreSession{}

func (s ExchangeRestoreSession) MarshalJSON() ([]byte, error) {
	type wrapper ExchangeRestoreSession
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExchangeRestoreSession: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExchangeRestoreSession: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.exchangeRestoreSession"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExchangeRestoreSession: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ExchangeRestoreSession{}

func (s *ExchangeRestoreSession) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		MailboxRestoreArtifacts *[]MailboxRestoreArtifact `json:"mailboxRestoreArtifacts,omitempty"`
		CompletedDateTime       nullable.Type[string]     `json:"completedDateTime,omitempty"`
		CreatedDateTime         nullable.Type[string]     `json:"createdDateTime,omitempty"`
		Error                   *PublicError              `json:"error,omitempty"`
		LastModifiedDateTime    nullable.Type[string]     `json:"lastModifiedDateTime,omitempty"`
		Status                  *RestoreSessionStatus     `json:"status,omitempty"`
		Id                      *string                   `json:"id,omitempty"`
		ODataId                 *string                   `json:"@odata.id,omitempty"`
		ODataType               *string                   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.MailboxRestoreArtifacts = decoded.MailboxRestoreArtifacts
	s.CompletedDateTime = decoded.CompletedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Error = decoded.Error
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ExchangeRestoreSession into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'ExchangeRestoreSession': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'ExchangeRestoreSession': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
