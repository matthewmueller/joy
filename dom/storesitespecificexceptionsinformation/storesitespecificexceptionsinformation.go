package storesitespecificexceptionsinformation

import "github.com/matthewmueller/golly/dom/storeexceptionsinformation"

type StoreSiteSpecificExceptionsInformation struct {
	*storeexceptionsinformation.StoreExceptionsInformation

	arrayOfDomainStrings *[]string
}
