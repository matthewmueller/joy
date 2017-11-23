package msfidocredentialparameters

import "github.com/matthewmueller/golly/dom2/mscredentialparameters"

type MSFIDOCredentialParameters struct {
	*mscredentialparameters.MSCredentialParameters

	algorithm      *interface{}
	authenticators *[]string
}
