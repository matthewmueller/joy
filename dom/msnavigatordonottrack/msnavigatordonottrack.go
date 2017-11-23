package msnavigatordonottrack

import (
	"github.com/matthewmueller/golly/dom/confirmsitespecificexceptionsinformation"
	"github.com/matthewmueller/golly/dom/exceptioninformation"
	"github.com/matthewmueller/golly/dom/storeexceptionsinformation"
	"github.com/matthewmueller/golly/dom/storesitespecificexceptionsinformation"
	"github.com/matthewmueller/golly/js"
)

// MSNavigatorDoNotTrack struct
// js:"MSNavigatorDoNotTrack,omit"
type MSNavigatorDoNotTrack struct {
}

// ConfirmSiteSpecificTrackingException fn
func (*MSNavigatorDoNotTrack) ConfirmSiteSpecificTrackingException(args *confirmsitespecificexceptionsinformation.ConfirmSiteSpecificExceptionsInformation) (b bool) {
	js.Rewrite("$<.confirmSiteSpecificTrackingException($1)", args)
	return b
}

// ConfirmWebWideTrackingException fn
func (*MSNavigatorDoNotTrack) ConfirmWebWideTrackingException(args *exceptioninformation.ExceptionInformation) (b bool) {
	js.Rewrite("$<.confirmWebWideTrackingException($1)", args)
	return b
}

// RemoveSiteSpecificTrackingException fn
func (*MSNavigatorDoNotTrack) RemoveSiteSpecificTrackingException(args *exceptioninformation.ExceptionInformation) {
	js.Rewrite("$<.removeSiteSpecificTrackingException($1)", args)
}

// RemoveWebWideTrackingException fn
func (*MSNavigatorDoNotTrack) RemoveWebWideTrackingException(args *exceptioninformation.ExceptionInformation) {
	js.Rewrite("$<.removeWebWideTrackingException($1)", args)
}

// StoreSiteSpecificTrackingException fn
func (*MSNavigatorDoNotTrack) StoreSiteSpecificTrackingException(args *storesitespecificexceptionsinformation.StoreSiteSpecificExceptionsInformation) {
	js.Rewrite("$<.storeSiteSpecificTrackingException($1)", args)
}

// StoreWebWideTrackingException fn
func (*MSNavigatorDoNotTrack) StoreWebWideTrackingException(args *storeexceptionsinformation.StoreExceptionsInformation) {
	js.Rewrite("$<.storeWebWideTrackingException($1)", args)
}
