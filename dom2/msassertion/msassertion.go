package msassertion

import "github.com/matthewmueller/golly/dom2/mscredentialtype"

// js:"MSAssertion,omit"
type MSAssertion interface {

	// ID
	ID() (id string)

	// Type
	Type() (kind *mscredentialtype.MSCredentialType)
}
