package scopedcredentialparameters

import "github.com/matthewmueller/golly/dom/scopedcredentialtype"

type ScopedCredentialParameters struct {
	algorithm interface{}
	kind      *scopedcredentialtype.ScopedCredentialType
}
