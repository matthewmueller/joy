package window

import (
	"github.com/matthewmueller/golly/dom/barprop"
	"github.com/matthewmueller/golly/dom/cachestorage"
	"github.com/matthewmueller/golly/dom/crypto"
	"github.com/matthewmueller/golly/dom/extensionscriptapis"
	"github.com/matthewmueller/golly/dom/external"
	"github.com/matthewmueller/golly/dom/focusnavigationorigin"
	"github.com/matthewmueller/golly/dom/history"
	"github.com/matthewmueller/golly/dom/location"
	"github.com/matthewmueller/golly/dom/mediaquery"
	"github.com/matthewmueller/golly/dom/mscredentials"
	"github.com/matthewmueller/golly/dom/navigationreason"
	"github.com/matthewmueller/golly/dom/performance"
	"github.com/matthewmueller/golly/dom/request"
	"github.com/matthewmueller/golly/dom/requestinit"
	"github.com/matthewmueller/golly/dom/response"
	"github.com/matthewmueller/golly/dom/storage"
	"github.com/matthewmueller/golly/dom/stylemedia"
	"github.com/matthewmueller/golly/dom/webkitpoint"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New() *Window {
	js.Rewrite("window")
	return &Window{}
}

// Window struct
// js:"Window,omit"
type Window struct {
	EventTarget
}

// ClearInterval fn
func (*Window) ClearInterval(handle int) {
	js.Rewrite("$<.clearInterval($1)", handle)
}

// ClearTimeout fn
func (*Window) ClearTimeout(handle int) {
	js.Rewrite("$<.clearTimeout($1)", handle)
}

// SetInterval fn
func (*Window) SetInterval(handler interface{}, timeout *interface{}, args interface{}) (i int) {
	js.Rewrite("$<.setInterval($1, $2, $3)", handler, timeout, args)
	return i
}

// SetTimeout fn
func (*Window) SetTimeout(handler interface{}, timeout *interface{}, args interface{}) (i int) {
	js.Rewrite("$<.setTimeout($1, $2, $3)", handler, timeout, args)
	return i
}

// Atob fn
func (*Window) Atob(encodedString string) (s string) {
	js.Rewrite("$<.atob($1)", encodedString)
	return s
}

// Btoa fn
func (*Window) Btoa(rawString string) (s string) {
	js.Rewrite("$<.btoa($1)", rawString)
	return s
}

// Fetch fn
func (*Window) Fetch(input *request.Request, init *requestinit.RequestInit) (r *response.Response) {
	js.Rewrite("await $<.fetch($1, $2)", input, init)
	return r
}

// Alert fn
func (*Window) Alert(message *string) {
	js.Rewrite("$<.alert($1)", message)
}

// Blur fn
func (*Window) Blur() {
	js.Rewrite("$<.blur()")
}

// CancelAnimationFrame fn
func (*Window) CancelAnimationFrame(handle int) {
	js.Rewrite("$<.cancelAnimationFrame($1)", handle)
}

// CaptureEvents fn
func (*Window) CaptureEvents() {
	js.Rewrite("$<.captureEvents()")
}

// Close fn
func (*Window) Close() {
	js.Rewrite("$<.close()")
}

// Confirm fn
func (*Window) Confirm(message *string) (b bool) {
	js.Rewrite("$<.confirm($1)", message)
	return b
}

// DepartFocus fn
func (*Window) DepartFocus(navigationReason *navigationreason.NavigationReason, origin *focusnavigationorigin.FocusNavigationOrigin) {
	js.Rewrite("$<.departFocus($1, $2)", navigationReason, origin)
}

// Focus fn
func (*Window) Focus() {
	js.Rewrite("$<.focus()")
}

// GetComputedStyle fn
func (*Window) GetComputedStyle(elt Element, pseudoElt *string) (c *CSSStyleDeclaration) {
	js.Rewrite("$<.getComputedStyle($1, $2)", elt, pseudoElt)
	return c
}

// GetMatchedCSSRules fn
func (*Window) GetMatchedCSSRules(elt Element, pseudoElt *string) (c *CSSRuleList) {
	js.Rewrite("$<.getMatchedCSSRules($1, $2)", elt, pseudoElt)
	return c
}

// GetSelection fn
func (*Window) GetSelection() (s *Selection) {
	js.Rewrite("$<.getSelection()")
	return s
}

// MatchMedia fn
func (*Window) MatchMedia(mediaQuery string) (m *mediaquery.MediaQueryList) {
	js.Rewrite("$<.matchMedia($1)", mediaQuery)
	return m
}

// MoveBy fn
func (*Window) MoveBy(x *int, y *int) {
	js.Rewrite("$<.moveBy($1, $2)", x, y)
}

// MoveTo fn
func (*Window) MoveTo(x *int, y *int) {
	js.Rewrite("$<.moveTo($1, $2)", x, y)
}

// MsWriteProfilerMark fn
func (*Window) MsWriteProfilerMark(profilerMarkName string) {
	js.Rewrite("$<.msWriteProfilerMark($1)", profilerMarkName)
}

// Open fn
func (*Window) Open(url *string, target *string, features *string, replace *bool) (w *Window) {
	js.Rewrite("$<.open($1, $2, $3, $4)", url, target, features, replace)
	return w
}

// PostMessage fn
func (*Window) PostMessage(message interface{}, targetOrigin string, transfer *[]interface{}) {
	js.Rewrite("$<.postMessage($1, $2, $3)", message, targetOrigin, transfer)
}

// Print fn
func (*Window) Print() {
	js.Rewrite("$<.print()")
}

// Prompt fn
func (*Window) Prompt(message *string, def *string) (s string) {
	js.Rewrite("$<.prompt($1, $2)", message, def)
	return s
}

// ReleaseEvents fn
func (*Window) ReleaseEvents() {
	js.Rewrite("$<.releaseEvents()")
}

// RequestAnimationFrame fn
func (*Window) RequestAnimationFrame(callback func(time int)) (i int) {
	js.Rewrite("$<.requestAnimationFrame($1)", callback)
	return i
}

// ResizeBy fn
func (*Window) ResizeBy(x *int, y *int) {
	js.Rewrite("$<.resizeBy($1, $2)", x, y)
}

// ResizeTo fn
func (*Window) ResizeTo(x *int, y *int) {
	js.Rewrite("$<.resizeTo($1, $2)", x, y)
}

// Scroll fn
func (*Window) Scroll(x *int, y *int) {
	js.Rewrite("$<.scroll($1, $2)", x, y)
}

// ScrollBy fn
func (*Window) ScrollBy(x *int, y *int) {
	js.Rewrite("$<.scrollBy($1, $2)", x, y)
}

// ScrollTo fn
func (*Window) ScrollTo(x *int, y *int) {
	js.Rewrite("$<.scrollTo($1, $2)", x, y)
}

// Stop fn
func (*Window) Stop() {
	js.Rewrite("$<.stop()")
}

// WebkitCancelAnimationFrame fn
func (*Window) WebkitCancelAnimationFrame(handle int) {
	js.Rewrite("$<.webkitCancelAnimationFrame($1)", handle)
}

// WebkitConvertPointFromNodeToPage fn
func (*Window) WebkitConvertPointFromNodeToPage(node Node, pt *webkitpoint.WebKitPoint) (w *webkitpoint.WebKitPoint) {
	js.Rewrite("$<.webkitConvertPointFromNodeToPage($1, $2)", node, pt)
	return w
}

// WebkitConvertPointFromPageToNode fn
func (*Window) WebkitConvertPointFromPageToNode(node Node, pt *webkitpoint.WebKitPoint) (w *webkitpoint.WebKitPoint) {
	js.Rewrite("$<.webkitConvertPointFromPageToNode($1, $2)", node, pt)
	return w
}

// WebkitRequestAnimationFrame fn
func (*Window) WebkitRequestAnimationFrame(callback func(time int)) (i int) {
	js.Rewrite("$<.webkitRequestAnimationFrame($1)", callback)
	return i
}

// SessionStorage prop
func (*Window) SessionStorage() (sessionStorage *storage.Storage) {
	js.Rewrite("$<.sessionStorage")
	return sessionStorage
}

// LocalStorage prop
func (*Window) LocalStorage() (localStorage *storage.Storage) {
	js.Rewrite("$<.localStorage")
	return localStorage
}

// Console prop
func (*Window) Console() (console *Console) {
	js.Rewrite("$<.console")
	return console
}

// Onpointercancel prop
func (*Window) Onpointercancel() (onpointercancel func(Event)) {
	js.Rewrite("$<.onpointercancel")
	return onpointercancel
}

// Onpointercancel prop
func (*Window) SetOnpointercancel(onpointercancel func(Event)) {
	js.Rewrite("$<.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
func (*Window) Onpointerdown() (onpointerdown func(Event)) {
	js.Rewrite("$<.onpointerdown")
	return onpointerdown
}

// Onpointerdown prop
func (*Window) SetOnpointerdown(onpointerdown func(Event)) {
	js.Rewrite("$<.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
func (*Window) Onpointerenter() (onpointerenter func(Event)) {
	js.Rewrite("$<.onpointerenter")
	return onpointerenter
}

// Onpointerenter prop
func (*Window) SetOnpointerenter(onpointerenter func(Event)) {
	js.Rewrite("$<.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
func (*Window) Onpointerleave() (onpointerleave func(Event)) {
	js.Rewrite("$<.onpointerleave")
	return onpointerleave
}

// Onpointerleave prop
func (*Window) SetOnpointerleave(onpointerleave func(Event)) {
	js.Rewrite("$<.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
func (*Window) Onpointermove() (onpointermove func(Event)) {
	js.Rewrite("$<.onpointermove")
	return onpointermove
}

// Onpointermove prop
func (*Window) SetOnpointermove(onpointermove func(Event)) {
	js.Rewrite("$<.onpointermove = $1", onpointermove)
}

// Onpointerout prop
func (*Window) Onpointerout() (onpointerout func(Event)) {
	js.Rewrite("$<.onpointerout")
	return onpointerout
}

// Onpointerout prop
func (*Window) SetOnpointerout(onpointerout func(Event)) {
	js.Rewrite("$<.onpointerout = $1", onpointerout)
}

// Onpointerover prop
func (*Window) Onpointerover() (onpointerover func(Event)) {
	js.Rewrite("$<.onpointerover")
	return onpointerover
}

// Onpointerover prop
func (*Window) SetOnpointerover(onpointerover func(Event)) {
	js.Rewrite("$<.onpointerover = $1", onpointerover)
}

// Onpointerup prop
func (*Window) Onpointerup() (onpointerup func(Event)) {
	js.Rewrite("$<.onpointerup")
	return onpointerup
}

// Onpointerup prop
func (*Window) SetOnpointerup(onpointerup func(Event)) {
	js.Rewrite("$<.onpointerup = $1", onpointerup)
}

// Onwheel prop
func (*Window) Onwheel() (onwheel func(Event)) {
	js.Rewrite("$<.onwheel")
	return onwheel
}

// Onwheel prop
func (*Window) SetOnwheel(onwheel func(Event)) {
	js.Rewrite("$<.onwheel = $1", onwheel)
}

// IndexedDB prop
func (*Window) IndexedDB() (indexedDB *IDBFactory) {
	js.Rewrite("$<.indexedDB")
	return indexedDB
}

// ApplicationCache prop
func (*Window) ApplicationCache() (applicationCache *ApplicationCache) {
	js.Rewrite("$<.applicationCache")
	return applicationCache
}

// Caches prop
func (*Window) Caches() (caches *cachestorage.CacheStorage) {
	js.Rewrite("$<.caches")
	return caches
}

// ClientInformation prop
func (*Window) ClientInformation() (clientInformation *Navigator) {
	js.Rewrite("$<.clientInformation")
	return clientInformation
}

// Closed prop
func (*Window) Closed() (closed bool) {
	js.Rewrite("$<.closed")
	return closed
}

// Crypto prop
func (*Window) Crypto() (crypto *crypto.Crypto) {
	js.Rewrite("$<.crypto")
	return crypto
}

// DefaultStatus prop
func (*Window) DefaultStatus() (defaultStatus string) {
	js.Rewrite("$<.defaultStatus")
	return defaultStatus
}

// DefaultStatus prop
func (*Window) SetDefaultStatus(defaultStatus string) {
	js.Rewrite("$<.defaultStatus = $1", defaultStatus)
}

// DevicePixelRatio prop
func (*Window) DevicePixelRatio() (devicePixelRatio float32) {
	js.Rewrite("$<.devicePixelRatio")
	return devicePixelRatio
}

// Document prop
func (*Window) Document() (document Document) {
	js.Rewrite("$<.document")
	return document
}

// DoNotTrack prop
func (*Window) DoNotTrack() (doNotTrack string) {
	js.Rewrite("$<.doNotTrack")
	return doNotTrack
}

// Event prop
func (*Window) Event() (event Event) {
	js.Rewrite("$<.event")
	return event
}

// Event prop
func (*Window) SetEvent(event Event) {
	js.Rewrite("$<.event = $1", event)
}

// External prop
func (*Window) External() (external *external.External) {
	js.Rewrite("$<.external")
	return external
}

// FrameElement prop
func (*Window) FrameElement() (frameElement Element) {
	js.Rewrite("$<.frameElement")
	return frameElement
}

// Frames prop
func (*Window) Frames() (frames *Window) {
	js.Rewrite("$<.frames")
	return frames
}

// History prop
func (*Window) History() (history *history.History) {
	js.Rewrite("$<.history")
	return history
}

// InnerHeight prop
func (*Window) InnerHeight() (innerHeight int) {
	js.Rewrite("$<.innerHeight")
	return innerHeight
}

// InnerWidth prop
func (*Window) InnerWidth() (innerWidth int) {
	js.Rewrite("$<.innerWidth")
	return innerWidth
}

// IsSecureContext prop
func (*Window) IsSecureContext() (isSecureContext bool) {
	js.Rewrite("$<.isSecureContext")
	return isSecureContext
}

// Length prop
func (*Window) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}

// Location prop
func (*Window) Location() (location *location.Location) {
	js.Rewrite("$<.location")
	return location
}

// Locationbar prop
func (*Window) Locationbar() (locationbar *barprop.BarProp) {
	js.Rewrite("$<.locationbar")
	return locationbar
}

// Menubar prop
func (*Window) Menubar() (menubar *barprop.BarProp) {
	js.Rewrite("$<.menubar")
	return menubar
}

// MsContentScript prop
func (*Window) MsContentScript() (msContentScript *extensionscriptapis.ExtensionScriptApis) {
	js.Rewrite("$<.msContentScript")
	return msContentScript
}

// MsCredentials prop
func (*Window) MsCredentials() (msCredentials *mscredentials.MSCredentials) {
	js.Rewrite("$<.msCredentials")
	return msCredentials
}

// Name prop
func (*Window) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop
func (*Window) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// Navigator prop
func (*Window) Navigator() (navigator *Navigator) {
	js.Rewrite("$<.navigator")
	return navigator
}

// OffscreenBuffering prop
func (*Window) OffscreenBuffering() (offscreenBuffering interface{}) {
	js.Rewrite("$<.offscreenBuffering")
	return offscreenBuffering
}

// OffscreenBuffering prop
func (*Window) SetOffscreenBuffering(offscreenBuffering interface{}) {
	js.Rewrite("$<.offscreenBuffering = $1", offscreenBuffering)
}

// Onabort prop
func (*Window) Onabort() (onabort func(Event)) {
	js.Rewrite("$<.onabort")
	return onabort
}

// Onabort prop
func (*Window) SetOnabort(onabort func(Event)) {
	js.Rewrite("$<.onabort = $1", onabort)
}

// Onafterprint prop
func (*Window) Onafterprint() (onafterprint func(Event)) {
	js.Rewrite("$<.onafterprint")
	return onafterprint
}

// Onafterprint prop
func (*Window) SetOnafterprint(onafterprint func(Event)) {
	js.Rewrite("$<.onafterprint = $1", onafterprint)
}

// Onbeforeprint prop
func (*Window) Onbeforeprint() (onbeforeprint func(Event)) {
	js.Rewrite("$<.onbeforeprint")
	return onbeforeprint
}

// Onbeforeprint prop
func (*Window) SetOnbeforeprint(onbeforeprint func(Event)) {
	js.Rewrite("$<.onbeforeprint = $1", onbeforeprint)
}

// Onbeforeunload prop
func (*Window) Onbeforeunload() (onbeforeunload func(*BeforeUnloadEvent)) {
	js.Rewrite("$<.onbeforeunload")
	return onbeforeunload
}

// Onbeforeunload prop
func (*Window) SetOnbeforeunload(onbeforeunload func(*BeforeUnloadEvent)) {
	js.Rewrite("$<.onbeforeunload = $1", onbeforeunload)
}

// Onblur prop
func (*Window) Onblur() (onblur func(*FocusEvent)) {
	js.Rewrite("$<.onblur")
	return onblur
}

// Onblur prop
func (*Window) SetOnblur(onblur func(*FocusEvent)) {
	js.Rewrite("$<.onblur = $1", onblur)
}

// Oncanplay prop
func (*Window) Oncanplay() (oncanplay func(Event)) {
	js.Rewrite("$<.oncanplay")
	return oncanplay
}

// Oncanplay prop
func (*Window) SetOncanplay(oncanplay func(Event)) {
	js.Rewrite("$<.oncanplay = $1", oncanplay)
}

// Oncanplaythrough prop
func (*Window) Oncanplaythrough() (oncanplaythrough func(Event)) {
	js.Rewrite("$<.oncanplaythrough")
	return oncanplaythrough
}

// Oncanplaythrough prop
func (*Window) SetOncanplaythrough(oncanplaythrough func(Event)) {
	js.Rewrite("$<.oncanplaythrough = $1", oncanplaythrough)
}

// Onchange prop
func (*Window) Onchange() (onchange func(Event)) {
	js.Rewrite("$<.onchange")
	return onchange
}

// Onchange prop
func (*Window) SetOnchange(onchange func(Event)) {
	js.Rewrite("$<.onchange = $1", onchange)
}

// Onclick prop
func (*Window) Onclick() (onclick func(Event)) {
	js.Rewrite("$<.onclick")
	return onclick
}

// Onclick prop
func (*Window) SetOnclick(onclick func(Event)) {
	js.Rewrite("$<.onclick = $1", onclick)
}

// Oncompassneedscalibration prop
func (*Window) Oncompassneedscalibration() (oncompassneedscalibration func(Event)) {
	js.Rewrite("$<.oncompassneedscalibration")
	return oncompassneedscalibration
}

// Oncompassneedscalibration prop
func (*Window) SetOncompassneedscalibration(oncompassneedscalibration func(Event)) {
	js.Rewrite("$<.oncompassneedscalibration = $1", oncompassneedscalibration)
}

// Oncontextmenu prop
func (*Window) Oncontextmenu() (oncontextmenu func(Event)) {
	js.Rewrite("$<.oncontextmenu")
	return oncontextmenu
}

// Oncontextmenu prop
func (*Window) SetOncontextmenu(oncontextmenu func(Event)) {
	js.Rewrite("$<.oncontextmenu = $1", oncontextmenu)
}

// Ondblclick prop
func (*Window) Ondblclick() (ondblclick func(Event)) {
	js.Rewrite("$<.ondblclick")
	return ondblclick
}

// Ondblclick prop
func (*Window) SetOndblclick(ondblclick func(Event)) {
	js.Rewrite("$<.ondblclick = $1", ondblclick)
}

// Ondevicelight prop
func (*Window) Ondevicelight() (ondevicelight func(*DeviceLightEvent)) {
	js.Rewrite("$<.ondevicelight")
	return ondevicelight
}

// Ondevicelight prop
func (*Window) SetOndevicelight(ondevicelight func(*DeviceLightEvent)) {
	js.Rewrite("$<.ondevicelight = $1", ondevicelight)
}

// Ondevicemotion prop
func (*Window) Ondevicemotion() (ondevicemotion func(*DeviceMotionEvent)) {
	js.Rewrite("$<.ondevicemotion")
	return ondevicemotion
}

// Ondevicemotion prop
func (*Window) SetOndevicemotion(ondevicemotion func(*DeviceMotionEvent)) {
	js.Rewrite("$<.ondevicemotion = $1", ondevicemotion)
}

// Ondeviceorientation prop
func (*Window) Ondeviceorientation() (ondeviceorientation func(*DeviceOrientationEvent)) {
	js.Rewrite("$<.ondeviceorientation")
	return ondeviceorientation
}

// Ondeviceorientation prop
func (*Window) SetOndeviceorientation(ondeviceorientation func(*DeviceOrientationEvent)) {
	js.Rewrite("$<.ondeviceorientation = $1", ondeviceorientation)
}

// Ondrag prop
func (*Window) Ondrag() (ondrag func(Event)) {
	js.Rewrite("$<.ondrag")
	return ondrag
}

// Ondrag prop
func (*Window) SetOndrag(ondrag func(Event)) {
	js.Rewrite("$<.ondrag = $1", ondrag)
}

// Ondragend prop
func (*Window) Ondragend() (ondragend func(Event)) {
	js.Rewrite("$<.ondragend")
	return ondragend
}

// Ondragend prop
func (*Window) SetOndragend(ondragend func(Event)) {
	js.Rewrite("$<.ondragend = $1", ondragend)
}

// Ondragenter prop
func (*Window) Ondragenter() (ondragenter func(Event)) {
	js.Rewrite("$<.ondragenter")
	return ondragenter
}

// Ondragenter prop
func (*Window) SetOndragenter(ondragenter func(Event)) {
	js.Rewrite("$<.ondragenter = $1", ondragenter)
}

// Ondragleave prop
func (*Window) Ondragleave() (ondragleave func(Event)) {
	js.Rewrite("$<.ondragleave")
	return ondragleave
}

// Ondragleave prop
func (*Window) SetOndragleave(ondragleave func(Event)) {
	js.Rewrite("$<.ondragleave = $1", ondragleave)
}

// Ondragover prop
func (*Window) Ondragover() (ondragover func(Event)) {
	js.Rewrite("$<.ondragover")
	return ondragover
}

// Ondragover prop
func (*Window) SetOndragover(ondragover func(Event)) {
	js.Rewrite("$<.ondragover = $1", ondragover)
}

// Ondragstart prop
func (*Window) Ondragstart() (ondragstart func(Event)) {
	js.Rewrite("$<.ondragstart")
	return ondragstart
}

// Ondragstart prop
func (*Window) SetOndragstart(ondragstart func(Event)) {
	js.Rewrite("$<.ondragstart = $1", ondragstart)
}

// Ondrop prop
func (*Window) Ondrop() (ondrop func(Event)) {
	js.Rewrite("$<.ondrop")
	return ondrop
}

// Ondrop prop
func (*Window) SetOndrop(ondrop func(Event)) {
	js.Rewrite("$<.ondrop = $1", ondrop)
}

// Ondurationchange prop
func (*Window) Ondurationchange() (ondurationchange func(Event)) {
	js.Rewrite("$<.ondurationchange")
	return ondurationchange
}

// Ondurationchange prop
func (*Window) SetOndurationchange(ondurationchange func(Event)) {
	js.Rewrite("$<.ondurationchange = $1", ondurationchange)
}

// Onemptied prop
func (*Window) Onemptied() (onemptied func(Event)) {
	js.Rewrite("$<.onemptied")
	return onemptied
}

// Onemptied prop
func (*Window) SetOnemptied(onemptied func(Event)) {
	js.Rewrite("$<.onemptied = $1", onemptied)
}

// Onended prop
func (*Window) Onended() (onended func(Event)) {
	js.Rewrite("$<.onended")
	return onended
}

// Onended prop
func (*Window) SetOnended(onended func(Event)) {
	js.Rewrite("$<.onended = $1", onended)
}

// Onerror prop
func (*Window) Onerror() (onerror func(columnNumber *uint, event interface{}, fileno *uint, source *string)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*Window) SetOnerror(onerror func(columnNumber *uint, event interface{}, fileno *uint, source *string)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onfocus prop
func (*Window) Onfocus() (onfocus func(*FocusEvent)) {
	js.Rewrite("$<.onfocus")
	return onfocus
}

// Onfocus prop
func (*Window) SetOnfocus(onfocus func(*FocusEvent)) {
	js.Rewrite("$<.onfocus = $1", onfocus)
}

// Onhashchange prop
func (*Window) Onhashchange() (onhashchange func(*HashChangeEvent)) {
	js.Rewrite("$<.onhashchange")
	return onhashchange
}

// Onhashchange prop
func (*Window) SetOnhashchange(onhashchange func(*HashChangeEvent)) {
	js.Rewrite("$<.onhashchange = $1", onhashchange)
}

// Oninput prop
func (*Window) Oninput() (oninput func(Event)) {
	js.Rewrite("$<.oninput")
	return oninput
}

// Oninput prop
func (*Window) SetOninput(oninput func(Event)) {
	js.Rewrite("$<.oninput = $1", oninput)
}

// Oninvalid prop
func (*Window) Oninvalid() (oninvalid func(Event)) {
	js.Rewrite("$<.oninvalid")
	return oninvalid
}

// Oninvalid prop
func (*Window) SetOninvalid(oninvalid func(Event)) {
	js.Rewrite("$<.oninvalid = $1", oninvalid)
}

// Onkeydown prop
func (*Window) Onkeydown() (onkeydown func(Event)) {
	js.Rewrite("$<.onkeydown")
	return onkeydown
}

// Onkeydown prop
func (*Window) SetOnkeydown(onkeydown func(Event)) {
	js.Rewrite("$<.onkeydown = $1", onkeydown)
}

// Onkeypress prop
func (*Window) Onkeypress() (onkeypress func(Event)) {
	js.Rewrite("$<.onkeypress")
	return onkeypress
}

// Onkeypress prop
func (*Window) SetOnkeypress(onkeypress func(Event)) {
	js.Rewrite("$<.onkeypress = $1", onkeypress)
}

// Onkeyup prop
func (*Window) Onkeyup() (onkeyup func(Event)) {
	js.Rewrite("$<.onkeyup")
	return onkeyup
}

// Onkeyup prop
func (*Window) SetOnkeyup(onkeyup func(Event)) {
	js.Rewrite("$<.onkeyup = $1", onkeyup)
}

// Onload prop
func (*Window) Onload() (onload func(Event)) {
	js.Rewrite("$<.onload")
	return onload
}

// Onload prop
func (*Window) SetOnload(onload func(Event)) {
	js.Rewrite("$<.onload = $1", onload)
}

// Onloadeddata prop
func (*Window) Onloadeddata() (onloadeddata func(Event)) {
	js.Rewrite("$<.onloadeddata")
	return onloadeddata
}

// Onloadeddata prop
func (*Window) SetOnloadeddata(onloadeddata func(Event)) {
	js.Rewrite("$<.onloadeddata = $1", onloadeddata)
}

// Onloadedmetadata prop
func (*Window) Onloadedmetadata() (onloadedmetadata func(Event)) {
	js.Rewrite("$<.onloadedmetadata")
	return onloadedmetadata
}

// Onloadedmetadata prop
func (*Window) SetOnloadedmetadata(onloadedmetadata func(Event)) {
	js.Rewrite("$<.onloadedmetadata = $1", onloadedmetadata)
}

// Onloadstart prop
func (*Window) Onloadstart() (onloadstart func(Event)) {
	js.Rewrite("$<.onloadstart")
	return onloadstart
}

// Onloadstart prop
func (*Window) SetOnloadstart(onloadstart func(Event)) {
	js.Rewrite("$<.onloadstart = $1", onloadstart)
}

// Onmessage prop
func (*Window) Onmessage() (onmessage func(*MessageEvent)) {
	js.Rewrite("$<.onmessage")
	return onmessage
}

// Onmessage prop
func (*Window) SetOnmessage(onmessage func(*MessageEvent)) {
	js.Rewrite("$<.onmessage = $1", onmessage)
}

// Onmousedown prop
func (*Window) Onmousedown() (onmousedown func(Event)) {
	js.Rewrite("$<.onmousedown")
	return onmousedown
}

// Onmousedown prop
func (*Window) SetOnmousedown(onmousedown func(Event)) {
	js.Rewrite("$<.onmousedown = $1", onmousedown)
}

// Onmouseenter prop
func (*Window) Onmouseenter() (onmouseenter func(Event)) {
	js.Rewrite("$<.onmouseenter")
	return onmouseenter
}

// Onmouseenter prop
func (*Window) SetOnmouseenter(onmouseenter func(Event)) {
	js.Rewrite("$<.onmouseenter = $1", onmouseenter)
}

// Onmouseleave prop
func (*Window) Onmouseleave() (onmouseleave func(Event)) {
	js.Rewrite("$<.onmouseleave")
	return onmouseleave
}

// Onmouseleave prop
func (*Window) SetOnmouseleave(onmouseleave func(Event)) {
	js.Rewrite("$<.onmouseleave = $1", onmouseleave)
}

// Onmousemove prop
func (*Window) Onmousemove() (onmousemove func(Event)) {
	js.Rewrite("$<.onmousemove")
	return onmousemove
}

// Onmousemove prop
func (*Window) SetOnmousemove(onmousemove func(Event)) {
	js.Rewrite("$<.onmousemove = $1", onmousemove)
}

// Onmouseout prop
func (*Window) Onmouseout() (onmouseout func(Event)) {
	js.Rewrite("$<.onmouseout")
	return onmouseout
}

// Onmouseout prop
func (*Window) SetOnmouseout(onmouseout func(Event)) {
	js.Rewrite("$<.onmouseout = $1", onmouseout)
}

// Onmouseover prop
func (*Window) Onmouseover() (onmouseover func(Event)) {
	js.Rewrite("$<.onmouseover")
	return onmouseover
}

// Onmouseover prop
func (*Window) SetOnmouseover(onmouseover func(Event)) {
	js.Rewrite("$<.onmouseover = $1", onmouseover)
}

// Onmouseup prop
func (*Window) Onmouseup() (onmouseup func(Event)) {
	js.Rewrite("$<.onmouseup")
	return onmouseup
}

// Onmouseup prop
func (*Window) SetOnmouseup(onmouseup func(Event)) {
	js.Rewrite("$<.onmouseup = $1", onmouseup)
}

// Onmousewheel prop
func (*Window) Onmousewheel() (onmousewheel func(Event)) {
	js.Rewrite("$<.onmousewheel")
	return onmousewheel
}

// Onmousewheel prop
func (*Window) SetOnmousewheel(onmousewheel func(Event)) {
	js.Rewrite("$<.onmousewheel = $1", onmousewheel)
}

// Onmsgesturechange prop
func (*Window) Onmsgesturechange() (onmsgesturechange func(Event)) {
	js.Rewrite("$<.onmsgesturechange")
	return onmsgesturechange
}

// Onmsgesturechange prop
func (*Window) SetOnmsgesturechange(onmsgesturechange func(Event)) {
	js.Rewrite("$<.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
func (*Window) Onmsgesturedoubletap() (onmsgesturedoubletap func(Event)) {
	js.Rewrite("$<.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// Onmsgesturedoubletap prop
func (*Window) SetOnmsgesturedoubletap(onmsgesturedoubletap func(Event)) {
	js.Rewrite("$<.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
func (*Window) Onmsgestureend() (onmsgestureend func(Event)) {
	js.Rewrite("$<.onmsgestureend")
	return onmsgestureend
}

// Onmsgestureend prop
func (*Window) SetOnmsgestureend(onmsgestureend func(Event)) {
	js.Rewrite("$<.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
func (*Window) Onmsgesturehold() (onmsgesturehold func(Event)) {
	js.Rewrite("$<.onmsgesturehold")
	return onmsgesturehold
}

// Onmsgesturehold prop
func (*Window) SetOnmsgesturehold(onmsgesturehold func(Event)) {
	js.Rewrite("$<.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
func (*Window) Onmsgesturestart() (onmsgesturestart func(Event)) {
	js.Rewrite("$<.onmsgesturestart")
	return onmsgesturestart
}

// Onmsgesturestart prop
func (*Window) SetOnmsgesturestart(onmsgesturestart func(Event)) {
	js.Rewrite("$<.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
func (*Window) Onmsgesturetap() (onmsgesturetap func(Event)) {
	js.Rewrite("$<.onmsgesturetap")
	return onmsgesturetap
}

// Onmsgesturetap prop
func (*Window) SetOnmsgesturetap(onmsgesturetap func(Event)) {
	js.Rewrite("$<.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsinertiastart prop
func (*Window) Onmsinertiastart() (onmsinertiastart func(Event)) {
	js.Rewrite("$<.onmsinertiastart")
	return onmsinertiastart
}

// Onmsinertiastart prop
func (*Window) SetOnmsinertiastart(onmsinertiastart func(Event)) {
	js.Rewrite("$<.onmsinertiastart = $1", onmsinertiastart)
}

// Onmspointercancel prop
func (*Window) Onmspointercancel() (onmspointercancel func(Event)) {
	js.Rewrite("$<.onmspointercancel")
	return onmspointercancel
}

// Onmspointercancel prop
func (*Window) SetOnmspointercancel(onmspointercancel func(Event)) {
	js.Rewrite("$<.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
func (*Window) Onmspointerdown() (onmspointerdown func(Event)) {
	js.Rewrite("$<.onmspointerdown")
	return onmspointerdown
}

// Onmspointerdown prop
func (*Window) SetOnmspointerdown(onmspointerdown func(Event)) {
	js.Rewrite("$<.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
func (*Window) Onmspointerenter() (onmspointerenter func(Event)) {
	js.Rewrite("$<.onmspointerenter")
	return onmspointerenter
}

// Onmspointerenter prop
func (*Window) SetOnmspointerenter(onmspointerenter func(Event)) {
	js.Rewrite("$<.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
func (*Window) Onmspointerleave() (onmspointerleave func(Event)) {
	js.Rewrite("$<.onmspointerleave")
	return onmspointerleave
}

// Onmspointerleave prop
func (*Window) SetOnmspointerleave(onmspointerleave func(Event)) {
	js.Rewrite("$<.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
func (*Window) Onmspointermove() (onmspointermove func(Event)) {
	js.Rewrite("$<.onmspointermove")
	return onmspointermove
}

// Onmspointermove prop
func (*Window) SetOnmspointermove(onmspointermove func(Event)) {
	js.Rewrite("$<.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
func (*Window) Onmspointerout() (onmspointerout func(Event)) {
	js.Rewrite("$<.onmspointerout")
	return onmspointerout
}

// Onmspointerout prop
func (*Window) SetOnmspointerout(onmspointerout func(Event)) {
	js.Rewrite("$<.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
func (*Window) Onmspointerover() (onmspointerover func(Event)) {
	js.Rewrite("$<.onmspointerover")
	return onmspointerover
}

// Onmspointerover prop
func (*Window) SetOnmspointerover(onmspointerover func(Event)) {
	js.Rewrite("$<.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
func (*Window) Onmspointerup() (onmspointerup func(Event)) {
	js.Rewrite("$<.onmspointerup")
	return onmspointerup
}

// Onmspointerup prop
func (*Window) SetOnmspointerup(onmspointerup func(Event)) {
	js.Rewrite("$<.onmspointerup = $1", onmspointerup)
}

// Onoffline prop
func (*Window) Onoffline() (onoffline func(Event)) {
	js.Rewrite("$<.onoffline")
	return onoffline
}

// Onoffline prop
func (*Window) SetOnoffline(onoffline func(Event)) {
	js.Rewrite("$<.onoffline = $1", onoffline)
}

// Ononline prop
func (*Window) Ononline() (ononline func(Event)) {
	js.Rewrite("$<.ononline")
	return ononline
}

// Ononline prop
func (*Window) SetOnonline(ononline func(Event)) {
	js.Rewrite("$<.ononline = $1", ononline)
}

// Onorientationchange prop
func (*Window) Onorientationchange() (onorientationchange func(Event)) {
	js.Rewrite("$<.onorientationchange")
	return onorientationchange
}

// Onorientationchange prop
func (*Window) SetOnorientationchange(onorientationchange func(Event)) {
	js.Rewrite("$<.onorientationchange = $1", onorientationchange)
}

// Onpagehide prop
func (*Window) Onpagehide() (onpagehide func(*PageTransitionEvent)) {
	js.Rewrite("$<.onpagehide")
	return onpagehide
}

// Onpagehide prop
func (*Window) SetOnpagehide(onpagehide func(*PageTransitionEvent)) {
	js.Rewrite("$<.onpagehide = $1", onpagehide)
}

// Onpageshow prop
func (*Window) Onpageshow() (onpageshow func(*PageTransitionEvent)) {
	js.Rewrite("$<.onpageshow")
	return onpageshow
}

// Onpageshow prop
func (*Window) SetOnpageshow(onpageshow func(*PageTransitionEvent)) {
	js.Rewrite("$<.onpageshow = $1", onpageshow)
}

// Onpause prop
func (*Window) Onpause() (onpause func(Event)) {
	js.Rewrite("$<.onpause")
	return onpause
}

// Onpause prop
func (*Window) SetOnpause(onpause func(Event)) {
	js.Rewrite("$<.onpause = $1", onpause)
}

// Onplay prop
func (*Window) Onplay() (onplay func(Event)) {
	js.Rewrite("$<.onplay")
	return onplay
}

// Onplay prop
func (*Window) SetOnplay(onplay func(Event)) {
	js.Rewrite("$<.onplay = $1", onplay)
}

// Onplaying prop
func (*Window) Onplaying() (onplaying func(Event)) {
	js.Rewrite("$<.onplaying")
	return onplaying
}

// Onplaying prop
func (*Window) SetOnplaying(onplaying func(Event)) {
	js.Rewrite("$<.onplaying = $1", onplaying)
}

// Onpopstate prop
func (*Window) Onpopstate() (onpopstate func(*PopStateEvent)) {
	js.Rewrite("$<.onpopstate")
	return onpopstate
}

// Onpopstate prop
func (*Window) SetOnpopstate(onpopstate func(*PopStateEvent)) {
	js.Rewrite("$<.onpopstate = $1", onpopstate)
}

// Onprogress prop
func (*Window) Onprogress() (onprogress func(Event)) {
	js.Rewrite("$<.onprogress")
	return onprogress
}

// Onprogress prop
func (*Window) SetOnprogress(onprogress func(Event)) {
	js.Rewrite("$<.onprogress = $1", onprogress)
}

// Onratechange prop
func (*Window) Onratechange() (onratechange func(Event)) {
	js.Rewrite("$<.onratechange")
	return onratechange
}

// Onratechange prop
func (*Window) SetOnratechange(onratechange func(Event)) {
	js.Rewrite("$<.onratechange = $1", onratechange)
}

// Onreadystatechange prop
func (*Window) Onreadystatechange() (onreadystatechange func(Event)) {
	js.Rewrite("$<.onreadystatechange")
	return onreadystatechange
}

// Onreadystatechange prop
func (*Window) SetOnreadystatechange(onreadystatechange func(Event)) {
	js.Rewrite("$<.onreadystatechange = $1", onreadystatechange)
}

// Onreset prop
func (*Window) Onreset() (onreset func(Event)) {
	js.Rewrite("$<.onreset")
	return onreset
}

// Onreset prop
func (*Window) SetOnreset(onreset func(Event)) {
	js.Rewrite("$<.onreset = $1", onreset)
}

// Onresize prop
func (*Window) Onresize() (onresize func(UIEvent)) {
	js.Rewrite("$<.onresize")
	return onresize
}

// Onresize prop
func (*Window) SetOnresize(onresize func(UIEvent)) {
	js.Rewrite("$<.onresize = $1", onresize)
}

// Onscroll prop
func (*Window) Onscroll() (onscroll func(Event)) {
	js.Rewrite("$<.onscroll")
	return onscroll
}

// Onscroll prop
func (*Window) SetOnscroll(onscroll func(Event)) {
	js.Rewrite("$<.onscroll = $1", onscroll)
}

// Onseeked prop
func (*Window) Onseeked() (onseeked func(Event)) {
	js.Rewrite("$<.onseeked")
	return onseeked
}

// Onseeked prop
func (*Window) SetOnseeked(onseeked func(Event)) {
	js.Rewrite("$<.onseeked = $1", onseeked)
}

// Onseeking prop
func (*Window) Onseeking() (onseeking func(Event)) {
	js.Rewrite("$<.onseeking")
	return onseeking
}

// Onseeking prop
func (*Window) SetOnseeking(onseeking func(Event)) {
	js.Rewrite("$<.onseeking = $1", onseeking)
}

// Onselect prop
func (*Window) Onselect() (onselect func(Event)) {
	js.Rewrite("$<.onselect")
	return onselect
}

// Onselect prop
func (*Window) SetOnselect(onselect func(Event)) {
	js.Rewrite("$<.onselect = $1", onselect)
}

// Onstalled prop
func (*Window) Onstalled() (onstalled func(Event)) {
	js.Rewrite("$<.onstalled")
	return onstalled
}

// Onstalled prop
func (*Window) SetOnstalled(onstalled func(Event)) {
	js.Rewrite("$<.onstalled = $1", onstalled)
}

// Onstorage prop
func (*Window) Onstorage() (onstorage func(*StorageEvent)) {
	js.Rewrite("$<.onstorage")
	return onstorage
}

// Onstorage prop
func (*Window) SetOnstorage(onstorage func(*StorageEvent)) {
	js.Rewrite("$<.onstorage = $1", onstorage)
}

// Onsubmit prop
func (*Window) Onsubmit() (onsubmit func(Event)) {
	js.Rewrite("$<.onsubmit")
	return onsubmit
}

// Onsubmit prop
func (*Window) SetOnsubmit(onsubmit func(Event)) {
	js.Rewrite("$<.onsubmit = $1", onsubmit)
}

// Onsuspend prop
func (*Window) Onsuspend() (onsuspend func(Event)) {
	js.Rewrite("$<.onsuspend")
	return onsuspend
}

// Onsuspend prop
func (*Window) SetOnsuspend(onsuspend func(Event)) {
	js.Rewrite("$<.onsuspend = $1", onsuspend)
}

// Ontimeupdate prop
func (*Window) Ontimeupdate() (ontimeupdate func(Event)) {
	js.Rewrite("$<.ontimeupdate")
	return ontimeupdate
}

// Ontimeupdate prop
func (*Window) SetOntimeupdate(ontimeupdate func(Event)) {
	js.Rewrite("$<.ontimeupdate = $1", ontimeupdate)
}

// Ontouchcancel prop
func (*Window) Ontouchcancel() (ontouchcancel interface{}) {
	js.Rewrite("$<.ontouchcancel")
	return ontouchcancel
}

// Ontouchcancel prop
func (*Window) SetOntouchcancel(ontouchcancel interface{}) {
	js.Rewrite("$<.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
func (*Window) Ontouchend() (ontouchend interface{}) {
	js.Rewrite("$<.ontouchend")
	return ontouchend
}

// Ontouchend prop
func (*Window) SetOntouchend(ontouchend interface{}) {
	js.Rewrite("$<.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
func (*Window) Ontouchmove() (ontouchmove interface{}) {
	js.Rewrite("$<.ontouchmove")
	return ontouchmove
}

// Ontouchmove prop
func (*Window) SetOntouchmove(ontouchmove interface{}) {
	js.Rewrite("$<.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
func (*Window) Ontouchstart() (ontouchstart interface{}) {
	js.Rewrite("$<.ontouchstart")
	return ontouchstart
}

// Ontouchstart prop
func (*Window) SetOntouchstart(ontouchstart interface{}) {
	js.Rewrite("$<.ontouchstart = $1", ontouchstart)
}

// Onunload prop
func (*Window) Onunload() (onunload func(Event)) {
	js.Rewrite("$<.onunload")
	return onunload
}

// Onunload prop
func (*Window) SetOnunload(onunload func(Event)) {
	js.Rewrite("$<.onunload = $1", onunload)
}

// Onvolumechange prop
func (*Window) Onvolumechange() (onvolumechange func(Event)) {
	js.Rewrite("$<.onvolumechange")
	return onvolumechange
}

// Onvolumechange prop
func (*Window) SetOnvolumechange(onvolumechange func(Event)) {
	js.Rewrite("$<.onvolumechange = $1", onvolumechange)
}

// Onwaiting prop
func (*Window) Onwaiting() (onwaiting func(Event)) {
	js.Rewrite("$<.onwaiting")
	return onwaiting
}

// Onwaiting prop
func (*Window) SetOnwaiting(onwaiting func(Event)) {
	js.Rewrite("$<.onwaiting = $1", onwaiting)
}

// Opener prop
func (*Window) Opener() (opener *Window) {
	js.Rewrite("$<.opener")
	return opener
}

// Orientation prop
func (*Window) Orientation() (orientation string) {
	js.Rewrite("$<.orientation")
	return orientation
}

// OuterHeight prop
func (*Window) OuterHeight() (outerHeight int) {
	js.Rewrite("$<.outerHeight")
	return outerHeight
}

// OuterWidth prop
func (*Window) OuterWidth() (outerWidth int) {
	js.Rewrite("$<.outerWidth")
	return outerWidth
}

// PageXOffset prop
func (*Window) PageXOffset() (pageXOffset int) {
	js.Rewrite("$<.pageXOffset")
	return pageXOffset
}

// PageYOffset prop
func (*Window) PageYOffset() (pageYOffset int) {
	js.Rewrite("$<.pageYOffset")
	return pageYOffset
}

// Parent prop
func (*Window) Parent() (parent *Window) {
	js.Rewrite("$<.parent")
	return parent
}

// Performance prop
func (*Window) Performance() (performance *performance.Performance) {
	js.Rewrite("$<.performance")
	return performance
}

// Personalbar prop
func (*Window) Personalbar() (personalbar *barprop.BarProp) {
	js.Rewrite("$<.personalbar")
	return personalbar
}

// Screen prop
func (*Window) Screen() (screen *Screen) {
	js.Rewrite("$<.screen")
	return screen
}

// ScreenLeft prop
func (*Window) ScreenLeft() (screenLeft int) {
	js.Rewrite("$<.screenLeft")
	return screenLeft
}

// ScreenTop prop
func (*Window) ScreenTop() (screenTop int) {
	js.Rewrite("$<.screenTop")
	return screenTop
}

// ScreenX prop
func (*Window) ScreenX() (screenX int) {
	js.Rewrite("$<.screenX")
	return screenX
}

// ScreenY prop
func (*Window) ScreenY() (screenY int) {
	js.Rewrite("$<.screenY")
	return screenY
}

// Scrollbars prop
func (*Window) Scrollbars() (scrollbars *barprop.BarProp) {
	js.Rewrite("$<.scrollbars")
	return scrollbars
}

// ScrollX prop
func (*Window) ScrollX() (scrollX int) {
	js.Rewrite("$<.scrollX")
	return scrollX
}

// ScrollY prop
func (*Window) ScrollY() (scrollY int) {
	js.Rewrite("$<.scrollY")
	return scrollY
}

// Self prop
func (*Window) Self() (self *Window) {
	js.Rewrite("$<.self")
	return self
}

// SpeechSynthesis prop
func (*Window) SpeechSynthesis() (speechSynthesis *SpeechSynthesis) {
	js.Rewrite("$<.speechSynthesis")
	return speechSynthesis
}

// Status prop
func (*Window) Status() (status string) {
	js.Rewrite("$<.status")
	return status
}

// Status prop
func (*Window) SetStatus(status string) {
	js.Rewrite("$<.status = $1", status)
}

// Statusbar prop
func (*Window) Statusbar() (statusbar *barprop.BarProp) {
	js.Rewrite("$<.statusbar")
	return statusbar
}

// StyleMedia prop
func (*Window) StyleMedia() (styleMedia *stylemedia.StyleMedia) {
	js.Rewrite("$<.styleMedia")
	return styleMedia
}

// Toolbar prop
func (*Window) Toolbar() (toolbar *barprop.BarProp) {
	js.Rewrite("$<.toolbar")
	return toolbar
}

// Top prop
func (*Window) Top() (top *Window) {
	js.Rewrite("$<.top")
	return top
}

// Window prop
func (*Window) Window() (window *Window) {
	js.Rewrite("$<.window")
	return window
}
