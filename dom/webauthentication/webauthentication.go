package webauthentication

import (
	"github.com/matthewmueller/golly/dom/account"
	"github.com/matthewmueller/golly/dom/assertionoptions"
	"github.com/matthewmueller/golly/dom/scopedcredentialinfo"
	"github.com/matthewmueller/golly/dom/scopedcredentialoptions"
	"github.com/matthewmueller/golly/dom/scopedcredentialparameters"
	"github.com/matthewmueller/golly/dom/webauthnassertion"
	"github.com/matthewmueller/golly/js"
)

// WebAuthentication struct
// js:"WebAuthentication,omit"
type WebAuthentication struct {
}

// GetAssertion fn
// js:"getAssertion"
func (*WebAuthentication) GetAssertion(assertionChallenge []byte, options *assertionoptions.AssertionOptions) (w *webauthnassertion.WebAuthnAssertion) {
	js.Rewrite("await $_.getAssertion($1, $2)", assertionChallenge, options)
	return w
}

// MakeCredential fn
// js:"makeCredential"
func (*WebAuthentication) MakeCredential(accountInformation *account.Account, cryptoParameters []*scopedcredentialparameters.ScopedCredentialParameters, attestationChallenge []byte, options *scopedcredentialoptions.ScopedCredentialOptions) (s *scopedcredentialinfo.ScopedCredentialInfo) {
	js.Rewrite("await $_.makeCredential($1, $2, $3, $4)", accountInformation, cryptoParameters, attestationChallenge, options)
	return s
}
