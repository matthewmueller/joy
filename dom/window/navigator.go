package window

import (
	"github.com/matthewmueller/joy/dom/confirmsitespecificexceptionsinformation"
	"github.com/matthewmueller/joy/dom/exceptioninformation"
	"github.com/matthewmueller/joy/dom/gamepad"
	"github.com/matthewmueller/joy/dom/gamepadinputemulationtype"
	"github.com/matthewmueller/joy/dom/geolocation"
	"github.com/matthewmueller/joy/dom/mediakeysystemconfiguration"
	"github.com/matthewmueller/joy/dom/mediastreamconstraints"
	"github.com/matthewmueller/joy/dom/mediastreamerror"
	"github.com/matthewmueller/joy/dom/mimetypearray"
	"github.com/matthewmueller/joy/dom/pluginarray"
	"github.com/matthewmueller/joy/dom/storeexceptionsinformation"
	"github.com/matthewmueller/joy/dom/storesitespecificexceptionsinformation"
	"github.com/matthewmueller/joy/dom/webauthentication"
	"github.com/matthewmueller/joy/macro"
)

// Navigator struct
// js:"Navigator,omit"
type Navigator struct {
}

// GetGamepads fn
// js:"getGamepads"
func (*Navigator) GetGamepads() (g []*gamepad.Gamepad) {
	macro.Rewrite("$_.getGamepads()")
	return g
}

// JavaEnabled fn
// js:"javaEnabled"
func (*Navigator) JavaEnabled() (b bool) {
	macro.Rewrite("$_.javaEnabled()")
	return b
}

// MsLaunchURI fn
// js:"msLaunchUri"
func (*Navigator) MsLaunchURI(uri string, successCallback *func(), noHandlerCallback *func()) {
	macro.Rewrite("$_.msLaunchUri($1, $2, $3)", uri, successCallback, noHandlerCallback)
}

// RequestMediaKeySystemAccess fn
// js:"requestMediaKeySystemAccess"
func (*Navigator) RequestMediaKeySystemAccess(keySystem string, supportedConfigurations []*mediakeysystemconfiguration.MediaKeySystemConfiguration) (m *MediaKeySystemAccess) {
	macro.Rewrite("await $_.requestMediaKeySystemAccess($1, $2)", keySystem, supportedConfigurations)
	return m
}

// ConfirmSiteSpecificTrackingException fn
// js:"confirmSiteSpecificTrackingException"
func (*Navigator) ConfirmSiteSpecificTrackingException(args *confirmsitespecificexceptionsinformation.ConfirmSiteSpecificExceptionsInformation) (b bool) {
	macro.Rewrite("$_.confirmSiteSpecificTrackingException($1)", args)
	return b
}

// ConfirmWebWideTrackingException fn
// js:"confirmWebWideTrackingException"
func (*Navigator) ConfirmWebWideTrackingException(args *exceptioninformation.ExceptionInformation) (b bool) {
	macro.Rewrite("$_.confirmWebWideTrackingException($1)", args)
	return b
}

// RemoveSiteSpecificTrackingException fn
// js:"removeSiteSpecificTrackingException"
func (*Navigator) RemoveSiteSpecificTrackingException(args *exceptioninformation.ExceptionInformation) {
	macro.Rewrite("$_.removeSiteSpecificTrackingException($1)", args)
}

// RemoveWebWideTrackingException fn
// js:"removeWebWideTrackingException"
func (*Navigator) RemoveWebWideTrackingException(args *exceptioninformation.ExceptionInformation) {
	macro.Rewrite("$_.removeWebWideTrackingException($1)", args)
}

// StoreSiteSpecificTrackingException fn
// js:"storeSiteSpecificTrackingException"
func (*Navigator) StoreSiteSpecificTrackingException(args *storesitespecificexceptionsinformation.StoreSiteSpecificExceptionsInformation) {
	macro.Rewrite("$_.storeSiteSpecificTrackingException($1)", args)
}

// StoreWebWideTrackingException fn
// js:"storeWebWideTrackingException"
func (*Navigator) StoreWebWideTrackingException(args *storeexceptionsinformation.StoreExceptionsInformation) {
	macro.Rewrite("$_.storeWebWideTrackingException($1)", args)
}

// MsSaveBlob fn
// js:"msSaveBlob"
func (*Navigator) MsSaveBlob(blob interface{}, defaultName *string) (b bool) {
	macro.Rewrite("$_.msSaveBlob($1, $2)", blob, defaultName)
	return b
}

// MsSaveOrOpenBlob fn
// js:"msSaveOrOpenBlob"
func (*Navigator) MsSaveOrOpenBlob(blob interface{}, defaultName *string) (b bool) {
	macro.Rewrite("$_.msSaveOrOpenBlob($1, $2)", blob, defaultName)
	return b
}

// SendBeacon fn
// js:"sendBeacon"
func (*Navigator) SendBeacon(url string, data *interface{}) (b bool) {
	macro.Rewrite("$_.sendBeacon($1, $2)", url, data)
	return b
}

// GetUserMedia fn
// js:"getUserMedia"
func (*Navigator) GetUserMedia(constraints *mediastreamconstraints.MediaStreamConstraints, successCallback func(stream *MediaStream), errorCallback func(err *mediastreamerror.MediaStreamError)) {
	macro.Rewrite("$_.getUserMedia($1, $2, $3)", constraints, successCallback, errorCallback)
}

