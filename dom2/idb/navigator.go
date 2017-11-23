package idb

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

// ConfirmSiteSpecificTrackingException
func (*Navigator) ConfirmSiteSpecificTrackingException(args *confirmsitespecificexceptionsinformation.ConfirmSiteSpecificExceptionsInformation) (b bool) {
	js.Rewrite("$<.ConfirmSiteSpecificTrackingException($1)", args)
	return b
}

// ConfirmWebWideTrackingException
func (*Navigator) ConfirmWebWideTrackingException(args *exceptioninformation.ExceptionInformation) (b bool) {
	js.Rewrite("$<.ConfirmWebWideTrackingException($1)", args)
	return b
}

// RemoveSiteSpecificTrackingException
func (*Navigator) RemoveSiteSpecificTrackingException(args *exceptioninformation.ExceptionInformation) {
	js.Rewrite("$<.RemoveSiteSpecificTrackingException($1)", args)
}

// RemoveWebWideTrackingException
func (*Navigator) RemoveWebWideTrackingException(args *exceptioninformation.ExceptionInformation) {
	js.Rewrite("$<.RemoveWebWideTrackingException($1)", args)
}

// StoreSiteSpecificTrackingException
func (*Navigator) StoreSiteSpecificTrackingException(args *storesitespecificexceptionsinformation.StoreSiteSpecificExceptionsInformation) {
	js.Rewrite("$<.StoreSiteSpecificTrackingException($1)", args)
}

// StoreWebWideTrackingException
func (*Navigator) StoreWebWideTrackingException(args *storeexceptionsinformation.StoreExceptionsInformation) {
	js.Rewrite("$<.StoreWebWideTrackingException($1)", args)
}

// MsSaveBlob
func (*Navigator) MsSaveBlob(blob interface{}, defaultName *string) (b bool) {
	js.Rewrite("$<.MsSaveBlob($1, $2)", blob, defaultName)
	return b
}

// MsSaveOrOpenBlob
func (*Navigator) MsSaveOrOpenBlob(blob interface{}, defaultName *string) (b bool) {
	js.Rewrite("$<.MsSaveOrOpenBlob($1, $2)", blob, defaultName)
	return b
}

// SendBeacon
func (*Navigator) SendBeacon(url string, data *interface{}) (b bool) {
	js.Rewrite("$<.SendBeacon($1, $2)", url, data)
	return b
}

// GetUserMedia
func (*Navigator) GetUserMedia(constraints *mediastreamconstraints.MediaStreamConstraints, successCallback func(stream *MediaStream), errorCallback func(err *mediastreamerror.MediaStreamError)) {
	js.Rewrite("$<.GetUserMedia($1, $2, $3)", constraints, successCallback, errorCallback)
}

// GetGamepads
func (*Navigator) GetGamepads() (g []*gamepad.Gamepad) {
	js.Rewrite("$<.GetGamepads()")
	return g
}

// JavaEnabled
func (*Navigator) JavaEnabled() (b bool) {
	js.Rewrite("$<.JavaEnabled()")
	return b
}

// MsLaunchURI
func (*Navigator) MsLaunchURI(uri string, successCallback *func(), noHandlerCallback *func()) {
	js.Rewrite("$<.MsLaunchURI($1, $2, $3)", uri, successCallback, noHandlerCallback)
}

// RequestMediaKeySystemAccess
func (*Navigator) RequestMediaKeySystemAccess(keySystem string, supportedConfigurations []*mediakeysystemconfiguration.MediaKeySystemConfiguration) (m *MediaKeySystemAccess) {
	js.Rewrite("await $<.RequestMediaKeySystemAccess($1, $2)", keySystem, supportedConfigurations)
	return m
}

// AppCodeName
func (*Navigator) AppCodeName() (appCodeName string) {
	js.Rewrite("$<.AppCodeName")
	return appCodeName
}

// AppName
func (*Navigator) AppName() (appName string) {
	js.Rewrite("$<.AppName")
	return appName
}

// AppVersion
func (*Navigator) AppVersion() (appVersion string) {
	js.Rewrite("$<.AppVersion")
	return appVersion
}

// Platform
func (*Navigator) Platform() (platform string) {
	js.Rewrite("$<.Platform")
	return platform
}

