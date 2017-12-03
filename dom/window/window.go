package window

import (
	"github.com/matthewmueller/joy/dom/barprop"
	"github.com/matthewmueller/joy/dom/cachestorage"
	"github.com/matthewmueller/joy/dom/crypto"
	"github.com/matthewmueller/joy/dom/extensionscriptapis"
	"github.com/matthewmueller/joy/dom/external"
	"github.com/matthewmueller/joy/dom/focusnavigationorigin"
	"github.com/matthewmueller/joy/dom/history"
	"github.com/matthewmueller/joy/dom/location"
	"github.com/matthewmueller/joy/dom/mediaquery"
	"github.com/matthewmueller/joy/dom/mscredentials"
	"github.com/matthewmueller/joy/dom/navigationreason"
	"github.com/matthewmueller/joy/dom/performance"
	"github.com/matthewmueller/joy/dom/request"
	"github.com/matthewmueller/joy/dom/requestinit"
	"github.com/matthewmueller/joy/dom/response"
	"github.com/matthewmueller/joy/dom/storage"
	"github.com/matthewmueller/joy/dom/stylemedia"
	"github.com/matthewmueller/joy/dom/webkitpoint"
	"github.com/matthewmueller/joy/js"
)

var _ EventTarget = (*Window)(nil)

// New fn
func New() *Window {
	js.Rewrite("window")
	return &Window{}
}

// Window struct
// js:"Window,omit"
type Window struct {
}

// Alert fn
// js:"alert"
func (*Window) Alert(message *string) {
	js.Rewrite("$_.alert($1)", message)
}

// Blur fn
// js:"blur"
func (*Window) Blur() {
	js.Rewrite("$_.blur()")
}

// CancelAnimationFrame fn
// js:"cancelAnimationFrame"
func (*Window) CancelAnimationFrame(handle int) {
	js.Rewrite("$_.cancelAnimationFrame($1)", handle)
}

// CaptureEvents fn
// js:"captureEvents"
func (*Window) CaptureEvents() {
	js.Rewrite("$_.captureEvents()")
}

// Close fn
// js:"close"
func (*Window) Close() {
	js.Rewrite("$_.close()")
}

// Confirm fn
// js:"confirm"
func (*Window) Confirm(message *string) (b bool) {
	js.Rewrite("$_.confirm($1)", message)
	return b
}

// DepartFocus fn
// js:"departFocus"
func (*Window) DepartFocus(navigationReason *navigationreason.NavigationReason, origin *focusnavigationorigin.FocusNavigationOrigin) {
	js.Rewrite("$_.departFocus($1, $2)", navigationReason, origin)
}

// Focus fn
// js:"focus"
func (*Window) Focus() {
	js.Rewrite("$_.focus()")
}

// GetComputedStyle fn
// js:"getComputedStyle"
func (*Window) GetComputedStyle(elt Element, pseudoElt *string) (c *CSSStyleDeclaration) {
	js.Rewrite("$_.getComputedStyle($1, $2)", elt, pseudoElt)
	return c
}

// GetMatchedCSSRules fn
// js:"getMatchedCSSRules"
func (*Window) GetMatchedCSSRules(elt Element, pseudoElt *string) (c *CSSRuleList) {
	js.Rewrite("$_.getMatchedCSSRules($1, $2)", elt, pseudoElt)
	return c
}

// GetSelection fn
// js:"getSelection"
func (*Window) GetSelection() (s *Selection) {
	js.Rewrite("$_.getSelection()")
	return s
}

// MatchMedia fn
// js:"matchMedia"
func (*Window) MatchMedia(mediaQuery string) (m *mediaquery.MediaQueryList) {
	js.Rewrite("$_.matchMedia($1)", mediaQuery)
	return m
}

// MoveBy fn
// js:"moveBy"
func (*Window) MoveBy(x *int, y *int) {
	js.Rewrite("$_.moveBy($1, $2)", x, y)
}

// MoveTo fn
// js:"moveTo"
func (*Window) MoveTo(x *int, y *int) {
	js.Rewrite("$_.moveTo($1, $2)", x, y)
}

// MsWriteProfilerMark fn
// js:"msWriteProfilerMark"
func (*Window) MsWriteProfilerMark(profilerMarkName string) {
	js.Rewrite("$_.msWriteProfilerMark($1)", profilerMarkName)
}

// Open fn
// js:"open"
func (*Window) Open(url *string, target *string, features *string, replace *bool) (w *Window) {
	js.Rewrite("$_.open($1, $2, $3, $4)", url, target, features, replace)
	return w
}

// PostMessage fn
// js:"postMessage"
func (*Window) PostMessage(message interface{}, targetOrigin string, transfer *[]interface{}) {
	js.Rewrite("$_.postMessage($1, $2, $3)", message, targetOrigin, transfer)
}

// Print fn
// js:"print"
func (*Window) Print() {
	js.Rewrite("$_.print()")
}

// Prompt fn
// js:"prompt"
func (*Window) Prompt(message *string, def *string) (s string) {
	js.Rewrite("$_.prompt($1, $2)", message, def)
	return s
}

// ReleaseEvents fn
// js:"releaseEvents"
func (*Window) ReleaseEvents() {
	js.Rewrite("$_.releaseEvents()")
}

// RequestAnimationFrame fn
// js:"requestAnimationFrame"
func (*Window) RequestAnimationFrame(callback func(time int)) (i int) {
	js.Rewrite("$_.requestAnimationFrame($1)", callback)
	return i
}

// ResizeBy fn
// js:"resizeBy"
func (*Window) ResizeBy(x *int, y *int) {
	js.Rewrite("$_.resizeBy($1, $2)", x, y)
}

// ResizeTo fn
// js:"resizeTo"
func (*Window) ResizeTo(x *int, y *int) {
	js.Rewrite("$_.resizeTo($1, $2)", x, y)
}

// Scroll fn
// js:"scroll"
func (*Window) Scroll(x *int, y *int) {
	js.Rewrite("$_.scroll($1, $2)", x, y)
}

// ScrollBy fn
// js:"scrollBy"
func (*Window) ScrollBy(x *int, y *int) {
	js.Rewrite("$_.scrollBy($1, $2)", x, y)
}

// ScrollTo fn
// js:"scrollTo"
func (*Window) ScrollTo(x *int, y *int) {
	js.Rewrite("$_.scrollTo($1, $2)", x, y)
}

