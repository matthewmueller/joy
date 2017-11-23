package storesitespecificexceptionsinformation

import "github.com/matthewmueller/golly/dom2/storeexceptionsinformation"

type StoreSiteSpecificExceptionsInformation struct {
	*storeexceptionsinformation.StoreExceptionsInformation

	arrayOfDomainStrings *[]string
}
