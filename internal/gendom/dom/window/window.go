package window

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/applicationcache"
	"github.com/matthewmueller/golly/internal/gendom/dom/barprop"
	"github.com/matthewmueller/golly/internal/gendom/dom/beforeunloadevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/cachestorage"
	"github.com/matthewmueller/golly/internal/gendom/dom/crypto"
	"github.com/matthewmueller/golly/internal/gendom/dom/cssrulelist"
	"github.com/matthewmueller/golly/internal/gendom/dom/cssstyledeclaration"
	"github.com/matthewmueller/golly/internal/gendom/dom/devicelightevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/devicemotionevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/deviceorientationevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/document"
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/extensionscriptapis"
	"github.com/matthewmueller/golly/internal/gendom/dom/external"
	"github.com/matthewmueller/golly/internal/gendom/dom/focusevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/focusnavigationorigin"
	"github.com/matthewmueller/golly/internal/gendom/dom/globaleventhandlers"
	"github.com/matthewmueller/golly/internal/gendom/dom/globalfetch"
	"github.com/matthewmueller/golly/internal/gendom/dom/hashchangeevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/history"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbenvironment"
	"github.com/matthewmueller/golly/internal/gendom/dom/location"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediaquerylist"
	"github.com/matthewmueller/golly/internal/gendom/dom/messageevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/mscredentials"
	"github.com/matthewmueller/golly/internal/gendom/dom/navigationreason"
	"github.com/matthewmueller/golly/internal/gendom/dom/navigator"
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/internal/gendom/dom/pagetransitionevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/performance"
	"github.com/matthewmueller/golly/internal/gendom/dom/popstateevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/screen"
	"github.com/matthewmueller/golly/internal/gendom/dom/selection"
	"github.com/matthewmueller/golly/internal/gendom/dom/speechsynthesis"
	"github.com/matthewmueller/golly/internal/gendom/dom/storageevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/stylemedia"
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/internal/gendom/dom/webkitpoint"
	"github.com/matthewmueller/golly/internal/gendom/dom/windowbase64"
	"github.com/matthewmueller/golly/internal/gendom/dom/windowconsole"
	"github.com/matthewmueller/golly/internal/gendom/dom/windowlocalstorage"
	"github.com/matthewmueller/golly/internal/gendom/dom/windowsessionstorage"
	"github.com/matthewmueller/golly/internal/gendom/dom/windowtimers"
	"github.com/matthewmueller/golly/js"
)

type Window struct {
	*eventtarget.EventTarget
	*windowtimers.WindowTimers
	*windowsessionstorage.WindowSessionStorage
	*windowlocalstorage.WindowLocalStorage
	*windowconsole.WindowConsole
	*globaleventhandlers.GlobalEventHandlers
	*idbenvironment.IDBEnvironment
	*windowbase64.WindowBase64
	*globalfetch.GlobalFetch
}

func (*Window) Alert(message string) {
	js.Rewrite("$<.alert($1)", message)
}

func (*Window) Blur() {
	js.Rewrite("$<.blur()")
}

func (*Window) CancelAnimationFrame(handle int) {
	js.Rewrite("$<.cancelAnimationFrame($1)", handle)
}

func (*Window) CaptureEvents() {
	js.Rewrite("$<.captureEvents()")
}

func (*Window) Close() {
	js.Rewrite("$<.close()")
}

func (*Window) Confirm(message string) (b bool) {
	js.Rewrite("$<.confirm($1)", message)
	return b
}

func (*Window) DepartFocus(navigationReason *navigationreason.NavigationReason, origin *focusnavigationorigin.FocusNavigationOrigin) {
	js.Rewrite("$<.departFocus($1, $2)", navigationReason, origin)
}

func (*Window) Focus() {
	js.Rewrite("$<.focus()")
}

func (*Window) GetComputedStyle(elt *element.Element, pseudoElt string) (c *cssstyledeclaration.CSSStyleDeclaration) {
	js.Rewrite("$<.getComputedStyle($1, $2)", elt, pseudoElt)
	return c
}

func (*Window) GetMatchedCSSRules(elt *element.Element, pseudoElt string) (c *cssrulelist.CSSRuleList) {
	js.Rewrite("$<.getMatchedCSSRules($1, $2)", elt, pseudoElt)
	return c
}

func (*Window) GetSelection() (s *selection.Selection) {
	js.Rewrite("$<.getSelection()")
	return s
}

func (*Window) MatchMedia(mediaQuery string) (m *mediaquerylist.MediaQueryList) {
	js.Rewrite("$<.matchMedia($1)", mediaQuery)
	return m
}

func (*Window) MoveBy(x int, y int) {
	js.Rewrite("$<.moveBy($1, $2)", x, y)
}

func (*Window) MoveTo(x int, y int) {
	js.Rewrite("$<.moveTo($1, $2)", x, y)
}

