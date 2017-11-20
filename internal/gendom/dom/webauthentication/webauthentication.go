package webauthentication

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/account"
	"github.com/matthewmueller/golly/internal/gendom/dom/assertionoptions"
	"github.com/matthewmueller/golly/internal/gendom/dom/scopedcredentialinfo"
	"github.com/matthewmueller/golly/internal/gendom/dom/scopedcredentialoptions"
	"github.com/matthewmueller/golly/internal/gendom/dom/scopedcredentialparameters"
	"github.com/matthewmueller/golly/internal/gendom/dom/webauthnassertion"
	"github.com/matthewmueller/golly/js"
)

type WebAuthentication struct {
}

func (*WebAuthentication) GetAssertion(assertionChallenge []byte, options *assertionoptions.AssertionOptions) (w *webauthnassertion.WebAuthnAssertion) {
	js.Rewrite("await $<.getAssertion($1, $2)", assertionChallenge, options)
	return w
}

func (*WebAuthentication) MakeCredential(accountInformation *account.Account, cryptoParameters []*scopedcredentialparameters.ScopedCredentialParameters, attestationChallenge []byte, options *scopedcredentialoptions.ScopedCredentialOptions) (s *scopedcredentialinfo.ScopedCredentialInfo) {
	js.Rewrite("await $<.makeCredential($1, $2, $3, $4)", accountInformation, cryptoParameters, attestationChallenge, options)
	return s
}
