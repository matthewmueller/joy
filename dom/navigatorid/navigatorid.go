package navigatorid

// NavigatorID interface
// js:"NavigatorID"
type NavigatorID interface {

	// AppCodeName prop
	// js:"appCodeName"
	AppCodeName() (appCodeName string)

	// AppName prop
	// js:"appName"
	AppName() (appName string)

	// AppVersion prop
	// js:"appVersion"
	AppVersion() (appVersion string)

	// Platform prop
	// js:"platform"
	Platform() (platform string)

	// Product prop
	// js:"product"
	Product() (product string)

	// ProductSub prop
	// js:"productSub"
	ProductSub() (productSub string)

	// UserAgent prop
	// js:"userAgent"
	UserAgent() (userAgent string)

	// Vendor prop
	// js:"vendor"
	Vendor() (vendor string)

	// VendorSub prop
	// js:"vendorSub"
	VendorSub() (vendorSub string)
}
