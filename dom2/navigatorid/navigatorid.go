package navigatorid

import "github.com/matthewmueller/golly/js"

// NavigatorID struct
// js:"NavigatorID,omit"
type NavigatorID struct {
}

// AppCodeName
func (*NavigatorID) AppCodeName() (appCodeName string) {
	js.Rewrite("$<.AppCodeName")
	return appCodeName
}

// AppName
func (*NavigatorID) AppName() (appName string) {
	js.Rewrite("$<.AppName")
	return appName
}

// AppVersion
func (*NavigatorID) AppVersion() (appVersion string) {
	js.Rewrite("$<.AppVersion")
	return appVersion
}

// Platform
func (*NavigatorID) Platform() (platform string) {
	js.Rewrite("$<.Platform")
	return platform
}

// Product
func (*NavigatorID) Product() (product string) {
	js.Rewrite("$<.Product")
	return product
}

// ProductSub
func (*NavigatorID) ProductSub() (productSub string) {
	js.Rewrite("$<.ProductSub")
	return productSub
}

// UserAgent
func (*NavigatorID) UserAgent() (userAgent string) {
	js.Rewrite("$<.UserAgent")
	return userAgent
}

// Vendor
func (*NavigatorID) Vendor() (vendor string) {
	js.Rewrite("$<.Vendor")
	return vendor
}

// VendorSub
func (*NavigatorID) VendorSub() (vendorSub string) {
	js.Rewrite("$<.VendorSub")
	return vendorSub
}
