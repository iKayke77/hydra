// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package consent

import (
	"context"

	"github.com/ory/fosite/handler/openid"
	"github.com/ory/hydra/v2/client"
	"github.com/ory/hydra/v2/jwk"
	"github.com/ory/hydra/v2/x"
)

type InternalRegistry interface {
	x.RegistryWriter
	x.RegistryCookieStore
	x.RegistryLogger
	x.HTTPClientProvider
	Registry
	client.Registry

	KeyCipher() *jwk.AEAD
	OAuth2Storage() x.FositeStorer
	OpenIDConnectRequestValidator() *openid.OpenIDConnectRequestValidator
}

type Registry interface {
	ConsentManager() Manager
	ConsentStrategy() Strategy
	SubjectIdentifierAlgorithm(ctx context.Context) map[string]SubjectIdentifierAlgorithm
}
