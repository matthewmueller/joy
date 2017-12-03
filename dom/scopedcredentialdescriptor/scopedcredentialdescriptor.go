package scopedcredentialdescriptor

import (
	"github.com/matthewmueller/joy/dom/scopedcredentialtype"
	"github.com/matthewmueller/joy/dom/transport"
)

type ScopedCredentialDescriptor struct {
	id         []byte
	transports *[]*transport.Transport
	kind       *scopedcredentialtype.ScopedCredentialType
}