func (*Window) MsWriteProfilerMark(profilerMarkName string) {
	js.Rewrite("$<.msWriteProfilerMark($1)", profilerMarkName)
}

func (*Window) Open(url string, target string, features string, replace bool) (w *Window) {
	js.Rewrite("$<.open($1, $2, $3, $4)", url, target, features, replace)
	return w
}

func (*Window) PostMessage(message interface{}, targetOrigin string, transfer []interface{}) {
	js.Rewrite("$<.postMessage($1, $2, $3)", message, targetOrigin, transfer)
}

func (*Window) Print() {
	js.Rewrite("$<.print()")
}

func (*Window) Prompt(message string, def string) (s string) {
	js.Rewrite("$<.prompt($1, $2)", message, def)
	return s
}

func (*Window) ReleaseEvents() {
	js.Rewrite("$<.releaseEvents()")
}

func (*Window) RequestAnimationFrame(callback func(time int)) (i int) {
	js.Rewrite("$<.requestAnimationFrame($1)", callback)
	return i
}

func (*Window) ResizeBy(x int, y int) {
	js.Rewrite("$<.resizeBy($1, $2)", x, y)
}

func (*Window) ResizeTo(x int, y int) {
	js.Rewrite("$<.resizeTo($1, $2)", x, y)
}

func (*Window) Scroll(x int, y int) {
	js.Rewrite("$<.scroll($1, $2)", x, y)
}

func (*Window) ScrollBy(x int, y int) {
	js.Rewrite("$<.scrollBy($1, $2)", x, y)
}

func (*Window) ScrollTo(x int, y int) {
	js.Rewrite("$<.scrollTo($1, $2)", x, y)
}

func (*Window) Stop() {
	js.Rewrite("$<.stop()")
}

func (*Window) WebkitCancelAnimationFrame(handle int) {
	js.Rewrite("$<.webkitCancelAnimationFrame($1)", handle)
}

func (*Window) WebkitConvertPointFromNodeToPage(node *node.Node, pt *webkitpoint.WebKitPoint) (w *webkitpoint.WebKitPoint) {
	js.Rewrite("$<.webkitConvertPointFromNodeToPage($1, $2)", node, pt)
	return w
}

func (*Window) WebkitConvertPointFromPageToNode(node *node.Node, pt *webkitpoint.WebKitPoint) (w *webkitpoint.WebKitPoint) {
	js.Rewrite("$<.webkitConvertPointFromPageToNode($1, $2)", node, pt)
	return w
}

func (*Window) WebkitRequestAnimationFrame(callback func(time int)) (i int) {
	js.Rewrite("$<.webkitRequestAnimationFrame($1)", callback)
	return i
}

func (*Window) GetApplicationCache() (applicationCache *applicationcache.ApplicationCache) {
	js.Rewrite("$<.applicationCache")
	return applicationCache
}

func (*Window) GetCaches() (caches *cachestorage.CacheStorage) {
	js.Rewrite("$<.caches")
	return caches
}

func (*Window) GetClientInformation() (clientInformation *navigator.Navigator) {
	js.Rewrite("$<.clientInformation")
	return clientInformation
}

func (*Window) GetClosed() (closed bool) {
	js.Rewrite("$<.closed")
	return closed
}

func (*Window) GetCrypto() (crypto *crypto.Crypto) {
	js.Rewrite("$<.crypto")
	return crypto
}

func (*Window) GetDefaultStatus() (defaultStatus string) {
	js.Rewrite("$<.defaultStatus")
	return defaultStatus
}

func (*Window) SetDefaultStatus(defaultStatus string) {
	js.Rewrite("$<.defaultStatus = $1", defaultStatus)
}

func (*Window) GetDevicePixelRatio() (devicePixelRatio float32) {
	js.Rewrite("$<.devicePixelRatio")
	return devicePixelRatio
}

func (*Window) GetDocument() (document *document.Document) {
	js.Rewrite("$<.document")
	return document
}

func (*Window) GetDoNotTrack() (doNotTrack string) {
	js.Rewrite("$<.doNotTrack")
	return doNotTrack
}

func (*Window) GetEvent() (event *event.Event) {
	js.Rewrite("$<.event")
	return event
}

func (*Window) SetEvent(event *event.Event) {
	js.Rewrite("$<.event = $1", event)
}

func (*Window) GetExternal() (external *external.External) {
	js.Rewrite("$<.external")
	return external
}

func (*Window) GetFrameElement() (frameElement *element.Element) {
	js.Rewrite("$<.frameElement")
	return frameElement
}

func (*Window) GetFrames() (frames *Window) {
	js.Rewrite("$<.frames")
	return frames
}

func (*Window) GetHistory() (history *history.History) {
	js.Rewrite("$<.history")
	return history
}

func (*Window) GetInnerHeight() (innerHeight int) {
	js.Rewrite("$<.innerHeight")
	return innerHeight
}

