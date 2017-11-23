package window

import (
	"github.com/matthewmueller/golly/dom2/barprop"
	"github.com/matthewmueller/golly/dom2/cachestorage"
	"github.com/matthewmueller/golly/dom2/crypto"
	"github.com/matthewmueller/golly/dom2/extensionscriptapis"
	"github.com/matthewmueller/golly/dom2/external"
	"github.com/matthewmueller/golly/dom2/focusnavigationorigin"
	"github.com/matthewmueller/golly/dom2/globalfetch"
	"github.com/matthewmueller/golly/dom2/history"
	"github.com/matthewmueller/golly/dom2/location"
	"github.com/matthewmueller/golly/dom2/mediaquery"
	"github.com/matthewmueller/golly/dom2/mscredentials"
	"github.com/matthewmueller/golly/dom2/navigationreason"
	"github.com/matthewmueller/golly/dom2/performance"
	"github.com/matthewmueller/golly/dom2/stylemedia"
	"github.com/matthewmueller/golly/dom2/webkitpoint"
	"github.com/matthewmueller/golly/dom2/windowbase64"
	"github.com/matthewmueller/golly/dom2/windowlocalstorage"
	"github.com/matthewmueller/golly/dom2/windowsessionstorage"
	"github.com/matthewmueller/golly/dom2/windowtimers"
	"github.com/matthewmueller/golly/js"
)

// js:"Window,omit"
type Window struct {
	EventTarget
	windowtimers.WindowTimers
	windowsessionstorage.WindowSessionStorage
	windowlocalstorage.WindowLocalStorage
	WindowConsole
	GlobalEventHandlers
	IDBEnvironment
	windowbase64.WindowBase64
	globalfetch.GlobalFetch
}

// Alert
func (*Window) Alert(message *string) {
	js.Rewrite("$<.Alert($1)", message)
}

// Blur
func (*Window) Blur() {
	js.Rewrite("$<.Blur()")
}

// CancelAnimationFrame
func (*Window) CancelAnimationFrame(handle int) {
	js.Rewrite("$<.CancelAnimationFrame($1)", handle)
}

// CaptureEvents
func (*Window) CaptureEvents() {
	js.Rewrite("$<.CaptureEvents()")
}

// Close
func (*Window) Close() {
	js.Rewrite("$<.Close()")
}

// Confirm
func (*Window) Confirm(message *string) (b bool) {
	js.Rewrite("$<.Confirm($1)", message)
	return b
}

// DepartFocus
func (*Window) DepartFocus(navigationReason *navigationreason.NavigationReason, origin *focusnavigationorigin.FocusNavigationOrigin) {
	js.Rewrite("$<.DepartFocus($1, $2)", navigationReason, origin)
}

// Focus
func (*Window) Focus() {
	js.Rewrite("$<.Focus()")
}

// GetComputedStyle
func (*Window) GetComputedStyle(elt Element, pseudoElt *string) (c *CSSStyleDeclaration) {
	js.Rewrite("$<.GetComputedStyle($1, $2)", elt, pseudoElt)
	return c
}

// GetMatchedCSSRules
func (*Window) GetMatchedCSSRules(elt Element, pseudoElt *string) (c *CSSRuleList) {
	js.Rewrite("$<.GetMatchedCSSRules($1, $2)", elt, pseudoElt)
	return c
}

// GetSelection
func (*Window) GetSelection() (s *Selection) {
	js.Rewrite("$<.GetSelection()")
	return s
}

// MatchMedia
func (*Window) MatchMedia(mediaQuery string) (m *mediaquery.MediaQueryList) {
	js.Rewrite("$<.MatchMedia($1)", mediaQuery)
	return m
}

// MoveBy
func (*Window) MoveBy(x *int, y *int) {
	js.Rewrite("$<.MoveBy($1, $2)", x, y)
}

// MoveTo
func (*Window) MoveTo(x *int, y *int) {
	js.Rewrite("$<.MoveTo($1, $2)", x, y)
}

// MsWriteProfilerMark
func (*Window) MsWriteProfilerMark(profilerMarkName string) {
	js.Rewrite("$<.MsWriteProfilerMark($1)", profilerMarkName)
}

// Open
func (*Window) Open(url *string, target *string, features *string, replace *bool) (w *Window) {
	js.Rewrite("$<.Open($1, $2, $3, $4)", url, target, features, replace)
	return w
}

// PostMessage
func (*Window) PostMessage(message interface{}, targetOrigin string, transfer *[]interface{}) {
	js.Rewrite("$<.PostMessage($1, $2, $3)", message, targetOrigin, transfer)
}

// Print
func (*Window) Print() {
	js.Rewrite("$<.Print()")
}

// Prompt
func (*Window) Prompt(message *string, def *string) (s string) {
	js.Rewrite("$<.Prompt($1, $2)", message, def)
	return s
}

// ReleaseEvents
func (*Window) ReleaseEvents() {
	js.Rewrite("$<.ReleaseEvents()")
}

// RequestAnimationFrame
func (*Window) RequestAnimationFrame(callback func(time int)) (i int) {
	js.Rewrite("$<.RequestAnimationFrame($1)", callback)
	return i
}

// ResizeBy
func (*Window) ResizeBy(x *int, y *int) {
	js.Rewrite("$<.ResizeBy($1, $2)", x, y)
}

// ResizeTo
func (*Window) ResizeTo(x *int, y *int) {
	js.Rewrite("$<.ResizeTo($1, $2)", x, y)
}

// Scroll
func (*Window) Scroll(x *int, y *int) {
	js.Rewrite("$<.Scroll($1, $2)", x, y)
}

