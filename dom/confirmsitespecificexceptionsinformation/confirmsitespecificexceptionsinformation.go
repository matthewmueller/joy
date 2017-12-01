package confirmsitespecificexceptionsinformation

import "github.com/matthewmueller/golly/dom/exceptioninformation"

type ConfirmSiteSpecificExceptionsInformation struct {
	*exceptioninformation.ExceptionInformation

	arrayOfDomainStrings *[]string
}
