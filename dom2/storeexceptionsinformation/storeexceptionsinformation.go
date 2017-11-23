package storeexceptionsinformation

type StoreExceptionsInformation struct {
	*exceptioninformation.ExceptionInformation

	detailURI         *string
	explanationString *string
	siteName          *string
}
