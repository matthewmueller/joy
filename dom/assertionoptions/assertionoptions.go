package assertionoptions

import (
	"github.com/matthewmueller/golly/dom/scopedcredentialdescriptor"
	"github.com/matthewmueller/golly/dom/webauthnextensions"
)

type AssertionOptions struct {
	allowList      *[]*scopedcredentialdescriptor.ScopedCredentialDescriptor
	extensions     *webauthnextensions.WebAuthnExtensions
	rpId           *string
	timeoutSeconds *uint
}
