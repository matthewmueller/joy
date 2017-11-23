package assertionoptions

import (
	"github.com/matthewmueller/golly/dom2/scopedcredentialdescriptor"
	"github.com/matthewmueller/golly/dom2/webauthnextensions"
)

type AssertionOptions struct {
	allowList      *[]*scopedcredentialdescriptor.ScopedCredentialDescriptor
	extensions     *webauthnextensions.WebAuthnExtensions
	rpId           *string
	timeoutSeconds *uint
}
