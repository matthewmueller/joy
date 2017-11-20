package storesitespecificexceptionsinformation

type StoreSiteSpecificExceptionsInformation struct {
	*StoreExceptionsInformation

	arrayOfDomainStrings *[]string
}
