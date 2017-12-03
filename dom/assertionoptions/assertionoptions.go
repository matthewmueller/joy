package assertionoptions

import (
	"github.com/matthewmueller/joy/dom/scopedcredentialdescriptor"
	"github.com/matthewmueller/joy/dom/webauthnextensions"
)

type AssertionOptions struct {
	allowList      *[]*scopedcredentialdescriptor.ScopedCredentialDescriptor
	extensions     *webauthnextensions.WebAuthnExtensions
	rpId           *string
	timeoutSeconds *uint
}