func (*Window) GetInnerWidth() (innerWidth int) {
	js.Rewrite("$<.innerWidth")
	return innerWidth
}

func (*Window) GetIsSecureContext() (isSecureContext bool) {
	js.Rewrite("$<.isSecureContext")
	return isSecureContext
}

func (*Window) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*Window) GetLocation() (location *location.Location) {
	js.Rewrite("$<.location")
	return location
}

func (*Window) GetLocationbar() (locationbar *barprop.BarProp) {
	js.Rewrite("$<.locationbar")
	return locationbar
}

func (*Window) GetMenubar() (menubar *barprop.BarProp) {
	js.Rewrite("$<.menubar")
	return menubar
}

func (*Window) GetMsContentScript() (msContentScript *extensionscriptapis.ExtensionScriptApis) {
	js.Rewrite("$<.msContentScript")
	return msContentScript
}

func (*Window) GetMsCredentials() (msCredentials *mscredentials.MSCredentials) {
	js.Rewrite("$<.msCredentials")
	return msCredentials
}

func (*Window) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*Window) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*Window) GetNavigator() (navigator *navigator.Navigator) {
	js.Rewrite("$<.navigator")
	return navigator
}

func (*Window) GetOffscreenBuffering() (offscreenBuffering interface{}) {
	js.Rewrite("$<.offscreenBuffering")
	return offscreenBuffering
}

func (*Window) SetOffscreenBuffering(offscreenBuffering interface{}) {
	js.Rewrite("$<.offscreenBuffering = $1", offscreenBuffering)
}

func (*Window) GetOnabort() (e *event.Event) {
	js.Rewrite("$<.onabort")
	return e
}

func (*Window) SetOnabort(e *event.Event) {
	js.Rewrite("$<.onabort = $1", e)
}

func (*Window) GetOnafterprint() (afterprint *event.Event) {
	js.Rewrite("$<.onafterprint")
	return afterprint
}

func (*Window) SetOnafterprint(afterprint *event.Event) {
	js.Rewrite("$<.onafterprint = $1", afterprint)
}

func (*Window) GetOnbeforeprint() (beforeprint *event.Event) {
	js.Rewrite("$<.onbeforeprint")
	return beforeprint
}

func (*Window) SetOnbeforeprint(beforeprint *event.Event) {
	js.Rewrite("$<.onbeforeprint = $1", beforeprint)
}

func (*Window) GetOnbeforeunload() (beforeunload *beforeunloadevent.BeforeUnloadEvent) {
	js.Rewrite("$<.onbeforeunload")
	return beforeunload
}

func (*Window) SetOnbeforeunload(beforeunload *beforeunloadevent.BeforeUnloadEvent) {
	js.Rewrite("$<.onbeforeunload = $1", beforeunload)
}

func (*Window) GetOnblur() (blur *focusevent.FocusEvent) {
	js.Rewrite("$<.onblur")
	return blur
}

func (*Window) SetOnblur(blur *focusevent.FocusEvent) {
	js.Rewrite("$<.onblur = $1", blur)
}

func (*Window) GetOncanplay() (e *event.Event) {
	js.Rewrite("$<.oncanplay")
	return e
}

func (*Window) SetOncanplay(e *event.Event) {
	js.Rewrite("$<.oncanplay = $1", e)
}

func (*Window) GetOncanplaythrough() (e *event.Event) {
	js.Rewrite("$<.oncanplaythrough")
	return e
}

func (*Window) SetOncanplaythrough(e *event.Event) {
	js.Rewrite("$<.oncanplaythrough = $1", e)
}

func (*Window) GetOnchange() (e *event.Event) {
	js.Rewrite("$<.onchange")
	return e
}

func (*Window) SetOnchange(e *event.Event) {
	js.Rewrite("$<.onchange = $1", e)
}

func (*Window) GetOnclick() (e *event.Event) {
	js.Rewrite("$<.onclick")
	return e
}

func (*Window) SetOnclick(e *event.Event) {
	js.Rewrite("$<.onclick = $1", e)
}

func (*Window) GetOncompassneedscalibration() (compassneedscalibration *event.Event) {
	js.Rewrite("$<.oncompassneedscalibration")
	return compassneedscalibration
}

func (*Window) SetOncompassneedscalibration(compassneedscalibration *event.Event) {
	js.Rewrite("$<.oncompassneedscalibration = $1", compassneedscalibration)
}

func (*Window) GetOncontextmenu() (e *event.Event) {
	js.Rewrite("$<.oncontextmenu")
	return e
}

func (*Window) SetOncontextmenu(e *event.Event) {
	js.Rewrite("$<.oncontextmenu = $1", e)
}

func (*Window) GetOndblclick() (e *event.Event) {
	js.Rewrite("$<.ondblclick")
	return e
}

func (*Window) SetOndblclick(e *event.Event) {
	js.Rewrite("$<.ondblclick = $1", e)
}

