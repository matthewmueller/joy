package scopedcredentialoptions

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/scopedcredentialdescriptor"
	"github.com/matthewmueller/golly/internal/gendom/dom/webauthnextensions"
)

type ScopedCredentialOptions struct {
	excludeList    *[]*scopedcredentialdescriptor.ScopedCredentialDescriptor
	extensions     *webauthnextensions.WebAuthnExtensions
	rpId           *string
	timeoutSeconds *uint
}
