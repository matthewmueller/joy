package msassertion

import (
	"github.com/matthewmueller/golly/dom/mscredentialtype"
	"github.com/matthewmueller/golly/js"
)

// js:"MSAssertion,omit"
type MSAssertion interface {

	// ID prop
	// js:"id,rewrite=id"
	ID() (id string)

	// Type prop
	// js:"type,rewrite=kind"
	Type() (kind *mscredentialtype.MSCredentialType)
}

// id prop
func id() (id string) {
	js.Rewrite("$<.id")
	return id
}

// kind prop
func kind() (kind *mscredentialtype.MSCredentialType) {
	js.Rewrite("$<.type")
	return kind
}
