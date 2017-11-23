package scopedcredentialoptions

import "github.com/matthewmueller/golly/dom2/scopedcredentialdescriptor"

type ScopedCredentialOptions struct {
	excludeList    *[]*scopedcredentialdescriptor.ScopedCredentialDescriptor
	extensions     *webauthnextensions.WebAuthnExtensions
	rpId           *string
	timeoutSeconds *uint
}
