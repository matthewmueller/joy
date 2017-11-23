package scopedcredentialoptions

import (
	"github.com/matthewmueller/golly/dom/scopedcredentialdescriptor"
	"github.com/matthewmueller/golly/dom/webauthnextensions"
)

type ScopedCredentialOptions struct {
	excludeList    *[]*scopedcredentialdescriptor.ScopedCredentialDescriptor
	extensions     *webauthnextensions.WebAuthnExtensions
	rpId           *string
	timeoutSeconds *uint
}