// Stop fn
// js:"stop"
func (*Window) Stop() {
	js.Rewrite("$_.stop()")
}

// WebkitCancelAnimationFrame fn
// js:"webkitCancelAnimationFrame"
func (*Window) WebkitCancelAnimationFrame(handle int) {
	js.Rewrite("$_.webkitCancelAnimationFrame($1)", handle)
}

// WebkitConvertPointFromNodeToPage fn
// js:"webkitConvertPointFromNodeToPage"
func (*Window) WebkitConvertPointFromNodeToPage(node Node, pt *webkitpoint.WebKitPoint) (w *webkitpoint.WebKitPoint) {
	js.Rewrite("$_.webkitConvertPointFromNodeToPage($1, $2)", node, pt)
	return w
}

// WebkitConvertPointFromPageToNode fn
// js:"webkitConvertPointFromPageToNode"
func (*Window) WebkitConvertPointFromPageToNode(node Node, pt *webkitpoint.WebKitPoint) (w *webkitpoint.WebKitPoint) {
	js.Rewrite("$_.webkitConvertPointFromPageToNode($1, $2)", node, pt)
	return w
}

// WebkitRequestAnimationFrame fn
// js:"webkitRequestAnimationFrame"
func (*Window) WebkitRequestAnimationFrame(callback func(time int)) (i int) {
	js.Rewrite("$_.webkitRequestAnimationFrame($1)", callback)
	return i
}

// ClearInterval fn
// js:"clearInterval"
func (*Window) ClearInterval(handle int) {
	js.Rewrite("$_.clearInterval($1)", handle)
}

// ClearTimeout fn
// js:"clearTimeout"
func (*Window) ClearTimeout(handle int) {
	js.Rewrite("$_.clearTimeout($1)", handle)
}

// SetInterval fn
// js:"setInterval"
func (*Window) SetInterval(handler interface{}, timeout *interface{}, args interface{}) (i int) {
	js.Rewrite("$_.setInterval($1, $2, $3)", handler, timeout, args)
	return i
}

// SetTimeout fn
// js:"setTimeout"
func (*Window) SetTimeout(handler interface{}, timeout *interface{}, args interface{}) (i int) {
	js.Rewrite("$_.setTimeout($1, $2, $3)", handler, timeout, args)
	return i
}

// Atob fn
// js:"atob"
func (*Window) Atob(encodedString string) (s string) {
	js.Rewrite("$_.atob($1)", encodedString)
	return s
}

// Btoa fn
// js:"btoa"
func (*Window) Btoa(rawString string) (s string) {
	js.Rewrite("$_.btoa($1)", rawString)
	return s
}

// Fetch fn
// js:"fetch"
func (*Window) Fetch(input *request.Request, init *requestinit.RequestInit) (r *response.Response) {
	js.Rewrite("await $_.fetch($1, $2)", input, init)
	return r
}

// AddEventListener fn
// js:"addEventListener"
func (*Window) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*Window) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*Window) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// ApplicationCache prop
// js:"applicationCache"
func (*Window) ApplicationCache() (applicationCache *ApplicationCache) {
	js.Rewrite("$_.applicationCache")
	return applicationCache
}

// Caches prop
// js:"caches"
func (*Window) Caches() (caches *cachestorage.CacheStorage) {
	js.Rewrite("$_.caches")
	return caches
}

// ClientInformation prop
// js:"clientInformation"
func (*Window) ClientInformation() (clientInformation *Navigator) {
	js.Rewrite("$_.clientInformation")
	return clientInformation
}

// Closed prop
// js:"closed"
func (*Window) Closed() (closed bool) {
	js.Rewrite("$_.closed")
	return closed
}

// Crypto prop
// js:"crypto"
func (*Window) Crypto() (crypto *crypto.Crypto) {
	js.Rewrite("$_.crypto")
	return crypto
}

// DefaultStatus prop
// js:"defaultStatus"
func (*Window) DefaultStatus() (defaultStatus string) {
	js.Rewrite("$_.defaultStatus")
	return defaultStatus
}

// SetDefaultStatus prop
// js:"defaultStatus"
func (*Window) SetDefaultStatus(defaultStatus string) {
	js.Rewrite("$_.defaultStatus = $1", defaultStatus)
}

// DevicePixelRatio prop
// js:"devicePixelRatio"
func (*Window) DevicePixelRatio() (devicePixelRatio float32) {
	js.Rewrite("$_.devicePixelRatio")
	return devicePixelRatio
}

// Document prop
// js:"document"
func (*Window) Document() (document Document) {
	js.Rewrite("$_.document")
	return document
}

// DoNotTrack prop
// js:"doNotTrack"
func (*Window) DoNotTrack() (doNotTrack string) {
	js.Rewrite("$_.doNotTrack")
	return doNotTrack
}

// Event prop
// js:"event"
func (*Window) Event() (event Event) {
	js.Rewrite("$_.event")
	return event
}

// SetEvent prop
// js:"event"
func (*Window) SetEvent(event Event) {
	js.Rewrite("$_.event = $1", event)
}

// External prop
// js:"external"
func (*Window) External() (external *external.External) {
	js.Rewrite("$_.external")
	return external
}

// FrameElement prop
// js:"frameElement"
func (*Window) FrameElement() (frameElement Element) {
	js.Rewrite("$_.frameElement")
	return frameElement
}

// Frames prop
// js:"frames"
func (*Window) Frames() (frames *Window) {
	js.Rewrite("$_.frames")
	return frames
}

// History prop
// js:"history"
func (*Window) History() (history *history.History) {
	js.Rewrite("$_.history")
	return history
}

// InnerHeight prop
// js:"innerHeight"
func (*Window) InnerHeight() (innerHeight int) {
	js.Rewrite("$_.innerHeight")
	return innerHeight
}

// InnerWidth prop
// js:"innerWidth"
func (*Window) InnerWidth() (innerWidth int) {
	js.Rewrite("$_.innerWidth")
	return innerWidth
}

// IsSecureContext prop
// js:"isSecureContext"
func (*Window) IsSecureContext() (isSecureContext bool) {
	js.Rewrite("$_.isSecureContext")
	return isSecureContext
}

// Length prop
// js:"length"
func (*Window) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}

// Location prop
// js:"location"
func (*Window) Location() (location *location.Location) {
	js.Rewrite("$_.location")
	return location
}

