package window

import (
	"github.com/matthewmueller/golly/dom2/confirmsitespecificexceptionsinformation"
	"github.com/matthewmueller/golly/dom2/exceptioninformation"
	"github.com/matthewmueller/golly/dom2/gamepad"
	"github.com/matthewmueller/golly/dom2/gamepadinputemulationtype"
	"github.com/matthewmueller/golly/dom2/geolocation"
	"github.com/matthewmueller/golly/dom2/mediakeysystemconfiguration"
	"github.com/matthewmueller/golly/dom2/mediastreamconstraints"
	"github.com/matthewmueller/golly/dom2/mediastreamerror"
	"github.com/matthewmueller/golly/dom2/mimetypearray"
	"github.com/matthewmueller/golly/dom2/pluginarray"
	"github.com/matthewmueller/golly/dom2/storeexceptionsinformation"
	"github.com/matthewmueller/golly/dom2/storesitespecificexceptionsinformation"
	"github.com/matthewmueller/golly/dom2/webauthentication"
	"github.com/matthewmueller/golly/js"
)

// Navigator struct
// js:"Navigator,omit"
type Navigator struct {
}

// ConfirmSiteSpecificTrackingException fn
func (*Navigator) ConfirmSiteSpecificTrackingException(args *confirmsitespecificexceptionsinformation.ConfirmSiteSpecificExceptionsInformation) (b bool) {
	js.Rewrite("$<.confirmSiteSpecificTrackingException($1)", args)
	return b
}

// ConfirmWebWideTrackingException fn
func (*Navigator) ConfirmWebWideTrackingException(args *exceptioninformation.ExceptionInformation) (b bool) {
	js.Rewrite("$<.confirmWebWideTrackingException($1)", args)
	return b
}

// RemoveSiteSpecificTrackingException fn
func (*Navigator) RemoveSiteSpecificTrackingException(args *exceptioninformation.ExceptionInformation) {
	js.Rewrite("$<.removeSiteSpecificTrackingException($1)", args)
}

// RemoveWebWideTrackingException fn
func (*Navigator) RemoveWebWideTrackingException(args *exceptioninformation.ExceptionInformation) {
	js.Rewrite("$<.removeWebWideTrackingException($1)", args)
}

// StoreSiteSpecificTrackingException fn
func (*Navigator) StoreSiteSpecificTrackingException(args *storesitespecificexceptionsinformation.StoreSiteSpecificExceptionsInformation) {
	js.Rewrite("$<.storeSiteSpecificTrackingException($1)", args)
}

// StoreWebWideTrackingException fn
func (*Navigator) StoreWebWideTrackingException(args *storeexceptionsinformation.StoreExceptionsInformation) {
	js.Rewrite("$<.storeWebWideTrackingException($1)", args)
}

// MsSaveBlob fn
func (*Navigator) MsSaveBlob(blob interface{}, defaultName *string) (b bool) {
	js.Rewrite("$<.msSaveBlob($1, $2)", blob, defaultName)
	return b
}

// MsSaveOrOpenBlob fn
func (*Navigator) MsSaveOrOpenBlob(blob interface{}, defaultName *string) (b bool) {
	js.Rewrite("$<.msSaveOrOpenBlob($1, $2)", blob, defaultName)
	return b
}

// SendBeacon fn
func (*Navigator) SendBeacon(url string, data *interface{}) (b bool) {
	js.Rewrite("$<.sendBeacon($1, $2)", url, data)
	return b
}

// GetUserMedia fn
func (*Navigator) GetUserMedia(constraints *mediastreamconstraints.MediaStreamConstraints, successCallback func(stream *MediaStream), errorCallback func(err *mediastreamerror.MediaStreamError)) {
	js.Rewrite("$<.getUserMedia($1, $2, $3)", constraints, successCallback, errorCallback)
}

// GetGamepads fn
func (*Navigator) GetGamepads() (g []*gamepad.Gamepad) {
	js.Rewrite("$<.getGamepads()")
	return g
}

// JavaEnabled fn
func (*Navigator) JavaEnabled() (b bool) {
	js.Rewrite("$<.javaEnabled()")
	return b
}

// MsLaunchURI fn
func (*Navigator) MsLaunchURI(uri string, successCallback *func(), noHandlerCallback *func()) {
	js.Rewrite("$<.msLaunchUri($1, $2, $3)", uri, successCallback, noHandlerCallback)
}

// RequestMediaKeySystemAccess fn
func (*Navigator) RequestMediaKeySystemAccess(keySystem string, supportedConfigurations []*mediakeysystemconfiguration.MediaKeySystemConfiguration) (m *MediaKeySystemAccess) {
	js.Rewrite("await $<.requestMediaKeySystemAccess($1, $2)", keySystem, supportedConfigurations)
	return m
}