// ScrollBy
func (*Window) ScrollBy(x *int, y *int) {
	js.Rewrite("$<.ScrollBy($1, $2)", x, y)
}

// ScrollTo
func (*Window) ScrollTo(x *int, y *int) {
	js.Rewrite("$<.ScrollTo($1, $2)", x, y)
}

// Stop
func (*Window) Stop() {
	js.Rewrite("$<.Stop()")
}

// WebkitCancelAnimationFrame
func (*Window) WebkitCancelAnimationFrame(handle int) {
	js.Rewrite("$<.WebkitCancelAnimationFrame($1)", handle)
}

// WebkitConvertPointFromNodeToPage
func (*Window) WebkitConvertPointFromNodeToPage(node Node, pt *webkitpoint.WebKitPoint) (w *webkitpoint.WebKitPoint) {
	js.Rewrite("$<.WebkitConvertPointFromNodeToPage($1, $2)", node, pt)
	return w
}

// WebkitConvertPointFromPageToNode
func (*Window) WebkitConvertPointFromPageToNode(node Node, pt *webkitpoint.WebKitPoint) (w *webkitpoint.WebKitPoint) {
	js.Rewrite("$<.WebkitConvertPointFromPageToNode($1, $2)", node, pt)
	return w
}

// WebkitRequestAnimationFrame
func (*Window) WebkitRequestAnimationFrame(callback func(time int)) (i int) {
	js.Rewrite("$<.WebkitRequestAnimationFrame($1)", callback)
	return i
}

// ApplicationCache
func (*Window) ApplicationCache() (applicationCache *ApplicationCache) {
	js.Rewrite("$<.ApplicationCache")
	return applicationCache
}

// Caches
func (*Window) Caches() (caches *cachestorage.CacheStorage) {
	js.Rewrite("$<.Caches")
	return caches
}

// ClientInformation
func (*Window) ClientInformation() (clientInformation *Navigator) {
	js.Rewrite("$<.ClientInformation")
	return clientInformation
}

// Closed
func (*Window) Closed() (closed bool) {
	js.Rewrite("$<.Closed")
	return closed
}

// Crypto
func (*Window) Crypto() (crypto *crypto.Crypto) {
	js.Rewrite("$<.Crypto")
	return crypto
}

// DefaultStatus
func (*Window) DefaultStatus() (defaultStatus string) {
	js.Rewrite("$<.DefaultStatus")
	return defaultStatus
}

// DefaultStatus
func (*Window) SetDefaultStatus(defaultStatus string) {
	js.Rewrite("$<.DefaultStatus = $1", defaultStatus)
}

// DevicePixelRatio
func (*Window) DevicePixelRatio() (devicePixelRatio float32) {
	js.Rewrite("$<.DevicePixelRatio")
	return devicePixelRatio
}

// Document
func (*Window) Document() (document Document) {
	js.Rewrite("$<.Document")
	return document
}

// DoNotTrack
func (*Window) DoNotTrack() (doNotTrack string) {
	js.Rewrite("$<.DoNotTrack")
	return doNotTrack
}

// Event
func (*Window) Event() (event Event) {
	js.Rewrite("$<.Event")
	return event
}

// Event
func (*Window) SetEvent(event Event) {
	js.Rewrite("$<.Event = $1", event)
}

// External
func (*Window) External() (external *external.External) {
	js.Rewrite("$<.External")
	return external
}

// FrameElement
func (*Window) FrameElement() (frameElement Element) {
	js.Rewrite("$<.FrameElement")
	return frameElement
}

// Frames
func (*Window) Frames() (frames *Window) {
	js.Rewrite("$<.Frames")
	return frames
}

// History
func (*Window) History() (history *history.History) {
	js.Rewrite("$<.History")
	return history
}

// InnerHeight
func (*Window) InnerHeight() (innerHeight int) {
	js.Rewrite("$<.InnerHeight")
	return innerHeight
}

// InnerWidth
func (*Window) InnerWidth() (innerWidth int) {
	js.Rewrite("$<.InnerWidth")
	return innerWidth
}

// IsSecureContext
func (*Window) IsSecureContext() (isSecureContext bool) {
	js.Rewrite("$<.IsSecureContext")
	return isSecureContext
}

// Length
func (*Window) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}

// Location
func (*Window) Location() (location *location.Location) {
	js.Rewrite("$<.Location")
	return location
}

// Locationbar
func (*Window) Locationbar() (locationbar *barprop.BarProp) {
	js.Rewrite("$<.Locationbar")
	return locationbar
}

// Menubar
func (*Window) Menubar() (menubar *barprop.BarProp) {
	js.Rewrite("$<.Menubar")
	return menubar
}

// MsContentScript
func (*Window) MsContentScript() (msContentScript *extensionscriptapis.ExtensionScriptApis) {
	js.Rewrite("$<.MsContentScript")
	return msContentScript
}

// MsCredentials
func (*Window) MsCredentials() (msCredentials *mscredentials.MSCredentials) {
	js.Rewrite("$<.MsCredentials")
	return msCredentials
}

// Name
func (*Window) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name
func (*Window) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// Navigator
func (*Window) Navigator() (navigator *Navigator) {
	js.Rewrite("$<.Navigator")
	return navigator
}

// OffscreenBuffering
func (*Window) OffscreenBuffering() (offscreenBuffering interface{}) {
	js.Rewrite("$<.OffscreenBuffering")
	return offscreenBuffering
}

// OffscreenBuffering
func (*Window) SetOffscreenBuffering(offscreenBuffering interface{}) {
	js.Rewrite("$<.OffscreenBuffering = $1", offscreenBuffering)
}

