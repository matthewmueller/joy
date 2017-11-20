package msassertion

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/mscredentialtype"
	"github.com/matthewmueller/golly/js"
)

type MSAssertion struct {
}

func (*MSAssertion) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*MSAssertion) GetType() (kind *mscredentialtype.MSCredentialType) {
	js.Rewrite("$<.type")
	return kind
}
