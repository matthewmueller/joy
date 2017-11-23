package navigatorid

import "github.com/matthewmueller/golly/js"

// NavigatorID struct
// js:"NavigatorID,omit"
type NavigatorID struct {
}

// AppCodeName prop
func (*NavigatorID) AppCodeName() (appCodeName string) {
	js.Rewrite("$<.appCodeName")
	return appCodeName
}

// AppName prop
func (*NavigatorID) AppName() (appName string) {
	js.Rewrite("$<.appName")
	return appName
}

// AppVersion prop
func (*NavigatorID) AppVersion() (appVersion string) {
	js.Rewrite("$<.appVersion")
	return appVersion
}

// Platform prop
func (*NavigatorID) Platform() (platform string) {
	js.Rewrite("$<.platform")
	return platform
}

// Product prop
func (*NavigatorID) Product() (product string) {
	js.Rewrite("$<.product")
	return product
}

// ProductSub prop
func (*NavigatorID) ProductSub() (productSub string) {
	js.Rewrite("$<.productSub")
	return productSub
}

// UserAgent prop
func (*NavigatorID) UserAgent() (userAgent string) {
	js.Rewrite("$<.userAgent")
	return userAgent
}

// Vendor prop
func (*NavigatorID) Vendor() (vendor string) {
	js.Rewrite("$<.vendor")
	return vendor
}

// VendorSub prop
func (*NavigatorID) VendorSub() (vendorSub string) {
	js.Rewrite("$<.vendorSub")
	return vendorSub
}
