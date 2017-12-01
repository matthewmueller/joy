package scopedcredentialdescriptor

import (
	"github.com/matthewmueller/golly/dom/scopedcredentialtype"
	"github.com/matthewmueller/golly/dom/transport"
)

type ScopedCredentialDescriptor struct {
	id         []byte
	transports *[]*transport.Transport
	kind       *scopedcredentialtype.ScopedCredentialType
}