// Authentication prop
// js:"authentication"
func (*Navigator) Authentication() (authentication *webauthentication.WebAuthentication) {
	macro.Rewrite("$_.authentication")
	return authentication
}

// CookieEnabled prop
// js:"cookieEnabled"
func (*Navigator) CookieEnabled() (cookieEnabled bool) {
	macro.Rewrite("$_.cookieEnabled")
	return cookieEnabled
}

// GamepadInputEmulation prop
// js:"gamepadInputEmulation"
func (*Navigator) GamepadInputEmulation() (gamepadInputEmulation *gamepadinputemulationtype.GamepadInputEmulationType) {
	macro.Rewrite("$_.gamepadInputEmulation")
	return gamepadInputEmulation
}

// SetGamepadInputEmulation prop
// js:"gamepadInputEmulation"
func (*Navigator) SetGamepadInputEmulation(gamepadInputEmulation *gamepadinputemulationtype.GamepadInputEmulationType) {
	macro.Rewrite("$_.gamepadInputEmulation = $1", gamepadInputEmulation)
}

// Language prop
// js:"language"
func (*Navigator) Language() (language string) {
	macro.Rewrite("$_.language")
	return language
}

// MaxTouchPoints prop
// js:"maxTouchPoints"
func (*Navigator) MaxTouchPoints() (maxTouchPoints int) {
	macro.Rewrite("$_.maxTouchPoints")
	return maxTouchPoints
}

// MimeTypes prop
// js:"mimeTypes"
func (*Navigator) MimeTypes() (mimeTypes *mimetypearray.MimeTypeArray) {
	macro.Rewrite("$_.mimeTypes")
	return mimeTypes
}

// MsManipulationViewsEnabled prop
// js:"msManipulationViewsEnabled"
func (*Navigator) MsManipulationViewsEnabled() (msManipulationViewsEnabled bool) {
	macro.Rewrite("$_.msManipulationViewsEnabled")
	return msManipulationViewsEnabled
}

// MsMaxTouchPoints prop
// js:"msMaxTouchPoints"
func (*Navigator) MsMaxTouchPoints() (msMaxTouchPoints int) {
	macro.Rewrite("$_.msMaxTouchPoints")
	return msMaxTouchPoints
}

// MsPointerEnabled prop
// js:"msPointerEnabled"
func (*Navigator) MsPointerEnabled() (msPointerEnabled bool) {
	macro.Rewrite("$_.msPointerEnabled")
	return msPointerEnabled
}

// Plugins prop
// js:"plugins"
func (*Navigator) Plugins() (plugins *pluginarray.PluginArray) {
	macro.Rewrite("$_.plugins")
	return plugins
}

// PointerEnabled prop
// js:"pointerEnabled"
func (*Navigator) PointerEnabled() (pointerEnabled bool) {
	macro.Rewrite("$_.pointerEnabled")
	return pointerEnabled
}

// ServiceWorker prop
// js:"serviceWorker"
func (*Navigator) ServiceWorker() (serviceWorker *ServiceWorkerContainer) {
	macro.Rewrite("$_.serviceWorker")
	return serviceWorker
}

// Webdriver prop
// js:"webdriver"
func (*Navigator) Webdriver() (webdriver bool) {
	macro.Rewrite("$_.webdriver")
	return webdriver
}

// AppCodeName prop
// js:"appCodeName"
func (*Navigator) AppCodeName() (appCodeName string) {
	macro.Rewrite("$_.appCodeName")
	return appCodeName
}

// AppName prop
// js:"appName"
func (*Navigator) AppName() (appName string) {
	macro.Rewrite("$_.appName")
	return appName
}

// AppVersion prop
// js:"appVersion"
func (*Navigator) AppVersion() (appVersion string) {
	macro.Rewrite("$_.appVersion")
	return appVersion
}

// Platform prop
// js:"platform"
func (*Navigator) Platform() (platform string) {
	macro.Rewrite("$_.platform")
	return platform
}

// Product prop
// js:"product"
func (*Navigator) Product() (product string) {
	macro.Rewrite("$_.product")
	return product
}

// ProductSub prop
// js:"productSub"
func (*Navigator) ProductSub() (productSub string) {
	macro.Rewrite("$_.productSub")
	return productSub
}

// UserAgent prop
// js:"userAgent"
func (*Navigator) UserAgent() (userAgent string) {
	macro.Rewrite("$_.userAgent")
	return userAgent
}

// Vendor prop
// js:"vendor"
func (*Navigator) Vendor() (vendor string) {
	macro.Rewrite("$_.vendor")
	return vendor
}

// VendorSub prop
// js:"vendorSub"
func (*Navigator) VendorSub() (vendorSub string) {
	macro.Rewrite("$_.vendorSub")
	return vendorSub
}

// OnLine prop
// js:"onLine"
func (*Navigator) OnLine() (onLine bool) {
	macro.Rewrite("$_.onLine")
	return onLine
}

// Geolocation prop
// js:"geolocation"
func (*Navigator) Geolocation() (geolocation *geolocation.Geolocation) {
	macro.Rewrite("$_.geolocation")
	return geolocation
}

// HardwareConcurrency prop
// js:"hardwareConcurrency"
func (*Navigator) HardwareConcurrency() (hardwareConcurrency uint64) {
	macro.Rewrite("$_.hardwareConcurrency")
	return hardwareConcurrency
}

// MediaDevices prop
// js:"mediaDevices"
func (*Navigator) MediaDevices() (mediaDevices *MediaDevices) {
	macro.Rewrite("$_.mediaDevices")
	return mediaDevices
}
