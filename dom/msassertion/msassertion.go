package msassertion

import "github.com/matthewmueller/golly/dom/mscredentialtype"

// MSAssertion interface
// js:"MSAssertion"
type MSAssertion interface {

	// ID prop
	// js:"id"
	// jsrewrite:"$_.id"
	ID() (id string)

	// Type prop
	// js:"type"
	// jsrewrite:"$_.type"
	Type() (kind *mscredentialtype.MSCredentialType)
}
