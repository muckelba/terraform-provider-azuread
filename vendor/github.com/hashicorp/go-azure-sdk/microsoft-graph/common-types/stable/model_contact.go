package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OutlookItem = Contact{}

type Contact struct {
	// The name of the contact's assistant.
	AssistantName nullable.Type[string] `json:"assistantName,omitempty"`

	// The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always
	// in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	Birthday nullable.Type[string] `json:"birthday,omitempty"`

	// The contact's business address.
	BusinessAddress *PhysicalAddress `json:"businessAddress,omitempty"`

	// The business home page of the contact.
	BusinessHomePage nullable.Type[string] `json:"businessHomePage,omitempty"`

	// The contact's business phone numbers.
	BusinessPhones *[]string `json:"businessPhones,omitempty"`

	// The names of the contact's children.
	Children *[]string `json:"children,omitempty"`

	// The name of the contact's company.
	CompanyName nullable.Type[string] `json:"companyName,omitempty"`

	// The contact's department.
	Department nullable.Type[string] `json:"department,omitempty"`

	// The contact's display name. You can specify the display name in a create or update operation. Note that later updates
	// to other properties may cause an automatically generated value to overwrite the displayName value you have specified.
	// To preserve a pre-existing value, always include it as displayName in an update operation.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The contact's email addresses.
	EmailAddresses *[]EmailAddress `json:"emailAddresses,omitempty"`

	// The collection of open extensions defined for the contact. Read-only. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`

	// The name the contact is filed under.
	FileAs nullable.Type[string] `json:"fileAs,omitempty"`

	// The contact's suffix.
	Generation nullable.Type[string] `json:"generation,omitempty"`

	// The contact's given name.
	GivenName nullable.Type[string] `json:"givenName,omitempty"`

	// The contact's home address.
	HomeAddress *PhysicalAddress `json:"homeAddress,omitempty"`

	// The contact's home phone numbers.
	HomePhones *[]string `json:"homePhones,omitempty"`

	// The contact's instant messaging (IM) addresses.
	ImAddresses *[]string `json:"imAddresses,omitempty"`

	// The contact's initials.
	Initials nullable.Type[string] `json:"initials,omitempty"`

	// The contact’s job title.
	JobTitle nullable.Type[string] `json:"jobTitle,omitempty"`

	// The name of the contact's manager.
	Manager nullable.Type[string] `json:"manager,omitempty"`

	// The contact's middle name.
	MiddleName nullable.Type[string] `json:"middleName,omitempty"`

	// The contact's mobile phone number.
	MobilePhone nullable.Type[string] `json:"mobilePhone,omitempty"`

	// The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`

	// The contact's nickname.
	NickName nullable.Type[string] `json:"nickName,omitempty"`

	// The location of the contact's office.
	OfficeLocation nullable.Type[string] `json:"officeLocation,omitempty"`

	// Other addresses for the contact.
	OtherAddress *PhysicalAddress `json:"otherAddress,omitempty"`

	// The ID of the contact's parent folder.
	ParentFolderId nullable.Type[string] `json:"parentFolderId,omitempty"`

	// The user's notes about the contact.
	PersonalNotes nullable.Type[string] `json:"personalNotes,omitempty"`

	// Optional contact picture. You can get or set a photo for a contact.
	Photo *ProfilePhoto `json:"photo,omitempty"`

	// The contact's profession.
	Profession nullable.Type[string] `json:"profession,omitempty"`

	// The collection of single-value extended properties defined for the contact. Read-only. Nullable.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`

	// The name of the contact's spouse/partner.
	SpouseName nullable.Type[string] `json:"spouseName,omitempty"`

	// The contact's surname.
	Surname nullable.Type[string] `json:"surname,omitempty"`

	// The contact's title.
	Title nullable.Type[string] `json:"title,omitempty"`

	// The phonetic Japanese company name of the contact.
	YomiCompanyName nullable.Type[string] `json:"yomiCompanyName,omitempty"`

	// The phonetic Japanese given name (first name) of the contact.
	YomiGivenName nullable.Type[string] `json:"yomiGivenName,omitempty"`

	// The phonetic Japanese surname (last name) of the contact.
	YomiSurname nullable.Type[string] `json:"yomiSurname,omitempty"`

	// Fields inherited from OutlookItem

	// The categories associated with the item
	Categories *[]string `json:"categories,omitempty"`

	// Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange
	// to apply changes to the correct version of the object. Read-only.
	ChangeKey nullable.Type[string] `json:"changeKey,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

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

func (s Contact) OutlookItem() BaseOutlookItemImpl {
	return BaseOutlookItemImpl{
		Categories:           s.Categories,
		ChangeKey:            s.ChangeKey,
		CreatedDateTime:      s.CreatedDateTime,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s Contact) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Contact{}

func (s Contact) MarshalJSON() ([]byte, error) {
	type wrapper Contact
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Contact: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Contact: %+v", err)
	}

	delete(decoded, "extensions")
	delete(decoded, "multiValueExtendedProperties")
	delete(decoded, "singleValueExtendedProperties")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.contact"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Contact: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Contact{}

func (s *Contact) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssistantName                 nullable.Type[string]                `json:"assistantName,omitempty"`
		Birthday                      nullable.Type[string]                `json:"birthday,omitempty"`
		BusinessAddress               *PhysicalAddress                     `json:"businessAddress,omitempty"`
		BusinessHomePage              nullable.Type[string]                `json:"businessHomePage,omitempty"`
		BusinessPhones                *[]string                            `json:"businessPhones,omitempty"`
		Children                      *[]string                            `json:"children,omitempty"`
		CompanyName                   nullable.Type[string]                `json:"companyName,omitempty"`
		Department                    nullable.Type[string]                `json:"department,omitempty"`
		DisplayName                   nullable.Type[string]                `json:"displayName,omitempty"`
		EmailAddresses                *[]EmailAddress                      `json:"emailAddresses,omitempty"`
		FileAs                        nullable.Type[string]                `json:"fileAs,omitempty"`
		Generation                    nullable.Type[string]                `json:"generation,omitempty"`
		GivenName                     nullable.Type[string]                `json:"givenName,omitempty"`
		HomeAddress                   *PhysicalAddress                     `json:"homeAddress,omitempty"`
		HomePhones                    *[]string                            `json:"homePhones,omitempty"`
		ImAddresses                   *[]string                            `json:"imAddresses,omitempty"`
		Initials                      nullable.Type[string]                `json:"initials,omitempty"`
		JobTitle                      nullable.Type[string]                `json:"jobTitle,omitempty"`
		Manager                       nullable.Type[string]                `json:"manager,omitempty"`
		MiddleName                    nullable.Type[string]                `json:"middleName,omitempty"`
		MobilePhone                   nullable.Type[string]                `json:"mobilePhone,omitempty"`
		MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
		NickName                      nullable.Type[string]                `json:"nickName,omitempty"`
		OfficeLocation                nullable.Type[string]                `json:"officeLocation,omitempty"`
		OtherAddress                  *PhysicalAddress                     `json:"otherAddress,omitempty"`
		ParentFolderId                nullable.Type[string]                `json:"parentFolderId,omitempty"`
		PersonalNotes                 nullable.Type[string]                `json:"personalNotes,omitempty"`
		Photo                         *ProfilePhoto                        `json:"photo,omitempty"`
		Profession                    nullable.Type[string]                `json:"profession,omitempty"`
		SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
		SpouseName                    nullable.Type[string]                `json:"spouseName,omitempty"`
		Surname                       nullable.Type[string]                `json:"surname,omitempty"`
		Title                         nullable.Type[string]                `json:"title,omitempty"`
		YomiCompanyName               nullable.Type[string]                `json:"yomiCompanyName,omitempty"`
		YomiGivenName                 nullable.Type[string]                `json:"yomiGivenName,omitempty"`
		YomiSurname                   nullable.Type[string]                `json:"yomiSurname,omitempty"`
		Categories                    *[]string                            `json:"categories,omitempty"`
		ChangeKey                     nullable.Type[string]                `json:"changeKey,omitempty"`
		CreatedDateTime               nullable.Type[string]                `json:"createdDateTime,omitempty"`
		LastModifiedDateTime          nullable.Type[string]                `json:"lastModifiedDateTime,omitempty"`
		Id                            *string                              `json:"id,omitempty"`
		ODataId                       *string                              `json:"@odata.id,omitempty"`
		ODataType                     *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssistantName = decoded.AssistantName
	s.Birthday = decoded.Birthday
	s.BusinessAddress = decoded.BusinessAddress
	s.BusinessHomePage = decoded.BusinessHomePage
	s.BusinessPhones = decoded.BusinessPhones
	s.Children = decoded.Children
	s.CompanyName = decoded.CompanyName
	s.Department = decoded.Department
	s.DisplayName = decoded.DisplayName
	s.EmailAddresses = decoded.EmailAddresses
	s.FileAs = decoded.FileAs
	s.Generation = decoded.Generation
	s.GivenName = decoded.GivenName
	s.HomeAddress = decoded.HomeAddress
	s.HomePhones = decoded.HomePhones
	s.ImAddresses = decoded.ImAddresses
	s.Initials = decoded.Initials
	s.JobTitle = decoded.JobTitle
	s.Manager = decoded.Manager
	s.MiddleName = decoded.MiddleName
	s.MobilePhone = decoded.MobilePhone
	s.MultiValueExtendedProperties = decoded.MultiValueExtendedProperties
	s.NickName = decoded.NickName
	s.OfficeLocation = decoded.OfficeLocation
	s.OtherAddress = decoded.OtherAddress
	s.ParentFolderId = decoded.ParentFolderId
	s.PersonalNotes = decoded.PersonalNotes
	s.Photo = decoded.Photo
	s.Profession = decoded.Profession
	s.SingleValueExtendedProperties = decoded.SingleValueExtendedProperties
	s.SpouseName = decoded.SpouseName
	s.Surname = decoded.Surname
	s.Title = decoded.Title
	s.YomiCompanyName = decoded.YomiCompanyName
	s.YomiGivenName = decoded.YomiGivenName
	s.YomiSurname = decoded.YomiSurname
	s.Categories = decoded.Categories
	s.ChangeKey = decoded.ChangeKey
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Contact into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["extensions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Extensions into list []json.RawMessage: %+v", err)
		}

		output := make([]Extension, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalExtensionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Extensions' for 'Contact': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Extensions = &output
	}

	return nil
}
