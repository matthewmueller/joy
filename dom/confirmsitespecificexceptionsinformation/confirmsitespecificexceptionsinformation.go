package confirmsitespecificexceptionsinformation

import "github.com/matthewmueller/golly/dom2/exceptioninformation"

type ConfirmSiteSpecificExceptionsInformation struct {
	*exceptioninformation.ExceptionInformation

	arrayOfDomainStrings *[]string
}