func (*Window) GetOndevicelight() (devicelight *devicelightevent.DeviceLightEvent) {
	js.Rewrite("$<.ondevicelight")
	return devicelight
}

func (*Window) SetOndevicelight(devicelight *devicelightevent.DeviceLightEvent) {
	js.Rewrite("$<.ondevicelight = $1", devicelight)
}

func (*Window) GetOndevicemotion() (devicemotion *devicemotionevent.DeviceMotionEvent) {
	js.Rewrite("$<.ondevicemotion")
	return devicemotion
}

func (*Window) SetOndevicemotion(devicemotion *devicemotionevent.DeviceMotionEvent) {
	js.Rewrite("$<.ondevicemotion = $1", devicemotion)
}

func (*Window) GetOndeviceorientation() (deviceorientation *deviceorientationevent.DeviceOrientationEvent) {
	js.Rewrite("$<.ondeviceorientation")
	return deviceorientation
}

func (*Window) SetOndeviceorientation(deviceorientation *deviceorientationevent.DeviceOrientationEvent) {
	js.Rewrite("$<.ondeviceorientation = $1", deviceorientation)
}

func (*Window) GetOndrag() (e *event.Event) {
	js.Rewrite("$<.ondrag")
	return e
}

func (*Window) SetOndrag(e *event.Event) {
	js.Rewrite("$<.ondrag = $1", e)
}

func (*Window) GetOndragend() (e *event.Event) {
	js.Rewrite("$<.ondragend")
	return e
}

func (*Window) SetOndragend(e *event.Event) {
	js.Rewrite("$<.ondragend = $1", e)
}

func (*Window) GetOndragenter() (e *event.Event) {
	js.Rewrite("$<.ondragenter")
	return e
}

func (*Window) SetOndragenter(e *event.Event) {
	js.Rewrite("$<.ondragenter = $1", e)
}

func (*Window) GetOndragleave() (e *event.Event) {
	js.Rewrite("$<.ondragleave")
	return e
}

func (*Window) SetOndragleave(e *event.Event) {
	js.Rewrite("$<.ondragleave = $1", e)
}

func (*Window) GetOndragover() (e *event.Event) {
	js.Rewrite("$<.ondragover")
	return e
}

func (*Window) SetOndragover(e *event.Event) {
	js.Rewrite("$<.ondragover = $1", e)
}

func (*Window) GetOndragstart() (e *event.Event) {
	js.Rewrite("$<.ondragstart")
	return e
}

func (*Window) SetOndragstart(e *event.Event) {
	js.Rewrite("$<.ondragstart = $1", e)
}

func (*Window) GetOndrop() (e *event.Event) {
	js.Rewrite("$<.ondrop")
	return e
}

func (*Window) SetOndrop(e *event.Event) {
	js.Rewrite("$<.ondrop = $1", e)
}

func (*Window) GetOndurationchange() (e *event.Event) {
	js.Rewrite("$<.ondurationchange")
	return e
}

func (*Window) SetOndurationchange(e *event.Event) {
	js.Rewrite("$<.ondurationchange = $1", e)
}

func (*Window) GetOnemptied() (e *event.Event) {
	js.Rewrite("$<.onemptied")
	return e
}

func (*Window) SetOnemptied(e *event.Event) {
	js.Rewrite("$<.onemptied = $1", e)
}

func (*Window) GetOnended() (e *event.Event) {
	js.Rewrite("$<.onended")
	return e
}

func (*Window) SetOnended(e *event.Event) {
	js.Rewrite("$<.onended = $1", e)
}

func (*Window) GetOnerror() (ErrorEventHandler func(columnNumber uint, event interface{}, fileno uint, source string)) {
	js.Rewrite("$<.onerror")
	return ErrorEventHandler
}

func (*Window) SetOnerror(ErrorEventHandler func(columnNumber uint, event interface{}, fileno uint, source string)) {
	js.Rewrite("$<.onerror = $1", ErrorEventHandler)
}

func (*Window) GetOnfocus() (focus *focusevent.FocusEvent) {
	js.Rewrite("$<.onfocus")
	return focus
}

func (*Window) SetOnfocus(focus *focusevent.FocusEvent) {
	js.Rewrite("$<.onfocus = $1", focus)
}

func (*Window) GetOnhashchange() (hashchange *hashchangeevent.HashChangeEvent) {
	js.Rewrite("$<.onhashchange")
	return hashchange
}

func (*Window) SetOnhashchange(hashchange *hashchangeevent.HashChangeEvent) {
	js.Rewrite("$<.onhashchange = $1", hashchange)
}

func (*Window) GetOninput() (e *event.Event) {
	js.Rewrite("$<.oninput")
	return e
}

func (*Window) SetOninput(e *event.Event) {
	js.Rewrite("$<.oninput = $1", e)
}

func (*Window) GetOninvalid() (e *event.Event) {
	js.Rewrite("$<.oninvalid")
	return e
}