// Onabort
func (*Window) Onabort() (onabort func(Event)) {
	js.Rewrite("$<.Onabort")
	return onabort
}

// Onabort
func (*Window) SetOnabort(onabort func(Event)) {
	js.Rewrite("$<.Onabort = $1", onabort)
}

// Onafterprint
func (*Window) Onafterprint() (onafterprint func(Event)) {
	js.Rewrite("$<.Onafterprint")
	return onafterprint
}

// Onafterprint
func (*Window) SetOnafterprint(onafterprint func(Event)) {
	js.Rewrite("$<.Onafterprint = $1", onafterprint)
}

// Onbeforeprint
func (*Window) Onbeforeprint() (onbeforeprint func(Event)) {
	js.Rewrite("$<.Onbeforeprint")
	return onbeforeprint
}

// Onbeforeprint
func (*Window) SetOnbeforeprint(onbeforeprint func(Event)) {
	js.Rewrite("$<.Onbeforeprint = $1", onbeforeprint)
}

// Onbeforeunload
func (*Window) Onbeforeunload() (onbeforeunload func(*BeforeUnloadEvent)) {
	js.Rewrite("$<.Onbeforeunload")
	return onbeforeunload
}

// Onbeforeunload
func (*Window) SetOnbeforeunload(onbeforeunload func(*BeforeUnloadEvent)) {
	js.Rewrite("$<.Onbeforeunload = $1", onbeforeunload)
}

// Onblur
func (*Window) Onblur() (onblur func(*FocusEvent)) {
	js.Rewrite("$<.Onblur")
	return onblur
}

// Onblur
func (*Window) SetOnblur(onblur func(*FocusEvent)) {
	js.Rewrite("$<.Onblur = $1", onblur)
}

// Oncanplay
func (*Window) Oncanplay() (oncanplay func(Event)) {
	js.Rewrite("$<.Oncanplay")
	return oncanplay
}

// Oncanplay
func (*Window) SetOncanplay(oncanplay func(Event)) {
	js.Rewrite("$<.Oncanplay = $1", oncanplay)
}

// Oncanplaythrough
func (*Window) Oncanplaythrough() (oncanplaythrough func(Event)) {
	js.Rewrite("$<.Oncanplaythrough")
	return oncanplaythrough
}

// Oncanplaythrough
func (*Window) SetOncanplaythrough(oncanplaythrough func(Event)) {
	js.Rewrite("$<.Oncanplaythrough = $1", oncanplaythrough)
}

// Onchange
func (*Window) Onchange() (onchange func(Event)) {
	js.Rewrite("$<.Onchange")
	return onchange
}

// Onchange
func (*Window) SetOnchange(onchange func(Event)) {
	js.Rewrite("$<.Onchange = $1", onchange)
}

// Onclick
func (*Window) Onclick() (onclick func(Event)) {
	js.Rewrite("$<.Onclick")
	return onclick
}

// Onclick
func (*Window) SetOnclick(onclick func(Event)) {
	js.Rewrite("$<.Onclick = $1", onclick)
}

// Oncompassneedscalibration
func (*Window) Oncompassneedscalibration() (oncompassneedscalibration func(Event)) {
	js.Rewrite("$<.Oncompassneedscalibration")
	return oncompassneedscalibration
}

// Oncompassneedscalibration
func (*Window) SetOncompassneedscalibration(oncompassneedscalibration func(Event)) {
	js.Rewrite("$<.Oncompassneedscalibration = $1", oncompassneedscalibration)
}

// Oncontextmenu
func (*Window) Oncontextmenu() (oncontextmenu func(Event)) {
	js.Rewrite("$<.Oncontextmenu")
	return oncontextmenu
}

// Oncontextmenu
func (*Window) SetOncontextmenu(oncontextmenu func(Event)) {
	js.Rewrite("$<.Oncontextmenu = $1", oncontextmenu)
}

// Ondblclick
func (*Window) Ondblclick() (ondblclick func(Event)) {
	js.Rewrite("$<.Ondblclick")
	return ondblclick
}

// Ondblclick
func (*Window) SetOndblclick(ondblclick func(Event)) {
	js.Rewrite("$<.Ondblclick = $1", ondblclick)
}

// Ondevicelight
func (*Window) Ondevicelight() (ondevicelight func(*DeviceLightEvent)) {
	js.Rewrite("$<.Ondevicelight")
	return ondevicelight
}

// Ondevicelight
func (*Window) SetOndevicelight(ondevicelight func(*DeviceLightEvent)) {
	js.Rewrite("$<.Ondevicelight = $1", ondevicelight)
}

// Ondevicemotion
func (*Window) Ondevicemotion() (ondevicemotion func(*DeviceMotionEvent)) {
	js.Rewrite("$<.Ondevicemotion")
	return ondevicemotion
}

// Ondevicemotion
func (*Window) SetOndevicemotion(ondevicemotion func(*DeviceMotionEvent)) {
	js.Rewrite("$<.Ondevicemotion = $1", ondevicemotion)
}

// Ondeviceorientation
func (*Window) Ondeviceorientation() (ondeviceorientation func(*DeviceOrientationEvent)) {
	js.Rewrite("$<.Ondeviceorientation")
	return ondeviceorientation
}

// Ondeviceorientation
func (*Window) SetOndeviceorientation(ondeviceorientation func(*DeviceOrientationEvent)) {
	js.Rewrite("$<.Ondeviceorientation = $1", ondeviceorientation)
}

// Ondrag
func (*Window) Ondrag() (ondrag func(Event)) {
	js.Rewrite("$<.Ondrag")
	return ondrag
}

// Ondrag
func (*Window) SetOndrag(ondrag func(Event)) {
	js.Rewrite("$<.Ondrag = $1", ondrag)
}

