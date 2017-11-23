package scopedcredentialdescriptor

import (
	"github.com/matthewmueller/golly/dom2/scopedcredentialtype"
	"github.com/matthewmueller/golly/dom2/transport"
)

type ScopedCredentialDescriptor struct {
	id         []byte
	transports *[]*transport.Transport
	kind       *scopedcredentialtype.ScopedCredentialType
}