func (*Window) SetOninvalid(e *event.Event) {
	js.Rewrite("$<.oninvalid = $1", e)
}

func (*Window) GetOnkeydown() (e *event.Event) {
	js.Rewrite("$<.onkeydown")
	return e
}

func (*Window) SetOnkeydown(e *event.Event) {
	js.Rewrite("$<.onkeydown = $1", e)
}

func (*Window) GetOnkeypress() (e *event.Event) {
	js.Rewrite("$<.onkeypress")
	return e
}

func (*Window) SetOnkeypress(e *event.Event) {
	js.Rewrite("$<.onkeypress = $1", e)
}

func (*Window) GetOnkeyup() (e *event.Event) {
	js.Rewrite("$<.onkeyup")
	return e
}

func (*Window) SetOnkeyup(e *event.Event) {
	js.Rewrite("$<.onkeyup = $1", e)
}

func (*Window) GetOnload() (load *event.Event) {
	js.Rewrite("$<.onload")
	return load
}

func (*Window) SetOnload(load *event.Event) {
	js.Rewrite("$<.onload = $1", load)
}

func (*Window) GetOnloadeddata() (e *event.Event) {
	js.Rewrite("$<.onloadeddata")
	return e
}

func (*Window) SetOnloadeddata(e *event.Event) {
	js.Rewrite("$<.onloadeddata = $1", e)
}

func (*Window) GetOnloadedmetadata() (e *event.Event) {
	js.Rewrite("$<.onloadedmetadata")
	return e
}

func (*Window) SetOnloadedmetadata(e *event.Event) {
	js.Rewrite("$<.onloadedmetadata = $1", e)
}

func (*Window) GetOnloadstart() (e *event.Event) {
	js.Rewrite("$<.onloadstart")
	return e
}

func (*Window) SetOnloadstart(e *event.Event) {
	js.Rewrite("$<.onloadstart = $1", e)
}

func (*Window) GetOnmessage() (message *messageevent.MessageEvent) {
	js.Rewrite("$<.onmessage")
	return message
}

func (*Window) SetOnmessage(message *messageevent.MessageEvent) {
	js.Rewrite("$<.onmessage = $1", message)
}

func (*Window) GetOnmousedown() (e *event.Event) {
	js.Rewrite("$<.onmousedown")
	return e
}

func (*Window) SetOnmousedown(e *event.Event) {
	js.Rewrite("$<.onmousedown = $1", e)
}

func (*Window) GetOnmouseenter() (e *event.Event) {
	js.Rewrite("$<.onmouseenter")
	return e
}

func (*Window) SetOnmouseenter(e *event.Event) {
	js.Rewrite("$<.onmouseenter = $1", e)
}

func (*Window) GetOnmouseleave() (e *event.Event) {
	js.Rewrite("$<.onmouseleave")
	return e
}

func (*Window) SetOnmouseleave(e *event.Event) {
	js.Rewrite("$<.onmouseleave = $1", e)
}

func (*Window) GetOnmousemove() (e *event.Event) {
	js.Rewrite("$<.onmousemove")
	return e
}

func (*Window) SetOnmousemove(e *event.Event) {
	js.Rewrite("$<.onmousemove = $1", e)
}

func (*Window) GetOnmouseout() (e *event.Event) {
	js.Rewrite("$<.onmouseout")
	return e
}

func (*Window) SetOnmouseout(e *event.Event) {
	js.Rewrite("$<.onmouseout = $1", e)
}

func (*Window) GetOnmouseover() (e *event.Event) {
	js.Rewrite("$<.onmouseover")
	return e
}

func (*Window) SetOnmouseover(e *event.Event) {
	js.Rewrite("$<.onmouseover = $1", e)
}

func (*Window) GetOnmouseup() (e *event.Event) {
	js.Rewrite("$<.onmouseup")
	return e
}

func (*Window) SetOnmouseup(e *event.Event) {
	js.Rewrite("$<.onmouseup = $1", e)
}

func (*Window) GetOnmousewheel() (e *event.Event) {
	js.Rewrite("$<.onmousewheel")
	return e
}

func (*Window) SetOnmousewheel(e *event.Event) {
	js.Rewrite("$<.onmousewheel = $1", e)
}

func (*Window) GetOnmsgesturechange() (e *event.Event) {
	js.Rewrite("$<.onmsgesturechange")
	return e
}

func (*Window) SetOnmsgesturechange(e *event.Event) {
	js.Rewrite("$<.onmsgesturechange = $1", e)
}

func (*Window) GetOnmsgesturedoubletap() (e *event.Event) {
	js.Rewrite("$<.onmsgesturedoubletap")
	return e
}

func (*Window) SetOnmsgesturedoubletap(e *event.Event) {
	js.Rewrite("$<.onmsgesturedoubletap = $1", e)
}

func (*Window) GetOnmsgestureend() (e *event.Event) {
	js.Rewrite("$<.onmsgestureend")
	return e
}

