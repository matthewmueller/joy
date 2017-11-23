package mscredentials

import (
	"github.com/matthewmueller/golly/dom/msaccountinfo"
	"github.com/matthewmueller/golly/dom/msassertion"
	"github.com/matthewmueller/golly/dom/mscredentialfilter"
	"github.com/matthewmueller/golly/dom/mscredentialparameters"
	"github.com/matthewmueller/golly/dom/mssignatureparameters"
	"github.com/matthewmueller/golly/js"
)

// MSCredentials struct
// js:"MSCredentials,omit"
type MSCredentials struct {
}

// GetAssertion fn
func (*MSCredentials) GetAssertion(challenge string, filter *mscredentialfilter.MSCredentialFilter, params *mssignatureparameters.MSSignatureParameters) (m msassertion.MSAssertion) {
	js.Rewrite("await $<.getAssertion($1, $2, $3)", challenge, filter, params)
	return m
}

// MakeCredential fn
func (*MSCredentials) MakeCredential(accountInfo *msaccountinfo.MSAccountInfo, params []*mscredentialparameters.MSCredentialParameters, challenge *string) (m msassertion.MSAssertion) {
	js.Rewrite("await $<.makeCredential($1, $2, $3)", accountInfo, params, challenge)
	return m
}
