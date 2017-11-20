package msnavigatordonottrack

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/confirmsitespecificexceptionsinformation"
	"github.com/matthewmueller/golly/internal/gendom/dom/exceptioninformation"
	"github.com/matthewmueller/golly/internal/gendom/dom/storeexceptionsinformation"
	"github.com/matthewmueller/golly/internal/gendom/dom/storesitespecificexceptionsinformation"
	"github.com/matthewmueller/golly/js"
)

type MSNavigatorDoNotTrack struct {
}

func (*MSNavigatorDoNotTrack) ConfirmSiteSpecificTrackingException(args *confirmsitespecificexceptionsinformation.ConfirmSiteSpecificExceptionsInformation) (b bool) {
	js.Rewrite("$<.confirmSiteSpecificTrackingException($1)", args)
	return b
}

func (*MSNavigatorDoNotTrack) ConfirmWebWideTrackingException(args *exceptioninformation.ExceptionInformation) (b bool) {
	js.Rewrite("$<.confirmWebWideTrackingException($1)", args)
	return b
}

func (*MSNavigatorDoNotTrack) RemoveSiteSpecificTrackingException(args *exceptioninformation.ExceptionInformation) {
	js.Rewrite("$<.removeSiteSpecificTrackingException($1)", args)
}

func (*MSNavigatorDoNotTrack) RemoveWebWideTrackingException(args *exceptioninformation.ExceptionInformation) {
	js.Rewrite("$<.removeWebWideTrackingException($1)", args)
}

func (*MSNavigatorDoNotTrack) StoreSiteSpecificTrackingException(args *storesitespecificexceptionsinformation.StoreSiteSpecificExceptionsInformation) {
	js.Rewrite("$<.storeSiteSpecificTrackingException($1)", args)
}

func (*MSNavigatorDoNotTrack) StoreWebWideTrackingException(args *storeexceptionsinformation.StoreExceptionsInformation) {
	js.Rewrite("$<.storeWebWideTrackingException($1)", args)
}
