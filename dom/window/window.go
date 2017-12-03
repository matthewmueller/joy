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
	"github.com/matthewmueller/joy/macro"
)

var _ EventTarget = (*Window)(nil)

// New fn
func New() *Window {
	macro.Rewrite("window")
	return &Window{}
}

// Window struct
// js:"Window,omit"
type Window struct {
}

// Alert fn
// js:"alert"
func (*Window) Alert(message *string) {
	macro.Rewrite("$_.alert($1)", message)
}

// Blur fn
// js:"blur"
func (*Window) Blur() {
	macro.Rewrite("$_.blur()")
}

// CancelAnimationFrame fn
// js:"cancelAnimationFrame"
func (*Window) CancelAnimationFrame(handle int) {
	macro.Rewrite("$_.cancelAnimationFrame($1)", handle)
}

// CaptureEvents fn
// js:"captureEvents"
func (*Window) CaptureEvents() {
	macro.Rewrite("$_.captureEvents()")
}

// Close fn
// js:"close"
func (*Window) Close() {
	macro.Rewrite("$_.close()")
}

// Confirm fn
// js:"confirm"
func (*Window) Confirm(message *string) (b bool) {
	macro.Rewrite("$_.confirm($1)", message)
	return b
}

// DepartFocus fn
// js:"departFocus"
func (*Window) DepartFocus(navigationReason *navigationreason.NavigationReason, origin *focusnavigationorigin.FocusNavigationOrigin) {
	macro.Rewrite("$_.departFocus($1, $2)", navigationReason, origin)
}

// Focus fn
// js:"focus"
func (*Window) Focus() {
	macro.Rewrite("$_.focus()")
}

// GetComputedStyle fn
// js:"getComputedStyle"
func (*Window) GetComputedStyle(elt Element, pseudoElt *string) (c *CSSStyleDeclaration) {
	macro.Rewrite("$_.getComputedStyle($1, $2)", elt, pseudoElt)
	return c
}

// GetMatchedCSSRules fn
// js:"getMatchedCSSRules"
func (*Window) GetMatchedCSSRules(elt Element, pseudoElt *string) (c *CSSRuleList) {
	macro.Rewrite("$_.getMatchedCSSRules($1, $2)", elt, pseudoElt)
	return c
}

// GetSelection fn
// js:"getSelection"
func (*Window) GetSelection() (s *Selection) {
	macro.Rewrite("$_.getSelection()")
	return s
}

// MatchMedia fn
// js:"matchMedia"
func (*Window) MatchMedia(mediaQuery string) (m *mediaquery.MediaQueryList) {
	macro.Rewrite("$_.matchMedia($1)", mediaQuery)
	return m
}

// MoveBy fn
// js:"moveBy"
func (*Window) MoveBy(x *int, y *int) {
	macro.Rewrite("$_.moveBy($1, $2)", x, y)
}

// MoveTo fn
// js:"moveTo"
func (*Window) MoveTo(x *int, y *int) {
	macro.Rewrite("$_.moveTo($1, $2)", x, y)
}

// MsWriteProfilerMark fn
// js:"msWriteProfilerMark"
func (*Window) MsWriteProfilerMark(profilerMarkName string) {
	macro.Rewrite("$_.msWriteProfilerMark($1)", profilerMarkName)
}

// Open fn
// js:"open"
func (*Window) Open(url *string, target *string, features *string, replace *bool) (w *Window) {
	macro.Rewrite("$_.open($1, $2, $3, $4)", url, target, features, replace)
	return w
}

// PostMessage fn
// js:"postMessage"
func (*Window) PostMessage(message interface{}, targetOrigin string, transfer *[]interface{}) {
	macro.Rewrite("$_.postMessage($1, $2, $3)", message, targetOrigin, transfer)
}

// Print fn
// js:"print"
func (*Window) Print() {
	macro.Rewrite("$_.print()")
}

// Prompt fn
// js:"prompt"
func (*Window) Prompt(message *string, def *string) (s string) {
	macro.Rewrite("$_.prompt($1, $2)", message, def)
	return s
}

// ReleaseEvents fn
// js:"releaseEvents"
func (*Window) ReleaseEvents() {
	macro.Rewrite("$_.releaseEvents()")
}

// RequestAnimationFrame fn
// js:"requestAnimationFrame"
func (*Window) RequestAnimationFrame(callback func(time int)) (i int) {
	macro.Rewrite("$_.requestAnimationFrame($1)", callback)
	return i
}

// ResizeBy fn
// js:"resizeBy"
func (*Window) ResizeBy(x *int, y *int) {
	macro.Rewrite("$_.resizeBy($1, $2)", x, y)
}

// ResizeTo fn
// js:"resizeTo"
func (*Window) ResizeTo(x *int, y *int) {
	macro.Rewrite("$_.resizeTo($1, $2)", x, y)
}

// Scroll fn
// js:"scroll"
func (*Window) Scroll(x *int, y *int) {
	macro.Rewrite("$_.scroll($1, $2)", x, y)
}

// ScrollBy fn
// js:"scrollBy"
func (*Window) ScrollBy(x *int, y *int) {
	macro.Rewrite("$_.scrollBy($1, $2)", x, y)
}

// ScrollTo fn
// js:"scrollTo"
func (*Window) ScrollTo(x *int, y *int) {
	macro.Rewrite("$_.scrollTo($1, $2)", x, y)
}

// Stop fn
// js:"stop"
func (*Window) Stop() {
	macro.Rewrite("$_.stop()")
}

// WebkitCancelAnimationFrame fn
// js:"webkitCancelAnimationFrame"
func (*Window) WebkitCancelAnimationFrame(handle int) {
	macro.Rewrite("$_.webkitCancelAnimationFrame($1)", handle)
}

// WebkitConvertPointFromNodeToPage fn
// js:"webkitConvertPointFromNodeToPage"
func (*Window) WebkitConvertPointFromNodeToPage(node Node, pt *webkitpoint.WebKitPoint) (w *webkitpoint.WebKitPoint) {
	macro.Rewrite("$_.webkitConvertPointFromNodeToPage($1, $2)", node, pt)
	return w
}

