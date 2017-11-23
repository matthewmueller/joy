package scopedcredentialdescriptor

import "github.com/src-d/go-git/plumbing/transport"

type ScopedCredentialDescriptor struct {
	id         []byte
	transports *[]*transport.Transport
	kind       *scopedcredentialtype.ScopedCredentialType
}
