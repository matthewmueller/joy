package navigator

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/gamepad"
	"github.com/matthewmueller/golly/internal/gendom/dom/gamepadinputemulationtype"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediakeysystemaccess"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediakeysystemconfiguration"
	"github.com/matthewmueller/golly/internal/gendom/dom/mimetypearray"
	"github.com/matthewmueller/golly/internal/gendom/dom/msfilesaver"
	"github.com/matthewmueller/golly/internal/gendom/dom/msnavigatordonottrack"
	"github.com/matthewmueller/golly/internal/gendom/dom/navigatorbeacon"
	"github.com/matthewmueller/golly/internal/gendom/dom/navigatorconcurrenthardware"
	"github.com/matthewmueller/golly/internal/gendom/dom/navigatorcontentutils"
	"github.com/matthewmueller/golly/internal/gendom/dom/navigatorgeolocation"
	"github.com/matthewmueller/golly/internal/gendom/dom/navigatorid"
	"github.com/matthewmueller/golly/internal/gendom/dom/navigatoronline"
	"github.com/matthewmueller/golly/internal/gendom/dom/navigatorstorageutils"
	"github.com/matthewmueller/golly/internal/gendom/dom/navigatorusermedia"
	"github.com/matthewmueller/golly/internal/gendom/dom/pluginarray"
	"github.com/matthewmueller/golly/internal/gendom/dom/serviceworkercontainer"
	"github.com/matthewmueller/golly/internal/gendom/dom/webauthentication"
	"github.com/matthewmueller/golly/js"
)

type Navigator struct {
	*navigatorid.NavigatorID
	*navigatoronline.NavigatorOnLine
	*navigatorcontentutils.NavigatorContentUtils
	*navigatorstorageutils.NavigatorStorageUtils
	*navigatorgeolocation.NavigatorGeolocation
	*msnavigatordonottrack.MSNavigatorDoNotTrack
	*msfilesaver.MSFileSaver
	*navigatorbeacon.NavigatorBeacon
	*navigatorconcurrenthardware.NavigatorConcurrentHardware
	*navigatorusermedia.NavigatorUserMedia
}

func (*Navigator) GetGamepads() (g []*gamepad.Gamepad) {
	js.Rewrite("$<.getGamepads()")
	return g
}

func (*Navigator) JavaEnabled() (b bool) {
	js.Rewrite("$<.javaEnabled()")
	return b
}

func (*Navigator) MsLaunchURI(uri string, successCallback func(), noHandlerCallback func()) {
	js.Rewrite("$<.msLaunchUri($1, $2, $3)", uri, successCallback, noHandlerCallback)
}

func (*Navigator) RequestMediaKeySystemAccess(keySystem string, supportedConfigurations []*mediakeysystemconfiguration.MediaKeySystemConfiguration) (m *mediakeysystemaccess.MediaKeySystemAccess) {
	js.Rewrite("await $<.requestMediaKeySystemAccess($1, $2)", keySystem, supportedConfigurations)
	return m
}

func (*Navigator) GetAuthentication() (authentication *webauthentication.WebAuthentication) {
	js.Rewrite("$<.authentication")
	return authentication
}

func (*Navigator) GetCookieEnabled() (cookieEnabled bool) {
	js.Rewrite("$<.cookieEnabled")
	return cookieEnabled
}

func (*Navigator) GetGamepadInputEmulation() (gamepadInputEmulation *gamepadinputemulationtype.GamepadInputEmulationType) {
	js.Rewrite("$<.gamepadInputEmulation")
	return gamepadInputEmulation
}

func (*Navigator) SetGamepadInputEmulation(gamepadInputEmulation *gamepadinputemulationtype.GamepadInputEmulationType) {
	js.Rewrite("$<.gamepadInputEmulation = $1", gamepadInputEmulation)
}

func (*Navigator) GetLanguage() (language string) {
	js.Rewrite("$<.language")
	return language
}

func (*Navigator) GetMaxTouchPoints() (maxTouchPoints int) {
	js.Rewrite("$<.maxTouchPoints")
	return maxTouchPoints
}

func (*Navigator) GetMimeTypes() (mimeTypes *mimetypearray.MimeTypeArray) {
	js.Rewrite("$<.mimeTypes")
	return mimeTypes
}

func (*Navigator) GetMsManipulationViewsEnabled() (msManipulationViewsEnabled bool) {
	js.Rewrite("$<.msManipulationViewsEnabled")
	return msManipulationViewsEnabled
}

func (*Navigator) GetMsMaxTouchPoints() (msMaxTouchPoints int) {
	js.Rewrite("$<.msMaxTouchPoints")
	return msMaxTouchPoints
}

func (*Navigator) GetMsPointerEnabled() (msPointerEnabled bool) {
	js.Rewrite("$<.msPointerEnabled")
	return msPointerEnabled
}

func (*Navigator) GetPlugins() (plugins *pluginarray.PluginArray) {
	js.Rewrite("$<.plugins")
	return plugins
}

func (*Navigator) GetPointerEnabled() (pointerEnabled bool) {
	js.Rewrite("$<.pointerEnabled")
	return pointerEnabled
}

func (*Navigator) GetServiceWorker() (serviceWorker *serviceworkercontainer.ServiceWorkerContainer) {
	js.Rewrite("$<.serviceWorker")
	return serviceWorker
}

func (*Navigator) GetWebdriver() (webdriver bool) {
	js.Rewrite("$<.webdriver")
	return webdriver
}