// Locationbar prop
// js:"locationbar"
func (*Window) Locationbar() (locationbar *barprop.BarProp) {
	js.Rewrite("$_.locationbar")
	return locationbar
}

// Menubar prop
// js:"menubar"
func (*Window) Menubar() (menubar *barprop.BarProp) {
	js.Rewrite("$_.menubar")
	return menubar
}

// MsContentScript prop
// js:"msContentScript"
func (*Window) MsContentScript() (msContentScript *extensionscriptapis.ExtensionScriptApis) {
	js.Rewrite("$_.msContentScript")
	return msContentScript
}

// MsCredentials prop
// js:"msCredentials"
func (*Window) MsCredentials() (msCredentials *mscredentials.MSCredentials) {
	js.Rewrite("$_.msCredentials")
	return msCredentials
}

// Name prop
// js:"name"
func (*Window) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// SetName prop
// js:"name"
func (*Window) SetName(name string) {
	js.Rewrite("$_.name = $1", name)
}

// Navigator prop
// js:"navigator"
func (*Window) Navigator() (navigator *Navigator) {
	js.Rewrite("$_.navigator")
	return navigator
}

// OffscreenBuffering prop
// js:"offscreenBuffering"
func (*Window) OffscreenBuffering() (offscreenBuffering interface{}) {
	js.Rewrite("$_.offscreenBuffering")
	return offscreenBuffering
}

// SetOffscreenBuffering prop
// js:"offscreenBuffering"
func (*Window) SetOffscreenBuffering(offscreenBuffering interface{}) {
	js.Rewrite("$_.offscreenBuffering = $1", offscreenBuffering)
}

// Onabort prop
// js:"onabort"
func (*Window) Onabort() (onabort func(Event)) {
	js.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*Window) SetOnabort(onabort func(Event)) {
	js.Rewrite("$_.onabort = $1", onabort)
}

// Onafterprint prop
// js:"onafterprint"
func (*Window) Onafterprint() (onafterprint func(Event)) {
	js.Rewrite("$_.onafterprint")
	return onafterprint
}

// SetOnafterprint prop
// js:"onafterprint"
func (*Window) SetOnafterprint(onafterprint func(Event)) {
	js.Rewrite("$_.onafterprint = $1", onafterprint)
}

// Onbeforeprint prop
// js:"onbeforeprint"
func (*Window) Onbeforeprint() (onbeforeprint func(Event)) {
	js.Rewrite("$_.onbeforeprint")
	return onbeforeprint
}

// SetOnbeforeprint prop
// js:"onbeforeprint"
func (*Window) SetOnbeforeprint(onbeforeprint func(Event)) {
	js.Rewrite("$_.onbeforeprint = $1", onbeforeprint)
}

// Onbeforeunload prop
// js:"onbeforeunload"
func (*Window) Onbeforeunload() (onbeforeunload func(*BeforeUnloadEvent)) {
	js.Rewrite("$_.onbeforeunload")
	return onbeforeunload
}

// SetOnbeforeunload prop
// js:"onbeforeunload"
func (*Window) SetOnbeforeunload(onbeforeunload func(*BeforeUnloadEvent)) {
	js.Rewrite("$_.onbeforeunload = $1", onbeforeunload)
}

// Onblur prop
// js:"onblur"
func (*Window) Onblur() (onblur func(*FocusEvent)) {
	js.Rewrite("$_.onblur")
	return onblur
}

// SetOnblur prop
// js:"onblur"
func (*Window) SetOnblur(onblur func(*FocusEvent)) {
	js.Rewrite("$_.onblur = $1", onblur)
}

// Oncanplay prop
// js:"oncanplay"
func (*Window) Oncanplay() (oncanplay func(Event)) {
	js.Rewrite("$_.oncanplay")
	return oncanplay
}

// SetOncanplay prop
// js:"oncanplay"
func (*Window) SetOncanplay(oncanplay func(Event)) {
	js.Rewrite("$_.oncanplay = $1", oncanplay)
}

