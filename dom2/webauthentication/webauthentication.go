package webauthentication

import (
	"github.com/matthewmueller/golly/dom2/account"
	"github.com/matthewmueller/golly/dom2/assertionoptions"
	"github.com/matthewmueller/golly/dom2/scopedcredentialinfo"
	"github.com/matthewmueller/golly/dom2/scopedcredentialoptions"
	"github.com/matthewmueller/golly/dom2/scopedcredentialparameters"
	"github.com/matthewmueller/golly/dom2/webauthnassertion"
	"github.com/matthewmueller/golly/js"
)

// js:"WebAuthentication,omit"
type WebAuthentication struct {
}

// GetAssertion
func (*WebAuthentication) GetAssertion(assertionChallenge []byte, options *assertionoptions.AssertionOptions) (w *webauthnassertion.WebAuthnAssertion) {
	js.Rewrite("await $<.GetAssertion($1, $2)", assertionChallenge, options)
	return w
}

// MakeCredential
func (*WebAuthentication) MakeCredential(accountInformation *account.Account, cryptoParameters []*scopedcredentialparameters.ScopedCredentialParameters, attestationChallenge []byte, options *scopedcredentialoptions.ScopedCredentialOptions) (s *scopedcredentialinfo.ScopedCredentialInfo) {
	js.Rewrite("await $<.MakeCredential($1, $2, $3, $4)", accountInformation, cryptoParameters, attestationChallenge, options)
	return s
}
