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
	// jsrewrite:"$_.confirmSiteSpecificTrackingException($1)"
	ConfirmSiteSpecificTrackingException(args *confirmsitespecificexceptionsinformation.ConfirmSiteSpecificExceptionsInformation) (b bool)

	// ConfirmWebWideTrackingException
	// js:"confirmWebWideTrackingException"
	// jsrewrite:"$_.confirmWebWideTrackingException($1)"
	ConfirmWebWideTrackingException(args *exceptioninformation.ExceptionInformation) (b bool)

	// RemoveSiteSpecificTrackingException
	// js:"removeSiteSpecificTrackingException"
	// jsrewrite:"$_.removeSiteSpecificTrackingException($1)"
	RemoveSiteSpecificTrackingException(args *exceptioninformation.ExceptionInformation)

	// RemoveWebWideTrackingException
	// js:"removeWebWideTrackingException"
	// jsrewrite:"$_.removeWebWideTrackingException($1)"
	RemoveWebWideTrackingException(args *exceptioninformation.ExceptionInformation)

	// StoreSiteSpecificTrackingException
	// js:"storeSiteSpecificTrackingException"
	// jsrewrite:"$_.storeSiteSpecificTrackingException($1)"
	StoreSiteSpecificTrackingException(args *storesitespecificexceptionsinformation.StoreSiteSpecificExceptionsInformation)

	// StoreWebWideTrackingException
	// js:"storeWebWideTrackingException"
	// jsrewrite:"$_.storeWebWideTrackingException($1)"
	StoreWebWideTrackingException(args *storeexceptionsinformation.StoreExceptionsInformation)
}
