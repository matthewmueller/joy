package scopedcredentialparameters

import "github.com/matthewmueller/golly/internal/gendom/dom/scopedcredentialtype"

type ScopedCredentialParameters struct {
	algorithm interface{}
	kind      *scopedcredentialtype.ScopedCredentialType
}