func (*Window) SetOnmsgestureend(e *event.Event) {
	js.Rewrite("$<.onmsgestureend = $1", e)
}

func (*Window) GetOnmsgesturehold() (e *event.Event) {
	js.Rewrite("$<.onmsgesturehold")
	return e
}

func (*Window) SetOnmsgesturehold(e *event.Event) {
	js.Rewrite("$<.onmsgesturehold = $1", e)
}

func (*Window) GetOnmsgesturestart() (e *event.Event) {
	js.Rewrite("$<.onmsgesturestart")
	return e
}

func (*Window) SetOnmsgesturestart(e *event.Event) {
	js.Rewrite("$<.onmsgesturestart = $1", e)
}

func (*Window) GetOnmsgesturetap() (e *event.Event) {
	js.Rewrite("$<.onmsgesturetap")
	return e
}

func (*Window) SetOnmsgesturetap(e *event.Event) {
	js.Rewrite("$<.onmsgesturetap = $1", e)
}

func (*Window) GetOnmsinertiastart() (e *event.Event) {
	js.Rewrite("$<.onmsinertiastart")
	return e
}

func (*Window) SetOnmsinertiastart(e *event.Event) {
	js.Rewrite("$<.onmsinertiastart = $1", e)
}

func (*Window) GetOnmspointercancel() (e *event.Event) {
	js.Rewrite("$<.onmspointercancel")
	return e
}

func (*Window) SetOnmspointercancel(e *event.Event) {
	js.Rewrite("$<.onmspointercancel = $1", e)
}

func (*Window) GetOnmspointerdown() (e *event.Event) {
	js.Rewrite("$<.onmspointerdown")
	return e
}

func (*Window) SetOnmspointerdown(e *event.Event) {
	js.Rewrite("$<.onmspointerdown = $1", e)
}

func (*Window) GetOnmspointerenter() (e *event.Event) {
	js.Rewrite("$<.onmspointerenter")
	return e
}

func (*Window) SetOnmspointerenter(e *event.Event) {
	js.Rewrite("$<.onmspointerenter = $1", e)
}

func (*Window) GetOnmspointerleave() (e *event.Event) {
	js.Rewrite("$<.onmspointerleave")
	return e
}

func (*Window) SetOnmspointerleave(e *event.Event) {
	js.Rewrite("$<.onmspointerleave = $1", e)
}

func (*Window) GetOnmspointermove() (e *event.Event) {
	js.Rewrite("$<.onmspointermove")
	return e
}

func (*Window) SetOnmspointermove(e *event.Event) {
	js.Rewrite("$<.onmspointermove = $1", e)
}

func (*Window) GetOnmspointerout() (e *event.Event) {
	js.Rewrite("$<.onmspointerout")
	return e
}

func (*Window) SetOnmspointerout(e *event.Event) {
	js.Rewrite("$<.onmspointerout = $1", e)
}

func (*Window) GetOnmspointerover() (e *event.Event) {
	js.Rewrite("$<.onmspointerover")
	return e
}

func (*Window) SetOnmspointerover(e *event.Event) {
	js.Rewrite("$<.onmspointerover = $1", e)
}

func (*Window) GetOnmspointerup() (e *event.Event) {
	js.Rewrite("$<.onmspointerup")
	return e
}

func (*Window) SetOnmspointerup(e *event.Event) {
	js.Rewrite("$<.onmspointerup = $1", e)
}

func (*Window) GetOnoffline() (e *event.Event) {
	js.Rewrite("$<.onoffline")
	return e
}

func (*Window) SetOnoffline(e *event.Event) {
	js.Rewrite("$<.onoffline = $1", e)
}

func (*Window) GetOnonline() (e *event.Event) {
	js.Rewrite("$<.ononline")
	return e
}

func (*Window) SetOnonline(e *event.Event) {
	js.Rewrite("$<.ononline = $1", e)
}

func (*Window) GetOnorientationchange() (orientationchange *event.Event) {
	js.Rewrite("$<.onorientationchange")
	return orientationchange
}

func (*Window) SetOnorientationchange(orientationchange *event.Event) {
	js.Rewrite("$<.onorientationchange = $1", orientationchange)
}

func (*Window) GetOnpagehide() (pagehide *pagetransitionevent.PageTransitionEvent) {
	js.Rewrite("$<.onpagehide")
	return pagehide
}

func (*Window) SetOnpagehide(pagehide *pagetransitionevent.PageTransitionEvent) {
	js.Rewrite("$<.onpagehide = $1", pagehide)
}

func (*Window) GetOnpageshow() (pageshow *pagetransitionevent.PageTransitionEvent) {
	js.Rewrite("$<.onpageshow")
	return pageshow
}

func (*Window) SetOnpageshow(pageshow *pagetransitionevent.PageTransitionEvent) {
	js.Rewrite("$<.onpageshow = $1", pageshow)
}

