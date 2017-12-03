package mscredentials

import (
	"github.com/matthewmueller/joy/dom/msaccountinfo"
	"github.com/matthewmueller/joy/dom/msassertion"
	"github.com/matthewmueller/joy/dom/mscredentialfilter"
	"github.com/matthewmueller/joy/dom/mscredentialparameters"
	"github.com/matthewmueller/joy/dom/mssignatureparameters"
	"github.com/matthewmueller/joy/macro"
)

// MSCredentials struct
// js:"MSCredentials,omit"
type MSCredentials struct {
}

// GetAssertion fn
// js:"getAssertion"
func (*MSCredentials) GetAssertion(challenge string, filter *mscredentialfilter.MSCredentialFilter, params *mssignatureparameters.MSSignatureParameters) (m msassertion.MSAssertion) {
	macro.Rewrite("await $_.getAssertion($1, $2, $3)", challenge, filter, params)
	return m
}

// MakeCredential fn
// js:"makeCredential"
func (*MSCredentials) MakeCredential(accountInfo *msaccountinfo.MSAccountInfo, params []*mscredentialparameters.MSCredentialParameters, challenge *string) (m msassertion.MSAssertion) {
	macro.Rewrite("await $_.makeCredential($1, $2, $3)", accountInfo, params, challenge)
	return m
}
