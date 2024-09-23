package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Message = CalendarSharingMessage{}

type CalendarSharingMessage struct {
	CanAccept             nullable.Type[bool]             `json:"canAccept,omitempty"`
	SharingMessageAction  *CalendarSharingMessageAction   `json:"sharingMessageAction,omitempty"`
	SharingMessageActions *[]CalendarSharingMessageAction `json:"sharingMessageActions,omitempty"`
	SuggestedCalendarName nullable.Type[string]           `json:"suggestedCalendarName,omitempty"`

	// Fields inherited from Message

	// The fileAttachment and itemAttachment attachments for the message.
	Attachments *[]Attachment `json:"attachments,omitempty"`

	// The Bcc: recipients for the message.
	BccRecipients *[]Recipient `json:"bccRecipients,omitempty"`

	// The body of the message. It can be in HTML or text format. Find out about safe HTML in a message body.
	Body *ItemBody `json:"body,omitempty"`

	// The first 255 characters of the message body. It is in text format. If the message contains instances of mention,
	// this property would contain a concatenation of these mentions as well.
	BodyPreview nullable.Type[string] `json:"bodyPreview,omitempty"`

	// The Cc: recipients for the message.
	CcRecipients *[]Recipient `json:"ccRecipients,omitempty"`

	// The ID of the conversation the email belongs to.
	ConversationId nullable.Type[string] `json:"conversationId,omitempty"`

	// Indicates the position of the message within the conversation.
	ConversationIndex nullable.Type[string] `json:"conversationIndex,omitempty"`

	// The collection of open extensions defined for the message. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`

	// The flag value that indicates the status, start date, due date, or completion date for the message.
	Flag *FollowupFlag `json:"flag,omitempty"`

	// The owner of the mailbox from which the message is sent. In most cases, this value is the same as the sender
	// property, except for sharing or delegation scenarios. The value must correspond to the actual mailbox used. Find out
	// more about setting the from and sender properties of a message.
	From Recipient `json:"from"`

	// Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message
	// contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the
	// body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
	HasAttachments nullable.Type[bool] `json:"hasAttachments,omitempty"`

	// The importance of the message. The possible values are: low, normal, and high.
	Importance *Importance `json:"importance,omitempty"`

	// The classification of the message for the user, based on inferred relevance or importance, or on an explicit
	// override. Possible values are: focused, other.
	InferenceClassification *InferenceClassificationType `json:"inferenceClassification,omitempty"`

	// A collection of message headers defined by RFC5322. The set includes message headers indicating the network path
	// taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data
	// for the message. Returned only on applying a $select query option. Read-only.
	InternetMessageHeaders *[]InternetMessageHeader `json:"internetMessageHeaders,omitempty"`

	// The message ID in the format specified by RFC5322. Updatable only if isDraft is true.
	InternetMessageId nullable.Type[string] `json:"internetMessageId,omitempty"`

	// Indicates whether a read receipt is requested for the message.
	IsDeliveryReceiptRequested nullable.Type[bool] `json:"isDeliveryReceiptRequested,omitempty"`

	// Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.
	IsDraft nullable.Type[bool] `json:"isDraft,omitempty"`

	// Indicates whether the message has been read.
	IsRead nullable.Type[bool] `json:"isRead,omitempty"`

	// Indicates whether a read receipt is requested for the message.
	IsReadReceiptRequested nullable.Type[bool] `json:"isReadReceiptRequested,omitempty"`

	// A collection of mentions in the message, ordered by the createdDateTime from the newest to the oldest. By default, a
	// GET /messages does not return this property unless you apply $expand on the property.
	Mentions *[]Mention `json:"mentions,omitempty"`

	// Information about mentions in the message. When processing a GET /messages request, the server sets this property and
	// includes it in the response by default. The server returns null if there are no mentions in the message. Optional.
	MentionsPreview *MentionsPreview `json:"mentionsPreview,omitempty"`

	// The collection of multi-value extended properties defined for the message. Nullable.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`

	// The unique identifier for the message's parent mailFolder.
	ParentFolderId nullable.Type[string] `json:"parentFolderId,omitempty"`

	// The date and time the message was received. The date and time information uses ISO 8601 format and is always in UTC
	// time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ReceivedDateTime nullable.Type[string] `json:"receivedDateTime,omitempty"`

	// The email addresses to use when replying.
	ReplyTo *[]Recipient `json:"replyTo,omitempty"`

	// The account that is actually used to generate the message. In most cases, this value is the same as the from
	// property. You can set this property to a different value when sending a message from a shared mailbox, for a shared
	// calendar, or as a delegate. In any case, the value must correspond to the actual mailbox used. Find out more about
	// setting the from and sender properties of a message.
	Sender Recipient `json:"sender"`

	// The date and time the message was sent. The date and time information uses ISO 8601 format and is always in UTC time.
	// For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	SentDateTime nullable.Type[string] `json:"sentDateTime,omitempty"`

	// The collection of single-value extended properties defined for the message. Nullable.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`

	// The subject of the message.
	Subject nullable.Type[string] `json:"subject,omitempty"`

	// The To: recipients for the message.
	ToRecipients *[]Recipient `json:"toRecipients,omitempty"`

	// The part of the body of the message that is unique to the current message. uniqueBody is not returned by default but
	// can be retrieved for a given message by use of the ?$select=uniqueBody query. It can be in HTML or text format.
	UniqueBody *ItemBody `json:"uniqueBody,omitempty"`

	// The valid entries parsed from the List-Unsubscribe header. This is the data for the mail command in the
	// List-Unsubscribe header if UnsubscribeEnabled property is true.
	UnsubscribeData *[]string `json:"unsubscribeData,omitempty"`

	// Indicates whether the message is enabled for unsubscribe. Its valueTrue if the list-Unsubscribe header conforms to
	// rfc-2369.
	UnsubscribeEnabled nullable.Type[bool] `json:"unsubscribeEnabled,omitempty"`

	// The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change
	// how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout
	// window. If ispopout is set to 0, the browser shows the message in the Outlook on the web review pane.The message
	// opens in the browser if you are signed in to your mailbox via Outlook on the web. You are prompted to sign in if you
	// are not already signed in with the browser.This URL cannot be accessed from within an iFrame.
	WebLink nullable.Type[string] `json:"webLink,omitempty"`

	// Fields inherited from OutlookItem

	// The categories associated with the item.
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

func (s CalendarSharingMessage) Message() BaseMessageImpl {
	return BaseMessageImpl{
		Attachments:                   s.Attachments,
		BccRecipients:                 s.BccRecipients,
		Body:                          s.Body,
		BodyPreview:                   s.BodyPreview,
		CcRecipients:                  s.CcRecipients,
		ConversationId:                s.ConversationId,
		ConversationIndex:             s.ConversationIndex,
		Extensions:                    s.Extensions,
		Flag:                          s.Flag,
		From:                          s.From,
		HasAttachments:                s.HasAttachments,
		Importance:                    s.Importance,
		InferenceClassification:       s.InferenceClassification,
		InternetMessageHeaders:        s.InternetMessageHeaders,
		InternetMessageId:             s.InternetMessageId,
		IsDeliveryReceiptRequested:    s.IsDeliveryReceiptRequested,
		IsDraft:                       s.IsDraft,
		IsRead:                        s.IsRead,
		IsReadReceiptRequested:        s.IsReadReceiptRequested,
		Mentions:                      s.Mentions,
		MentionsPreview:               s.MentionsPreview,
		MultiValueExtendedProperties:  s.MultiValueExtendedProperties,
		ParentFolderId:                s.ParentFolderId,
		ReceivedDateTime:              s.ReceivedDateTime,
		ReplyTo:                       s.ReplyTo,
		Sender:                        s.Sender,
		SentDateTime:                  s.SentDateTime,
		SingleValueExtendedProperties: s.SingleValueExtendedProperties,
		Subject:                       s.Subject,
		ToRecipients:                  s.ToRecipients,
		UniqueBody:                    s.UniqueBody,
		UnsubscribeData:               s.UnsubscribeData,
		UnsubscribeEnabled:            s.UnsubscribeEnabled,
		WebLink:                       s.WebLink,
		Categories:                    s.Categories,
		ChangeKey:                     s.ChangeKey,
		CreatedDateTime:               s.CreatedDateTime,
		LastModifiedDateTime:          s.LastModifiedDateTime,
		Id:                            s.Id,
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
	}
}

func (s CalendarSharingMessage) OutlookItem() BaseOutlookItemImpl {
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

func (s CalendarSharingMessage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CalendarSharingMessage{}

func (s CalendarSharingMessage) MarshalJSON() ([]byte, error) {
	type wrapper CalendarSharingMessage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CalendarSharingMessage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CalendarSharingMessage: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.calendarSharingMessage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CalendarSharingMessage: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CalendarSharingMessage{}

func (s *CalendarSharingMessage) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CanAccept                     nullable.Type[bool]                  `json:"canAccept,omitempty"`
		SharingMessageAction          *CalendarSharingMessageAction        `json:"sharingMessageAction,omitempty"`
		SharingMessageActions         *[]CalendarSharingMessageAction      `json:"sharingMessageActions,omitempty"`
		SuggestedCalendarName         nullable.Type[string]                `json:"suggestedCalendarName,omitempty"`
		Body                          *ItemBody                            `json:"body,omitempty"`
		BodyPreview                   nullable.Type[string]                `json:"bodyPreview,omitempty"`
		ConversationId                nullable.Type[string]                `json:"conversationId,omitempty"`
		ConversationIndex             nullable.Type[string]                `json:"conversationIndex,omitempty"`
		Flag                          *FollowupFlag                        `json:"flag,omitempty"`
		HasAttachments                nullable.Type[bool]                  `json:"hasAttachments,omitempty"`
		Importance                    *Importance                          `json:"importance,omitempty"`
		InferenceClassification       *InferenceClassificationType         `json:"inferenceClassification,omitempty"`
		InternetMessageHeaders        *[]InternetMessageHeader             `json:"internetMessageHeaders,omitempty"`
		InternetMessageId             nullable.Type[string]                `json:"internetMessageId,omitempty"`
		IsDeliveryReceiptRequested    nullable.Type[bool]                  `json:"isDeliveryReceiptRequested,omitempty"`
		IsDraft                       nullable.Type[bool]                  `json:"isDraft,omitempty"`
		IsRead                        nullable.Type[bool]                  `json:"isRead,omitempty"`
		IsReadReceiptRequested        nullable.Type[bool]                  `json:"isReadReceiptRequested,omitempty"`
		Mentions                      *[]Mention                           `json:"mentions,omitempty"`
		MentionsPreview               *MentionsPreview                     `json:"mentionsPreview,omitempty"`
		MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
		ParentFolderId                nullable.Type[string]                `json:"parentFolderId,omitempty"`
		ReceivedDateTime              nullable.Type[string]                `json:"receivedDateTime,omitempty"`
		SentDateTime                  nullable.Type[string]                `json:"sentDateTime,omitempty"`
		SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
		Subject                       nullable.Type[string]                `json:"subject,omitempty"`
		UniqueBody                    *ItemBody                            `json:"uniqueBody,omitempty"`
		UnsubscribeData               *[]string                            `json:"unsubscribeData,omitempty"`
		UnsubscribeEnabled            nullable.Type[bool]                  `json:"unsubscribeEnabled,omitempty"`
		WebLink                       nullable.Type[string]                `json:"webLink,omitempty"`
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

	s.CanAccept = decoded.CanAccept
	s.SharingMessageAction = decoded.SharingMessageAction
	s.SharingMessageActions = decoded.SharingMessageActions
	s.SuggestedCalendarName = decoded.SuggestedCalendarName
	s.Body = decoded.Body
	s.BodyPreview = decoded.BodyPreview
	s.Categories = decoded.Categories
	s.ChangeKey = decoded.ChangeKey
	s.ConversationId = decoded.ConversationId
	s.ConversationIndex = decoded.ConversationIndex
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Flag = decoded.Flag
	s.HasAttachments = decoded.HasAttachments
	s.Id = decoded.Id
	s.Importance = decoded.Importance
	s.InferenceClassification = decoded.InferenceClassification
	s.InternetMessageHeaders = decoded.InternetMessageHeaders
	s.InternetMessageId = decoded.InternetMessageId
	s.IsDeliveryReceiptRequested = decoded.IsDeliveryReceiptRequested
	s.IsDraft = decoded.IsDraft
	s.IsRead = decoded.IsRead
	s.IsReadReceiptRequested = decoded.IsReadReceiptRequested
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Mentions = decoded.Mentions
	s.MentionsPreview = decoded.MentionsPreview
	s.MultiValueExtendedProperties = decoded.MultiValueExtendedProperties
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ParentFolderId = decoded.ParentFolderId
	s.ReceivedDateTime = decoded.ReceivedDateTime
	s.SentDateTime = decoded.SentDateTime
	s.SingleValueExtendedProperties = decoded.SingleValueExtendedProperties
	s.Subject = decoded.Subject
	s.UniqueBody = decoded.UniqueBody
	s.UnsubscribeData = decoded.UnsubscribeData
	s.UnsubscribeEnabled = decoded.UnsubscribeEnabled
	s.WebLink = decoded.WebLink

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CalendarSharingMessage into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["attachments"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Attachments into list []json.RawMessage: %+v", err)
		}

		output := make([]Attachment, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAttachmentImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Attachments' for 'CalendarSharingMessage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Attachments = &output
	}

	if v, ok := temp["bccRecipients"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling BccRecipients into list []json.RawMessage: %+v", err)
		}

		output := make([]Recipient, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRecipientImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'BccRecipients' for 'CalendarSharingMessage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.BccRecipients = &output
	}

	if v, ok := temp["ccRecipients"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling CcRecipients into list []json.RawMessage: %+v", err)
		}

		output := make([]Recipient, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRecipientImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'CcRecipients' for 'CalendarSharingMessage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CcRecipients = &output
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
				return fmt.Errorf("unmarshaling index %d field 'Extensions' for 'CalendarSharingMessage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Extensions = &output
	}

	if v, ok := temp["from"]; ok {
		impl, err := UnmarshalRecipientImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'From' for 'CalendarSharingMessage': %+v", err)
		}
		s.From = impl
	}

	if v, ok := temp["replyTo"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ReplyTo into list []json.RawMessage: %+v", err)
		}

		output := make([]Recipient, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRecipientImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ReplyTo' for 'CalendarSharingMessage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ReplyTo = &output
	}

	if v, ok := temp["sender"]; ok {
		impl, err := UnmarshalRecipientImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Sender' for 'CalendarSharingMessage': %+v", err)
		}
		s.Sender = impl
	}

	if v, ok := temp["toRecipients"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ToRecipients into list []json.RawMessage: %+v", err)
		}

		output := make([]Recipient, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRecipientImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ToRecipients' for 'CalendarSharingMessage': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ToRecipients = &output
	}

	return nil
}
