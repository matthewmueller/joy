package scopedcredentialparameters

import "github.com/matthewmueller/golly/dom2/scopedcredentialtype"

type ScopedCredentialParameters struct {
	algorithm interface{}
	kind      *scopedcredentialtype.ScopedCredentialType
}
