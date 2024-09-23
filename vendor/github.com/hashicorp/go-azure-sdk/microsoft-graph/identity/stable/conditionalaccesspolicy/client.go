package conditionalaccesspolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessPolicyClient struct {
	Client *msgraph.Client
}

func NewConditionalAccessPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*ConditionalAccessPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "conditionalaccesspolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConditionalAccessPolicyClient: %+v", err)
	}

	return &ConditionalAccessPolicyClient{
		Client: client,
	}, nil
}
