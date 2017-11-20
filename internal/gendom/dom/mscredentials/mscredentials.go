package mscredentials

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/msaccountinfo"
	"github.com/matthewmueller/golly/internal/gendom/dom/msassertion"
	"github.com/matthewmueller/golly/internal/gendom/dom/mscredentialfilter"
	"github.com/matthewmueller/golly/internal/gendom/dom/mscredentialparameters"
	"github.com/matthewmueller/golly/internal/gendom/dom/mssignatureparameters"
	"github.com/matthewmueller/golly/js"
)

type MSCredentials struct {
}

func (*MSCredentials) GetAssertion(challenge string, filter *mscredentialfilter.MSCredentialFilter, params *mssignatureparameters.MSSignatureParameters) (m *msassertion.MSAssertion) {
	js.Rewrite("await $<.getAssertion($1, $2, $3)", challenge, filter, params)
	return m
}

func (*MSCredentials) MakeCredential(accountInfo *msaccountinfo.MSAccountInfo, params []*mscredentialparameters.MSCredentialParameters, challenge string) (m *msassertion.MSAssertion) {
	js.Rewrite("await $<.makeCredential($1, $2, $3)", accountInfo, params, challenge)
	return m
}
