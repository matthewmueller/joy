package storeexceptionsinformation

import "github.com/matthewmueller/golly/dom2/exceptioninformation"

type StoreExceptionsInformation struct {
	*exceptioninformation.ExceptionInformation

	detailURI         *string
	explanationString *string
	siteName          *string
}
