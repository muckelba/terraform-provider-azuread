package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{}

// GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId is a struct representing the Resource ID for a Group Id Onenote Notebook Id Section Group Id Section Id Page
type GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId struct {
	GroupId          string
	NotebookId       string
	SectionGroupId   string
	OnenoteSectionId string
	OnenotePageId    string
}

// NewGroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageID returns a new GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId struct
func NewGroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageID(groupId string, notebookId string, sectionGroupId string, onenoteSectionId string, onenotePageId string) GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId {
	return GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{
		GroupId:          groupId,
		NotebookId:       notebookId,
		SectionGroupId:   sectionGroupId,
		OnenoteSectionId: onenoteSectionId,
		OnenotePageId:    onenotePageId,
	}
}

// ParseGroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageID parses 'input' into a GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId
func ParseGroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageID(input string) (*GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageIDInsensitively parses 'input' case-insensitively into a GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId
// note: this method should only be used for API response data and not user input
func ParseGroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageIDInsensitively(input string) (*GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.NotebookId, ok = input.Parsed["notebookId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notebookId", input)
	}

	if id.SectionGroupId, ok = input.Parsed["sectionGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sectionGroupId", input)
	}

	if id.OnenoteSectionId, ok = input.Parsed["onenoteSectionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteSectionId", input)
	}

	if id.OnenotePageId, ok = input.Parsed["onenotePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenotePageId", input)
	}

	return nil
}

// ValidateGroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageID checks that 'input' can be parsed as a Group Id Onenote Notebook Id Section Group Id Section Id Page ID
func ValidateGroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Onenote Notebook Id Section Group Id Section Id Page ID
func (id GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId) ID() string {
	fmtString := "/groups/%s/onenote/notebooks/%s/sectionGroups/%s/sections/%s/pages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.NotebookId, id.SectionGroupId, id.OnenoteSectionId, id.OnenotePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Onenote Notebook Id Section Group Id Section Id Page ID
func (id GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("notebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookId", "notebookId"),
		resourceids.StaticSegment("sectionGroups", "sectionGroups", "sectionGroups"),
		resourceids.UserSpecifiedSegment("sectionGroupId", "sectionGroupId"),
		resourceids.StaticSegment("sections", "sections", "sections"),
		resourceids.UserSpecifiedSegment("onenoteSectionId", "onenoteSectionId"),
		resourceids.StaticSegment("pages", "pages", "pages"),
		resourceids.UserSpecifiedSegment("onenotePageId", "onenotePageId"),
	}
}

// String returns a human-readable description of this Group Id Onenote Notebook Id Section Group Id Section Id Page ID
func (id GroupIdOnenoteNotebookIdSectionGroupIdSectionIdPageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Notebook: %q", id.NotebookId),
		fmt.Sprintf("Section Group: %q", id.SectionGroupId),
		fmt.Sprintf("Onenote Section: %q", id.OnenoteSectionId),
		fmt.Sprintf("Onenote Page: %q", id.OnenotePageId),
	}
	return fmt.Sprintf("Group Id Onenote Notebook Id Section Group Id Section Id Page (%s)", strings.Join(components, "\n"))
}