// Product
func (*Navigator) Product() (product string) {
	js.Rewrite("$<.Product")
	return product
}

// ProductSub
func (*Navigator) ProductSub() (productSub string) {
	js.Rewrite("$<.ProductSub")
	return productSub
}

// UserAgent
func (*Navigator) UserAgent() (userAgent string) {
	js.Rewrite("$<.UserAgent")
	return userAgent
}

// Vendor
func (*Navigator) Vendor() (vendor string) {
	js.Rewrite("$<.Vendor")
	return vendor
}

// VendorSub
func (*Navigator) VendorSub() (vendorSub string) {
	js.Rewrite("$<.VendorSub")
	return vendorSub
}

// OnLine
func (*Navigator) OnLine() (onLine bool) {
	js.Rewrite("$<.OnLine")
	return onLine
}

// Geolocation
func (*Navigator) Geolocation() (geolocation *geolocation.Geolocation) {
	js.Rewrite("$<.Geolocation")
	return geolocation
}

// HardwareConcurrency
func (*Navigator) HardwareConcurrency() (hardwareConcurrency uint64) {
	js.Rewrite("$<.HardwareConcurrency")
	return hardwareConcurrency
}

// MediaDevices
func (*Navigator) MediaDevices() (mediaDevices *MediaDevices) {
	js.Rewrite("$<.MediaDevices")
	return mediaDevices
}

// Authentication
func (*Navigator) Authentication() (authentication *webauthentication.WebAuthentication) {
	js.Rewrite("$<.Authentication")
	return authentication
}

// CookieEnabled
func (*Navigator) CookieEnabled() (cookieEnabled bool) {
	js.Rewrite("$<.CookieEnabled")
	return cookieEnabled
}

// GamepadInputEmulation
func (*Navigator) GamepadInputEmulation() (gamepadInputEmulation *gamepadinputemulationtype.GamepadInputEmulationType) {
	js.Rewrite("$<.GamepadInputEmulation")
	return gamepadInputEmulation
}

// GamepadInputEmulation
func (*Navigator) SetGamepadInputEmulation(gamepadInputEmulation *gamepadinputemulationtype.GamepadInputEmulationType) {
	js.Rewrite("$<.GamepadInputEmulation = $1", gamepadInputEmulation)
}

// Language
func (*Navigator) Language() (language string) {
	js.Rewrite("$<.Language")
	return language
}

// MaxTouchPoints
func (*Navigator) MaxTouchPoints() (maxTouchPoints int) {
	js.Rewrite("$<.MaxTouchPoints")
	return maxTouchPoints
}

// MimeTypes
func (*Navigator) MimeTypes() (mimeTypes *mimetypearray.MimeTypeArray) {
	js.Rewrite("$<.MimeTypes")
	return mimeTypes
}

// MsManipulationViewsEnabled
func (*Navigator) MsManipulationViewsEnabled() (msManipulationViewsEnabled bool) {
	js.Rewrite("$<.MsManipulationViewsEnabled")
	return msManipulationViewsEnabled
}

// MsMaxTouchPoints
func (*Navigator) MsMaxTouchPoints() (msMaxTouchPoints int) {
	js.Rewrite("$<.MsMaxTouchPoints")
	return msMaxTouchPoints
}

// MsPointerEnabled
func (*Navigator) MsPointerEnabled() (msPointerEnabled bool) {
	js.Rewrite("$<.MsPointerEnabled")
	return msPointerEnabled
}

// Plugins
func (*Navigator) Plugins() (plugins *pluginarray.PluginArray) {
	js.Rewrite("$<.Plugins")
	return plugins
}

// PointerEnabled
func (*Navigator) PointerEnabled() (pointerEnabled bool) {
	js.Rewrite("$<.PointerEnabled")
	return pointerEnabled
}

// ServiceWorker
func (*Navigator) ServiceWorker() (serviceWorker *ServiceWorkerContainer) {
	js.Rewrite("$<.ServiceWorker")
	return serviceWorker
}

// Webdriver
func (*Navigator) Webdriver() (webdriver bool) {
	js.Rewrite("$<.Webdriver")
	return webdriver
}
