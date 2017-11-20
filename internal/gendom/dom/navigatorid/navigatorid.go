package navigatorid

import "github.com/matthewmueller/golly/js"

type NavigatorID struct {
}

func (*NavigatorID) GetAppCodeName() (appCodeName string) {
	js.Rewrite("$<.appCodeName")
	return appCodeName
}

func (*NavigatorID) GetAppName() (appName string) {
	js.Rewrite("$<.appName")
	return appName
}

func (*NavigatorID) GetAppVersion() (appVersion string) {
	js.Rewrite("$<.appVersion")
	return appVersion
}

func (*NavigatorID) GetPlatform() (platform string) {
	js.Rewrite("$<.platform")
	return platform
}

func (*NavigatorID) GetProduct() (product string) {
	js.Rewrite("$<.product")
	return product
}

func (*NavigatorID) GetProductSub() (productSub string) {
	js.Rewrite("$<.productSub")
	return productSub
}

func (*NavigatorID) GetUserAgent() (userAgent string) {
	js.Rewrite("$<.userAgent")
	return userAgent
}

func (*NavigatorID) GetVendor() (vendor string) {
	js.Rewrite("$<.vendor")
	return vendor
}

func (*NavigatorID) GetVendorSub() (vendorSub string) {
	js.Rewrite("$<.vendorSub")
	return vendorSub
}
