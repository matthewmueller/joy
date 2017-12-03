package scopedcredentialparameters

import "github.com/matthewmueller/joy/dom/scopedcredentialtype"

type ScopedCredentialParameters struct {
	algorithm interface{}
	kind      *scopedcredentialtype.ScopedCredentialType
}
