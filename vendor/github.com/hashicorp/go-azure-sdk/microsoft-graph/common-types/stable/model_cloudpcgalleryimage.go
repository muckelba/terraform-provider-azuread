package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCGalleryImage{}

type CloudPCGalleryImage struct {
	// The display name of this gallery image. For example, Windows 11 Enterprise + Microsoft 365 Apps 22H2. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The date when the status of the image becomes supportedWithWarning. Users can still provision new Cloud PCs if the
	// current time is later than endDate and earlier than expirationDate. For example, assume the endDate of a gallery
	// image is 2023-9-14 and expirationDate is 2024-3-14, users are able to provision new Cloud PCs if today is 2023-10-01.
	// Read-only.
	EndDate nullable.Type[string] `json:"endDate,omitempty"`

	// The date when the image is no longer available. Users are unable to provision new Cloud PCs if the current time is
	// later than expirationDate. The value is usually endDate plus six months. For example, if the startDate is 2025-10-14,
	// the expirationDate is usually 2026-04-14. Read-only.
	ExpirationDate nullable.Type[string] `json:"expirationDate,omitempty"`

	// The offer name of this gallery image that is passed to Azure Resource Manager (ARM) to retrieve the image resource.
	// Read-only.
	OfferName nullable.Type[string] `json:"offerName,omitempty"`

	// The publisher name of this gallery image that is passed to Azure Resource Manager (ARM) to retrieve the image
	// resource. Read-only.
	PublisherName nullable.Type[string] `json:"publisherName,omitempty"`

	// Indicates the size of this image in gigabytes. For example, 64. Read-only.
	SizeInGB nullable.Type[int64] `json:"sizeInGB,omitempty"`

	// The SKU name of this image that is passed to Azure Resource Manager (ARM) to retrieve the image resource. Read-only.
	SkuName nullable.Type[string] `json:"skuName,omitempty"`

	// The date when the Cloud PC image is available for provisioning new Cloud PCs. For example, 2022-09-20. Read-only.
	StartDate nullable.Type[string] `json:"startDate,omitempty"`

	// The status of the gallery image on the Cloud PC. Possible values are: supported, supportedWithWarning, notSupported,
	// unknownFutureValue. The default value is supported. Read-only.
	Status *CloudPCGalleryImageStatus `json:"status,omitempty"`

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

func (s CloudPCGalleryImage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCGalleryImage{}

func (s CloudPCGalleryImage) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCGalleryImage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCGalleryImage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCGalleryImage: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "endDate")
	delete(decoded, "expirationDate")
	delete(decoded, "offerName")
	delete(decoded, "publisherName")
	delete(decoded, "sizeInGB")
	delete(decoded, "skuName")
	delete(decoded, "startDate")
	delete(decoded, "status")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcGalleryImage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCGalleryImage: %+v", err)
	}

	return encoded, nil
}