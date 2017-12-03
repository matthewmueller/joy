package confirmsitespecificexceptionsinformation

import "github.com/matthewmueller/joy/dom/exceptioninformation"

type ConfirmSiteSpecificExceptionsInformation struct {
	*exceptioninformation.ExceptionInformation

	arrayOfDomainStrings *[]string
}
