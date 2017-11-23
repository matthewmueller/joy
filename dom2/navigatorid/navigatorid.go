package navigatorid

// js:"NavigatorID,omit"
type NavigatorID interface {

	// AppCodeName
	AppCodeName() (appCodeName string)

	// AppName
	AppName() (appName string)

	// AppVersion
	AppVersion() (appVersion string)

	// Platform
	Platform() (platform string)

	// Product
	Product() (product string)

	// ProductSub
	ProductSub() (productSub string)

	// UserAgent
	UserAgent() (userAgent string)

	// Vendor
	Vendor() (vendor string)

	// VendorSub
	VendorSub() (vendorSub string)
}
