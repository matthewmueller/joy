package msnavigatordonottrack

import (
	"github.com/matthewmueller/golly/dom2/confirmsitespecificexceptionsinformation"
	"github.com/matthewmueller/golly/dom2/exceptioninformation"
	"github.com/matthewmueller/golly/dom2/storeexceptionsinformation"
	"github.com/matthewmueller/golly/dom2/storesitespecificexceptionsinformation"
)

// js:"MSNavigatorDoNotTrack,omit"
type MSNavigatorDoNotTrack interface {

	// ConfirmSiteSpecificTrackingException
	ConfirmSiteSpecificTrackingException(args *confirmsitespecificexceptionsinformation.ConfirmSiteSpecificExceptionsInformation) (b bool)

	// ConfirmWebWideTrackingException
	ConfirmWebWideTrackingException(args *exceptioninformation.ExceptionInformation) (b bool)

	// RemoveSiteSpecificTrackingException
	RemoveSiteSpecificTrackingException(args *exceptioninformation.ExceptionInformation)

	// RemoveWebWideTrackingException
	RemoveWebWideTrackingException(args *exceptioninformation.ExceptionInformation)

	// StoreSiteSpecificTrackingException
	StoreSiteSpecificTrackingException(args *storesitespecificexceptionsinformation.StoreSiteSpecificExceptionsInformation)

	// StoreWebWideTrackingException
	StoreWebWideTrackingException(args *storeexceptionsinformation.StoreExceptionsInformation)
}
