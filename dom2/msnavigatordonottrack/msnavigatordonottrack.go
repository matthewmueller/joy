package msnavigatordonottrack

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