// Ondragend
func (*Window) Ondragend() (ondragend func(Event)) {
	js.Rewrite("$<.Ondragend")
	return ondragend
}

// Ondragend
func (*Window) SetOndragend(ondragend func(Event)) {
	js.Rewrite("$<.Ondragend = $1", ondragend)
}

// Ondragenter
func (*Window) Ondragenter() (ondragenter func(Event)) {
	js.Rewrite("$<.Ondragenter")
	return ondragenter
}

// Ondragenter
func (*Window) SetOndragenter(ondragenter func(Event)) {
	js.Rewrite("$<.Ondragenter = $1", ondragenter)
}

// Ondragleave
func (*Window) Ondragleave() (ondragleave func(Event)) {
	js.Rewrite("$<.Ondragleave")
	return ondragleave
}

// Ondragleave
func (*Window) SetOndragleave(ondragleave func(Event)) {
	js.Rewrite("$<.Ondragleave = $1", ondragleave)
}

// Ondragover
func (*Window) Ondragover() (ondragover func(Event)) {
	js.Rewrite("$<.Ondragover")
	return ondragover
}

// Ondragover
func (*Window) SetOndragover(ondragover func(Event)) {
	js.Rewrite("$<.Ondragover = $1", ondragover)
}

// Ondragstart
func (*Window) Ondragstart() (ondragstart func(Event)) {
	js.Rewrite("$<.Ondragstart")
	return ondragstart
}

// Ondragstart
func (*Window) SetOndragstart(ondragstart func(Event)) {
	js.Rewrite("$<.Ondragstart = $1", ondragstart)
}

// Ondrop
func (*Window) Ondrop() (ondrop func(Event)) {
	js.Rewrite("$<.Ondrop")
	return ondrop
}

// Ondrop
func (*Window) SetOndrop(ondrop func(Event)) {
	js.Rewrite("$<.Ondrop = $1", ondrop)
}

// Ondurationchange
func (*Window) Ondurationchange() (ondurationchange func(Event)) {
	js.Rewrite("$<.Ondurationchange")
	return ondurationchange
}

// Ondurationchange
func (*Window) SetOndurationchange(ondurationchange func(Event)) {
	js.Rewrite("$<.Ondurationchange = $1", ondurationchange)
}

// Onemptied
func (*Window) Onemptied() (onemptied func(Event)) {
	js.Rewrite("$<.Onemptied")
	return onemptied
}

// Onemptied
func (*Window) SetOnemptied(onemptied func(Event)) {
	js.Rewrite("$<.Onemptied = $1", onemptied)
}

// Onended
func (*Window) Onended() (onended func(Event)) {
	js.Rewrite("$<.Onended")
	return onended
}

// Onended
func (*Window) SetOnended(onended func(Event)) {
	js.Rewrite("$<.Onended = $1", onended)
}

