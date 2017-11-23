package scopedcredentialoptions

import (
	"github.com/matthewmueller/golly/dom2/scopedcredentialdescriptor"
	"github.com/matthewmueller/golly/dom2/webauthnextensions"
)

type ScopedCredentialOptions struct {
	excludeList    *[]*scopedcredentialdescriptor.ScopedCredentialDescriptor
	extensions     *webauthnextensions.WebAuthnExtensions
	rpId           *string
	timeoutSeconds *uint
}
