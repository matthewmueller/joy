package webauthentication

import (
	"github.com/matthewmueller/joy/dom/account"
	"github.com/matthewmueller/joy/dom/assertionoptions"
	"github.com/matthewmueller/joy/dom/scopedcredentialinfo"
	"github.com/matthewmueller/joy/dom/scopedcredentialoptions"
	"github.com/matthewmueller/joy/dom/scopedcredentialparameters"
	"github.com/matthewmueller/joy/dom/webauthnassertion"
	"github.com/matthewmueller/joy/macro"
)

// WebAuthentication struct
// js:"WebAuthentication,omit"
type WebAuthentication struct {
}

// GetAssertion fn
// js:"getAssertion"
func (*WebAuthentication) GetAssertion(assertionChallenge []byte, options *assertionoptions.AssertionOptions) (w *webauthnassertion.WebAuthnAssertion) {
	macro.Rewrite("await $_.getAssertion($1, $2)", assertionChallenge, options)
	return w
}

// MakeCredential fn
// js:"makeCredential"
func (*WebAuthentication) MakeCredential(accountInformation *account.Account, cryptoParameters []*scopedcredentialparameters.ScopedCredentialParameters, attestationChallenge []byte, options *scopedcredentialoptions.ScopedCredentialOptions) (s *scopedcredentialinfo.ScopedCredentialInfo) {
	macro.Rewrite("await $_.makeCredential($1, $2, $3, $4)", accountInformation, cryptoParameters, attestationChallenge, options)
	return s
}
