package storeexceptionsinformation

import "github.com/matthewmueller/joy/dom/exceptioninformation"

type StoreExceptionsInformation struct {
	*exceptioninformation.ExceptionInformation

	detailURI         *string
	explanationString *string
	siteName          *string
}
