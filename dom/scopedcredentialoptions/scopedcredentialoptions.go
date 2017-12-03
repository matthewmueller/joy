package scopedcredentialoptions

import (
	"github.com/matthewmueller/joy/dom/scopedcredentialdescriptor"
	"github.com/matthewmueller/joy/dom/webauthnextensions"
)

type ScopedCredentialOptions struct {
	excludeList    *[]*scopedcredentialdescriptor.ScopedCredentialDescriptor
	extensions     *webauthnextensions.WebAuthnExtensions
	rpId           *string
	timeoutSeconds *uint
}
