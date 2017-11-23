package msfidocredentialparameters

type MSFIDOCredentialParameters struct {
	*mscredentialparameters.MSCredentialParameters

	algorithm      *interface{}
	authenticators *[]string
}
