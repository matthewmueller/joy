package assertionoptions

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/scopedcredentialdescriptor"
	"github.com/matthewmueller/golly/internal/gendom/dom/webauthnextensions"
)

type AssertionOptions struct {
	allowList      *[]*scopedcredentialdescriptor.ScopedCredentialDescriptor
	extensions     *webauthnextensions.WebAuthnExtensions
	rpId           *string
	timeoutSeconds *uint
}
