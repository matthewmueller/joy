package navigatorid

// NavigatorID interface
// js:"NavigatorID"
type NavigatorID interface {

	// AppCodeName prop
	// js:"appCodeName"
	// jsrewrite:"$_.appCodeName"
	AppCodeName() (appCodeName string)

	// AppName prop
	// js:"appName"
	// jsrewrite:"$_.appName"
	AppName() (appName string)

	// AppVersion prop
	// js:"appVersion"
	// jsrewrite:"$_.appVersion"
	AppVersion() (appVersion string)

	// Platform prop
	// js:"platform"
	// jsrewrite:"$_.platform"
	Platform() (platform string)

	// Product prop
	// js:"product"
	// jsrewrite:"$_.product"
	Product() (product string)

	// ProductSub prop
	// js:"productSub"
	// jsrewrite:"$_.productSub"
	ProductSub() (productSub string)

	// UserAgent prop
	// js:"userAgent"
	// jsrewrite:"$_.userAgent"
	UserAgent() (userAgent string)

	// Vendor prop
	// js:"vendor"
	// jsrewrite:"$_.vendor"
	Vendor() (vendor string)

	// VendorSub prop
	// js:"vendorSub"
	// jsrewrite:"$_.vendorSub"
	VendorSub() (vendorSub string)
}
