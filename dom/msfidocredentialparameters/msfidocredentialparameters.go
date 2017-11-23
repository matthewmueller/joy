package msfidocredentialparameters

import "github.com/matthewmueller/golly/dom/mscredentialparameters"

type MSFIDOCredentialParameters struct {
	*mscredentialparameters.MSCredentialParameters

	algorithm      *interface{}
	authenticators *[]string
}
