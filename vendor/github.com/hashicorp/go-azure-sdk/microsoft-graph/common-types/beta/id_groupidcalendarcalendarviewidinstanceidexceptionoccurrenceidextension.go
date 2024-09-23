package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{}

// GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId is a struct representing the Resource ID for a Group Id Calendar Calendar View Id Instance Id Exception Occurrence Id Extension
type GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId struct {
	GroupId     string
	EventId     string
	EventId1    string
	EventId2    string
	ExtensionId string
}

// NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID returns a new GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId struct
func NewGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID(groupId string, eventId string, eventId1 string, eventId2 string, extensionId string) GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId {
	return GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		EventId1:    eventId1,
		EventId2:    eventId2,
		ExtensionId: extensionId,
	}
}

// ParseGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID parses 'input' into a GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId
func ParseGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID(input string) (*GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionIDInsensitively parses 'input' case-insensitively into a GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionIDInsensitively(input string) (*GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	if id.EventId2, ok = input.Parsed["eventId2"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId2", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID checks that 'input' can be parsed as a Group Id Calendar Calendar View Id Instance Id Exception Occurrence Id Extension ID
func ValidateGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Calendar Calendar View Id Instance Id Exception Occurrence Id Extension ID
func (id GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId) ID() string {
	fmtString := "/groups/%s/calendar/calendarView/%s/instances/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.EventId2, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Calendar Calendar View Id Instance Id Exception Occurrence Id Extension ID
func (id GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionId"),
	}
}

// String returns a human-readable description of this Group Id Calendar Calendar View Id Instance Id Exception Occurrence Id Extension ID
func (id GroupIdCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Id Calendar Calendar View Id Instance Id Exception Occurrence Id Extension (%s)", strings.Join(components, "\n"))
}
