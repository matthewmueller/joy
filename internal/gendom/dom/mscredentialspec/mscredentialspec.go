package mscredentialspec

import "github.com/matthewmueller/golly/internal/gendom/dom/mscredentialtype"

type MSCredentialSpec struct {
	id   *string
	kind *mscredentialtype.MSCredentialType
}
