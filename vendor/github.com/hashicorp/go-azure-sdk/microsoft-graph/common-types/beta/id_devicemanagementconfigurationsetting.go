package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementConfigurationSettingId{}

// DeviceManagementConfigurationSettingId is a struct representing the Resource ID for a Device Management Configuration Setting
type DeviceManagementConfigurationSettingId struct {
	DeviceManagementConfigurationSettingDefinitionId string
}

// NewDeviceManagementConfigurationSettingID returns a new DeviceManagementConfigurationSettingId struct
func NewDeviceManagementConfigurationSettingID(deviceManagementConfigurationSettingDefinitionId string) DeviceManagementConfigurationSettingId {
	return DeviceManagementConfigurationSettingId{
		DeviceManagementConfigurationSettingDefinitionId: deviceManagementConfigurationSettingDefinitionId,
	}
}

// ParseDeviceManagementConfigurationSettingID parses 'input' into a DeviceManagementConfigurationSettingId
func ParseDeviceManagementConfigurationSettingID(input string) (*DeviceManagementConfigurationSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementConfigurationSettingIDInsensitively parses 'input' case-insensitively into a DeviceManagementConfigurationSettingId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementConfigurationSettingIDInsensitively(input string) (*DeviceManagementConfigurationSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementConfigurationSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementConfigurationSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementConfigurationSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DeviceManagementConfigurationSettingDefinitionId, ok = input.Parsed["deviceManagementConfigurationSettingDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementConfigurationSettingDefinitionId", input)
	}

	return nil
}

// ValidateDeviceManagementConfigurationSettingID checks that 'input' can be parsed as a Device Management Configuration Setting ID
func ValidateDeviceManagementConfigurationSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementConfigurationSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Configuration Setting ID
func (id DeviceManagementConfigurationSettingId) ID() string {
	fmtString := "/deviceManagement/configurationSettings/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementConfigurationSettingDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Configuration Setting ID
func (id DeviceManagementConfigurationSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("configurationSettings", "configurationSettings", "configurationSettings"),
		resourceids.UserSpecifiedSegment("deviceManagementConfigurationSettingDefinitionId", "deviceManagementConfigurationSettingDefinitionId"),
	}
}

// String returns a human-readable description of this Device Management Configuration Setting ID
func (id DeviceManagementConfigurationSettingId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Configuration Setting Definition: %q", id.DeviceManagementConfigurationSettingDefinitionId),
	}
	return fmt.Sprintf("Device Management Configuration Setting (%s)", strings.Join(components, "\n"))
}
