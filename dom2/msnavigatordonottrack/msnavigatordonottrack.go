package msnavigatordonottrack

import (
	"github.com/matthewmueller/golly/dom2/confirmsitespecificexceptionsinformation"
	"github.com/matthewmueller/golly/dom2/exceptioninformation"
	"github.com/matthewmueller/golly/dom2/storeexceptionsinformation"
	"github.com/matthewmueller/golly/dom2/storesitespecificexceptionsinformation"
	"github.com/matthewmueller/golly/js"
)

// MSNavigatorDoNotTrack struct
// js:"MSNavigatorDoNotTrack,omit"
type MSNavigatorDoNotTrack struct {
}

// ConfirmSiteSpecificTrackingException
func (*MSNavigatorDoNotTrack) ConfirmSiteSpecificTrackingException(args *confirmsitespecificexceptionsinformation.ConfirmSiteSpecificExceptionsInformation) (b bool) {
	js.Rewrite("$<.ConfirmSiteSpecificTrackingException($1)", args)
	return b
}

// ConfirmWebWideTrackingException
func (*MSNavigatorDoNotTrack) ConfirmWebWideTrackingException(args *exceptioninformation.ExceptionInformation) (b bool) {
	js.Rewrite("$<.ConfirmWebWideTrackingException($1)", args)
	return b
}

// RemoveSiteSpecificTrackingException
func (*MSNavigatorDoNotTrack) RemoveSiteSpecificTrackingException(args *exceptioninformation.ExceptionInformation) {
	js.Rewrite("$<.RemoveSiteSpecificTrackingException($1)", args)
}

// RemoveWebWideTrackingException
func (*MSNavigatorDoNotTrack) RemoveWebWideTrackingException(args *exceptioninformation.ExceptionInformation) {
	js.Rewrite("$<.RemoveWebWideTrackingException($1)", args)
}

// StoreSiteSpecificTrackingException
func (*MSNavigatorDoNotTrack) StoreSiteSpecificTrackingException(args *storesitespecificexceptionsinformation.StoreSiteSpecificExceptionsInformation) {
	js.Rewrite("$<.StoreSiteSpecificTrackingException($1)", args)
}

// StoreWebWideTrackingException
func (*MSNavigatorDoNotTrack) StoreWebWideTrackingException(args *storeexceptionsinformation.StoreExceptionsInformation) {
	js.Rewrite("$<.StoreWebWideTrackingException($1)", args)
}
