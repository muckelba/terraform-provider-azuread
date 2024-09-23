package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityCaseOperation = SecurityEdiscoveryHoldOperation{}

type SecurityEdiscoveryHoldOperation struct {

	// Fields inherited from SecurityCaseOperation

	// The type of action the operation represents. Possible values are:
	// addToReviewSet,applyTags,contentExport,convertToPdf,estimateStatistics, purgeData
	Action *SecurityCaseAction `json:"action,omitempty"`

	// The date and time the operation was completed.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The user that created the operation.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time the operation was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The progress of the operation.
	PercentProgress nullable.Type[int64] `json:"percentProgress,omitempty"`

	// Contains success and failure-specific result information.
	ResultInfo *ResultInfo `json:"resultInfo,omitempty"`

	// The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded,
	// partiallySucceeded, failed.
	Status *SecurityCaseOperationStatus `json:"status,omitempty"`

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

func (s SecurityEdiscoveryHoldOperation) SecurityCaseOperation() BaseSecurityCaseOperationImpl {
	return BaseSecurityCaseOperationImpl{
		Action:            s.Action,
		CompletedDateTime: s.CompletedDateTime,
		CreatedBy:         s.CreatedBy,
		CreatedDateTime:   s.CreatedDateTime,
		PercentProgress:   s.PercentProgress,
		ResultInfo:        s.ResultInfo,
		Status:            s.Status,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s SecurityEdiscoveryHoldOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEdiscoveryHoldOperation{}

func (s SecurityEdiscoveryHoldOperation) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEdiscoveryHoldOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEdiscoveryHoldOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEdiscoveryHoldOperation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.ediscoveryHoldOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEdiscoveryHoldOperation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityEdiscoveryHoldOperation{}

func (s *SecurityEdiscoveryHoldOperation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Action            *SecurityCaseAction          `json:"action,omitempty"`
		CompletedDateTime nullable.Type[string]        `json:"completedDateTime,omitempty"`
		CreatedDateTime   nullable.Type[string]        `json:"createdDateTime,omitempty"`
		PercentProgress   nullable.Type[int64]         `json:"percentProgress,omitempty"`
		ResultInfo        *ResultInfo                  `json:"resultInfo,omitempty"`
		Status            *SecurityCaseOperationStatus `json:"status,omitempty"`
		Id                *string                      `json:"id,omitempty"`
		ODataId           *string                      `json:"@odata.id,omitempty"`
		ODataType         *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Action = decoded.Action
	s.CompletedDateTime = decoded.CompletedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PercentProgress = decoded.PercentProgress
	s.ResultInfo = decoded.ResultInfo
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityEdiscoveryHoldOperation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SecurityEdiscoveryHoldOperation': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