// WebkitConvertPointFromPageToNode fn
// js:"webkitConvertPointFromPageToNode"
func (*Window) WebkitConvertPointFromPageToNode(node Node, pt *webkitpoint.WebKitPoint) (w *webkitpoint.WebKitPoint) {
	macro.Rewrite("$_.webkitConvertPointFromPageToNode($1, $2)", node, pt)
	return w
}

// WebkitRequestAnimationFrame fn
// js:"webkitRequestAnimationFrame"
func (*Window) WebkitRequestAnimationFrame(callback func(time int)) (i int) {
	macro.Rewrite("$_.webkitRequestAnimationFrame($1)", callback)
	return i
}

// ClearInterval fn
// js:"clearInterval"
func (*Window) ClearInterval(handle int) {
	macro.Rewrite("$_.clearInterval($1)", handle)
}

// ClearTimeout fn
// js:"clearTimeout"
func (*Window) ClearTimeout(handle int) {
	macro.Rewrite("$_.clearTimeout($1)", handle)
}

// SetInterval fn
// js:"setInterval"
func (*Window) SetInterval(handler interface{}, timeout *interface{}, args interface{}) (i int) {
	macro.Rewrite("$_.setInterval($1, $2, $3)", handler, timeout, args)
	return i
}

// SetTimeout fn
// js:"setTimeout"
func (*Window) SetTimeout(handler interface{}, timeout *interface{}, args interface{}) (i int) {
	macro.Rewrite("$_.setTimeout($1, $2, $3)", handler, timeout, args)
	return i
}

// Atob fn
// js:"atob"
func (*Window) Atob(encodedString string) (s string) {
	macro.Rewrite("$_.atob($1)", encodedString)
	return s
}

// Btoa fn
// js:"btoa"
func (*Window) Btoa(rawString string) (s string) {
	macro.Rewrite("$_.btoa($1)", rawString)
	return s
}

// Fetch fn
// js:"fetch"
func (*Window) Fetch(input *request.Request, init *requestinit.RequestInit) (r *response.Response) {
	macro.Rewrite("await $_.fetch($1, $2)", input, init)
	return r
}

// AddEventListener fn
// js:"addEventListener"
func (*Window) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*Window) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*Window) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// ApplicationCache prop
// js:"applicationCache"
func (*Window) ApplicationCache() (applicationCache *ApplicationCache) {
	macro.Rewrite("$_.applicationCache")
	return applicationCache
}

// Caches prop
// js:"caches"
func (*Window) Caches() (caches *cachestorage.CacheStorage) {
	macro.Rewrite("$_.caches")
	return caches
}

// ClientInformation prop
// js:"clientInformation"
func (*Window) ClientInformation() (clientInformation *Navigator) {
	macro.Rewrite("$_.clientInformation")
	return clientInformation
}

// Closed prop
// js:"closed"
func (*Window) Closed() (closed bool) {
	macro.Rewrite("$_.closed")
	return closed
}

// Crypto prop
// js:"crypto"
func (*Window) Crypto() (crypto *crypto.Crypto) {
	macro.Rewrite("$_.crypto")
	return crypto
}

// DefaultStatus prop
// js:"defaultStatus"
func (*Window) DefaultStatus() (defaultStatus string) {
	macro.Rewrite("$_.defaultStatus")
	return defaultStatus
}

// SetDefaultStatus prop
// js:"defaultStatus"
func (*Window) SetDefaultStatus(defaultStatus string) {
	macro.Rewrite("$_.defaultStatus = $1", defaultStatus)
}

// DevicePixelRatio prop
// js:"devicePixelRatio"
func (*Window) DevicePixelRatio() (devicePixelRatio float32) {
	macro.Rewrite("$_.devicePixelRatio")
	return devicePixelRatio
}

// Document prop
// js:"document"
func (*Window) Document() (document Document) {
	macro.Rewrite("$_.document")
	return document
}

// DoNotTrack prop
// js:"doNotTrack"
func (*Window) DoNotTrack() (doNotTrack string) {
	macro.Rewrite("$_.doNotTrack")
	return doNotTrack
}

// Event prop
// js:"event"
func (*Window) Event() (event Event) {
	macro.Rewrite("$_.event")
	return event
}

// SetEvent prop
// js:"event"
func (*Window) SetEvent(event Event) {
	macro.Rewrite("$_.event = $1", event)
}

// External prop
// js:"external"
func (*Window) External() (external *external.External) {
	macro.Rewrite("$_.external")
	return external
}

// FrameElement prop
// js:"frameElement"
func (*Window) FrameElement() (frameElement Element) {
	macro.Rewrite("$_.frameElement")
	return frameElement
}

// Frames prop
// js:"frames"
func (*Window) Frames() (frames *Window) {
	macro.Rewrite("$_.frames")
	return frames
}

// History prop
// js:"history"
func (*Window) History() (history *history.History) {
	macro.Rewrite("$_.history")
	return history
}

// InnerHeight prop
// js:"innerHeight"
func (*Window) InnerHeight() (innerHeight int) {
	macro.Rewrite("$_.innerHeight")
	return innerHeight
}

// InnerWidth prop
// js:"innerWidth"
func (*Window) InnerWidth() (innerWidth int) {
	macro.Rewrite("$_.innerWidth")
	return innerWidth
}

// IsSecureContext prop
// js:"isSecureContext"
func (*Window) IsSecureContext() (isSecureContext bool) {
	macro.Rewrite("$_.isSecureContext")
	return isSecureContext
}

// Length prop
// js:"length"
func (*Window) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

// Location prop
// js:"location"
func (*Window) Location() (location *location.Location) {
	macro.Rewrite("$_.location")
	return location
}

// Locationbar prop
// js:"locationbar"
func (*Window) Locationbar() (locationbar *barprop.BarProp) {
	macro.Rewrite("$_.locationbar")
	return locationbar
}

// Menubar prop
// js:"menubar"
func (*Window) Menubar() (menubar *barprop.BarProp) {
	macro.Rewrite("$_.menubar")
	return menubar
}

// MsContentScript prop
// js:"msContentScript"
func (*Window) MsContentScript() (msContentScript *extensionscriptapis.ExtensionScriptApis) {
	macro.Rewrite("$_.msContentScript")
	return msContentScript
}

// MsCredentials prop
// js:"msCredentials"
func (*Window) MsCredentials() (msCredentials *mscredentials.MSCredentials) {
	macro.Rewrite("$_.msCredentials")
	return msCredentials
}

