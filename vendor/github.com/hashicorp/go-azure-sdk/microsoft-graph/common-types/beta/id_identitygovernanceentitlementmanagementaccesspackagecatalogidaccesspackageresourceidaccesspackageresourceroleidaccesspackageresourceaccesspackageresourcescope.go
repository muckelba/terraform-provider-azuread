package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId{}

// IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Id Access Package Resource Role Id Access Package Resource Access Package Resource Scope
type IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId struct {
	AccessPackageCatalogId       string
	AccessPackageResourceId      string
	AccessPackageResourceRoleId  string
	AccessPackageResourceScopeId string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeID returns a new IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeID(accessPackageCatalogId string, accessPackageResourceId string, accessPackageResourceRoleId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId {
	return IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId{
		AccessPackageCatalogId:       accessPackageCatalogId,
		AccessPackageResourceId:      accessPackageResourceId,
		AccessPackageResourceRoleId:  accessPackageResourceRoleId,
		AccessPackageResourceScopeId: accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId
func ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageCatalogId, ok = input.Parsed["accessPackageCatalogId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageCatalogId", input)
	}

	if id.AccessPackageResourceId, ok = input.Parsed["accessPackageResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Id Access Package Resource Role Id Access Package Resource Access Package Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Id Access Package Resource Role Id Access Package Resource Access Package Resource Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageCatalogs/%s/accessPackageResources/%s/accessPackageResourceRoles/%s/accessPackageResource/accessPackageResourceScopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageCatalogId, id.AccessPackageResourceId, id.AccessPackageResourceRoleId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Id Access Package Resource Role Id Access Package Resource Access Package Resource Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageCatalogs", "accessPackageCatalogs", "accessPackageCatalogs"),
		resourceids.UserSpecifiedSegment("accessPackageCatalogId", "accessPackageCatalogId"),
		resourceids.StaticSegment("accessPackageResources", "accessPackageResources", "accessPackageResources"),
		resourceids.UserSpecifiedSegment("accessPackageResourceId", "accessPackageResourceId"),
		resourceids.StaticSegment("accessPackageResourceRoles", "accessPackageResourceRoles", "accessPackageResourceRoles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
		resourceids.StaticSegment("accessPackageResource", "accessPackageResource", "accessPackageResource"),
		resourceids.StaticSegment("accessPackageResourceScopes", "accessPackageResourceScopes", "accessPackageResourceScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Id Access Package Resource Role Id Access Package Resource Access Package Resource Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceRoleIdAccessPackageResourceAccessPackageResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Catalog: %q", id.AccessPackageCatalogId),
		fmt.Sprintf("Access Package Resource: %q", id.AccessPackageResourceId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Catalog Id Access Package Resource Id Access Package Resource Role Id Access Package Resource Access Package Resource Scope (%s)", strings.Join(components, "\n"))
}
