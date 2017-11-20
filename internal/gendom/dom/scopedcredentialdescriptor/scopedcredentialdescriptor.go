package scopedcredentialdescriptor

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/scopedcredentialtype"
	"github.com/matthewmueller/golly/internal/gendom/dom/transport"
)

type ScopedCredentialDescriptor struct {
	id         []byte
	transports *[]*transport.Transport
	kind       *scopedcredentialtype.ScopedCredentialType
}
