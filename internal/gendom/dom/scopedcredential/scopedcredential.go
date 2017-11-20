package scopedcredential

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/scopedcredentialtype"
	"github.com/matthewmueller/golly/js"
)

type ScopedCredential struct {
}

func (*ScopedCredential) GetID() (id []byte) {
	js.Rewrite("$<.id")
	return id
}

func (*ScopedCredential) GetType() (kind *scopedcredentialtype.ScopedCredentialType) {
	js.Rewrite("$<.type")
	return kind
}
