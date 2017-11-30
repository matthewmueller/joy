package msnavigatordonottrack

import (
	"github.com/matthewmueller/golly/dom/confirmsitespecificexceptionsinformation"
	"github.com/matthewmueller/golly/dom/exceptioninformation"
	"github.com/matthewmueller/golly/dom/storeexceptionsinformation"
	"github.com/matthewmueller/golly/dom/storesitespecificexceptionsinformation"
)

// MSNavigatorDoNotTrack interface
// js:"MSNavigatorDoNotTrack"
type MSNavigatorDoNotTrack interface {

	// ConfirmSiteSpecificTrackingException
	// js:"confirmSiteSpecificTrackingException"
	ConfirmSiteSpecificTrackingException(args *confirmsitespecificexceptionsinformation.ConfirmSiteSpecificExceptionsInformation) (b bool)

	// ConfirmWebWideTrackingException
	// js:"confirmWebWideTrackingException"
	ConfirmWebWideTrackingException(args *exceptioninformation.ExceptionInformation) (b bool)

	// RemoveSiteSpecificTrackingException
	// js:"removeSiteSpecificTrackingException"
	RemoveSiteSpecificTrackingException(args *exceptioninformation.ExceptionInformation)

	// RemoveWebWideTrackingException
	// js:"removeWebWideTrackingException"
	RemoveWebWideTrackingException(args *exceptioninformation.ExceptionInformation)

	// StoreSiteSpecificTrackingException
	// js:"storeSiteSpecificTrackingException"
	StoreSiteSpecificTrackingException(args *storesitespecificexceptionsinformation.StoreSiteSpecificExceptionsInformation)

	// StoreWebWideTrackingException
	// js:"storeWebWideTrackingException"
	StoreWebWideTrackingException(args *storeexceptionsinformation.StoreExceptionsInformation)
}
