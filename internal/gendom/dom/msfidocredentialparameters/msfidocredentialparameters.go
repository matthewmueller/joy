package msfidocredentialparameters

type MSFIDOCredentialParameters struct {
	*MSCredentialParameters

	algorithm      *interface{}
	authenticators *[]string
}
