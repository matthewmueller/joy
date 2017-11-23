package mscredentialspec

import "github.com/matthewmueller/golly/dom/mscredentialtype"

type MSCredentialSpec struct {
	id   *string
	kind *mscredentialtype.MSCredentialType
}