func (*Window) GetOnpause() (e *event.Event) {
	js.Rewrite("$<.onpause")
	return e
}

func (*Window) SetOnpause(e *event.Event) {
	js.Rewrite("$<.onpause = $1", e)
}

func (*Window) GetOnplay() (e *event.Event) {
	js.Rewrite("$<.onplay")
	return e
}

func (*Window) SetOnplay(e *event.Event) {
	js.Rewrite("$<.onplay = $1", e)
}

func (*Window) GetOnplaying() (e *event.Event) {
	js.Rewrite("$<.onplaying")
	return e
}

func (*Window) SetOnplaying(e *event.Event) {
	js.Rewrite("$<.onplaying = $1", e)
}

func (*Window) GetOnpopstate() (popstate *popstateevent.PopStateEvent) {
	js.Rewrite("$<.onpopstate")
	return popstate
}

func (*Window) SetOnpopstate(popstate *popstateevent.PopStateEvent) {
	js.Rewrite("$<.onpopstate = $1", popstate)
}

func (*Window) GetOnprogress() (e *event.Event) {
	js.Rewrite("$<.onprogress")
	return e
}

func (*Window) SetOnprogress(e *event.Event) {
	js.Rewrite("$<.onprogress = $1", e)
}

func (*Window) GetOnratechange() (e *event.Event) {
	js.Rewrite("$<.onratechange")
	return e
}

func (*Window) SetOnratechange(e *event.Event) {
	js.Rewrite("$<.onratechange = $1", e)
}

func (*Window) GetOnreadystatechange() (e *event.Event) {
	js.Rewrite("$<.onreadystatechange")
	return e
}

func (*Window) SetOnreadystatechange(e *event.Event) {
	js.Rewrite("$<.onreadystatechange = $1", e)
}

func (*Window) GetOnreset() (e *event.Event) {
	js.Rewrite("$<.onreset")
	return e
}

func (*Window) SetOnreset(e *event.Event) {
	js.Rewrite("$<.onreset = $1", e)
}

func (*Window) GetOnresize() (resize *uievent.UIEvent) {
	js.Rewrite("$<.onresize")
	return resize
}

func (*Window) SetOnresize(resize *uievent.UIEvent) {
	js.Rewrite("$<.onresize = $1", resize)
}

func (*Window) GetOnscroll() (e *event.Event) {
	js.Rewrite("$<.onscroll")
	return e
}

func (*Window) SetOnscroll(e *event.Event) {
	js.Rewrite("$<.onscroll = $1", e)
}

func (*Window) GetOnseeked() (e *event.Event) {
	js.Rewrite("$<.onseeked")
	return e
}

func (*Window) SetOnseeked(e *event.Event) {
	js.Rewrite("$<.onseeked = $1", e)
}

func (*Window) GetOnseeking() (e *event.Event) {
	js.Rewrite("$<.onseeking")
	return e
}

func (*Window) SetOnseeking(e *event.Event) {
	js.Rewrite("$<.onseeking = $1", e)
}

func (*Window) GetOnselect() (e *event.Event) {
	js.Rewrite("$<.onselect")
	return e
}

func (*Window) SetOnselect(e *event.Event) {
	js.Rewrite("$<.onselect = $1", e)
}

func (*Window) GetOnstalled() (e *event.Event) {
	js.Rewrite("$<.onstalled")
	return e
}

func (*Window) SetOnstalled(e *event.Event) {
	js.Rewrite("$<.onstalled = $1", e)
}

func (*Window) GetOnstorage() (storage *storageevent.StorageEvent) {
	js.Rewrite("$<.onstorage")
	return storage
}

func (*Window) SetOnstorage(storage *storageevent.StorageEvent) {
	js.Rewrite("$<.onstorage = $1", storage)
}

func (*Window) GetOnsubmit() (e *event.Event) {
	js.Rewrite("$<.onsubmit")
	return e
}

func (*Window) SetOnsubmit(e *event.Event) {
	js.Rewrite("$<.onsubmit = $1", e)
}

func (*Window) GetOnsuspend() (e *event.Event) {
	js.Rewrite("$<.onsuspend")
	return e
}

func (*Window) SetOnsuspend(e *event.Event) {
	js.Rewrite("$<.onsuspend = $1", e)
}

func (*Window) GetOntimeupdate() (e *event.Event) {
	js.Rewrite("$<.ontimeupdate")
	return e
}

func (*Window) SetOntimeupdate(e *event.Event) {
	js.Rewrite("$<.ontimeupdate = $1", e)
}

func (*Window) GetOntouchcancel() (ontouchcancel interface{}) {
	js.Rewrite("$<.ontouchcancel")
	return ontouchcancel
}

func (*Window) SetOntouchcancel(ontouchcancel interface{}) {
	js.Rewrite("$<.ontouchcancel = $1", ontouchcancel)
}

func (*Window) GetOntouchend() (ontouchend interface{}) {
	js.Rewrite("$<.ontouchend")
	return ontouchend
}

