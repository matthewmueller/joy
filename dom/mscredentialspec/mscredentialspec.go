package mscredentialspec

import "github.com/matthewmueller/joy/dom/mscredentialtype"

type MSCredentialSpec struct {
	id   *string
	kind *mscredentialtype.MSCredentialType
}
