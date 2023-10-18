// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

type ApplicationRegistrationResource struct{}

func TestAccApplicationRegistration_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_registration", "test")
	r := ApplicationRegistrationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("client_id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-AppRegistration-%d", data.RandomInteger)),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationRegistration_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_registration", "test")
	r := ApplicationRegistrationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("client_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationRegistration_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_registration", "test")
	r := ApplicationRegistrationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("client_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("client_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("client_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func (r ApplicationRegistrationResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	id, err := parse.ParseApplicationID(state.ID)
	if err != nil {
		return nil, err
	}

	app, status, err := client.Get(ctx, id.ApplicationId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", id, err)
	}

	return pointer.To(app.ID() != nil && *app.ID() == id.ApplicationId), nil
}

func (ApplicationRegistrationResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-AppRegistration-%[1]d"
}
`, data.RandomInteger)
}

func (ApplicationRegistrationResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name     = "acctest-AppRegistration-complete-%[1]d"
  sign_in_audience = "AzureADandPersonalMicrosoftAccount"

  requested_access_token_version         = 2
  implicit_access_token_issuance_enabled = true
  implicit_id_token_issuance_enabled     = true

  description                  = "Acceptance testing application"
  notes                        = "Testing application"
  service_management_reference = "app-for-testing"

  homepage_url          = "https://app.hashitown-%[1]d.com/"
  logout_url            = "https://app.hashitown-%[1]d.com/logout"
  marketing_url         = "https://hashitown-%[1]d.com/"
  privacy_statement_url = "https://hashitown-%[1]d.com/privacy"
  support_url           = "https://support.hashitown-%[1]d.com/"
  terms_of_service_url  = "https://hashitown-%[1]d.com/terms"
}
`, data.RandomInteger)
}
