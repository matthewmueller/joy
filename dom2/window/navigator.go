package window

import (
	"github.com/matthewmueller/golly/dom2/gamepad"
	"github.com/matthewmueller/golly/dom2/gamepadinputemulationtype"
	"github.com/matthewmueller/golly/dom2/mediakeysystemconfiguration"
	"github.com/matthewmueller/golly/dom2/msfilesaver"
	"github.com/matthewmueller/golly/dom2/msnavigatordonottrack"
	"github.com/matthewmueller/golly/dom2/navigatorconcurrenthardware"
	"github.com/matthewmueller/golly/dom2/navigatorcontentutils"
	"github.com/matthewmueller/golly/dom2/navigatorgeolocation"
	"github.com/matthewmueller/golly/dom2/navigatorid"
	"github.com/matthewmueller/golly/dom2/navigatoronline"
	"github.com/matthewmueller/golly/dom2/navigatorstorageutils"
	"github.com/matthewmueller/golly/dom2/pluginarray"
	"github.com/matthewmueller/golly/dom2/webauthentication"
	"github.com/matthewmueller/golly/js"
)

// js:"Navigator,omit"
type Navigator struct {
	navigatorid.NavigatorID
	navigatoronline.NavigatorOnLine
	navigatorcontentutils.NavigatorContentUtils
	navigatorstorageutils.NavigatorStorageUtils
	navigatorgeolocation.NavigatorGeolocation
	msnavigatordonottrack.MSNavigatorDoNotTrack
	msfilesaver.MSFileSaver
	navigatorbeacon.NavigatorBeacon
	navigatorconcurrenthardware.NavigatorConcurrentHardware
	NavigatorUserMedia
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