// AppCodeName prop
func (*Navigator) AppCodeName() (appCodeName string) {
	js.Rewrite("$<.appCodeName")
	return appCodeName
}

// AppName prop
func (*Navigator) AppName() (appName string) {
	js.Rewrite("$<.appName")
	return appName
}

// AppVersion prop
func (*Navigator) AppVersion() (appVersion string) {
	js.Rewrite("$<.appVersion")
	return appVersion
}

// Platform prop
func (*Navigator) Platform() (platform string) {
	js.Rewrite("$<.platform")
	return platform
}

// Product prop
func (*Navigator) Product() (product string) {
	js.Rewrite("$<.product")
	return product
}

// ProductSub prop
func (*Navigator) ProductSub() (productSub string) {
	js.Rewrite("$<.productSub")
	return productSub
}

// UserAgent prop
func (*Navigator) UserAgent() (userAgent string) {
	js.Rewrite("$<.userAgent")
	return userAgent
}

// Vendor prop
func (*Navigator) Vendor() (vendor string) {
	js.Rewrite("$<.vendor")
	return vendor
}

// VendorSub prop
func (*Navigator) VendorSub() (vendorSub string) {
	js.Rewrite("$<.vendorSub")
	return vendorSub
}

// OnLine prop
func (*Navigator) OnLine() (onLine bool) {
	js.Rewrite("$<.onLine")
	return onLine
}

// Geolocation prop
func (*Navigator) Geolocation() (geolocation *geolocation.Geolocation) {
	js.Rewrite("$<.geolocation")
	return geolocation
}

// HardwareConcurrency prop
func (*Navigator) HardwareConcurrency() (hardwareConcurrency uint64) {
	js.Rewrite("$<.hardwareConcurrency")
	return hardwareConcurrency
}

// MediaDevices prop
func (*Navigator) MediaDevices() (mediaDevices *MediaDevices) {
	js.Rewrite("$<.mediaDevices")
	return mediaDevices
}

// Authentication prop
func (*Navigator) Authentication() (authentication *webauthentication.WebAuthentication) {
	js.Rewrite("$<.authentication")
	return authentication
}

// CookieEnabled prop
func (*Navigator) CookieEnabled() (cookieEnabled bool) {
	js.Rewrite("$<.cookieEnabled")
	return cookieEnabled
}

// GamepadInputEmulation prop
func (*Navigator) GamepadInputEmulation() (gamepadInputEmulation *gamepadinputemulationtype.GamepadInputEmulationType) {
	js.Rewrite("$<.gamepadInputEmulation")
	return gamepadInputEmulation
}

// GamepadInputEmulation prop
func (*Navigator) SetGamepadInputEmulation(gamepadInputEmulation *gamepadinputemulationtype.GamepadInputEmulationType) {
	js.Rewrite("$<.gamepadInputEmulation = $1", gamepadInputEmulation)
}

// Language prop
func (*Navigator) Language() (language string) {
	js.Rewrite("$<.language")
	return language
}

// MaxTouchPoints prop
func (*Navigator) MaxTouchPoints() (maxTouchPoints int) {
	js.Rewrite("$<.maxTouchPoints")
	return maxTouchPoints
}

// MimeTypes prop
func (*Navigator) MimeTypes() (mimeTypes *mimetypearray.MimeTypeArray) {
	js.Rewrite("$<.mimeTypes")
	return mimeTypes
}

// MsManipulationViewsEnabled prop
func (*Navigator) MsManipulationViewsEnabled() (msManipulationViewsEnabled bool) {
	js.Rewrite("$<.msManipulationViewsEnabled")
	return msManipulationViewsEnabled
}

// MsMaxTouchPoints prop
func (*Navigator) MsMaxTouchPoints() (msMaxTouchPoints int) {
	js.Rewrite("$<.msMaxTouchPoints")
	return msMaxTouchPoints
}

// MsPointerEnabled prop
func (*Navigator) MsPointerEnabled() (msPointerEnabled bool) {
	js.Rewrite("$<.msPointerEnabled")
	return msPointerEnabled
}

// Plugins prop
func (*Navigator) Plugins() (plugins *pluginarray.PluginArray) {
	js.Rewrite("$<.plugins")
	return plugins
}

// PointerEnabled prop
func (*Navigator) PointerEnabled() (pointerEnabled bool) {
	js.Rewrite("$<.pointerEnabled")
	return pointerEnabled
}

// ServiceWorker prop
func (*Navigator) ServiceWorker() (serviceWorker *ServiceWorkerContainer) {
	js.Rewrite("$<.serviceWorker")
	return serviceWorker
}

// Webdriver prop
func (*Navigator) Webdriver() (webdriver bool) {
	js.Rewrite("$<.webdriver")
	return webdriver
}