// Onerror
func (*Window) Onerror() (onerror func(columnNumber *uint, event interface{}, fileno *uint, source *string)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*Window) SetOnerror(onerror func(columnNumber *uint, event interface{}, fileno *uint, source *string)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onfocus
func (*Window) Onfocus() (onfocus func(*FocusEvent)) {
	js.Rewrite("$<.Onfocus")
	return onfocus
}

// Onfocus
func (*Window) SetOnfocus(onfocus func(*FocusEvent)) {
	js.Rewrite("$<.Onfocus = $1", onfocus)
}

// Onhashchange
func (*Window) Onhashchange() (onhashchange func(*HashChangeEvent)) {
	js.Rewrite("$<.Onhashchange")
	return onhashchange
}

// Onhashchange
func (*Window) SetOnhashchange(onhashchange func(*HashChangeEvent)) {
	js.Rewrite("$<.Onhashchange = $1", onhashchange)
}

// Oninput
func (*Window) Oninput() (oninput func(Event)) {
	js.Rewrite("$<.Oninput")
	return oninput
}

// Oninput
func (*Window) SetOninput(oninput func(Event)) {
	js.Rewrite("$<.Oninput = $1", oninput)
}

// Oninvalid
func (*Window) Oninvalid() (oninvalid func(Event)) {
	js.Rewrite("$<.Oninvalid")
	return oninvalid
}

// Oninvalid
func (*Window) SetOninvalid(oninvalid func(Event)) {
	js.Rewrite("$<.Oninvalid = $1", oninvalid)
}

// Onkeydown
func (*Window) Onkeydown() (onkeydown func(Event)) {
	js.Rewrite("$<.Onkeydown")
	return onkeydown
}

// Onkeydown
func (*Window) SetOnkeydown(onkeydown func(Event)) {
	js.Rewrite("$<.Onkeydown = $1", onkeydown)
}

// Onkeypress
func (*Window) Onkeypress() (onkeypress func(Event)) {
	js.Rewrite("$<.Onkeypress")
	return onkeypress
}

// Onkeypress
func (*Window) SetOnkeypress(onkeypress func(Event)) {
	js.Rewrite("$<.Onkeypress = $1", onkeypress)
}

// Onkeyup
func (*Window) Onkeyup() (onkeyup func(Event)) {
	js.Rewrite("$<.Onkeyup")
	return onkeyup
}

// Onkeyup
func (*Window) SetOnkeyup(onkeyup func(Event)) {
	js.Rewrite("$<.Onkeyup = $1", onkeyup)
}

// Onload
func (*Window) Onload() (onload func(Event)) {
	js.Rewrite("$<.Onload")
	return onload
}

// Onload
func (*Window) SetOnload(onload func(Event)) {
	js.Rewrite("$<.Onload = $1", onload)
}

// Onloadeddata
func (*Window) Onloadeddata() (onloadeddata func(Event)) {
	js.Rewrite("$<.Onloadeddata")
	return onloadeddata
}

// Onloadeddata
func (*Window) SetOnloadeddata(onloadeddata func(Event)) {
	js.Rewrite("$<.Onloadeddata = $1", onloadeddata)
}

// Onloadedmetadata
func (*Window) Onloadedmetadata() (onloadedmetadata func(Event)) {
	js.Rewrite("$<.Onloadedmetadata")
	return onloadedmetadata
}

// Onloadedmetadata
func (*Window) SetOnloadedmetadata(onloadedmetadata func(Event)) {
	js.Rewrite("$<.Onloadedmetadata = $1", onloadedmetadata)
}

// Onloadstart
func (*Window) Onloadstart() (onloadstart func(Event)) {
	js.Rewrite("$<.Onloadstart")
	return onloadstart
}

// Onloadstart
func (*Window) SetOnloadstart(onloadstart func(Event)) {
	js.Rewrite("$<.Onloadstart = $1", onloadstart)
}

// Onmessage
func (*Window) Onmessage() (onmessage func(*MessageEvent)) {
	js.Rewrite("$<.Onmessage")
	return onmessage
}

// Onmessage
func (*Window) SetOnmessage(onmessage func(*MessageEvent)) {
	js.Rewrite("$<.Onmessage = $1", onmessage)
}

// Onmousedown
func (*Window) Onmousedown() (onmousedown func(Event)) {
	js.Rewrite("$<.Onmousedown")
	return onmousedown
}

// Onmousedown
func (*Window) SetOnmousedown(onmousedown func(Event)) {
	js.Rewrite("$<.Onmousedown = $1", onmousedown)
}

// Onmouseenter
func (*Window) Onmouseenter() (onmouseenter func(Event)) {
	js.Rewrite("$<.Onmouseenter")
	return onmouseenter
}

// Onmouseenter
func (*Window) SetOnmouseenter(onmouseenter func(Event)) {
	js.Rewrite("$<.Onmouseenter = $1", onmouseenter)
}

// Onmouseleave
func (*Window) Onmouseleave() (onmouseleave func(Event)) {
	js.Rewrite("$<.Onmouseleave")
	return onmouseleave
}

// Onmouseleave
func (*Window) SetOnmouseleave(onmouseleave func(Event)) {
	js.Rewrite("$<.Onmouseleave = $1", onmouseleave)
}

// Onmousemove
func (*Window) Onmousemove() (onmousemove func(Event)) {
	js.Rewrite("$<.Onmousemove")
	return onmousemove
}

// Onmousemove
func (*Window) SetOnmousemove(onmousemove func(Event)) {
	js.Rewrite("$<.Onmousemove = $1", onmousemove)
}

// Onmouseout
func (*Window) Onmouseout() (onmouseout func(Event)) {
	js.Rewrite("$<.Onmouseout")
	return onmouseout
}

// Onmouseout
func (*Window) SetOnmouseout(onmouseout func(Event)) {
	js.Rewrite("$<.Onmouseout = $1", onmouseout)
}

// Onmouseover
func (*Window) Onmouseover() (onmouseover func(Event)) {
	js.Rewrite("$<.Onmouseover")
	return onmouseover
}

// Onmouseover
func (*Window) SetOnmouseover(onmouseover func(Event)) {
	js.Rewrite("$<.Onmouseover = $1", onmouseover)
}

// Onmouseup
func (*Window) Onmouseup() (onmouseup func(Event)) {
	js.Rewrite("$<.Onmouseup")
	return onmouseup
}

// Onmouseup
func (*Window) SetOnmouseup(onmouseup func(Event)) {
	js.Rewrite("$<.Onmouseup = $1", onmouseup)
}

// Onmousewheel
func (*Window) Onmousewheel() (onmousewheel func(Event)) {
	js.Rewrite("$<.Onmousewheel")
	return onmousewheel
}

// Onmousewheel
func (*Window) SetOnmousewheel(onmousewheel func(Event)) {
	js.Rewrite("$<.Onmousewheel = $1", onmousewheel)
}

// Onmsgesturechange
func (*Window) Onmsgesturechange() (onmsgesturechange func(Event)) {
	js.Rewrite("$<.Onmsgesturechange")
	return onmsgesturechange
}

// Onmsgesturechange
func (*Window) SetOnmsgesturechange(onmsgesturechange func(Event)) {
	js.Rewrite("$<.Onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap
func (*Window) Onmsgesturedoubletap() (onmsgesturedoubletap func(Event)) {
	js.Rewrite("$<.Onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// Onmsgesturedoubletap
func (*Window) SetOnmsgesturedoubletap(onmsgesturedoubletap func(Event)) {
	js.Rewrite("$<.Onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend
func (*Window) Onmsgestureend() (onmsgestureend func(Event)) {
	js.Rewrite("$<.Onmsgestureend")
	return onmsgestureend
}

// Onmsgestureend
func (*Window) SetOnmsgestureend(onmsgestureend func(Event)) {
	js.Rewrite("$<.Onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold
func (*Window) Onmsgesturehold() (onmsgesturehold func(Event)) {
	js.Rewrite("$<.Onmsgesturehold")
	return onmsgesturehold
}

// Onmsgesturehold
func (*Window) SetOnmsgesturehold(onmsgesturehold func(Event)) {
	js.Rewrite("$<.Onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart
func (*Window) Onmsgesturestart() (onmsgesturestart func(Event)) {
	js.Rewrite("$<.Onmsgesturestart")
	return onmsgesturestart
}

// Onmsgesturestart
func (*Window) SetOnmsgesturestart(onmsgesturestart func(Event)) {
	js.Rewrite("$<.Onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap
func (*Window) Onmsgesturetap() (onmsgesturetap func(Event)) {
	js.Rewrite("$<.Onmsgesturetap")
	return onmsgesturetap
}

// Onmsgesturetap
func (*Window) SetOnmsgesturetap(onmsgesturetap func(Event)) {
	js.Rewrite("$<.Onmsgesturetap = $1", onmsgesturetap)
}

// Onmsinertiastart
func (*Window) Onmsinertiastart() (onmsinertiastart func(Event)) {
	js.Rewrite("$<.Onmsinertiastart")
	return onmsinertiastart
}

// Onmsinertiastart
func (*Window) SetOnmsinertiastart(onmsinertiastart func(Event)) {
	js.Rewrite("$<.Onmsinertiastart = $1", onmsinertiastart)
}

// Onmspointercancel
func (*Window) Onmspointercancel() (onmspointercancel func(Event)) {
	js.Rewrite("$<.Onmspointercancel")
	return onmspointercancel
}

// Onmspointercancel
func (*Window) SetOnmspointercancel(onmspointercancel func(Event)) {
	js.Rewrite("$<.Onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown
func (*Window) Onmspointerdown() (onmspointerdown func(Event)) {
	js.Rewrite("$<.Onmspointerdown")
	return onmspointerdown
}

// Onmspointerdown
func (*Window) SetOnmspointerdown(onmspointerdown func(Event)) {
	js.Rewrite("$<.Onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter
func (*Window) Onmspointerenter() (onmspointerenter func(Event)) {
	js.Rewrite("$<.Onmspointerenter")
	return onmspointerenter
}

// Onmspointerenter
func (*Window) SetOnmspointerenter(onmspointerenter func(Event)) {
	js.Rewrite("$<.Onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave
func (*Window) Onmspointerleave() (onmspointerleave func(Event)) {
	js.Rewrite("$<.Onmspointerleave")
	return onmspointerleave
}

// Onmspointerleave
func (*Window) SetOnmspointerleave(onmspointerleave func(Event)) {
	js.Rewrite("$<.Onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove
func (*Window) Onmspointermove() (onmspointermove func(Event)) {
	js.Rewrite("$<.Onmspointermove")
	return onmspointermove
}

// Onmspointermove
func (*Window) SetOnmspointermove(onmspointermove func(Event)) {
	js.Rewrite("$<.Onmspointermove = $1", onmspointermove)
}

// Onmspointerout
func (*Window) Onmspointerout() (onmspointerout func(Event)) {
	js.Rewrite("$<.Onmspointerout")
	return onmspointerout
}

// Onmspointerout
func (*Window) SetOnmspointerout(onmspointerout func(Event)) {
	js.Rewrite("$<.Onmspointerout = $1", onmspointerout)
}

// Onmspointerover
func (*Window) Onmspointerover() (onmspointerover func(Event)) {
	js.Rewrite("$<.Onmspointerover")
	return onmspointerover
}

// Onmspointerover
func (*Window) SetOnmspointerover(onmspointerover func(Event)) {
	js.Rewrite("$<.Onmspointerover = $1", onmspointerover)
}

// Onmspointerup
func (*Window) Onmspointerup() (onmspointerup func(Event)) {
	js.Rewrite("$<.Onmspointerup")
	return onmspointerup
}

// Onmspointerup
func (*Window) SetOnmspointerup(onmspointerup func(Event)) {
	js.Rewrite("$<.Onmspointerup = $1", onmspointerup)
}

// Onoffline
func (*Window) Onoffline() (onoffline func(Event)) {
	js.Rewrite("$<.Onoffline")
	return onoffline
}

// Onoffline
func (*Window) SetOnoffline(onoffline func(Event)) {
	js.Rewrite("$<.Onoffline = $1", onoffline)
}

// Ononline
func (*Window) Ononline() (ononline func(Event)) {
	js.Rewrite("$<.Ononline")
	return ononline
}

// Ononline
func (*Window) SetOnonline(ononline func(Event)) {
	js.Rewrite("$<.Ononline = $1", ononline)
}

// Onorientationchange
func (*Window) Onorientationchange() (onorientationchange func(Event)) {
	js.Rewrite("$<.Onorientationchange")
	return onorientationchange
}

// Onorientationchange
func (*Window) SetOnorientationchange(onorientationchange func(Event)) {
	js.Rewrite("$<.Onorientationchange = $1", onorientationchange)
}

// Onpagehide
func (*Window) Onpagehide() (onpagehide func(*PageTransitionEvent)) {
	js.Rewrite("$<.Onpagehide")
	return onpagehide
}

// Onpagehide
func (*Window) SetOnpagehide(onpagehide func(*PageTransitionEvent)) {
	js.Rewrite("$<.Onpagehide = $1", onpagehide)
}

// Onpageshow
func (*Window) Onpageshow() (onpageshow func(*PageTransitionEvent)) {
	js.Rewrite("$<.Onpageshow")
	return onpageshow
}

// Onpageshow
func (*Window) SetOnpageshow(onpageshow func(*PageTransitionEvent)) {
	js.Rewrite("$<.Onpageshow = $1", onpageshow)
}

// Onpause
func (*Window) Onpause() (onpause func(Event)) {
	js.Rewrite("$<.Onpause")
	return onpause
}

// Onpause
func (*Window) SetOnpause(onpause func(Event)) {
	js.Rewrite("$<.Onpause = $1", onpause)
}

// Onplay
func (*Window) Onplay() (onplay func(Event)) {
	js.Rewrite("$<.Onplay")
	return onplay
}

// Onplay
func (*Window) SetOnplay(onplay func(Event)) {
	js.Rewrite("$<.Onplay = $1", onplay)
}

// Onplaying
func (*Window) Onplaying() (onplaying func(Event)) {
	js.Rewrite("$<.Onplaying")
	return onplaying
}

// Onplaying
func (*Window) SetOnplaying(onplaying func(Event)) {
	js.Rewrite("$<.Onplaying = $1", onplaying)
}

// Onpopstate
func (*Window) Onpopstate() (onpopstate func(*PopStateEvent)) {
	js.Rewrite("$<.Onpopstate")
	return onpopstate
}

// Onpopstate
func (*Window) SetOnpopstate(onpopstate func(*PopStateEvent)) {
	js.Rewrite("$<.Onpopstate = $1", onpopstate)
}

// Onprogress
func (*Window) Onprogress() (onprogress func(Event)) {
	js.Rewrite("$<.Onprogress")
	return onprogress
}

// Onprogress
func (*Window) SetOnprogress(onprogress func(Event)) {
	js.Rewrite("$<.Onprogress = $1", onprogress)
}

// Onratechange
func (*Window) Onratechange() (onratechange func(Event)) {
	js.Rewrite("$<.Onratechange")
	return onratechange
}

// Onratechange
func (*Window) SetOnratechange(onratechange func(Event)) {
	js.Rewrite("$<.Onratechange = $1", onratechange)
}

// Onreadystatechange
func (*Window) Onreadystatechange() (onreadystatechange func(Event)) {
	js.Rewrite("$<.Onreadystatechange")
	return onreadystatechange
}

// Onreadystatechange
func (*Window) SetOnreadystatechange(onreadystatechange func(Event)) {
	js.Rewrite("$<.Onreadystatechange = $1", onreadystatechange)
}

// Onreset
func (*Window) Onreset() (onreset func(Event)) {
	js.Rewrite("$<.Onreset")
	return onreset
}

// Onreset
func (*Window) SetOnreset(onreset func(Event)) {
	js.Rewrite("$<.Onreset = $1", onreset)
}

// Onresize
func (*Window) Onresize() (onresize func(UIEvent)) {
	js.Rewrite("$<.Onresize")
	return onresize
}

// Onresize
func (*Window) SetOnresize(onresize func(UIEvent)) {
	js.Rewrite("$<.Onresize = $1", onresize)
}

// Onscroll
func (*Window) Onscroll() (onscroll func(Event)) {
	js.Rewrite("$<.Onscroll")
	return onscroll
}

// Onscroll
func (*Window) SetOnscroll(onscroll func(Event)) {
	js.Rewrite("$<.Onscroll = $1", onscroll)
}

// Onseeked
func (*Window) Onseeked() (onseeked func(Event)) {
	js.Rewrite("$<.Onseeked")
	return onseeked
}

// Onseeked
func (*Window) SetOnseeked(onseeked func(Event)) {
	js.Rewrite("$<.Onseeked = $1", onseeked)
}

// Onseeking
func (*Window) Onseeking() (onseeking func(Event)) {
	js.Rewrite("$<.Onseeking")
	return onseeking
}

// Onseeking
func (*Window) SetOnseeking(onseeking func(Event)) {
	js.Rewrite("$<.Onseeking = $1", onseeking)
}

// Onselect
func (*Window) Onselect() (onselect func(Event)) {
	js.Rewrite("$<.Onselect")
	return onselect
}

// Onselect
func (*Window) SetOnselect(onselect func(Event)) {
	js.Rewrite("$<.Onselect = $1", onselect)
}

// Onstalled
func (*Window) Onstalled() (onstalled func(Event)) {
	js.Rewrite("$<.Onstalled")
	return onstalled
}

// Onstalled
func (*Window) SetOnstalled(onstalled func(Event)) {
	js.Rewrite("$<.Onstalled = $1", onstalled)
}

// Onstorage
func (*Window) Onstorage() (onstorage func(*StorageEvent)) {
	js.Rewrite("$<.Onstorage")
	return onstorage
}

// Onstorage
func (*Window) SetOnstorage(onstorage func(*StorageEvent)) {
	js.Rewrite("$<.Onstorage = $1", onstorage)
}

// Onsubmit
func (*Window) Onsubmit() (onsubmit func(Event)) {
	js.Rewrite("$<.Onsubmit")
	return onsubmit
}

// Onsubmit
func (*Window) SetOnsubmit(onsubmit func(Event)) {
	js.Rewrite("$<.Onsubmit = $1", onsubmit)
}

// Onsuspend
func (*Window) Onsuspend() (onsuspend func(Event)) {
	js.Rewrite("$<.Onsuspend")
	return onsuspend
}

// Onsuspend
func (*Window) SetOnsuspend(onsuspend func(Event)) {
	js.Rewrite("$<.Onsuspend = $1", onsuspend)
}

// Ontimeupdate
func (*Window) Ontimeupdate() (ontimeupdate func(Event)) {
	js.Rewrite("$<.Ontimeupdate")
	return ontimeupdate
}

// Ontimeupdate
func (*Window) SetOntimeupdate(ontimeupdate func(Event)) {
	js.Rewrite("$<.Ontimeupdate = $1", ontimeupdate)
}

// Ontouchcancel
func (*Window) Ontouchcancel() (ontouchcancel interface{}) {
	js.Rewrite("$<.Ontouchcancel")
	return ontouchcancel
}

// Ontouchcancel
func (*Window) SetOntouchcancel(ontouchcancel interface{}) {
	js.Rewrite("$<.Ontouchcancel = $1", ontouchcancel)
}

// Ontouchend
func (*Window) Ontouchend() (ontouchend interface{}) {
	js.Rewrite("$<.Ontouchend")
	return ontouchend
}

// Ontouchend
func (*Window) SetOntouchend(ontouchend interface{}) {
	js.Rewrite("$<.Ontouchend = $1", ontouchend)
}

// Ontouchmove
func (*Window) Ontouchmove() (ontouchmove interface{}) {
	js.Rewrite("$<.Ontouchmove")
	return ontouchmove
}

// Ontouchmove
func (*Window) SetOntouchmove(ontouchmove interface{}) {
	js.Rewrite("$<.Ontouchmove = $1", ontouchmove)
}

// Ontouchstart
func (*Window) Ontouchstart() (ontouchstart interface{}) {
	js.Rewrite("$<.Ontouchstart")
	return ontouchstart
}

// Ontouchstart
func (*Window) SetOntouchstart(ontouchstart interface{}) {
	js.Rewrite("$<.Ontouchstart = $1", ontouchstart)
}

// Onunload
func (*Window) Onunload() (onunload func(Event)) {
	js.Rewrite("$<.Onunload")
	return onunload
}

// Onunload
func (*Window) SetOnunload(onunload func(Event)) {
	js.Rewrite("$<.Onunload = $1", onunload)
}

// Onvolumechange
func (*Window) Onvolumechange() (onvolumechange func(Event)) {
	js.Rewrite("$<.Onvolumechange")
	return onvolumechange
}

// Onvolumechange
func (*Window) SetOnvolumechange(onvolumechange func(Event)) {
	js.Rewrite("$<.Onvolumechange = $1", onvolumechange)
}

// Onwaiting
func (*Window) Onwaiting() (onwaiting func(Event)) {
	js.Rewrite("$<.Onwaiting")
	return onwaiting
}

// Onwaiting
func (*Window) SetOnwaiting(onwaiting func(Event)) {
	js.Rewrite("$<.Onwaiting = $1", onwaiting)
}

// Opener
func (*Window) Opener() (opener *Window) {
	js.Rewrite("$<.Opener")
	return opener
}

// Orientation
func (*Window) Orientation() (orientation string) {
	js.Rewrite("$<.Orientation")
	return orientation
}

// OuterHeight
func (*Window) OuterHeight() (outerHeight int) {
	js.Rewrite("$<.OuterHeight")
	return outerHeight
}

// OuterWidth
func (*Window) OuterWidth() (outerWidth int) {
	js.Rewrite("$<.OuterWidth")
	return outerWidth
}

// PageXOffset
func (*Window) PageXOffset() (pageXOffset int) {
	js.Rewrite("$<.PageXOffset")
	return pageXOffset
}

// PageYOffset
func (*Window) PageYOffset() (pageYOffset int) {
	js.Rewrite("$<.PageYOffset")
	return pageYOffset
}

// Parent
func (*Window) Parent() (parent *Window) {
	js.Rewrite("$<.Parent")
	return parent
}

// Performance
func (*Window) Performance() (performance *performance.Performance) {
	js.Rewrite("$<.Performance")
	return performance
}

// Personalbar
func (*Window) Personalbar() (personalbar *barprop.BarProp) {
	js.Rewrite("$<.Personalbar")
	return personalbar
}

// Screen
func (*Window) Screen() (screen *Screen) {
	js.Rewrite("$<.Screen")
	return screen
}

// ScreenLeft
func (*Window) ScreenLeft() (screenLeft int) {
	js.Rewrite("$<.ScreenLeft")
	return screenLeft
}

// ScreenTop
func (*Window) ScreenTop() (screenTop int) {
	js.Rewrite("$<.ScreenTop")
	return screenTop
}

// ScreenX
func (*Window) ScreenX() (screenX int) {
	js.Rewrite("$<.ScreenX")
	return screenX
}

// ScreenY
func (*Window) ScreenY() (screenY int) {
	js.Rewrite("$<.ScreenY")
	return screenY
}

// Scrollbars
func (*Window) Scrollbars() (scrollbars *barprop.BarProp) {
	js.Rewrite("$<.Scrollbars")
	return scrollbars
}

// ScrollX
func (*Window) ScrollX() (scrollX int) {
	js.Rewrite("$<.ScrollX")
	return scrollX
}

// ScrollY
func (*Window) ScrollY() (scrollY int) {
	js.Rewrite("$<.ScrollY")
	return scrollY
}

// Self
func (*Window) Self() (self *Window) {
	js.Rewrite("$<.Self")
	return self
}

// SpeechSynthesis
func (*Window) SpeechSynthesis() (speechSynthesis *SpeechSynthesis) {
	js.Rewrite("$<.SpeechSynthesis")
	return speechSynthesis
}

// Status
func (*Window) Status() (status string) {
	js.Rewrite("$<.Status")
	return status
}

// Status
func (*Window) SetStatus(status string) {
	js.Rewrite("$<.Status = $1", status)
}

// Statusbar
func (*Window) Statusbar() (statusbar *barprop.BarProp) {
	js.Rewrite("$<.Statusbar")
	return statusbar
}

// StyleMedia
func (*Window) StyleMedia() (styleMedia *stylemedia.StyleMedia) {
	js.Rewrite("$<.StyleMedia")
	return styleMedia
}

// Toolbar
func (*Window) Toolbar() (toolbar *barprop.BarProp) {
	js.Rewrite("$<.Toolbar")
	return toolbar
}

// Top
func (*Window) Top() (top *Window) {
	js.Rewrite("$<.Top")
	return top
}

// Window
func (*Window) Window() (window *Window) {
	js.Rewrite("$<.Window")
	return window
}
