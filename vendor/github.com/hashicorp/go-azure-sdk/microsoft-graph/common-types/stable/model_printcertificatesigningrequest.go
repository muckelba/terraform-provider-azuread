package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintCertificateSigningRequest struct {
	// A base64-encoded pkcs10 certificate request. Read-only.
	Content *string `json:"content,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The base64-encoded public portion of an asymmetric key that is generated by the client. Read-only.
	TransportKey *string `json:"transportKey,omitempty"`
}

var _ json.Marshaler = PrintCertificateSigningRequest{}

func (s PrintCertificateSigningRequest) MarshalJSON() ([]byte, error) {
	type wrapper PrintCertificateSigningRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrintCertificateSigningRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrintCertificateSigningRequest: %+v", err)
	}

	delete(decoded, "content")
	delete(decoded, "transportKey")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrintCertificateSigningRequest: %+v", err)
	}

	return encoded, nil
}
