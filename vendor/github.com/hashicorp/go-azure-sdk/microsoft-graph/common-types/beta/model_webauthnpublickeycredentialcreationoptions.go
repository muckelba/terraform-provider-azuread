package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebauthnPublicKeyCredentialCreationOptions struct {
	// Attestation preference for the Microsoft Entra ID relying party.
	Attestation nullable.Type[string] `json:"attestation,omitempty"`

	// Properties of WebAuthn authenticators allowed to be used for authentication in Microsoft Entra ID.
	AuthenticatorSelection *WebauthnAuthenticatorSelectionCriteria `json:"authenticatorSelection,omitempty"`

	// A challenge generated by Microsoft Entra ID and sent back with the registration request to prevent replay attacks.
	Challenge nullable.Type[string] `json:"challenge,omitempty"`

	// List of credentials blocked for creations in Microsoft Entra ID.
	ExcludeCredentials *[]WebauthnPublicKeyCredentialDescriptor `json:"excludeCredentials,omitempty"`

	// Additional processing required by Microsoft Entra ID for the client and WebAuthn authenticator. For example,
	// Microsoft Entra ID might require that particular information be returned in the attestation object.
	Extensions *WebauthnAuthenticationExtensionsClientInputs `json:"extensions,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A set of preferred credential properties for the creation of a new public key credential in Microsoft Entra ID.
	PubKeyCredParams *[]WebauthnPublicKeyCredentialParameters `json:"pubKeyCredParams,omitempty"`

	// Information about the relying party (Microsoft Entra ID) responsible for the request.
	Rp *WebauthnPublicKeyCredentialRpEntity `json:"rp,omitempty"`

	// The time in milliseconds that the client is willing to wait for the credential creation operation to complete.
	Timeout nullable.Type[int64] `json:"timeout,omitempty"`

	// Information about the user account for which the credential is generated.
	User *WebauthnPublicKeyCredentialUserEntity `json:"user,omitempty"`
}