func (*Window) SetOntouchend(ontouchend interface{}) {
	js.Rewrite("$<.ontouchend = $1", ontouchend)
}

func (*Window) GetOntouchmove() (ontouchmove interface{}) {
	js.Rewrite("$<.ontouchmove")
	return ontouchmove
}

func (*Window) SetOntouchmove(ontouchmove interface{}) {
	js.Rewrite("$<.ontouchmove = $1", ontouchmove)
}

func (*Window) GetOntouchstart() (ontouchstart interface{}) {
	js.Rewrite("$<.ontouchstart")
	return ontouchstart
}

func (*Window) SetOntouchstart(ontouchstart interface{}) {
	js.Rewrite("$<.ontouchstart = $1", ontouchstart)
}

func (*Window) GetOnunload() (unload *event.Event) {
	js.Rewrite("$<.onunload")
	return unload
}

func (*Window) SetOnunload(unload *event.Event) {
	js.Rewrite("$<.onunload = $1", unload)
}

func (*Window) GetOnvolumechange() (e *event.Event) {
	js.Rewrite("$<.onvolumechange")
	return e
}

func (*Window) SetOnvolumechange(e *event.Event) {
	js.Rewrite("$<.onvolumechange = $1", e)
}

func (*Window) GetOnwaiting() (e *event.Event) {
	js.Rewrite("$<.onwaiting")
	return e
}

func (*Window) SetOnwaiting(e *event.Event) {
	js.Rewrite("$<.onwaiting = $1", e)
}

func (*Window) GetOpener() (opener *Window) {
	js.Rewrite("$<.opener")
	return opener
}

func (*Window) GetOrientation() (orientation string) {
	js.Rewrite("$<.orientation")
	return orientation
}

func (*Window) GetOuterHeight() (outerHeight int) {
	js.Rewrite("$<.outerHeight")
	return outerHeight
}

func (*Window) GetOuterWidth() (outerWidth int) {
	js.Rewrite("$<.outerWidth")
	return outerWidth
}

func (*Window) GetPageXOffset() (pageXOffset int) {
	js.Rewrite("$<.pageXOffset")
	return pageXOffset
}

func (*Window) GetPageYOffset() (pageYOffset int) {
	js.Rewrite("$<.pageYOffset")
	return pageYOffset
}

func (*Window) GetParent() (parent *Window) {
	js.Rewrite("$<.parent")
	return parent
}

func (*Window) GetPerformance() (performance *performance.Performance) {
	js.Rewrite("$<.performance")
	return performance
}

func (*Window) GetPersonalbar() (personalbar *barprop.BarProp) {
	js.Rewrite("$<.personalbar")
	return personalbar
}

func (*Window) GetScreen() (screen *screen.Screen) {
	js.Rewrite("$<.screen")
	return screen
}

func (*Window) GetScreenLeft() (screenLeft int) {
	js.Rewrite("$<.screenLeft")
	return screenLeft
}

func (*Window) GetScreenTop() (screenTop int) {
	js.Rewrite("$<.screenTop")
	return screenTop
}

func (*Window) GetScreenX() (screenX int) {
	js.Rewrite("$<.screenX")
	return screenX
}

func (*Window) GetScreenY() (screenY int) {
	js.Rewrite("$<.screenY")
	return screenY
}

func (*Window) GetScrollbars() (scrollbars *barprop.BarProp) {
	js.Rewrite("$<.scrollbars")
	return scrollbars
}

func (*Window) GetScrollX() (scrollX int) {
	js.Rewrite("$<.scrollX")
	return scrollX
}

func (*Window) GetScrollY() (scrollY int) {
	js.Rewrite("$<.scrollY")
	return scrollY
}

func (*Window) GetSelf() (self *Window) {
	js.Rewrite("$<.self")
	return self
}

func (*Window) GetSpeechSynthesis() (speechSynthesis *speechsynthesis.SpeechSynthesis) {
	js.Rewrite("$<.speechSynthesis")
	return speechSynthesis
}

func (*Window) GetStatus() (status string) {
	js.Rewrite("$<.status")
	return status
}

func (*Window) SetStatus(status string) {
	js.Rewrite("$<.status = $1", status)
}

func (*Window) GetStatusbar() (statusbar *barprop.BarProp) {
	js.Rewrite("$<.statusbar")
	return statusbar
}

func (*Window) GetStyleMedia() (styleMedia *stylemedia.StyleMedia) {
	js.Rewrite("$<.styleMedia")
	return styleMedia
}

func (*Window) GetToolbar() (toolbar *barprop.BarProp) {
	js.Rewrite("$<.toolbar")
	return toolbar
}

func (*Window) GetTop() (top *Window) {
	js.Rewrite("$<.top")
	return top
}

func (*Window) GetWindow() (window *Window) {
	js.Rewrite("$<.window")
	return window
}
