package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileLobApp = IosLobApp{}

type IosLobApp struct {
	// Contains properties of the possible iOS device types the mobile app can run on.
	ApplicableDeviceType *IosDeviceType `json:"applicableDeviceType,omitempty"`

	// The build number of iOS Line of Business (LoB) app.
	BuildNumber nullable.Type[string] `json:"buildNumber,omitempty"`

	// The Identity Name.
	BundleId nullable.Type[string] `json:"bundleId,omitempty"`

	// The expiration time.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// The value for the minimum applicable operating system.
	MinimumSupportedOperatingSystem *IosMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`

	// The version number of iOS Line of Business (LoB) app.
	VersionNumber nullable.Type[string] `json:"versionNumber,omitempty"`

	// Fields inherited from MobileLobApp

	// The internal committed content version.
	CommittedContentVersion nullable.Type[string] `json:"committedContentVersion,omitempty"`

	// The list of content versions for this app.
	ContentVersions *[]MobileAppContent `json:"contentVersions,omitempty"`

	// The name of the main Lob application file.
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// The total size, including all uploaded files.
	Size *int64 `json:"size,omitempty"`

	// Fields inherited from MobileApp

	// The list of group assignments for this mobile app.
	Assignments *[]MobileAppAssignment `json:"assignments,omitempty"`

	// The list of categories for this app.
	Categories *[]MobileAppCategory `json:"categories,omitempty"`

	// The date and time the app was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the app.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The developer of the app.
	Developer nullable.Type[string] `json:"developer,omitempty"`

	// The admin provided or imported title of the app.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The more information Url.
	InformationUrl nullable.Type[string] `json:"informationUrl,omitempty"`

	// The value indicating whether the app is marked as featured by the admin.
	IsFeatured *bool `json:"isFeatured,omitempty"`

	// The large icon, to be displayed in the app details and used for upload of the icon.
	LargeIcon *MimeContent `json:"largeIcon,omitempty"`

	// The date and time the app was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Notes for the app.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	// The owner of the app.
	Owner nullable.Type[string] `json:"owner,omitempty"`

	// The privacy statement Url.
	PrivacyInformationUrl nullable.Type[string] `json:"privacyInformationUrl,omitempty"`

	// The publisher of the app.
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	// Indicates the publishing state of an app.
	PublishingState *MobileAppPublishingState `json:"publishingState,omitempty"`

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

func (s IosLobApp) MobileLobApp() BaseMobileLobAppImpl {
	return BaseMobileLobAppImpl{
		CommittedContentVersion: s.CommittedContentVersion,
		ContentVersions:         s.ContentVersions,
		FileName:                s.FileName,
		Size:                    s.Size,
		Assignments:             s.Assignments,
		Categories:              s.Categories,
		CreatedDateTime:         s.CreatedDateTime,
		Description:             s.Description,
		Developer:               s.Developer,
		DisplayName:             s.DisplayName,
		InformationUrl:          s.InformationUrl,
		IsFeatured:              s.IsFeatured,
		LargeIcon:               s.LargeIcon,
		LastModifiedDateTime:    s.LastModifiedDateTime,
		Notes:                   s.Notes,
		Owner:                   s.Owner,
		PrivacyInformationUrl:   s.PrivacyInformationUrl,
		Publisher:               s.Publisher,
		PublishingState:         s.PublishingState,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s IosLobApp) MobileApp() BaseMobileAppImpl {
	return BaseMobileAppImpl{
		Assignments:           s.Assignments,
		Categories:            s.Categories,
		CreatedDateTime:       s.CreatedDateTime,
		Description:           s.Description,
		Developer:             s.Developer,
		DisplayName:           s.DisplayName,
		InformationUrl:        s.InformationUrl,
		IsFeatured:            s.IsFeatured,
		LargeIcon:             s.LargeIcon,
		LastModifiedDateTime:  s.LastModifiedDateTime,
		Notes:                 s.Notes,
		Owner:                 s.Owner,
		PrivacyInformationUrl: s.PrivacyInformationUrl,
		Publisher:             s.Publisher,
		PublishingState:       s.PublishingState,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s IosLobApp) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosLobApp{}

func (s IosLobApp) MarshalJSON() ([]byte, error) {
	type wrapper IosLobApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosLobApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosLobApp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosLobApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosLobApp: %+v", err)
	}

	return encoded, nil
}
