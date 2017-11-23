package mscredentialspec

import "github.com/matthewmueller/golly/dom2/mscredentialtype"

type MSCredentialSpec struct {
	id   *string
	kind *mscredentialtype.MSCredentialType
}
