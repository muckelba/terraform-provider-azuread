package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId{}

// MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId is a struct representing the Resource ID for a Me Joined Team Id Channel Id Shared With Team Id Allowed Member
type MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId struct {
	TeamId                      string
	ChannelId                   string
	SharedWithChannelTeamInfoId string
	ConversationMemberId        string
}

// NewMeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberID returns a new MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId struct
func NewMeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberID(teamId string, channelId string, sharedWithChannelTeamInfoId string, conversationMemberId string) MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId {
	return MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId{
		TeamId:                      teamId,
		ChannelId:                   channelId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
		ConversationMemberId:        conversationMemberId,
	}
}

// ParseMeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberID parses 'input' into a MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId
func ParseMeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberID(input string) (*MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberIDInsensitively(input string) (*MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TeamId, ok = input.Parsed["teamId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "teamId", input)
	}

	if id.ChannelId, ok = input.Parsed["channelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "channelId", input)
	}

	if id.SharedWithChannelTeamInfoId, ok = input.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", input)
	}

	if id.ConversationMemberId, ok = input.Parsed["conversationMemberId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", input)
	}

	return nil
}

// ValidateMeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberID checks that 'input' can be parsed as a Me Joined Team Id Channel Id Shared With Team Id Allowed Member ID
func ValidateMeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Id Channel Id Shared With Team Id Allowed Member ID
func (id MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId) ID() string {
	fmtString := "/me/joinedTeams/%s/channels/%s/sharedWithTeams/%s/allowedMembers/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChannelId, id.SharedWithChannelTeamInfoId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Id Channel Id Shared With Team Id Allowed Member ID
func (id MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("joinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamId"),
		resourceids.StaticSegment("channels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelId"),
		resourceids.StaticSegment("sharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoId"),
		resourceids.StaticSegment("allowedMembers", "allowedMembers", "allowedMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberId"),
	}
}

// String returns a human-readable description of this Me Joined Team Id Channel Id Shared With Team Id Allowed Member ID
func (id MeJoinedTeamIdChannelIdSharedWithTeamIdAllowedMemberId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Me Joined Team Id Channel Id Shared With Team Id Allowed Member (%s)", strings.Join(components, "\n"))
}