// Name prop
// js:"name"
func (*Window) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

// SetName prop
// js:"name"
func (*Window) SetName(name string) {
	macro.Rewrite("$_.name = $1", name)
}

// Navigator prop
// js:"navigator"
func (*Window) Navigator() (navigator *Navigator) {
	macro.Rewrite("$_.navigator")
	return navigator
}

// OffscreenBuffering prop
// js:"offscreenBuffering"
func (*Window) OffscreenBuffering() (offscreenBuffering interface{}) {
	macro.Rewrite("$_.offscreenBuffering")
	return offscreenBuffering
}

// SetOffscreenBuffering prop
// js:"offscreenBuffering"
func (*Window) SetOffscreenBuffering(offscreenBuffering interface{}) {
	macro.Rewrite("$_.offscreenBuffering = $1", offscreenBuffering)
}

// Onabort prop
// js:"onabort"
func (*Window) Onabort() (onabort func(Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*Window) SetOnabort(onabort func(Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

// Onafterprint prop
// js:"onafterprint"
func (*Window) Onafterprint() (onafterprint func(Event)) {
	macro.Rewrite("$_.onafterprint")
	return onafterprint
}

// SetOnafterprint prop
// js:"onafterprint"
func (*Window) SetOnafterprint(onafterprint func(Event)) {
	macro.Rewrite("$_.onafterprint = $1", onafterprint)
}

// Onbeforeprint prop
// js:"onbeforeprint"
func (*Window) Onbeforeprint() (onbeforeprint func(Event)) {
	macro.Rewrite("$_.onbeforeprint")
	return onbeforeprint
}

// SetOnbeforeprint prop
// js:"onbeforeprint"
func (*Window) SetOnbeforeprint(onbeforeprint func(Event)) {
	macro.Rewrite("$_.onbeforeprint = $1", onbeforeprint)
}

// Onbeforeunload prop
// js:"onbeforeunload"
func (*Window) Onbeforeunload() (onbeforeunload func(*BeforeUnloadEvent)) {
	macro.Rewrite("$_.onbeforeunload")
	return onbeforeunload
}

// SetOnbeforeunload prop
// js:"onbeforeunload"
func (*Window) SetOnbeforeunload(onbeforeunload func(*BeforeUnloadEvent)) {
	macro.Rewrite("$_.onbeforeunload = $1", onbeforeunload)
}

// Onblur prop
// js:"onblur"
func (*Window) Onblur() (onblur func(*FocusEvent)) {
	macro.Rewrite("$_.onblur")
	return onblur
}

// SetOnblur prop
// js:"onblur"
func (*Window) SetOnblur(onblur func(*FocusEvent)) {
	macro.Rewrite("$_.onblur = $1", onblur)
}

// Oncanplay prop
// js:"oncanplay"
func (*Window) Oncanplay() (oncanplay func(Event)) {
	macro.Rewrite("$_.oncanplay")
	return oncanplay
}

// SetOncanplay prop
// js:"oncanplay"
func (*Window) SetOncanplay(oncanplay func(Event)) {
	macro.Rewrite("$_.oncanplay = $1", oncanplay)
}

// Oncanplaythrough prop
// js:"oncanplaythrough"
func (*Window) Oncanplaythrough() (oncanplaythrough func(Event)) {
	macro.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

// SetOncanplaythrough prop
// js:"oncanplaythrough"
func (*Window) SetOncanplaythrough(oncanplaythrough func(Event)) {
	macro.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

// Onchange prop
// js:"onchange"
func (*Window) Onchange() (onchange func(Event)) {
	macro.Rewrite("$_.onchange")
	return onchange
}

// SetOnchange prop
// js:"onchange"
func (*Window) SetOnchange(onchange func(Event)) {
	macro.Rewrite("$_.onchange = $1", onchange)
}

// Onclick prop
// js:"onclick"
func (*Window) Onclick() (onclick func(Event)) {
	macro.Rewrite("$_.onclick")
	return onclick
}

// SetOnclick prop
// js:"onclick"
func (*Window) SetOnclick(onclick func(Event)) {
	macro.Rewrite("$_.onclick = $1", onclick)
}

// Oncompassneedscalibration prop
// js:"oncompassneedscalibration"
func (*Window) Oncompassneedscalibration() (oncompassneedscalibration func(Event)) {
	macro.Rewrite("$_.oncompassneedscalibration")
	return oncompassneedscalibration
}

// SetOncompassneedscalibration prop
// js:"oncompassneedscalibration"
func (*Window) SetOncompassneedscalibration(oncompassneedscalibration func(Event)) {
	macro.Rewrite("$_.oncompassneedscalibration = $1", oncompassneedscalibration)
}

// Oncontextmenu prop
// js:"oncontextmenu"
func (*Window) Oncontextmenu() (oncontextmenu func(Event)) {
	macro.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

// SetOncontextmenu prop
// js:"oncontextmenu"
func (*Window) SetOncontextmenu(oncontextmenu func(Event)) {
	macro.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

// Ondblclick prop
// js:"ondblclick"
func (*Window) Ondblclick() (ondblclick func(Event)) {
	macro.Rewrite("$_.ondblclick")
	return ondblclick
}

// SetOndblclick prop
// js:"ondblclick"
func (*Window) SetOndblclick(ondblclick func(Event)) {
	macro.Rewrite("$_.ondblclick = $1", ondblclick)
}

// Ondevicelight prop
// js:"ondevicelight"
func (*Window) Ondevicelight() (ondevicelight func(*DeviceLightEvent)) {
	macro.Rewrite("$_.ondevicelight")
	return ondevicelight
}

// SetOndevicelight prop
// js:"ondevicelight"
func (*Window) SetOndevicelight(ondevicelight func(*DeviceLightEvent)) {
	macro.Rewrite("$_.ondevicelight = $1", ondevicelight)
}

// Ondevicemotion prop
// js:"ondevicemotion"
func (*Window) Ondevicemotion() (ondevicemotion func(*DeviceMotionEvent)) {
	macro.Rewrite("$_.ondevicemotion")
	return ondevicemotion
}

// SetOndevicemotion prop
// js:"ondevicemotion"
func (*Window) SetOndevicemotion(ondevicemotion func(*DeviceMotionEvent)) {
	macro.Rewrite("$_.ondevicemotion = $1", ondevicemotion)
}

// Ondeviceorientation prop
// js:"ondeviceorientation"
func (*Window) Ondeviceorientation() (ondeviceorientation func(*DeviceOrientationEvent)) {
	macro.Rewrite("$_.ondeviceorientation")
	return ondeviceorientation
}

// SetOndeviceorientation prop
// js:"ondeviceorientation"
func (*Window) SetOndeviceorientation(ondeviceorientation func(*DeviceOrientationEvent)) {
	macro.Rewrite("$_.ondeviceorientation = $1", ondeviceorientation)
}

// Ondrag prop
// js:"ondrag"
func (*Window) Ondrag() (ondrag func(Event)) {
	macro.Rewrite("$_.ondrag")
	return ondrag
}

// SetOndrag prop
// js:"ondrag"
func (*Window) SetOndrag(ondrag func(Event)) {
	macro.Rewrite("$_.ondrag = $1", ondrag)
}

// Ondragend prop
// js:"ondragend"
func (*Window) Ondragend() (ondragend func(Event)) {
	macro.Rewrite("$_.ondragend")
	return ondragend
}

// SetOndragend prop
// js:"ondragend"
func (*Window) SetOndragend(ondragend func(Event)) {
	macro.Rewrite("$_.ondragend = $1", ondragend)
}

// Ondragenter prop
// js:"ondragenter"
func (*Window) Ondragenter() (ondragenter func(Event)) {
	macro.Rewrite("$_.ondragenter")
	return ondragenter
}

// SetOndragenter prop
// js:"ondragenter"
func (*Window) SetOndragenter(ondragenter func(Event)) {
	macro.Rewrite("$_.ondragenter = $1", ondragenter)
}

// Ondragleave prop
// js:"ondragleave"
func (*Window) Ondragleave() (ondragleave func(Event)) {
	macro.Rewrite("$_.ondragleave")
	return ondragleave
}

// SetOndragleave prop
// js:"ondragleave"
func (*Window) SetOndragleave(ondragleave func(Event)) {
	macro.Rewrite("$_.ondragleave = $1", ondragleave)
}

// Ondragover prop
// js:"ondragover"
func (*Window) Ondragover() (ondragover func(Event)) {
	macro.Rewrite("$_.ondragover")
	return ondragover
}

// SetOndragover prop
// js:"ondragover"
func (*Window) SetOndragover(ondragover func(Event)) {
	macro.Rewrite("$_.ondragover = $1", ondragover)
}

// Ondragstart prop
// js:"ondragstart"
func (*Window) Ondragstart() (ondragstart func(Event)) {
	macro.Rewrite("$_.ondragstart")
	return ondragstart
}

// SetOndragstart prop
// js:"ondragstart"
func (*Window) SetOndragstart(ondragstart func(Event)) {
	macro.Rewrite("$_.ondragstart = $1", ondragstart)
}

// Ondrop prop
// js:"ondrop"
func (*Window) Ondrop() (ondrop func(Event)) {
	macro.Rewrite("$_.ondrop")
	return ondrop
}

// SetOndrop prop
// js:"ondrop"
func (*Window) SetOndrop(ondrop func(Event)) {
	macro.Rewrite("$_.ondrop = $1", ondrop)
}

// Ondurationchange prop
// js:"ondurationchange"
func (*Window) Ondurationchange() (ondurationchange func(Event)) {
	macro.Rewrite("$_.ondurationchange")
	return ondurationchange
}

// SetOndurationchange prop
// js:"ondurationchange"
func (*Window) SetOndurationchange(ondurationchange func(Event)) {
	macro.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

// Onemptied prop
// js:"onemptied"
func (*Window) Onemptied() (onemptied func(Event)) {
	macro.Rewrite("$_.onemptied")
	return onemptied
}

// SetOnemptied prop
// js:"onemptied"
func (*Window) SetOnemptied(onemptied func(Event)) {
	macro.Rewrite("$_.onemptied = $1", onemptied)
}

// Onended prop
// js:"onended"
func (*Window) Onended() (onended func(Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

// SetOnended prop
// js:"onended"
func (*Window) SetOnended(onended func(Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

// Onerror prop
// js:"onerror"
func (*Window) Onerror() (onerror func(columnNumber *uint, event interface{}, fileno *uint, source *string)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*Window) SetOnerror(onerror func(columnNumber *uint, event interface{}, fileno *uint, source *string)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// Onfocus prop
// js:"onfocus"
func (*Window) Onfocus() (onfocus func(*FocusEvent)) {
	macro.Rewrite("$_.onfocus")
	return onfocus
}

// SetOnfocus prop
// js:"onfocus"
func (*Window) SetOnfocus(onfocus func(*FocusEvent)) {
	macro.Rewrite("$_.onfocus = $1", onfocus)
}

// Onhashchange prop
// js:"onhashchange"
func (*Window) Onhashchange() (onhashchange func(*HashChangeEvent)) {
	macro.Rewrite("$_.onhashchange")
	return onhashchange
}

// SetOnhashchange prop
// js:"onhashchange"
func (*Window) SetOnhashchange(onhashchange func(*HashChangeEvent)) {
	macro.Rewrite("$_.onhashchange = $1", onhashchange)
}

// Oninput prop
// js:"oninput"
func (*Window) Oninput() (oninput func(Event)) {
	macro.Rewrite("$_.oninput")
	return oninput
}

// SetOninput prop
// js:"oninput"
func (*Window) SetOninput(oninput func(Event)) {
	macro.Rewrite("$_.oninput = $1", oninput)
}

// Oninvalid prop
// js:"oninvalid"
func (*Window) Oninvalid() (oninvalid func(Event)) {
	macro.Rewrite("$_.oninvalid")
	return oninvalid
}

// SetOninvalid prop
// js:"oninvalid"
func (*Window) SetOninvalid(oninvalid func(Event)) {
	macro.Rewrite("$_.oninvalid = $1", oninvalid)
}

// Onkeydown prop
// js:"onkeydown"
func (*Window) Onkeydown() (onkeydown func(Event)) {
	macro.Rewrite("$_.onkeydown")
	return onkeydown
}

// SetOnkeydown prop
// js:"onkeydown"
func (*Window) SetOnkeydown(onkeydown func(Event)) {
	macro.Rewrite("$_.onkeydown = $1", onkeydown)
}

// Onkeypress prop
// js:"onkeypress"
func (*Window) Onkeypress() (onkeypress func(Event)) {
	macro.Rewrite("$_.onkeypress")
	return onkeypress
}

// SetOnkeypress prop
// js:"onkeypress"
func (*Window) SetOnkeypress(onkeypress func(Event)) {
	macro.Rewrite("$_.onkeypress = $1", onkeypress)
}

// Onkeyup prop
// js:"onkeyup"
func (*Window) Onkeyup() (onkeyup func(Event)) {
	macro.Rewrite("$_.onkeyup")
	return onkeyup
}

// SetOnkeyup prop
// js:"onkeyup"
func (*Window) SetOnkeyup(onkeyup func(Event)) {
	macro.Rewrite("$_.onkeyup = $1", onkeyup)
}

// Onload prop
// js:"onload"
func (*Window) Onload() (onload func(Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

// SetOnload prop
// js:"onload"
func (*Window) SetOnload(onload func(Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

// Onloadeddata prop
// js:"onloadeddata"
func (*Window) Onloadeddata() (onloadeddata func(Event)) {
	macro.Rewrite("$_.onloadeddata")
	return onloadeddata
}

// SetOnloadeddata prop
// js:"onloadeddata"
func (*Window) SetOnloadeddata(onloadeddata func(Event)) {
	macro.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

// Onloadedmetadata prop
// js:"onloadedmetadata"
func (*Window) Onloadedmetadata() (onloadedmetadata func(Event)) {
	macro.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

// SetOnloadedmetadata prop
// js:"onloadedmetadata"
func (*Window) SetOnloadedmetadata(onloadedmetadata func(Event)) {
	macro.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

// Onloadstart prop
// js:"onloadstart"
func (*Window) Onloadstart() (onloadstart func(Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

// SetOnloadstart prop
// js:"onloadstart"
func (*Window) SetOnloadstart(onloadstart func(Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

// Onmessage prop
// js:"onmessage"
func (*Window) Onmessage() (onmessage func(*MessageEvent)) {
	macro.Rewrite("$_.onmessage")
	return onmessage
}

// SetOnmessage prop
// js:"onmessage"
func (*Window) SetOnmessage(onmessage func(*MessageEvent)) {
	macro.Rewrite("$_.onmessage = $1", onmessage)
}

// Onmousedown prop
// js:"onmousedown"
func (*Window) Onmousedown() (onmousedown func(Event)) {
	macro.Rewrite("$_.onmousedown")
	return onmousedown
}

// SetOnmousedown prop
// js:"onmousedown"
func (*Window) SetOnmousedown(onmousedown func(Event)) {
	macro.Rewrite("$_.onmousedown = $1", onmousedown)
}

// Onmouseenter prop
// js:"onmouseenter"
func (*Window) Onmouseenter() (onmouseenter func(Event)) {
	macro.Rewrite("$_.onmouseenter")
	return onmouseenter
}

// SetOnmouseenter prop
// js:"onmouseenter"
func (*Window) SetOnmouseenter(onmouseenter func(Event)) {
	macro.Rewrite("$_.onmouseenter = $1", onmouseenter)
}

// Onmouseleave prop
// js:"onmouseleave"
func (*Window) Onmouseleave() (onmouseleave func(Event)) {
	macro.Rewrite("$_.onmouseleave")
	return onmouseleave
}

// SetOnmouseleave prop
// js:"onmouseleave"
func (*Window) SetOnmouseleave(onmouseleave func(Event)) {
	macro.Rewrite("$_.onmouseleave = $1", onmouseleave)
}

// Onmousemove prop
// js:"onmousemove"
func (*Window) Onmousemove() (onmousemove func(Event)) {
	macro.Rewrite("$_.onmousemove")
	return onmousemove
}

// SetOnmousemove prop
// js:"onmousemove"
func (*Window) SetOnmousemove(onmousemove func(Event)) {
	macro.Rewrite("$_.onmousemove = $1", onmousemove)
}

// Onmouseout prop
// js:"onmouseout"
func (*Window) Onmouseout() (onmouseout func(Event)) {
	macro.Rewrite("$_.onmouseout")
	return onmouseout
}

// SetOnmouseout prop
// js:"onmouseout"
func (*Window) SetOnmouseout(onmouseout func(Event)) {
	macro.Rewrite("$_.onmouseout = $1", onmouseout)
}

// Onmouseover prop
// js:"onmouseover"
func (*Window) Onmouseover() (onmouseover func(Event)) {
	macro.Rewrite("$_.onmouseover")
	return onmouseover
}

// SetOnmouseover prop
// js:"onmouseover"
func (*Window) SetOnmouseover(onmouseover func(Event)) {
	macro.Rewrite("$_.onmouseover = $1", onmouseover)
}

// Onmouseup prop
// js:"onmouseup"
func (*Window) Onmouseup() (onmouseup func(Event)) {
	macro.Rewrite("$_.onmouseup")
	return onmouseup
}

// SetOnmouseup prop
// js:"onmouseup"
func (*Window) SetOnmouseup(onmouseup func(Event)) {
	macro.Rewrite("$_.onmouseup = $1", onmouseup)
}

// Onmousewheel prop
// js:"onmousewheel"
func (*Window) Onmousewheel() (onmousewheel func(Event)) {
	macro.Rewrite("$_.onmousewheel")
	return onmousewheel
}

// SetOnmousewheel prop
// js:"onmousewheel"
func (*Window) SetOnmousewheel(onmousewheel func(Event)) {
	macro.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*Window) Onmsgesturechange() (onmsgesturechange func(Event)) {
	macro.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*Window) SetOnmsgesturechange(onmsgesturechange func(Event)) {
	macro.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*Window) Onmsgesturedoubletap() (onmsgesturedoubletap func(Event)) {
	macro.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*Window) SetOnmsgesturedoubletap(onmsgesturedoubletap func(Event)) {
	macro.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*Window) Onmsgestureend() (onmsgestureend func(Event)) {
	macro.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*Window) SetOnmsgestureend(onmsgestureend func(Event)) {
	macro.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*Window) Onmsgesturehold() (onmsgesturehold func(Event)) {
	macro.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*Window) SetOnmsgesturehold(onmsgesturehold func(Event)) {
	macro.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*Window) Onmsgesturestart() (onmsgesturestart func(Event)) {
	macro.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*Window) SetOnmsgesturestart(onmsgesturestart func(Event)) {
	macro.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*Window) Onmsgesturetap() (onmsgesturetap func(Event)) {
	macro.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*Window) SetOnmsgesturetap(onmsgesturetap func(Event)) {
	macro.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*Window) Onmsinertiastart() (onmsinertiastart func(Event)) {
	macro.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*Window) SetOnmsinertiastart(onmsinertiastart func(Event)) {
	macro.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*Window) Onmspointercancel() (onmspointercancel func(Event)) {
	macro.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*Window) SetOnmspointercancel(onmspointercancel func(Event)) {
	macro.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*Window) Onmspointerdown() (onmspointerdown func(Event)) {
	macro.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*Window) SetOnmspointerdown(onmspointerdown func(Event)) {
	macro.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*Window) Onmspointerenter() (onmspointerenter func(Event)) {
	macro.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*Window) SetOnmspointerenter(onmspointerenter func(Event)) {
	macro.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*Window) Onmspointerleave() (onmspointerleave func(Event)) {
	macro.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*Window) SetOnmspointerleave(onmspointerleave func(Event)) {
	macro.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*Window) Onmspointermove() (onmspointermove func(Event)) {
	macro.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*Window) SetOnmspointermove(onmspointermove func(Event)) {
	macro.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*Window) Onmspointerout() (onmspointerout func(Event)) {
	macro.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*Window) SetOnmspointerout(onmspointerout func(Event)) {
	macro.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*Window) Onmspointerover() (onmspointerover func(Event)) {
	macro.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*Window) SetOnmspointerover(onmspointerover func(Event)) {
	macro.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*Window) Onmspointerup() (onmspointerup func(Event)) {
	macro.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*Window) SetOnmspointerup(onmspointerup func(Event)) {
	macro.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Onoffline prop
// js:"onoffline"
func (*Window) Onoffline() (onoffline func(Event)) {
	macro.Rewrite("$_.onoffline")
	return onoffline
}

// SetOnoffline prop
// js:"onoffline"
func (*Window) SetOnoffline(onoffline func(Event)) {
	macro.Rewrite("$_.onoffline = $1", onoffline)
}

// Ononline prop
// js:"ononline"
func (*Window) Ononline() (ononline func(Event)) {
	macro.Rewrite("$_.ononline")
	return ononline
}

// SetOnonline prop
// js:"ononline"
func (*Window) SetOnonline(ononline func(Event)) {
	macro.Rewrite("$_.ononline = $1", ononline)
}

// Onorientationchange prop
// js:"onorientationchange"
func (*Window) Onorientationchange() (onorientationchange func(Event)) {
	macro.Rewrite("$_.onorientationchange")
	return onorientationchange
}

// SetOnorientationchange prop
// js:"onorientationchange"
func (*Window) SetOnorientationchange(onorientationchange func(Event)) {
	macro.Rewrite("$_.onorientationchange = $1", onorientationchange)
}

// Onpagehide prop
// js:"onpagehide"
func (*Window) Onpagehide() (onpagehide func(*PageTransitionEvent)) {
	macro.Rewrite("$_.onpagehide")
	return onpagehide
}

// SetOnpagehide prop
// js:"onpagehide"
func (*Window) SetOnpagehide(onpagehide func(*PageTransitionEvent)) {
	macro.Rewrite("$_.onpagehide = $1", onpagehide)
}

// Onpageshow prop
// js:"onpageshow"
func (*Window) Onpageshow() (onpageshow func(*PageTransitionEvent)) {
	macro.Rewrite("$_.onpageshow")
	return onpageshow
}

// SetOnpageshow prop
// js:"onpageshow"
func (*Window) SetOnpageshow(onpageshow func(*PageTransitionEvent)) {
	macro.Rewrite("$_.onpageshow = $1", onpageshow)
}

// Onpause prop
// js:"onpause"
func (*Window) Onpause() (onpause func(Event)) {
	macro.Rewrite("$_.onpause")
	return onpause
}

// SetOnpause prop
// js:"onpause"
func (*Window) SetOnpause(onpause func(Event)) {
	macro.Rewrite("$_.onpause = $1", onpause)
}

// Onplay prop
// js:"onplay"
func (*Window) Onplay() (onplay func(Event)) {
	macro.Rewrite("$_.onplay")
	return onplay
}

// SetOnplay prop
// js:"onplay"
func (*Window) SetOnplay(onplay func(Event)) {
	macro.Rewrite("$_.onplay = $1", onplay)
}

// Onplaying prop
// js:"onplaying"
func (*Window) Onplaying() (onplaying func(Event)) {
	macro.Rewrite("$_.onplaying")
	return onplaying
}

// SetOnplaying prop
// js:"onplaying"
func (*Window) SetOnplaying(onplaying func(Event)) {
	macro.Rewrite("$_.onplaying = $1", onplaying)
}

// Onpopstate prop
// js:"onpopstate"
func (*Window) Onpopstate() (onpopstate func(*PopStateEvent)) {
	macro.Rewrite("$_.onpopstate")
	return onpopstate
}

// SetOnpopstate prop
// js:"onpopstate"
func (*Window) SetOnpopstate(onpopstate func(*PopStateEvent)) {
	macro.Rewrite("$_.onpopstate = $1", onpopstate)
}

// Onprogress prop
// js:"onprogress"
func (*Window) Onprogress() (onprogress func(Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop
// js:"onprogress"
func (*Window) SetOnprogress(onprogress func(Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

// Onratechange prop
// js:"onratechange"
func (*Window) Onratechange() (onratechange func(Event)) {
	macro.Rewrite("$_.onratechange")
	return onratechange
}

// SetOnratechange prop
// js:"onratechange"
func (*Window) SetOnratechange(onratechange func(Event)) {
	macro.Rewrite("$_.onratechange = $1", onratechange)
}

// Onreadystatechange prop
// js:"onreadystatechange"
func (*Window) Onreadystatechange() (onreadystatechange func(Event)) {
	macro.Rewrite("$_.onreadystatechange")
	return onreadystatechange
}

// SetOnreadystatechange prop
// js:"onreadystatechange"
func (*Window) SetOnreadystatechange(onreadystatechange func(Event)) {
	macro.Rewrite("$_.onreadystatechange = $1", onreadystatechange)
}

// Onreset prop
// js:"onreset"
func (*Window) Onreset() (onreset func(Event)) {
	macro.Rewrite("$_.onreset")
	return onreset
}

// SetOnreset prop
// js:"onreset"
func (*Window) SetOnreset(onreset func(Event)) {
	macro.Rewrite("$_.onreset = $1", onreset)
}

// Onresize prop
// js:"onresize"
func (*Window) Onresize() (onresize func(UIEvent)) {
	macro.Rewrite("$_.onresize")
	return onresize
}

// SetOnresize prop
// js:"onresize"
func (*Window) SetOnresize(onresize func(UIEvent)) {
	macro.Rewrite("$_.onresize = $1", onresize)
}

// Onscroll prop
// js:"onscroll"
func (*Window) Onscroll() (onscroll func(Event)) {
	macro.Rewrite("$_.onscroll")
	return onscroll
}

// SetOnscroll prop
// js:"onscroll"
func (*Window) SetOnscroll(onscroll func(Event)) {
	macro.Rewrite("$_.onscroll = $1", onscroll)
}

// Onseeked prop
// js:"onseeked"
func (*Window) Onseeked() (onseeked func(Event)) {
	macro.Rewrite("$_.onseeked")
	return onseeked
}

// SetOnseeked prop
// js:"onseeked"
func (*Window) SetOnseeked(onseeked func(Event)) {
	macro.Rewrite("$_.onseeked = $1", onseeked)
}

// Onseeking prop
// js:"onseeking"
func (*Window) Onseeking() (onseeking func(Event)) {
	macro.Rewrite("$_.onseeking")
	return onseeking
}

// SetOnseeking prop
// js:"onseeking"
func (*Window) SetOnseeking(onseeking func(Event)) {
	macro.Rewrite("$_.onseeking = $1", onseeking)
}

// Onselect prop
// js:"onselect"
func (*Window) Onselect() (onselect func(Event)) {
	macro.Rewrite("$_.onselect")
	return onselect
}

// SetOnselect prop
// js:"onselect"
func (*Window) SetOnselect(onselect func(Event)) {
	macro.Rewrite("$_.onselect = $1", onselect)
}

// Onstalled prop
// js:"onstalled"
func (*Window) Onstalled() (onstalled func(Event)) {
	macro.Rewrite("$_.onstalled")
	return onstalled
}

// SetOnstalled prop
// js:"onstalled"
func (*Window) SetOnstalled(onstalled func(Event)) {
	macro.Rewrite("$_.onstalled = $1", onstalled)
}

// Onstorage prop
// js:"onstorage"
func (*Window) Onstorage() (onstorage func(*StorageEvent)) {
	macro.Rewrite("$_.onstorage")
	return onstorage
}

// SetOnstorage prop
// js:"onstorage"
func (*Window) SetOnstorage(onstorage func(*StorageEvent)) {
	macro.Rewrite("$_.onstorage = $1", onstorage)
}

// Onsubmit prop
// js:"onsubmit"
func (*Window) Onsubmit() (onsubmit func(Event)) {
	macro.Rewrite("$_.onsubmit")
	return onsubmit
}

// SetOnsubmit prop
// js:"onsubmit"
func (*Window) SetOnsubmit(onsubmit func(Event)) {
	macro.Rewrite("$_.onsubmit = $1", onsubmit)
}

// Onsuspend prop
// js:"onsuspend"
func (*Window) Onsuspend() (onsuspend func(Event)) {
	macro.Rewrite("$_.onsuspend")
	return onsuspend
}

// SetOnsuspend prop
// js:"onsuspend"
func (*Window) SetOnsuspend(onsuspend func(Event)) {
	macro.Rewrite("$_.onsuspend = $1", onsuspend)
}

// Ontimeupdate prop
// js:"ontimeupdate"
func (*Window) Ontimeupdate() (ontimeupdate func(Event)) {
	macro.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

// SetOntimeupdate prop
// js:"ontimeupdate"
func (*Window) SetOntimeupdate(ontimeupdate func(Event)) {
	macro.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*Window) Ontouchcancel() (ontouchcancel interface{}) {
	macro.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*Window) SetOntouchcancel(ontouchcancel interface{}) {
	macro.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*Window) Ontouchend() (ontouchend interface{}) {
	macro.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*Window) SetOntouchend(ontouchend interface{}) {
	macro.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*Window) Ontouchmove() (ontouchmove interface{}) {
	macro.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*Window) SetOntouchmove(ontouchmove interface{}) {
	macro.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*Window) Ontouchstart() (ontouchstart interface{}) {
	macro.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*Window) SetOntouchstart(ontouchstart interface{}) {
	macro.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onunload prop
// js:"onunload"
func (*Window) Onunload() (onunload func(Event)) {
	macro.Rewrite("$_.onunload")
	return onunload
}

// SetOnunload prop
// js:"onunload"
func (*Window) SetOnunload(onunload func(Event)) {
	macro.Rewrite("$_.onunload = $1", onunload)
}

// Onvolumechange prop
// js:"onvolumechange"
func (*Window) Onvolumechange() (onvolumechange func(Event)) {
	macro.Rewrite("$_.onvolumechange")
	return onvolumechange
}

// SetOnvolumechange prop
// js:"onvolumechange"
func (*Window) SetOnvolumechange(onvolumechange func(Event)) {
	macro.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

// Onwaiting prop
// js:"onwaiting"
func (*Window) Onwaiting() (onwaiting func(Event)) {
	macro.Rewrite("$_.onwaiting")
	return onwaiting
}

// SetOnwaiting prop
// js:"onwaiting"
func (*Window) SetOnwaiting(onwaiting func(Event)) {
	macro.Rewrite("$_.onwaiting = $1", onwaiting)
}

// Opener prop
// js:"opener"
func (*Window) Opener() (opener *Window) {
	macro.Rewrite("$_.opener")
	return opener
}

// Orientation prop
// js:"orientation"
func (*Window) Orientation() (orientation string) {
	macro.Rewrite("$_.orientation")
	return orientation
}

// OuterHeight prop
// js:"outerHeight"
func (*Window) OuterHeight() (outerHeight int) {
	macro.Rewrite("$_.outerHeight")
	return outerHeight
}

// OuterWidth prop
// js:"outerWidth"
func (*Window) OuterWidth() (outerWidth int) {
	macro.Rewrite("$_.outerWidth")
	return outerWidth
}

// PageXOffset prop
// js:"pageXOffset"
func (*Window) PageXOffset() (pageXOffset int) {
	macro.Rewrite("$_.pageXOffset")
	return pageXOffset
}

// PageYOffset prop
// js:"pageYOffset"
func (*Window) PageYOffset() (pageYOffset int) {
	macro.Rewrite("$_.pageYOffset")
	return pageYOffset
}

// Parent prop
// js:"parent"
func (*Window) Parent() (parent *Window) {
	macro.Rewrite("$_.parent")
	return parent
}

// Performance prop
// js:"performance"
func (*Window) Performance() (performance *performance.Performance) {
	macro.Rewrite("$_.performance")
	return performance
}

// Personalbar prop
// js:"personalbar"
func (*Window) Personalbar() (personalbar *barprop.BarProp) {
	macro.Rewrite("$_.personalbar")
	return personalbar
}

// Screen prop
// js:"screen"
func (*Window) Screen() (screen *Screen) {
	macro.Rewrite("$_.screen")
	return screen
}

// ScreenLeft prop
// js:"screenLeft"
func (*Window) ScreenLeft() (screenLeft int) {
	macro.Rewrite("$_.screenLeft")
	return screenLeft
}

// ScreenTop prop
// js:"screenTop"
func (*Window) ScreenTop() (screenTop int) {
	macro.Rewrite("$_.screenTop")
	return screenTop
}

// ScreenX prop
// js:"screenX"
func (*Window) ScreenX() (screenX int) {
	macro.Rewrite("$_.screenX")
	return screenX
}

// ScreenY prop
// js:"screenY"
func (*Window) ScreenY() (screenY int) {
	macro.Rewrite("$_.screenY")
	return screenY
}

// Scrollbars prop
// js:"scrollbars"
func (*Window) Scrollbars() (scrollbars *barprop.BarProp) {
	macro.Rewrite("$_.scrollbars")
	return scrollbars
}

// ScrollX prop
// js:"scrollX"
func (*Window) ScrollX() (scrollX int) {
	macro.Rewrite("$_.scrollX")
	return scrollX
}

// ScrollY prop
// js:"scrollY"
func (*Window) ScrollY() (scrollY int) {
	macro.Rewrite("$_.scrollY")
	return scrollY
}

// Self prop
// js:"self"
func (*Window) Self() (self *Window) {
	macro.Rewrite("$_.self")
	return self
}

// SpeechSynthesis prop
// js:"speechSynthesis"
func (*Window) SpeechSynthesis() (speechSynthesis *SpeechSynthesis) {
	macro.Rewrite("$_.speechSynthesis")
	return speechSynthesis
}

// Status prop
// js:"status"
func (*Window) Status() (status string) {
	macro.Rewrite("$_.status")
	return status
}

// SetStatus prop
// js:"status"
func (*Window) SetStatus(status string) {
	macro.Rewrite("$_.status = $1", status)
}

// Statusbar prop
// js:"statusbar"
func (*Window) Statusbar() (statusbar *barprop.BarProp) {
	macro.Rewrite("$_.statusbar")
	return statusbar
}

// StyleMedia prop
// js:"styleMedia"
func (*Window) StyleMedia() (styleMedia *stylemedia.StyleMedia) {
	macro.Rewrite("$_.styleMedia")
	return styleMedia
}

// Toolbar prop
// js:"toolbar"
func (*Window) Toolbar() (toolbar *barprop.BarProp) {
	macro.Rewrite("$_.toolbar")
	return toolbar
}

// Top prop
// js:"top"
func (*Window) Top() (top *Window) {
	macro.Rewrite("$_.top")
	return top
}

// Window prop
// js:"window"
func (*Window) Window() (window *Window) {
	macro.Rewrite("$_.window")
	return window
}

// SessionStorage prop
// js:"sessionStorage"
func (*Window) SessionStorage() (sessionStorage *storage.Storage) {
	macro.Rewrite("$_.sessionStorage")
	return sessionStorage
}

// LocalStorage prop
// js:"localStorage"
func (*Window) LocalStorage() (localStorage *storage.Storage) {
	macro.Rewrite("$_.localStorage")
	return localStorage
}

// Console prop
// js:"console"
func (*Window) Console() (console *Console) {
	macro.Rewrite("$_.console")
	return console
}

// Onpointercancel prop
// js:"onpointercancel"
func (*Window) Onpointercancel() (onpointercancel func(Event)) {
	macro.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*Window) SetOnpointercancel(onpointercancel func(Event)) {
	macro.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*Window) Onpointerdown() (onpointerdown func(Event)) {
	macro.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*Window) SetOnpointerdown(onpointerdown func(Event)) {
	macro.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*Window) Onpointerenter() (onpointerenter func(Event)) {
	macro.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*Window) SetOnpointerenter(onpointerenter func(Event)) {
	macro.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*Window) Onpointerleave() (onpointerleave func(Event)) {
	macro.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*Window) SetOnpointerleave(onpointerleave func(Event)) {
	macro.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*Window) Onpointermove() (onpointermove func(Event)) {
	macro.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*Window) SetOnpointermove(onpointermove func(Event)) {
	macro.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*Window) Onpointerout() (onpointerout func(Event)) {
	macro.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*Window) SetOnpointerout(onpointerout func(Event)) {
	macro.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*Window) Onpointerover() (onpointerover func(Event)) {
	macro.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*Window) SetOnpointerover(onpointerover func(Event)) {
	macro.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*Window) Onpointerup() (onpointerup func(Event)) {
	macro.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*Window) SetOnpointerup(onpointerup func(Event)) {
	macro.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*Window) Onwheel() (onwheel func(Event)) {
	macro.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*Window) SetOnwheel(onwheel func(Event)) {
	macro.Rewrite("$_.onwheel = $1", onwheel)
}

// IndexedDB prop
// js:"indexedDB"
func (*Window) IndexedDB() (indexedDB *IDBFactory) {
	macro.Rewrite("$_.indexedDB")
	return indexedDB
}