// Oncanplaythrough prop
// js:"oncanplaythrough"
func (*Window) Oncanplaythrough() (oncanplaythrough func(Event)) {
	js.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

// SetOncanplaythrough prop
// js:"oncanplaythrough"
func (*Window) SetOncanplaythrough(oncanplaythrough func(Event)) {
	js.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

// Onchange prop
// js:"onchange"
func (*Window) Onchange() (onchange func(Event)) {
	js.Rewrite("$_.onchange")
	return onchange
}

// SetOnchange prop
// js:"onchange"
func (*Window) SetOnchange(onchange func(Event)) {
	js.Rewrite("$_.onchange = $1", onchange)
}

// Onclick prop
// js:"onclick"
func (*Window) Onclick() (onclick func(Event)) {
	js.Rewrite("$_.onclick")
	return onclick
}

// SetOnclick prop
// js:"onclick"
func (*Window) SetOnclick(onclick func(Event)) {
	js.Rewrite("$_.onclick = $1", onclick)
}

// Oncompassneedscalibration prop
// js:"oncompassneedscalibration"
func (*Window) Oncompassneedscalibration() (oncompassneedscalibration func(Event)) {
	js.Rewrite("$_.oncompassneedscalibration")
	return oncompassneedscalibration
}

// SetOncompassneedscalibration prop
// js:"oncompassneedscalibration"
func (*Window) SetOncompassneedscalibration(oncompassneedscalibration func(Event)) {
	js.Rewrite("$_.oncompassneedscalibration = $1", oncompassneedscalibration)
}

// Oncontextmenu prop
// js:"oncontextmenu"
func (*Window) Oncontextmenu() (oncontextmenu func(Event)) {
	js.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

// SetOncontextmenu prop
// js:"oncontextmenu"
func (*Window) SetOncontextmenu(oncontextmenu func(Event)) {
	js.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

// Ondblclick prop
// js:"ondblclick"
func (*Window) Ondblclick() (ondblclick func(Event)) {
	js.Rewrite("$_.ondblclick")
	return ondblclick
}

// SetOndblclick prop
// js:"ondblclick"
func (*Window) SetOndblclick(ondblclick func(Event)) {
	js.Rewrite("$_.ondblclick = $1", ondblclick)
}

// Ondevicelight prop
// js:"ondevicelight"
func (*Window) Ondevicelight() (ondevicelight func(*DeviceLightEvent)) {
	js.Rewrite("$_.ondevicelight")
	return ondevicelight
}

// SetOndevicelight prop
// js:"ondevicelight"
func (*Window) SetOndevicelight(ondevicelight func(*DeviceLightEvent)) {
	js.Rewrite("$_.ondevicelight = $1", ondevicelight)
}

// Ondevicemotion prop
// js:"ondevicemotion"
func (*Window) Ondevicemotion() (ondevicemotion func(*DeviceMotionEvent)) {
	js.Rewrite("$_.ondevicemotion")
	return ondevicemotion
}

// SetOndevicemotion prop
// js:"ondevicemotion"
func (*Window) SetOndevicemotion(ondevicemotion func(*DeviceMotionEvent)) {
	js.Rewrite("$_.ondevicemotion = $1", ondevicemotion)
}

// Ondeviceorientation prop
// js:"ondeviceorientation"
func (*Window) Ondeviceorientation() (ondeviceorientation func(*DeviceOrientationEvent)) {
	js.Rewrite("$_.ondeviceorientation")
	return ondeviceorientation
}

// SetOndeviceorientation prop
// js:"ondeviceorientation"
func (*Window) SetOndeviceorientation(ondeviceorientation func(*DeviceOrientationEvent)) {
	js.Rewrite("$_.ondeviceorientation = $1", ondeviceorientation)
}

// Ondrag prop
// js:"ondrag"
func (*Window) Ondrag() (ondrag func(Event)) {
	js.Rewrite("$_.ondrag")
	return ondrag
}

// SetOndrag prop
// js:"ondrag"
func (*Window) SetOndrag(ondrag func(Event)) {
	js.Rewrite("$_.ondrag = $1", ondrag)
}

// Ondragend prop
// js:"ondragend"
func (*Window) Ondragend() (ondragend func(Event)) {
	js.Rewrite("$_.ondragend")
	return ondragend
}

// SetOndragend prop
// js:"ondragend"
func (*Window) SetOndragend(ondragend func(Event)) {
	js.Rewrite("$_.ondragend = $1", ondragend)
}

// Ondragenter prop
// js:"ondragenter"
func (*Window) Ondragenter() (ondragenter func(Event)) {
	js.Rewrite("$_.ondragenter")
	return ondragenter
}

// SetOndragenter prop
// js:"ondragenter"
func (*Window) SetOndragenter(ondragenter func(Event)) {
	js.Rewrite("$_.ondragenter = $1", ondragenter)
}

// Ondragleave prop
// js:"ondragleave"
func (*Window) Ondragleave() (ondragleave func(Event)) {
	js.Rewrite("$_.ondragleave")
	return ondragleave
}

// SetOndragleave prop
// js:"ondragleave"
func (*Window) SetOndragleave(ondragleave func(Event)) {
	js.Rewrite("$_.ondragleave = $1", ondragleave)
}

// Ondragover prop
// js:"ondragover"
func (*Window) Ondragover() (ondragover func(Event)) {
	js.Rewrite("$_.ondragover")
	return ondragover
}

// SetOndragover prop
// js:"ondragover"
func (*Window) SetOndragover(ondragover func(Event)) {
	js.Rewrite("$_.ondragover = $1", ondragover)
}

// Ondragstart prop
// js:"ondragstart"
func (*Window) Ondragstart() (ondragstart func(Event)) {
	js.Rewrite("$_.ondragstart")
	return ondragstart
}

// SetOndragstart prop
// js:"ondragstart"
func (*Window) SetOndragstart(ondragstart func(Event)) {
	js.Rewrite("$_.ondragstart = $1", ondragstart)
}

// Ondrop prop
// js:"ondrop"
func (*Window) Ondrop() (ondrop func(Event)) {
	js.Rewrite("$_.ondrop")
	return ondrop
}

// SetOndrop prop
// js:"ondrop"
func (*Window) SetOndrop(ondrop func(Event)) {
	js.Rewrite("$_.ondrop = $1", ondrop)
}

// Ondurationchange prop
// js:"ondurationchange"
func (*Window) Ondurationchange() (ondurationchange func(Event)) {
	js.Rewrite("$_.ondurationchange")
	return ondurationchange
}

// SetOndurationchange prop
// js:"ondurationchange"
func (*Window) SetOndurationchange(ondurationchange func(Event)) {
	js.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

// Onemptied prop
// js:"onemptied"
func (*Window) Onemptied() (onemptied func(Event)) {
	js.Rewrite("$_.onemptied")
	return onemptied
}

// SetOnemptied prop
// js:"onemptied"
func (*Window) SetOnemptied(onemptied func(Event)) {
	js.Rewrite("$_.onemptied = $1", onemptied)
}

// Onended prop
// js:"onended"
func (*Window) Onended() (onended func(Event)) {
	js.Rewrite("$_.onended")
	return onended
}

// SetOnended prop
// js:"onended"
func (*Window) SetOnended(onended func(Event)) {
	js.Rewrite("$_.onended = $1", onended)
}

// Onerror prop
// js:"onerror"
func (*Window) Onerror() (onerror func(columnNumber *uint, event interface{}, fileno *uint, source *string)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*Window) SetOnerror(onerror func(columnNumber *uint, event interface{}, fileno *uint, source *string)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Onfocus prop
// js:"onfocus"
func (*Window) Onfocus() (onfocus func(*FocusEvent)) {
	js.Rewrite("$_.onfocus")
	return onfocus
}

// SetOnfocus prop
// js:"onfocus"
func (*Window) SetOnfocus(onfocus func(*FocusEvent)) {
	js.Rewrite("$_.onfocus = $1", onfocus)
}

// Onhashchange prop
// js:"onhashchange"
func (*Window) Onhashchange() (onhashchange func(*HashChangeEvent)) {
	js.Rewrite("$_.onhashchange")
	return onhashchange
}

// SetOnhashchange prop
// js:"onhashchange"
func (*Window) SetOnhashchange(onhashchange func(*HashChangeEvent)) {
	js.Rewrite("$_.onhashchange = $1", onhashchange)
}

// Oninput prop
// js:"oninput"
func (*Window) Oninput() (oninput func(Event)) {
	js.Rewrite("$_.oninput")
	return oninput
}

// SetOninput prop
// js:"oninput"
func (*Window) SetOninput(oninput func(Event)) {
	js.Rewrite("$_.oninput = $1", oninput)
}

// Oninvalid prop
// js:"oninvalid"
func (*Window) Oninvalid() (oninvalid func(Event)) {
	js.Rewrite("$_.oninvalid")
	return oninvalid
}

// SetOninvalid prop
// js:"oninvalid"
func (*Window) SetOninvalid(oninvalid func(Event)) {
	js.Rewrite("$_.oninvalid = $1", oninvalid)
}

// Onkeydown prop
// js:"onkeydown"
func (*Window) Onkeydown() (onkeydown func(Event)) {
	js.Rewrite("$_.onkeydown")
	return onkeydown
}

// SetOnkeydown prop
// js:"onkeydown"
func (*Window) SetOnkeydown(onkeydown func(Event)) {
	js.Rewrite("$_.onkeydown = $1", onkeydown)
}

// Onkeypress prop
// js:"onkeypress"
func (*Window) Onkeypress() (onkeypress func(Event)) {
	js.Rewrite("$_.onkeypress")
	return onkeypress
}

// SetOnkeypress prop
// js:"onkeypress"
func (*Window) SetOnkeypress(onkeypress func(Event)) {
	js.Rewrite("$_.onkeypress = $1", onkeypress)
}

// Onkeyup prop
// js:"onkeyup"
func (*Window) Onkeyup() (onkeyup func(Event)) {
	js.Rewrite("$_.onkeyup")
	return onkeyup
}

// SetOnkeyup prop
// js:"onkeyup"
func (*Window) SetOnkeyup(onkeyup func(Event)) {
	js.Rewrite("$_.onkeyup = $1", onkeyup)
}

// Onload prop
// js:"onload"
func (*Window) Onload() (onload func(Event)) {
	js.Rewrite("$_.onload")
	return onload
}

// SetOnload prop
// js:"onload"
func (*Window) SetOnload(onload func(Event)) {
	js.Rewrite("$_.onload = $1", onload)
}

// Onloadeddata prop
// js:"onloadeddata"
func (*Window) Onloadeddata() (onloadeddata func(Event)) {
	js.Rewrite("$_.onloadeddata")
	return onloadeddata
}

// SetOnloadeddata prop
// js:"onloadeddata"
func (*Window) SetOnloadeddata(onloadeddata func(Event)) {
	js.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

// Onloadedmetadata prop
// js:"onloadedmetadata"
func (*Window) Onloadedmetadata() (onloadedmetadata func(Event)) {
	js.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

// SetOnloadedmetadata prop
// js:"onloadedmetadata"
func (*Window) SetOnloadedmetadata(onloadedmetadata func(Event)) {
	js.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

// Onloadstart prop
// js:"onloadstart"
func (*Window) Onloadstart() (onloadstart func(Event)) {
	js.Rewrite("$_.onloadstart")
	return onloadstart
}

// SetOnloadstart prop
// js:"onloadstart"
func (*Window) SetOnloadstart(onloadstart func(Event)) {
	js.Rewrite("$_.onloadstart = $1", onloadstart)
}

// Onmessage prop
// js:"onmessage"
func (*Window) Onmessage() (onmessage func(*MessageEvent)) {
	js.Rewrite("$_.onmessage")
	return onmessage
}

// SetOnmessage prop
// js:"onmessage"
func (*Window) SetOnmessage(onmessage func(*MessageEvent)) {
	js.Rewrite("$_.onmessage = $1", onmessage)
}

// Onmousedown prop
// js:"onmousedown"
func (*Window) Onmousedown() (onmousedown func(Event)) {
	js.Rewrite("$_.onmousedown")
	return onmousedown
}

// SetOnmousedown prop
// js:"onmousedown"
func (*Window) SetOnmousedown(onmousedown func(Event)) {
	js.Rewrite("$_.onmousedown = $1", onmousedown)
}

// Onmouseenter prop
// js:"onmouseenter"
func (*Window) Onmouseenter() (onmouseenter func(Event)) {
	js.Rewrite("$_.onmouseenter")
	return onmouseenter
}

// SetOnmouseenter prop
// js:"onmouseenter"
func (*Window) SetOnmouseenter(onmouseenter func(Event)) {
	js.Rewrite("$_.onmouseenter = $1", onmouseenter)
}

// Onmouseleave prop
// js:"onmouseleave"
func (*Window) Onmouseleave() (onmouseleave func(Event)) {
	js.Rewrite("$_.onmouseleave")
	return onmouseleave
}

// SetOnmouseleave prop
// js:"onmouseleave"
func (*Window) SetOnmouseleave(onmouseleave func(Event)) {
	js.Rewrite("$_.onmouseleave = $1", onmouseleave)
}

// Onmousemove prop
// js:"onmousemove"
func (*Window) Onmousemove() (onmousemove func(Event)) {
	js.Rewrite("$_.onmousemove")
	return onmousemove
}

// SetOnmousemove prop
// js:"onmousemove"
func (*Window) SetOnmousemove(onmousemove func(Event)) {
	js.Rewrite("$_.onmousemove = $1", onmousemove)
}

// Onmouseout prop
// js:"onmouseout"
func (*Window) Onmouseout() (onmouseout func(Event)) {
	js.Rewrite("$_.onmouseout")
	return onmouseout
}

// SetOnmouseout prop
// js:"onmouseout"
func (*Window) SetOnmouseout(onmouseout func(Event)) {
	js.Rewrite("$_.onmouseout = $1", onmouseout)
}

// Onmouseover prop
// js:"onmouseover"
func (*Window) Onmouseover() (onmouseover func(Event)) {
	js.Rewrite("$_.onmouseover")
	return onmouseover
}

// SetOnmouseover prop
// js:"onmouseover"
func (*Window) SetOnmouseover(onmouseover func(Event)) {
	js.Rewrite("$_.onmouseover = $1", onmouseover)
}

// Onmouseup prop
// js:"onmouseup"
func (*Window) Onmouseup() (onmouseup func(Event)) {
	js.Rewrite("$_.onmouseup")
	return onmouseup
}

// SetOnmouseup prop
// js:"onmouseup"
func (*Window) SetOnmouseup(onmouseup func(Event)) {
	js.Rewrite("$_.onmouseup = $1", onmouseup)
}

// Onmousewheel prop
// js:"onmousewheel"
func (*Window) Onmousewheel() (onmousewheel func(Event)) {
	js.Rewrite("$_.onmousewheel")
	return onmousewheel
}

// SetOnmousewheel prop
// js:"onmousewheel"
func (*Window) SetOnmousewheel(onmousewheel func(Event)) {
	js.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*Window) Onmsgesturechange() (onmsgesturechange func(Event)) {
	js.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*Window) SetOnmsgesturechange(onmsgesturechange func(Event)) {
	js.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*Window) Onmsgesturedoubletap() (onmsgesturedoubletap func(Event)) {
	js.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*Window) SetOnmsgesturedoubletap(onmsgesturedoubletap func(Event)) {
	js.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*Window) Onmsgestureend() (onmsgestureend func(Event)) {
	js.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*Window) SetOnmsgestureend(onmsgestureend func(Event)) {
	js.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*Window) Onmsgesturehold() (onmsgesturehold func(Event)) {
	js.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*Window) SetOnmsgesturehold(onmsgesturehold func(Event)) {
	js.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*Window) Onmsgesturestart() (onmsgesturestart func(Event)) {
	js.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*Window) SetOnmsgesturestart(onmsgesturestart func(Event)) {
	js.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*Window) Onmsgesturetap() (onmsgesturetap func(Event)) {
	js.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*Window) SetOnmsgesturetap(onmsgesturetap func(Event)) {
	js.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*Window) Onmsinertiastart() (onmsinertiastart func(Event)) {
	js.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*Window) SetOnmsinertiastart(onmsinertiastart func(Event)) {
	js.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*Window) Onmspointercancel() (onmspointercancel func(Event)) {
	js.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*Window) SetOnmspointercancel(onmspointercancel func(Event)) {
	js.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*Window) Onmspointerdown() (onmspointerdown func(Event)) {
	js.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*Window) SetOnmspointerdown(onmspointerdown func(Event)) {
	js.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*Window) Onmspointerenter() (onmspointerenter func(Event)) {
	js.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*Window) SetOnmspointerenter(onmspointerenter func(Event)) {
	js.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*Window) Onmspointerleave() (onmspointerleave func(Event)) {
	js.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*Window) SetOnmspointerleave(onmspointerleave func(Event)) {
	js.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*Window) Onmspointermove() (onmspointermove func(Event)) {
	js.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*Window) SetOnmspointermove(onmspointermove func(Event)) {
	js.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*Window) Onmspointerout() (onmspointerout func(Event)) {
	js.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*Window) SetOnmspointerout(onmspointerout func(Event)) {
	js.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*Window) Onmspointerover() (onmspointerover func(Event)) {
	js.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*Window) SetOnmspointerover(onmspointerover func(Event)) {
	js.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*Window) Onmspointerup() (onmspointerup func(Event)) {
	js.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*Window) SetOnmspointerup(onmspointerup func(Event)) {
	js.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Onoffline prop
// js:"onoffline"
func (*Window) Onoffline() (onoffline func(Event)) {
	js.Rewrite("$_.onoffline")
	return onoffline
}

// SetOnoffline prop
// js:"onoffline"
func (*Window) SetOnoffline(onoffline func(Event)) {
	js.Rewrite("$_.onoffline = $1", onoffline)
}

// Ononline prop
// js:"ononline"
func (*Window) Ononline() (ononline func(Event)) {
	js.Rewrite("$_.ononline")
	return ononline
}

// SetOnonline prop
// js:"ononline"
func (*Window) SetOnonline(ononline func(Event)) {
	js.Rewrite("$_.ononline = $1", ononline)
}

// Onorientationchange prop
// js:"onorientationchange"
func (*Window) Onorientationchange() (onorientationchange func(Event)) {
	js.Rewrite("$_.onorientationchange")
	return onorientationchange
}

// SetOnorientationchange prop
// js:"onorientationchange"
func (*Window) SetOnorientationchange(onorientationchange func(Event)) {
	js.Rewrite("$_.onorientationchange = $1", onorientationchange)
}

// Onpagehide prop
// js:"onpagehide"
func (*Window) Onpagehide() (onpagehide func(*PageTransitionEvent)) {
	js.Rewrite("$_.onpagehide")
	return onpagehide
}

// SetOnpagehide prop
// js:"onpagehide"
func (*Window) SetOnpagehide(onpagehide func(*PageTransitionEvent)) {
	js.Rewrite("$_.onpagehide = $1", onpagehide)
}

// Onpageshow prop
// js:"onpageshow"
func (*Window) Onpageshow() (onpageshow func(*PageTransitionEvent)) {
	js.Rewrite("$_.onpageshow")
	return onpageshow
}

// SetOnpageshow prop
// js:"onpageshow"
func (*Window) SetOnpageshow(onpageshow func(*PageTransitionEvent)) {
	js.Rewrite("$_.onpageshow = $1", onpageshow)
}

// Onpause prop
// js:"onpause"
func (*Window) Onpause() (onpause func(Event)) {
	js.Rewrite("$_.onpause")
	return onpause
}

// SetOnpause prop
// js:"onpause"
func (*Window) SetOnpause(onpause func(Event)) {
	js.Rewrite("$_.onpause = $1", onpause)
}

// Onplay prop
// js:"onplay"
func (*Window) Onplay() (onplay func(Event)) {
	js.Rewrite("$_.onplay")
	return onplay
}

// SetOnplay prop
// js:"onplay"
func (*Window) SetOnplay(onplay func(Event)) {
	js.Rewrite("$_.onplay = $1", onplay)
}

// Onplaying prop
// js:"onplaying"
func (*Window) Onplaying() (onplaying func(Event)) {
	js.Rewrite("$_.onplaying")
	return onplaying
}

// SetOnplaying prop
// js:"onplaying"
func (*Window) SetOnplaying(onplaying func(Event)) {
	js.Rewrite("$_.onplaying = $1", onplaying)
}

// Onpopstate prop
// js:"onpopstate"
func (*Window) Onpopstate() (onpopstate func(*PopStateEvent)) {
	js.Rewrite("$_.onpopstate")
	return onpopstate
}

// SetOnpopstate prop
// js:"onpopstate"
func (*Window) SetOnpopstate(onpopstate func(*PopStateEvent)) {
	js.Rewrite("$_.onpopstate = $1", onpopstate)
}

// Onprogress prop
// js:"onprogress"
func (*Window) Onprogress() (onprogress func(Event)) {
	js.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop
// js:"onprogress"
func (*Window) SetOnprogress(onprogress func(Event)) {
	js.Rewrite("$_.onprogress = $1", onprogress)
}

// Onratechange prop
// js:"onratechange"
func (*Window) Onratechange() (onratechange func(Event)) {
	js.Rewrite("$_.onratechange")
	return onratechange
}

// SetOnratechange prop
// js:"onratechange"
func (*Window) SetOnratechange(onratechange func(Event)) {
	js.Rewrite("$_.onratechange = $1", onratechange)
}

// Onreadystatechange prop
// js:"onreadystatechange"
func (*Window) Onreadystatechange() (onreadystatechange func(Event)) {
	js.Rewrite("$_.onreadystatechange")
	return onreadystatechange
}

// SetOnreadystatechange prop
// js:"onreadystatechange"
func (*Window) SetOnreadystatechange(onreadystatechange func(Event)) {
	js.Rewrite("$_.onreadystatechange = $1", onreadystatechange)
}

// Onreset prop
// js:"onreset"
func (*Window) Onreset() (onreset func(Event)) {
	js.Rewrite("$_.onreset")
	return onreset
}

// SetOnreset prop
// js:"onreset"
func (*Window) SetOnreset(onreset func(Event)) {
	js.Rewrite("$_.onreset = $1", onreset)
}

// Onresize prop
// js:"onresize"
func (*Window) Onresize() (onresize func(UIEvent)) {
	js.Rewrite("$_.onresize")
	return onresize
}

// SetOnresize prop
// js:"onresize"
func (*Window) SetOnresize(onresize func(UIEvent)) {
	js.Rewrite("$_.onresize = $1", onresize)
}

// Onscroll prop
// js:"onscroll"
func (*Window) Onscroll() (onscroll func(Event)) {
	js.Rewrite("$_.onscroll")
	return onscroll
}

// SetOnscroll prop
// js:"onscroll"
func (*Window) SetOnscroll(onscroll func(Event)) {
	js.Rewrite("$_.onscroll = $1", onscroll)
}

// Onseeked prop
// js:"onseeked"
func (*Window) Onseeked() (onseeked func(Event)) {
	js.Rewrite("$_.onseeked")
	return onseeked
}

// SetOnseeked prop
// js:"onseeked"
func (*Window) SetOnseeked(onseeked func(Event)) {
	js.Rewrite("$_.onseeked = $1", onseeked)
}

// Onseeking prop
// js:"onseeking"
func (*Window) Onseeking() (onseeking func(Event)) {
	js.Rewrite("$_.onseeking")
	return onseeking
}

// SetOnseeking prop
// js:"onseeking"
func (*Window) SetOnseeking(onseeking func(Event)) {
	js.Rewrite("$_.onseeking = $1", onseeking)
}

// Onselect prop
// js:"onselect"
func (*Window) Onselect() (onselect func(Event)) {
	js.Rewrite("$_.onselect")
	return onselect
}

// SetOnselect prop
// js:"onselect"
func (*Window) SetOnselect(onselect func(Event)) {
	js.Rewrite("$_.onselect = $1", onselect)
}

// Onstalled prop
// js:"onstalled"
func (*Window) Onstalled() (onstalled func(Event)) {
	js.Rewrite("$_.onstalled")
	return onstalled
}

// SetOnstalled prop
// js:"onstalled"
func (*Window) SetOnstalled(onstalled func(Event)) {
	js.Rewrite("$_.onstalled = $1", onstalled)
}

// Onstorage prop
// js:"onstorage"
func (*Window) Onstorage() (onstorage func(*StorageEvent)) {
	js.Rewrite("$_.onstorage")
	return onstorage
}

// SetOnstorage prop
// js:"onstorage"
func (*Window) SetOnstorage(onstorage func(*StorageEvent)) {
	js.Rewrite("$_.onstorage = $1", onstorage)
}

// Onsubmit prop
// js:"onsubmit"
func (*Window) Onsubmit() (onsubmit func(Event)) {
	js.Rewrite("$_.onsubmit")
	return onsubmit
}

// SetOnsubmit prop
// js:"onsubmit"
func (*Window) SetOnsubmit(onsubmit func(Event)) {
	js.Rewrite("$_.onsubmit = $1", onsubmit)
}

// Onsuspend prop
// js:"onsuspend"
func (*Window) Onsuspend() (onsuspend func(Event)) {
	js.Rewrite("$_.onsuspend")
	return onsuspend
}

// SetOnsuspend prop
// js:"onsuspend"
func (*Window) SetOnsuspend(onsuspend func(Event)) {
	js.Rewrite("$_.onsuspend = $1", onsuspend)
}

// Ontimeupdate prop
// js:"ontimeupdate"
func (*Window) Ontimeupdate() (ontimeupdate func(Event)) {
	js.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

// SetOntimeupdate prop
// js:"ontimeupdate"
func (*Window) SetOntimeupdate(ontimeupdate func(Event)) {
	js.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*Window) Ontouchcancel() (ontouchcancel interface{}) {
	js.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*Window) SetOntouchcancel(ontouchcancel interface{}) {
	js.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*Window) Ontouchend() (ontouchend interface{}) {
	js.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*Window) SetOntouchend(ontouchend interface{}) {
	js.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*Window) Ontouchmove() (ontouchmove interface{}) {
	js.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*Window) SetOntouchmove(ontouchmove interface{}) {
	js.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*Window) Ontouchstart() (ontouchstart interface{}) {
	js.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*Window) SetOntouchstart(ontouchstart interface{}) {
	js.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onunload prop
// js:"onunload"
func (*Window) Onunload() (onunload func(Event)) {
	js.Rewrite("$_.onunload")
	return onunload
}

// SetOnunload prop
// js:"onunload"
func (*Window) SetOnunload(onunload func(Event)) {
	js.Rewrite("$_.onunload = $1", onunload)
}

// Onvolumechange prop
// js:"onvolumechange"
func (*Window) Onvolumechange() (onvolumechange func(Event)) {
	js.Rewrite("$_.onvolumechange")
	return onvolumechange
}

// SetOnvolumechange prop
// js:"onvolumechange"
func (*Window) SetOnvolumechange(onvolumechange func(Event)) {
	js.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

// Onwaiting prop
// js:"onwaiting"
func (*Window) Onwaiting() (onwaiting func(Event)) {
	js.Rewrite("$_.onwaiting")
	return onwaiting
}

// SetOnwaiting prop
// js:"onwaiting"
func (*Window) SetOnwaiting(onwaiting func(Event)) {
	js.Rewrite("$_.onwaiting = $1", onwaiting)
}

// Opener prop
// js:"opener"
func (*Window) Opener() (opener *Window) {
	js.Rewrite("$_.opener")
	return opener
}

// Orientation prop
// js:"orientation"
func (*Window) Orientation() (orientation string) {
	js.Rewrite("$_.orientation")
	return orientation
}

// OuterHeight prop
// js:"outerHeight"
func (*Window) OuterHeight() (outerHeight int) {
	js.Rewrite("$_.outerHeight")
	return outerHeight
}

// OuterWidth prop
// js:"outerWidth"
func (*Window) OuterWidth() (outerWidth int) {
	js.Rewrite("$_.outerWidth")
	return outerWidth
}

// PageXOffset prop
// js:"pageXOffset"
func (*Window) PageXOffset() (pageXOffset int) {
	js.Rewrite("$_.pageXOffset")
	return pageXOffset
}

// PageYOffset prop
// js:"pageYOffset"
func (*Window) PageYOffset() (pageYOffset int) {
	js.Rewrite("$_.pageYOffset")
	return pageYOffset
}

// Parent prop
// js:"parent"
func (*Window) Parent() (parent *Window) {
	js.Rewrite("$_.parent")
	return parent
}

// Performance prop
// js:"performance"
func (*Window) Performance() (performance *performance.Performance) {
	js.Rewrite("$_.performance")
	return performance
}

// Personalbar prop
// js:"personalbar"
func (*Window) Personalbar() (personalbar *barprop.BarProp) {
	js.Rewrite("$_.personalbar")
	return personalbar
}

// Screen prop
// js:"screen"
func (*Window) Screen() (screen *Screen) {
	js.Rewrite("$_.screen")
	return screen
}

// ScreenLeft prop
// js:"screenLeft"
func (*Window) ScreenLeft() (screenLeft int) {
	js.Rewrite("$_.screenLeft")
	return screenLeft
}

// ScreenTop prop
// js:"screenTop"
func (*Window) ScreenTop() (screenTop int) {
	js.Rewrite("$_.screenTop")
	return screenTop
}

// ScreenX prop
// js:"screenX"
func (*Window) ScreenX() (screenX int) {
	js.Rewrite("$_.screenX")
	return screenX
}

// ScreenY prop
// js:"screenY"
func (*Window) ScreenY() (screenY int) {
	js.Rewrite("$_.screenY")
	return screenY
}

// Scrollbars prop
// js:"scrollbars"
func (*Window) Scrollbars() (scrollbars *barprop.BarProp) {
	js.Rewrite("$_.scrollbars")
	return scrollbars
}

// ScrollX prop
// js:"scrollX"
func (*Window) ScrollX() (scrollX int) {
	js.Rewrite("$_.scrollX")
	return scrollX
}

// ScrollY prop
// js:"scrollY"
func (*Window) ScrollY() (scrollY int) {
	js.Rewrite("$_.scrollY")
	return scrollY
}

// Self prop
// js:"self"
func (*Window) Self() (self *Window) {
	js.Rewrite("$_.self")
	return self
}

// SpeechSynthesis prop
// js:"speechSynthesis"
func (*Window) SpeechSynthesis() (speechSynthesis *SpeechSynthesis) {
	js.Rewrite("$_.speechSynthesis")
	return speechSynthesis
}

// Status prop
// js:"status"
func (*Window) Status() (status string) {
	js.Rewrite("$_.status")
	return status
}

// SetStatus prop
// js:"status"
func (*Window) SetStatus(status string) {
	js.Rewrite("$_.status = $1", status)
}

// Statusbar prop
// js:"statusbar"
func (*Window) Statusbar() (statusbar *barprop.BarProp) {
	js.Rewrite("$_.statusbar")
	return statusbar
}

// StyleMedia prop
// js:"styleMedia"
func (*Window) StyleMedia() (styleMedia *stylemedia.StyleMedia) {
	js.Rewrite("$_.styleMedia")
	return styleMedia
}

// Toolbar prop
// js:"toolbar"
func (*Window) Toolbar() (toolbar *barprop.BarProp) {
	js.Rewrite("$_.toolbar")
	return toolbar
}

// Top prop
// js:"top"
func (*Window) Top() (top *Window) {
	js.Rewrite("$_.top")
	return top
}

// Window prop
// js:"window"
func (*Window) Window() (window *Window) {
	js.Rewrite("$_.window")
	return window
}

// SessionStorage prop
// js:"sessionStorage"
func (*Window) SessionStorage() (sessionStorage *storage.Storage) {
	js.Rewrite("$_.sessionStorage")
	return sessionStorage
}

// LocalStorage prop
// js:"localStorage"
func (*Window) LocalStorage() (localStorage *storage.Storage) {
	js.Rewrite("$_.localStorage")
	return localStorage
}

// Console prop
// js:"console"
func (*Window) Console() (console *Console) {
	js.Rewrite("$_.console")
	return console
}

// Onpointercancel prop
// js:"onpointercancel"
func (*Window) Onpointercancel() (onpointercancel func(Event)) {
	js.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*Window) SetOnpointercancel(onpointercancel func(Event)) {
	js.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*Window) Onpointerdown() (onpointerdown func(Event)) {
	js.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*Window) SetOnpointerdown(onpointerdown func(Event)) {
	js.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*Window) Onpointerenter() (onpointerenter func(Event)) {
	js.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*Window) SetOnpointerenter(onpointerenter func(Event)) {
	js.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*Window) Onpointerleave() (onpointerleave func(Event)) {
	js.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*Window) SetOnpointerleave(onpointerleave func(Event)) {
	js.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*Window) Onpointermove() (onpointermove func(Event)) {
	js.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*Window) SetOnpointermove(onpointermove func(Event)) {
	js.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*Window) Onpointerout() (onpointerout func(Event)) {
	js.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*Window) SetOnpointerout(onpointerout func(Event)) {
	js.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*Window) Onpointerover() (onpointerover func(Event)) {
	js.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*Window) SetOnpointerover(onpointerover func(Event)) {
	js.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*Window) Onpointerup() (onpointerup func(Event)) {
	js.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*Window) SetOnpointerup(onpointerup func(Event)) {
	js.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*Window) Onwheel() (onwheel func(Event)) {
	js.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*Window) SetOnwheel(onwheel func(Event)) {
	js.Rewrite("$_.onwheel = $1", onwheel)
}

// IndexedDB prop
// js:"indexedDB"
func (*Window) IndexedDB() (indexedDB *IDBFactory) {
	js.Rewrite("$_.indexedDB")
	return indexedDB
}
