package htmlvideoelement

import (
	"github.com/matthewmueller/golly/dom/avtrack"
	"github.com/matthewmueller/golly/dom/childnode"
	"github.com/matthewmueller/golly/dom/clientrect"
	"github.com/matthewmueller/golly/dom/clientrectlist"
	"github.com/matthewmueller/golly/dom/domstringmap"
	"github.com/matthewmueller/golly/dom/domtokenlist"
	"github.com/matthewmueller/golly/dom/htmlmediaelement"
	"github.com/matthewmueller/golly/dom/mediaencryptedevent"
	"github.com/matthewmueller/golly/dom/mediaerror"
	"github.com/matthewmueller/golly/dom/msgraphicstrust"
	"github.com/matthewmueller/golly/dom/msmediakeyneededevent"
	"github.com/matthewmueller/golly/dom/msmediakeys"
	"github.com/matthewmueller/golly/dom/mszoomtooptions"
	"github.com/matthewmueller/golly/dom/texttrack"
	"github.com/matthewmueller/golly/dom/texttracklist"
	"github.com/matthewmueller/golly/dom/timeranges"
	"github.com/matthewmueller/golly/dom/videoplaybackquality"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ htmlmediaelement.HTMLMediaElement = (*HTMLVideoElement)(nil)
var _ window.HTMLElement = (*HTMLVideoElement)(nil)
var _ window.Element = (*HTMLVideoElement)(nil)
var _ window.GlobalEventHandlers = (*HTMLVideoElement)(nil)
var _ window.ElementTraversal = (*HTMLVideoElement)(nil)
var _ window.NodeSelector = (*HTMLVideoElement)(nil)
var _ childnode.ChildNode = (*HTMLVideoElement)(nil)
var _ window.Node = (*HTMLVideoElement)(nil)
var _ window.EventTarget = (*HTMLVideoElement)(nil)

// HTMLVideoElement struct
// js:"HTMLVideoElement,omit"
type HTMLVideoElement struct {
}

// GetVideoPlaybackQuality fn
// js:"getVideoPlaybackQuality"
func (*HTMLVideoElement) GetVideoPlaybackQuality() (v *videoplaybackquality.VideoPlaybackQuality) {
	js.Rewrite("$_.getVideoPlaybackQuality()")
	return v
}

// MsFrameStep fn
// js:"msFrameStep"
func (*HTMLVideoElement) MsFrameStep(forward bool) {
	js.Rewrite("$_.msFrameStep($1)", forward)
}

// MsInsertVideoEffect fn
// js:"msInsertVideoEffect"
func (*HTMLVideoElement) MsInsertVideoEffect(activatableClassId string, effectRequired bool, config *interface{}) {
	js.Rewrite("$_.msInsertVideoEffect($1, $2, $3)", activatableClassId, effectRequired, config)
}

// MsSetVideoRectangle fn
// js:"msSetVideoRectangle"
func (*HTMLVideoElement) MsSetVideoRectangle(left float32, top float32, right float32, bottom float32) {
	js.Rewrite("$_.msSetVideoRectangle($1, $2, $3, $4)", left, top, right, bottom)
}

// WebkitEnterFullscreen fn
// js:"webkitEnterFullscreen"
func (*HTMLVideoElement) WebkitEnterFullscreen() {
	js.Rewrite("$_.webkitEnterFullscreen()")
}

// WebkitEnterFullScreen fn
// js:"webkitEnterFullScreen"
func (*HTMLVideoElement) WebkitEnterFullScreen() {
	js.Rewrite("$_.webkitEnterFullScreen()")
}

// WebkitExitFullscreen fn
// js:"webkitExitFullscreen"
func (*HTMLVideoElement) WebkitExitFullscreen() {
	js.Rewrite("$_.webkitExitFullscreen()")
}

// WebkitExitFullScreen fn
// js:"webkitExitFullScreen"
func (*HTMLVideoElement) WebkitExitFullScreen() {
	js.Rewrite("$_.webkitExitFullScreen()")
}

// AddTextTrack fn
// js:"addTextTrack"
func (*HTMLVideoElement) AddTextTrack(kind string, label *string, language *string) (t *texttrack.TextTrack) {
	js.Rewrite("$_.addTextTrack($1, $2, $3)", kind, label, language)
	return t
}

// CanPlayType fn Returns a string that specifies whether the client can play a given media resource type.
// js:"canPlayType"
func (*HTMLVideoElement) CanPlayType(kind string) (s string) {
	js.Rewrite("$_.canPlayType($1)", kind)
	return s
}

// Load fn Resets the audio or video object and loads a new media resource.
// js:"load"
func (*HTMLVideoElement) Load() {
	js.Rewrite("$_.load()")
}

// MsClearEffects fn Clears all effects from the media pipeline.
// js:"msClearEffects"
func (*HTMLVideoElement) MsClearEffects() {
	js.Rewrite("$_.msClearEffects()")
}

// MsGetAsCastingSource fn
// js:"msGetAsCastingSource"
func (*HTMLVideoElement) MsGetAsCastingSource() (i interface{}) {
	js.Rewrite("$_.msGetAsCastingSource()")
	return i
}

// MsInsertAudioEffect fn Inserts the specified audio effect into media pipeline.
// js:"msInsertAudioEffect"
func (*HTMLVideoElement) MsInsertAudioEffect(activatableClassId string, effectRequired bool, config *interface{}) {
	js.Rewrite("$_.msInsertAudioEffect($1, $2, $3)", activatableClassId, effectRequired, config)
}

// MsSetMediaKeys fn
// js:"msSetMediaKeys"
func (*HTMLVideoElement) MsSetMediaKeys(mediaKeys *msmediakeys.MSMediaKeys) {
	js.Rewrite("$_.msSetMediaKeys($1)", mediaKeys)
}

// MsSetMediaProtectionManager fn Specifies the media protection manager for a given media pipeline.
// js:"msSetMediaProtectionManager"
func (*HTMLVideoElement) MsSetMediaProtectionManager(mediaProtectionManager *interface{}) {
	js.Rewrite("$_.msSetMediaProtectionManager($1)", mediaProtectionManager)
}

// Pause fn Pauses the current playback and sets paused to TRUE. This can be used to test whether the media is playing or paused. You can also use the pause or play events to tell whether the media is playing or not.
// js:"pause"
func (*HTMLVideoElement) Pause() {
	js.Rewrite("$_.pause()")
}

// Play fn Loads and starts playback of a media resource.
// js:"play"
func (*HTMLVideoElement) Play() {
	js.Rewrite("$_.play()")
}

// SetMediaKeys fn
// js:"setMediaKeys"
func (*HTMLVideoElement) SetMediaKeys(mediaKeys *window.MediaKeys) {
	js.Rewrite("await $_.setMediaKeys($1)", mediaKeys)
}

// Blur fn
// js:"blur"
func (*HTMLVideoElement) Blur() {
	js.Rewrite("$_.blur()")
}

// Click fn
// js:"click"
func (*HTMLVideoElement) Click() {
	js.Rewrite("$_.click()")
}

// DragDrop fn
// js:"dragDrop"
func (*HTMLVideoElement) DragDrop() (b bool) {
	js.Rewrite("$_.dragDrop()")
	return b
}

// Focus fn
// js:"focus"
func (*HTMLVideoElement) Focus() {
	js.Rewrite("$_.focus()")
}

// GetElementsByClassName fn
// js:"getElementsByClassName"
func (*HTMLVideoElement) GetElementsByClassName(classNames string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByClassName($1)", classNames)
	return w
}

// InsertAdjacentElement fn
// js:"insertAdjacentElement"
func (*HTMLVideoElement) InsertAdjacentElement(position string, insertedElement window.Element) (w window.Element) {
	js.Rewrite("$_.insertAdjacentElement($1, $2)", position, insertedElement)
	return w
}

// InsertAdjacentHTML fn
// js:"insertAdjacentHTML"
func (*HTMLVideoElement) InsertAdjacentHTML(where string, html string) {
	js.Rewrite("$_.insertAdjacentHTML($1, $2)", where, html)
}

// InsertAdjacentText fn
// js:"insertAdjacentText"
func (*HTMLVideoElement) InsertAdjacentText(where string, text string) {
	js.Rewrite("$_.insertAdjacentText($1, $2)", where, text)
}

// MsGetInputContext fn
// js:"msGetInputContext"
func (*HTMLVideoElement) MsGetInputContext() (w *window.MSInputMethodContext) {
	js.Rewrite("$_.msGetInputContext()")
	return w
}

// ScrollIntoView fn
// js:"scrollIntoView"
func (*HTMLVideoElement) ScrollIntoView(top *bool) {
	js.Rewrite("$_.scrollIntoView($1)", top)
}

// GetAttribute fn
// js:"getAttribute"
func (*HTMLVideoElement) GetAttribute(qualifiedName string) (s string) {
	js.Rewrite("$_.getAttribute($1)", qualifiedName)
	return s
}

// GetAttributeNode fn
// js:"getAttributeNode"
func (*HTMLVideoElement) GetAttributeNode(name string) (w *window.Attr) {
	js.Rewrite("$_.getAttributeNode($1)", name)
	return w
}

// GetAttributeNodeNS fn
// js:"getAttributeNodeNS"
func (*HTMLVideoElement) GetAttributeNodeNS(namespaceURI string, localName string) (w *window.Attr) {
	js.Rewrite("$_.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return w
}

// GetAttributeNS fn
// js:"getAttributeNS"
func (*HTMLVideoElement) GetAttributeNS(namespaceURI string, localName string) (s string) {
	js.Rewrite("$_.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

// GetBoundingClientRect fn
// js:"getBoundingClientRect"
func (*HTMLVideoElement) GetBoundingClientRect() (c *clientrect.ClientRect) {
	js.Rewrite("$_.getBoundingClientRect()")
	return c
}

// GetClientRects fn
// js:"getClientRects"
func (*HTMLVideoElement) GetClientRects() (c *clientrectlist.ClientRectList) {
	js.Rewrite("$_.getClientRects()")
	return c
}

// GetElementsByTagName fn
// js:"getElementsByTagName"
func (*HTMLVideoElement) GetElementsByTagName(name string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByTagName($1)", name)
	return w
}

// GetElementsByTagNameNS fn
// js:"getElementsByTagNameNS"
func (*HTMLVideoElement) GetElementsByTagNameNS(namespaceURI string, localName string) (w *window.NodeList) {
	js.Rewrite("$_.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return w
}

// HasAttribute fn
// js:"hasAttribute"
func (*HTMLVideoElement) HasAttribute(name string) (b bool) {
	js.Rewrite("$_.hasAttribute($1)", name)
	return b
}

// HasAttributeNS fn
// js:"hasAttributeNS"
func (*HTMLVideoElement) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	js.Rewrite("$_.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

// MsGetRegionContent fn
// js:"msGetRegionContent"
func (*HTMLVideoElement) MsGetRegionContent() (w *window.MSRangeCollection) {
	js.Rewrite("$_.msGetRegionContent()")
	return w
}

// MsGetUntransformedBounds fn
// js:"msGetUntransformedBounds"
func (*HTMLVideoElement) MsGetUntransformedBounds() (c *clientrect.ClientRect) {
	js.Rewrite("$_.msGetUntransformedBounds()")
	return c
}

// MsMatchesSelector fn
// js:"msMatchesSelector"
func (*HTMLVideoElement) MsMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$_.msMatchesSelector($1)", selectors)
	return b
}

// MsReleasePointerCapture fn
// js:"msReleasePointerCapture"
func (*HTMLVideoElement) MsReleasePointerCapture(pointerId int) {
	js.Rewrite("$_.msReleasePointerCapture($1)", pointerId)
}

// MsSetPointerCapture fn
// js:"msSetPointerCapture"
func (*HTMLVideoElement) MsSetPointerCapture(pointerId int) {
	js.Rewrite("$_.msSetPointerCapture($1)", pointerId)
}

// MsZoomTo fn
// js:"msZoomTo"
func (*HTMLVideoElement) MsZoomTo(args *mszoomtooptions.MsZoomToOptions) {
	js.Rewrite("$_.msZoomTo($1)", args)
}

// ReleasePointerCapture fn
// js:"releasePointerCapture"
func (*HTMLVideoElement) ReleasePointerCapture(pointerId int) {
	js.Rewrite("$_.releasePointerCapture($1)", pointerId)
}

// RemoveAttribute fn
// js:"removeAttribute"
func (*HTMLVideoElement) RemoveAttribute(qualifiedName string) {
	js.Rewrite("$_.removeAttribute($1)", qualifiedName)
}

// RemoveAttributeNode fn
// js:"removeAttributeNode"
func (*HTMLVideoElement) RemoveAttributeNode(oldAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.removeAttributeNode($1)", oldAttr)
	return w
}

// RemoveAttributeNS fn
// js:"removeAttributeNS"
func (*HTMLVideoElement) RemoveAttributeNS(namespaceURI string, localName string) {
	js.Rewrite("$_.removeAttributeNS($1, $2)", namespaceURI, localName)
}

// RequestFullscreen fn
// js:"requestFullscreen"
func (*HTMLVideoElement) RequestFullscreen() {
	js.Rewrite("$_.requestFullscreen()")
}

// RequestPointerLock fn
// js:"requestPointerLock"
func (*HTMLVideoElement) RequestPointerLock() {
	js.Rewrite("$_.requestPointerLock()")
}

// SetAttribute fn
// js:"setAttribute"
func (*HTMLVideoElement) SetAttribute(qualifiedName string, value string) {
	js.Rewrite("$_.setAttribute($1, $2)", qualifiedName, value)
}

// SetAttributeNode fn
// js:"setAttributeNode"
func (*HTMLVideoElement) SetAttributeNode(newAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.setAttributeNode($1)", newAttr)
	return w
}

// SetAttributeNodeNS fn
// js:"setAttributeNodeNS"
func (*HTMLVideoElement) SetAttributeNodeNS(newAttr *window.Attr) (w *window.Attr) {
	js.Rewrite("$_.setAttributeNodeNS($1)", newAttr)
	return w
}

// SetAttributeNS fn
// js:"setAttributeNS"
func (*HTMLVideoElement) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	js.Rewrite("$_.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

// SetPointerCapture fn
// js:"setPointerCapture"
func (*HTMLVideoElement) SetPointerCapture(pointerId int) {
	js.Rewrite("$_.setPointerCapture($1)", pointerId)
}

// WebkitMatchesSelector fn
// js:"webkitMatchesSelector"
func (*HTMLVideoElement) WebkitMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$_.webkitMatchesSelector($1)", selectors)
	return b
}

// WebkitRequestFullscreen fn
// js:"webkitRequestFullscreen"
func (*HTMLVideoElement) WebkitRequestFullscreen() {
	js.Rewrite("$_.webkitRequestFullscreen()")
}

// WebkitRequestFullScreen fn
// js:"webkitRequestFullScreen"
func (*HTMLVideoElement) WebkitRequestFullScreen() {
	js.Rewrite("$_.webkitRequestFullScreen()")
}

// QuerySelector fn
// js:"querySelector"
func (*HTMLVideoElement) QuerySelector(selectors string) (w window.Element) {
	js.Rewrite("$_.querySelector($1)", selectors)
	return w
}

// QuerySelectorAll fn
// js:"querySelectorAll"
func (*HTMLVideoElement) QuerySelectorAll(selectors string) (w *window.NodeList) {
	js.Rewrite("$_.querySelectorAll($1)", selectors)
	return w
}

// AppendChild fn
// js:"appendChild"
func (*HTMLVideoElement) AppendChild(newChild window.Node) (w window.Node) {
	js.Rewrite("$_.appendChild($1)", newChild)
	return w
}

// CloneNode fn
// js:"cloneNode"
func (*HTMLVideoElement) CloneNode(deep *bool) (w window.Node) {
	js.Rewrite("$_.cloneNode($1)", deep)
	return w
}

// CompareDocumentPosition fn
// js:"compareDocumentPosition"
func (*HTMLVideoElement) CompareDocumentPosition(other window.Node) (u uint8) {
	js.Rewrite("$_.compareDocumentPosition($1)", other)
	return u
}

// Contains fn
// js:"contains"
func (*HTMLVideoElement) Contains(child window.Node) (b bool) {
	js.Rewrite("$_.contains($1)", child)
	return b
}

// HasAttributes fn
// js:"hasAttributes"
func (*HTMLVideoElement) HasAttributes() (b bool) {
	js.Rewrite("$_.hasAttributes()")
	return b
}

// HasChildNodes fn
// js:"hasChildNodes"
func (*HTMLVideoElement) HasChildNodes() (b bool) {
	js.Rewrite("$_.hasChildNodes()")
	return b
}

// InsertBefore fn
// js:"insertBefore"
func (*HTMLVideoElement) InsertBefore(newChild window.Node, refChild *window.Node) (w window.Node) {
	js.Rewrite("$_.insertBefore($1, $2)", newChild, refChild)
	return w
}

// IsDefaultNamespace fn
// js:"isDefaultNamespace"
func (*HTMLVideoElement) IsDefaultNamespace(namespaceURI string) (b bool) {
	js.Rewrite("$_.isDefaultNamespace($1)", namespaceURI)
	return b
}

// IsEqualNode fn
// js:"isEqualNode"
func (*HTMLVideoElement) IsEqualNode(arg window.Node) (b bool) {
	js.Rewrite("$_.isEqualNode($1)", arg)
	return b
}

// IsSameNode fn
// js:"isSameNode"
func (*HTMLVideoElement) IsSameNode(other window.Node) (b bool) {
	js.Rewrite("$_.isSameNode($1)", other)
	return b
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*HTMLVideoElement) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}

// LookupPrefix fn
// js:"lookupPrefix"
func (*HTMLVideoElement) LookupPrefix(namespaceURI string) (s string) {
	js.Rewrite("$_.lookupPrefix($1)", namespaceURI)
	return s
}

// Normalize fn
// js:"normalize"
func (*HTMLVideoElement) Normalize() {
	js.Rewrite("$_.normalize()")
}

// RemoveChild fn
// js:"removeChild"
func (*HTMLVideoElement) RemoveChild(oldChild window.Node) (w window.Node) {
	js.Rewrite("$_.removeChild($1)", oldChild)
	return w
}

// ReplaceChild fn
// js:"replaceChild"
func (*HTMLVideoElement) ReplaceChild(newChild window.Node, oldChild window.Node) (w window.Node) {
	js.Rewrite("$_.replaceChild($1, $2)", newChild, oldChild)
	return w
}

// AddEventListener fn
// js:"addEventListener"
func (*HTMLVideoElement) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*HTMLVideoElement) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*HTMLVideoElement) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Height prop Gets or sets the height of the video element.
// js:"height"
func (*HTMLVideoElement) Height() (height uint) {
	js.Rewrite("$_.height")
	return height
}

// SetHeight prop Gets or sets the height of the video element.
// js:"height"
func (*HTMLVideoElement) SetHeight(height uint) {
	js.Rewrite("$_.height = $1", height)
}

// MsHorizontalMirror prop
// js:"msHorizontalMirror"
func (*HTMLVideoElement) MsHorizontalMirror() (msHorizontalMirror bool) {
	js.Rewrite("$_.msHorizontalMirror")
	return msHorizontalMirror
}

// SetMsHorizontalMirror prop
// js:"msHorizontalMirror"
func (*HTMLVideoElement) SetMsHorizontalMirror(msHorizontalMirror bool) {
	js.Rewrite("$_.msHorizontalMirror = $1", msHorizontalMirror)
}

// MsIsLayoutOptimalForPlayback prop
// js:"msIsLayoutOptimalForPlayback"
func (*HTMLVideoElement) MsIsLayoutOptimalForPlayback() (msIsLayoutOptimalForPlayback bool) {
	js.Rewrite("$_.msIsLayoutOptimalForPlayback")
	return msIsLayoutOptimalForPlayback
}

// MsIsStereo3d prop
// js:"msIsStereo3D"
func (*HTMLVideoElement) MsIsStereo3d() (msIsStereo3D bool) {
	js.Rewrite("$_.msIsStereo3D")
	return msIsStereo3D
}

// MsStereo3dPackingMode prop
// js:"msStereo3DPackingMode"
func (*HTMLVideoElement) MsStereo3dPackingMode() (msStereo3DPackingMode string) {
	js.Rewrite("$_.msStereo3DPackingMode")
	return msStereo3DPackingMode
}

// SetMsStereo3dPackingMode prop
// js:"msStereo3DPackingMode"
func (*HTMLVideoElement) SetMsStereo3dPackingMode(msStereo3DPackingMode string) {
	js.Rewrite("$_.msStereo3DPackingMode = $1", msStereo3DPackingMode)
}

// MsStereo3dRenderMode prop
// js:"msStereo3DRenderMode"
func (*HTMLVideoElement) MsStereo3dRenderMode() (msStereo3DRenderMode string) {
	js.Rewrite("$_.msStereo3DRenderMode")
	return msStereo3DRenderMode
}

// SetMsStereo3dRenderMode prop
// js:"msStereo3DRenderMode"
func (*HTMLVideoElement) SetMsStereo3dRenderMode(msStereo3DRenderMode string) {
	js.Rewrite("$_.msStereo3DRenderMode = $1", msStereo3DRenderMode)
}

// MsZoom prop
// js:"msZoom"
func (*HTMLVideoElement) MsZoom() (msZoom bool) {
	js.Rewrite("$_.msZoom")
	return msZoom
}

// SetMsZoom prop
// js:"msZoom"
func (*HTMLVideoElement) SetMsZoom(msZoom bool) {
	js.Rewrite("$_.msZoom = $1", msZoom)
}

// OnMSVideoFormatChanged prop
// js:"onMSVideoFormatChanged"
func (*HTMLVideoElement) OnMSVideoFormatChanged() (onmsvideoformatchanged func(window.Event)) {
	js.Rewrite("$_.onMSVideoFormatChanged")
	return onmsvideoformatchanged
}

// SetOnMSVideoFormatChanged prop
// js:"onMSVideoFormatChanged"
func (*HTMLVideoElement) SetOnMSVideoFormatChanged(onmsvideoformatchanged func(window.Event)) {
	js.Rewrite("$_.onMSVideoFormatChanged = $1", onmsvideoformatchanged)
}

// OnMSVideoFrameStepCompleted prop
// js:"onMSVideoFrameStepCompleted"
func (*HTMLVideoElement) OnMSVideoFrameStepCompleted() (onmsvideoframestepcompleted func(window.Event)) {
	js.Rewrite("$_.onMSVideoFrameStepCompleted")
	return onmsvideoframestepcompleted
}

// SetOnMSVideoFrameStepCompleted prop
// js:"onMSVideoFrameStepCompleted"
func (*HTMLVideoElement) SetOnMSVideoFrameStepCompleted(onmsvideoframestepcompleted func(window.Event)) {
	js.Rewrite("$_.onMSVideoFrameStepCompleted = $1", onmsvideoframestepcompleted)
}

// OnMSVideoOptimalLayoutChanged prop
// js:"onMSVideoOptimalLayoutChanged"
func (*HTMLVideoElement) OnMSVideoOptimalLayoutChanged() (onmsvideooptimallayoutchanged func(window.Event)) {
	js.Rewrite("$_.onMSVideoOptimalLayoutChanged")
	return onmsvideooptimallayoutchanged
}

// SetOnMSVideoOptimalLayoutChanged prop
// js:"onMSVideoOptimalLayoutChanged"
func (*HTMLVideoElement) SetOnMSVideoOptimalLayoutChanged(onmsvideooptimallayoutchanged func(window.Event)) {
	js.Rewrite("$_.onMSVideoOptimalLayoutChanged = $1", onmsvideooptimallayoutchanged)
}

// Poster prop Gets or sets a URL of an image to display, for example, like a movie poster. This can be a still frame from the video, or another image if no video data is available.
// js:"poster"
func (*HTMLVideoElement) Poster() (poster string) {
	js.Rewrite("$_.poster")
	return poster
}

// SetPoster prop Gets or sets a URL of an image to display, for example, like a movie poster. This can be a still frame from the video, or another image if no video data is available.
// js:"poster"
func (*HTMLVideoElement) SetPoster(poster string) {
	js.Rewrite("$_.poster = $1", poster)
}

// VideoHeight prop Gets the intrinsic height of a video in CSS pixels, or zero if the dimensions are not known.
// js:"videoHeight"
func (*HTMLVideoElement) VideoHeight() (videoHeight uint) {
	js.Rewrite("$_.videoHeight")
	return videoHeight
}

// VideoWidth prop Gets the intrinsic width of a video in CSS pixels, or zero if the dimensions are not known.
// js:"videoWidth"
func (*HTMLVideoElement) VideoWidth() (videoWidth uint) {
	js.Rewrite("$_.videoWidth")
	return videoWidth
}

// WebkitDisplayingFullscreen prop
// js:"webkitDisplayingFullscreen"
func (*HTMLVideoElement) WebkitDisplayingFullscreen() (webkitDisplayingFullscreen bool) {
	js.Rewrite("$_.webkitDisplayingFullscreen")
	return webkitDisplayingFullscreen
}

// WebkitSupportsFullscreen prop
// js:"webkitSupportsFullscreen"
func (*HTMLVideoElement) WebkitSupportsFullscreen() (webkitSupportsFullscreen bool) {
	js.Rewrite("$_.webkitSupportsFullscreen")
	return webkitSupportsFullscreen
}

// Width prop Gets or sets the width of the video element.
// js:"width"
func (*HTMLVideoElement) Width() (width uint) {
	js.Rewrite("$_.width")
	return width
}

// SetWidth prop Gets or sets the width of the video element.
// js:"width"
func (*HTMLVideoElement) SetWidth(width uint) {
	js.Rewrite("$_.width = $1", width)
}

// AudioTracks prop Returns an AudioTrackList object with the audio tracks for a given video element.
// js:"audioTracks"
func (*HTMLVideoElement) AudioTracks() (audioTracks *avtrack.AudioTrackList) {
	js.Rewrite("$_.audioTracks")
	return audioTracks
}

// Autoplay prop Gets or sets a value that indicates whether to start playing the media automatically.
// js:"autoplay"
func (*HTMLVideoElement) Autoplay() (autoplay bool) {
	js.Rewrite("$_.autoplay")
	return autoplay
}

// SetAutoplay prop Gets or sets a value that indicates whether to start playing the media automatically.
// js:"autoplay"
func (*HTMLVideoElement) SetAutoplay(autoplay bool) {
	js.Rewrite("$_.autoplay = $1", autoplay)
}

// Buffered prop Gets a collection of buffered time ranges.
// js:"buffered"
func (*HTMLVideoElement) Buffered() (buffered *timeranges.TimeRanges) {
	js.Rewrite("$_.buffered")
	return buffered
}

// Controls prop Gets or sets a flag that indicates whether the client provides a set of controls for the media (in case the developer does not include controls for the player).
// js:"controls"
func (*HTMLVideoElement) Controls() (controls bool) {
	js.Rewrite("$_.controls")
	return controls
}

// SetControls prop Gets or sets a flag that indicates whether the client provides a set of controls for the media (in case the developer does not include controls for the player).
// js:"controls"
func (*HTMLVideoElement) SetControls(controls bool) {
	js.Rewrite("$_.controls = $1", controls)
}

// CrossOrigin prop
// js:"crossOrigin"
func (*HTMLVideoElement) CrossOrigin() (crossOrigin string) {
	js.Rewrite("$_.crossOrigin")
	return crossOrigin
}

// SetCrossOrigin prop
// js:"crossOrigin"
func (*HTMLVideoElement) SetCrossOrigin(crossOrigin string) {
	js.Rewrite("$_.crossOrigin = $1", crossOrigin)
}

// CurrentSrc prop Gets the address or URL of the current media resource that is selected by IHTMLMediaElement.
// js:"currentSrc"
func (*HTMLVideoElement) CurrentSrc() (currentSrc string) {
	js.Rewrite("$_.currentSrc")
	return currentSrc
}

// CurrentTime prop Gets or sets the current playback position, in seconds.
// js:"currentTime"
func (*HTMLVideoElement) CurrentTime() (currentTime float32) {
	js.Rewrite("$_.currentTime")
	return currentTime
}

// SetCurrentTime prop Gets or sets the current playback position, in seconds.
// js:"currentTime"
func (*HTMLVideoElement) SetCurrentTime(currentTime float32) {
	js.Rewrite("$_.currentTime = $1", currentTime)
}

// DefaultMuted prop
// js:"defaultMuted"
func (*HTMLVideoElement) DefaultMuted() (defaultMuted bool) {
	js.Rewrite("$_.defaultMuted")
	return defaultMuted
}

// SetDefaultMuted prop
// js:"defaultMuted"
func (*HTMLVideoElement) SetDefaultMuted(defaultMuted bool) {
	js.Rewrite("$_.defaultMuted = $1", defaultMuted)
}

// DefaultPlaybackRate prop Gets or sets the default playback rate when the user is not using fast forward or reverse for a video or audio resource.
// js:"defaultPlaybackRate"
func (*HTMLVideoElement) DefaultPlaybackRate() (defaultPlaybackRate float32) {
	js.Rewrite("$_.defaultPlaybackRate")
	return defaultPlaybackRate
}

// SetDefaultPlaybackRate prop Gets or sets the default playback rate when the user is not using fast forward or reverse for a video or audio resource.
// js:"defaultPlaybackRate"
func (*HTMLVideoElement) SetDefaultPlaybackRate(defaultPlaybackRate float32) {
	js.Rewrite("$_.defaultPlaybackRate = $1", defaultPlaybackRate)
}

// Duration prop Returns the duration in seconds of the current media resource. A NaN value is returned if duration is not available, or Infinity if the media resource is streaming.
// js:"duration"
func (*HTMLVideoElement) Duration() (duration float32) {
	js.Rewrite("$_.duration")
	return duration
}

// Ended prop Gets information about whether the playback has ended or not.
// js:"ended"
func (*HTMLVideoElement) Ended() (ended bool) {
	js.Rewrite("$_.ended")
	return ended
}

// Error prop Returns an object representing the current error state of the audio or video element.
// js:"error"
func (*HTMLVideoElement) Error() (err *mediaerror.MediaError) {
	js.Rewrite("$_.error")
	return err
}

// Loop prop Gets or sets a flag to specify whether playback should restart after it completes.
// js:"loop"
func (*HTMLVideoElement) Loop() (loop bool) {
	js.Rewrite("$_.loop")
	return loop
}

// SetLoop prop Gets or sets a flag to specify whether playback should restart after it completes.
// js:"loop"
func (*HTMLVideoElement) SetLoop(loop bool) {
	js.Rewrite("$_.loop = $1", loop)
}

// MediaKeys prop
// js:"mediaKeys"
func (*HTMLVideoElement) MediaKeys() (mediaKeys *window.MediaKeys) {
	js.Rewrite("$_.mediaKeys")
	return mediaKeys
}

// MsAudioCategory prop Specifies the purpose of the audio or video media, such as background audio or alerts.
// js:"msAudioCategory"
func (*HTMLVideoElement) MsAudioCategory() (msAudioCategory string) {
	js.Rewrite("$_.msAudioCategory")
	return msAudioCategory
}

// SetMsAudioCategory prop Specifies the purpose of the audio or video media, such as background audio or alerts.
// js:"msAudioCategory"
func (*HTMLVideoElement) SetMsAudioCategory(msAudioCategory string) {
	js.Rewrite("$_.msAudioCategory = $1", msAudioCategory)
}

// MsAudioDeviceType prop Specifies the output device id that the audio will be sent to.
// js:"msAudioDeviceType"
func (*HTMLVideoElement) MsAudioDeviceType() (msAudioDeviceType string) {
	js.Rewrite("$_.msAudioDeviceType")
	return msAudioDeviceType
}

// SetMsAudioDeviceType prop Specifies the output device id that the audio will be sent to.
// js:"msAudioDeviceType"
func (*HTMLVideoElement) SetMsAudioDeviceType(msAudioDeviceType string) {
	js.Rewrite("$_.msAudioDeviceType = $1", msAudioDeviceType)
}

// MsGraphicsTrustStatus prop
// js:"msGraphicsTrustStatus"
func (*HTMLVideoElement) MsGraphicsTrustStatus() (msGraphicsTrustStatus *msgraphicstrust.MSGraphicsTrust) {
	js.Rewrite("$_.msGraphicsTrustStatus")
	return msGraphicsTrustStatus
}

// MsKeys prop Gets the MSMediaKeys object, which is used for decrypting media data, that is associated with this media element.
// js:"msKeys"
func (*HTMLVideoElement) MsKeys() (msKeys *msmediakeys.MSMediaKeys) {
	js.Rewrite("$_.msKeys")
	return msKeys
}

// MsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
// js:"msPlayToDisabled"
func (*HTMLVideoElement) MsPlayToDisabled() (msPlayToDisabled bool) {
	js.Rewrite("$_.msPlayToDisabled")
	return msPlayToDisabled
}

// SetMsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
// js:"msPlayToDisabled"
func (*HTMLVideoElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	js.Rewrite("$_.msPlayToDisabled = $1", msPlayToDisabled)
}

// MsPlayToPreferredSourceURI prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
// js:"msPlayToPreferredSourceUri"
func (*HTMLVideoElement) MsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$_.msPlayToPreferredSourceUri")
	return msPlayToPreferredSourceUri
}

// SetMsPlayToPreferredSourceURI prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
// js:"msPlayToPreferredSourceUri"
func (*HTMLVideoElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	js.Rewrite("$_.msPlayToPreferredSourceUri = $1", msPlayToPreferredSourceUri)
}

// MsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
// js:"msPlayToPrimary"
func (*HTMLVideoElement) MsPlayToPrimary() (msPlayToPrimary bool) {
	js.Rewrite("$_.msPlayToPrimary")
	return msPlayToPrimary
}

// SetMsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
// js:"msPlayToPrimary"
func (*HTMLVideoElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	js.Rewrite("$_.msPlayToPrimary = $1", msPlayToPrimary)
}

// MsPlayToSource prop Gets the source associated with the media element for use by the PlayToManager.
// js:"msPlayToSource"
func (*HTMLVideoElement) MsPlayToSource() (msPlayToSource interface{}) {
	js.Rewrite("$_.msPlayToSource")
	return msPlayToSource
}

// MsRealTime prop Specifies whether or not to enable low-latency playback on the media element.
// js:"msRealTime"
func (*HTMLVideoElement) MsRealTime() (msRealTime bool) {
	js.Rewrite("$_.msRealTime")
	return msRealTime
}

// SetMsRealTime prop Specifies whether or not to enable low-latency playback on the media element.
// js:"msRealTime"
func (*HTMLVideoElement) SetMsRealTime(msRealTime bool) {
	js.Rewrite("$_.msRealTime = $1", msRealTime)
}

// Muted prop Gets or sets a flag that indicates whether the audio (either audio or the audio track on video media) is muted.
// js:"muted"
func (*HTMLVideoElement) Muted() (muted bool) {
	js.Rewrite("$_.muted")
	return muted
}

// SetMuted prop Gets or sets a flag that indicates whether the audio (either audio or the audio track on video media) is muted.
// js:"muted"
func (*HTMLVideoElement) SetMuted(muted bool) {
	js.Rewrite("$_.muted = $1", muted)
}

// NetworkState prop Gets the current network activity for the element.
// js:"networkState"
func (*HTMLVideoElement) NetworkState() (networkState uint8) {
	js.Rewrite("$_.networkState")
	return networkState
}

// Onencrypted prop
// js:"onencrypted"
func (*HTMLVideoElement) Onencrypted() (onencrypted func(*mediaencryptedevent.MediaEncryptedEvent)) {
	js.Rewrite("$_.onencrypted")
	return onencrypted
}

// SetOnencrypted prop
// js:"onencrypted"
func (*HTMLVideoElement) SetOnencrypted(onencrypted func(*mediaencryptedevent.MediaEncryptedEvent)) {
	js.Rewrite("$_.onencrypted = $1", onencrypted)
}

// Onmsneedkey prop
// js:"onmsneedkey"
func (*HTMLVideoElement) Onmsneedkey() (onmsneedkey func(*msmediakeyneededevent.MSMediaKeyNeededEvent)) {
	js.Rewrite("$_.onmsneedkey")
	return onmsneedkey
}

// SetOnmsneedkey prop
// js:"onmsneedkey"
func (*HTMLVideoElement) SetOnmsneedkey(onmsneedkey func(*msmediakeyneededevent.MSMediaKeyNeededEvent)) {
	js.Rewrite("$_.onmsneedkey = $1", onmsneedkey)
}

// Paused prop Gets a flag that specifies whether playback is paused.
// js:"paused"
func (*HTMLVideoElement) Paused() (paused bool) {
	js.Rewrite("$_.paused")
	return paused
}

// PlaybackRate prop Gets or sets the current rate of speed for the media resource to play. This speed is expressed as a multiple of the normal speed of the media resource.
// js:"playbackRate"
func (*HTMLVideoElement) PlaybackRate() (playbackRate float32) {
	js.Rewrite("$_.playbackRate")
	return playbackRate
}

// SetPlaybackRate prop Gets or sets the current rate of speed for the media resource to play. This speed is expressed as a multiple of the normal speed of the media resource.
// js:"playbackRate"
func (*HTMLVideoElement) SetPlaybackRate(playbackRate float32) {
	js.Rewrite("$_.playbackRate = $1", playbackRate)
}

// Played prop Gets TimeRanges for the current media resource that has been played.
// js:"played"
func (*HTMLVideoElement) Played() (played *timeranges.TimeRanges) {
	js.Rewrite("$_.played")
	return played
}

// Preload prop Gets or sets the current playback position, in seconds.
// js:"preload"
func (*HTMLVideoElement) Preload() (preload string) {
	js.Rewrite("$_.preload")
	return preload
}

// SetPreload prop Gets or sets the current playback position, in seconds.
// js:"preload"
func (*HTMLVideoElement) SetPreload(preload string) {
	js.Rewrite("$_.preload = $1", preload)
}

// ReadyState prop
// js:"readyState"
func (*HTMLVideoElement) ReadyState() (readyState interface{}) {
	js.Rewrite("$_.readyState")
	return readyState
}

// Seekable prop Returns a TimeRanges object that represents the ranges of the current media resource that can be seeked.
// js:"seekable"
func (*HTMLVideoElement) Seekable() (seekable *timeranges.TimeRanges) {
	js.Rewrite("$_.seekable")
	return seekable
}

// Seeking prop Gets a flag that indicates whether the the client is currently moving to a new playback position in the media resource.
// js:"seeking"
func (*HTMLVideoElement) Seeking() (seeking bool) {
	js.Rewrite("$_.seeking")
	return seeking
}

// Src prop The address or URL of the a media resource that is to be considered.
// js:"src"
func (*HTMLVideoElement) Src() (src string) {
	js.Rewrite("$_.src")
	return src
}

// SetSrc prop The address or URL of the a media resource that is to be considered.
// js:"src"
func (*HTMLVideoElement) SetSrc(src string) {
	js.Rewrite("$_.src = $1", src)
}

// SrcObject prop
// js:"srcObject"
func (*HTMLVideoElement) SrcObject() (srcObject *window.MediaStream) {
	js.Rewrite("$_.srcObject")
	return srcObject
}

// SetSrcObject prop
// js:"srcObject"
func (*HTMLVideoElement) SetSrcObject(srcObject *window.MediaStream) {
	js.Rewrite("$_.srcObject = $1", srcObject)
}

// TextTracks prop
// js:"textTracks"
func (*HTMLVideoElement) TextTracks() (textTracks *texttracklist.TextTrackList) {
	js.Rewrite("$_.textTracks")
	return textTracks
}

// VideoTracks prop
// js:"videoTracks"
func (*HTMLVideoElement) VideoTracks() (videoTracks *avtrack.VideoTrackList) {
	js.Rewrite("$_.videoTracks")
	return videoTracks
}

// Volume prop Gets or sets the volume level for audio portions of the media element.
// js:"volume"
func (*HTMLVideoElement) Volume() (volume float32) {
	js.Rewrite("$_.volume")
	return volume
}

// SetVolume prop Gets or sets the volume level for audio portions of the media element.
// js:"volume"
func (*HTMLVideoElement) SetVolume(volume float32) {
	js.Rewrite("$_.volume = $1", volume)
}

// AccessKey prop
// js:"accessKey"
func (*HTMLVideoElement) AccessKey() (accessKey string) {
	js.Rewrite("$_.accessKey")
	return accessKey
}

// SetAccessKey prop
// js:"accessKey"
func (*HTMLVideoElement) SetAccessKey(accessKey string) {
	js.Rewrite("$_.accessKey = $1", accessKey)
}

// Children prop
// js:"children"
func (*HTMLVideoElement) Children() (children window.HTMLCollection) {
	js.Rewrite("$_.children")
	return children
}

// ContentEditable prop
// js:"contentEditable"
func (*HTMLVideoElement) ContentEditable() (contentEditable string) {
	js.Rewrite("$_.contentEditable")
	return contentEditable
}

// SetContentEditable prop
// js:"contentEditable"
func (*HTMLVideoElement) SetContentEditable(contentEditable string) {
	js.Rewrite("$_.contentEditable = $1", contentEditable)
}

// Dataset prop
// js:"dataset"
func (*HTMLVideoElement) Dataset() (dataset *domstringmap.DOMStringMap) {
	js.Rewrite("$_.dataset")
	return dataset
}

// Dir prop
// js:"dir"
func (*HTMLVideoElement) Dir() (dir string) {
	js.Rewrite("$_.dir")
	return dir
}

// SetDir prop
// js:"dir"
func (*HTMLVideoElement) SetDir(dir string) {
	js.Rewrite("$_.dir = $1", dir)
}

// Draggable prop
// js:"draggable"
func (*HTMLVideoElement) Draggable() (draggable bool) {
	js.Rewrite("$_.draggable")
	return draggable
}

// SetDraggable prop
// js:"draggable"
func (*HTMLVideoElement) SetDraggable(draggable bool) {
	js.Rewrite("$_.draggable = $1", draggable)
}

// Hidden prop
// js:"hidden"
func (*HTMLVideoElement) Hidden() (hidden bool) {
	js.Rewrite("$_.hidden")
	return hidden
}

// SetHidden prop
// js:"hidden"
func (*HTMLVideoElement) SetHidden(hidden bool) {
	js.Rewrite("$_.hidden = $1", hidden)
}

// HideFocus prop
// js:"hideFocus"
func (*HTMLVideoElement) HideFocus() (hideFocus bool) {
	js.Rewrite("$_.hideFocus")
	return hideFocus
}

// SetHideFocus prop
// js:"hideFocus"
func (*HTMLVideoElement) SetHideFocus(hideFocus bool) {
	js.Rewrite("$_.hideFocus = $1", hideFocus)
}

// InnerText prop
// js:"innerText"
func (*HTMLVideoElement) InnerText() (innerText string) {
	js.Rewrite("$_.innerText")
	return innerText
}

// SetInnerText prop
// js:"innerText"
func (*HTMLVideoElement) SetInnerText(innerText string) {
	js.Rewrite("$_.innerText = $1", innerText)
}

// IsContentEditable prop
// js:"isContentEditable"
func (*HTMLVideoElement) IsContentEditable() (isContentEditable bool) {
	js.Rewrite("$_.isContentEditable")
	return isContentEditable
}

// Lang prop
// js:"lang"
func (*HTMLVideoElement) Lang() (lang string) {
	js.Rewrite("$_.lang")
	return lang
}

// SetLang prop
// js:"lang"
func (*HTMLVideoElement) SetLang(lang string) {
	js.Rewrite("$_.lang = $1", lang)
}

// OffsetHeight prop
// js:"offsetHeight"
func (*HTMLVideoElement) OffsetHeight() (offsetHeight int) {
	js.Rewrite("$_.offsetHeight")
	return offsetHeight
}

// OffsetLeft prop
// js:"offsetLeft"
func (*HTMLVideoElement) OffsetLeft() (offsetLeft int) {
	js.Rewrite("$_.offsetLeft")
	return offsetLeft
}

// OffsetParent prop
// js:"offsetParent"
func (*HTMLVideoElement) OffsetParent() (offsetParent window.Element) {
	js.Rewrite("$_.offsetParent")
	return offsetParent
}

// OffsetTop prop
// js:"offsetTop"
func (*HTMLVideoElement) OffsetTop() (offsetTop int) {
	js.Rewrite("$_.offsetTop")
	return offsetTop
}

// OffsetWidth prop
// js:"offsetWidth"
func (*HTMLVideoElement) OffsetWidth() (offsetWidth int) {
	js.Rewrite("$_.offsetWidth")
	return offsetWidth
}

// Onabort prop
// js:"onabort"
func (*HTMLVideoElement) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*HTMLVideoElement) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$_.onabort = $1", onabort)
}

// Onactivate prop
// js:"onactivate"
func (*HTMLVideoElement) Onactivate() (onactivate func(window.UIEvent)) {
	js.Rewrite("$_.onactivate")
	return onactivate
}

// SetOnactivate prop
// js:"onactivate"
func (*HTMLVideoElement) SetOnactivate(onactivate func(window.UIEvent)) {
	js.Rewrite("$_.onactivate = $1", onactivate)
}

// Onbeforeactivate prop
// js:"onbeforeactivate"
func (*HTMLVideoElement) Onbeforeactivate() (onbeforeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforeactivate")
	return onbeforeactivate
}

// SetOnbeforeactivate prop
// js:"onbeforeactivate"
func (*HTMLVideoElement) SetOnbeforeactivate(onbeforeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforeactivate = $1", onbeforeactivate)
}

// Onbeforecopy prop
// js:"onbeforecopy"
func (*HTMLVideoElement) Onbeforecopy() (onbeforecopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecopy")
	return onbeforecopy
}

// SetOnbeforecopy prop
// js:"onbeforecopy"
func (*HTMLVideoElement) SetOnbeforecopy(onbeforecopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecopy = $1", onbeforecopy)
}

// Onbeforecut prop
// js:"onbeforecut"
func (*HTMLVideoElement) Onbeforecut() (onbeforecut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecut")
	return onbeforecut
}

// SetOnbeforecut prop
// js:"onbeforecut"
func (*HTMLVideoElement) SetOnbeforecut(onbeforecut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforecut = $1", onbeforecut)
}

// Onbeforedeactivate prop
// js:"onbeforedeactivate"
func (*HTMLVideoElement) Onbeforedeactivate() (onbeforedeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforedeactivate")
	return onbeforedeactivate
}

// SetOnbeforedeactivate prop
// js:"onbeforedeactivate"
func (*HTMLVideoElement) SetOnbeforedeactivate(onbeforedeactivate func(window.UIEvent)) {
	js.Rewrite("$_.onbeforedeactivate = $1", onbeforedeactivate)
}

// Onbeforepaste prop
// js:"onbeforepaste"
func (*HTMLVideoElement) Onbeforepaste() (onbeforepaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforepaste")
	return onbeforepaste
}

// SetOnbeforepaste prop
// js:"onbeforepaste"
func (*HTMLVideoElement) SetOnbeforepaste(onbeforepaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onbeforepaste = $1", onbeforepaste)
}

// Onblur prop
// js:"onblur"
func (*HTMLVideoElement) Onblur() (onblur func(*window.FocusEvent)) {
	js.Rewrite("$_.onblur")
	return onblur
}

// SetOnblur prop
// js:"onblur"
func (*HTMLVideoElement) SetOnblur(onblur func(*window.FocusEvent)) {
	js.Rewrite("$_.onblur = $1", onblur)
}

// Oncanplay prop
// js:"oncanplay"
func (*HTMLVideoElement) Oncanplay() (oncanplay func(window.Event)) {
	js.Rewrite("$_.oncanplay")
	return oncanplay
}

// SetOncanplay prop
// js:"oncanplay"
func (*HTMLVideoElement) SetOncanplay(oncanplay func(window.Event)) {
	js.Rewrite("$_.oncanplay = $1", oncanplay)
}

// Oncanplaythrough prop
// js:"oncanplaythrough"
func (*HTMLVideoElement) Oncanplaythrough() (oncanplaythrough func(window.Event)) {
	js.Rewrite("$_.oncanplaythrough")
	return oncanplaythrough
}

// SetOncanplaythrough prop
// js:"oncanplaythrough"
func (*HTMLVideoElement) SetOncanplaythrough(oncanplaythrough func(window.Event)) {
	js.Rewrite("$_.oncanplaythrough = $1", oncanplaythrough)
}

// Onchange prop
// js:"onchange"
func (*HTMLVideoElement) Onchange() (onchange func(window.Event)) {
	js.Rewrite("$_.onchange")
	return onchange
}

// SetOnchange prop
// js:"onchange"
func (*HTMLVideoElement) SetOnchange(onchange func(window.Event)) {
	js.Rewrite("$_.onchange = $1", onchange)
}

// Onclick prop
// js:"onclick"
func (*HTMLVideoElement) Onclick() (onclick func(window.MouseEvent)) {
	js.Rewrite("$_.onclick")
	return onclick
}

// SetOnclick prop
// js:"onclick"
func (*HTMLVideoElement) SetOnclick(onclick func(window.MouseEvent)) {
	js.Rewrite("$_.onclick = $1", onclick)
}

// Oncontextmenu prop
// js:"oncontextmenu"
func (*HTMLVideoElement) Oncontextmenu() (oncontextmenu func(*window.PointerEvent)) {
	js.Rewrite("$_.oncontextmenu")
	return oncontextmenu
}

// SetOncontextmenu prop
// js:"oncontextmenu"
func (*HTMLVideoElement) SetOncontextmenu(oncontextmenu func(*window.PointerEvent)) {
	js.Rewrite("$_.oncontextmenu = $1", oncontextmenu)
}

// Oncopy prop
// js:"oncopy"
func (*HTMLVideoElement) Oncopy() (oncopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncopy")
	return oncopy
}

// SetOncopy prop
// js:"oncopy"
func (*HTMLVideoElement) SetOncopy(oncopy func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncopy = $1", oncopy)
}

// Oncuechange prop
// js:"oncuechange"
func (*HTMLVideoElement) Oncuechange() (oncuechange func(window.Event)) {
	js.Rewrite("$_.oncuechange")
	return oncuechange
}

// SetOncuechange prop
// js:"oncuechange"
func (*HTMLVideoElement) SetOncuechange(oncuechange func(window.Event)) {
	js.Rewrite("$_.oncuechange = $1", oncuechange)
}

// Oncut prop
// js:"oncut"
func (*HTMLVideoElement) Oncut() (oncut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncut")
	return oncut
}

// SetOncut prop
// js:"oncut"
func (*HTMLVideoElement) SetOncut(oncut func(*window.ClipboardEvent)) {
	js.Rewrite("$_.oncut = $1", oncut)
}

// Ondblclick prop
// js:"ondblclick"
func (*HTMLVideoElement) Ondblclick() (ondblclick func(window.MouseEvent)) {
	js.Rewrite("$_.ondblclick")
	return ondblclick
}

// SetOndblclick prop
// js:"ondblclick"
func (*HTMLVideoElement) SetOndblclick(ondblclick func(window.MouseEvent)) {
	js.Rewrite("$_.ondblclick = $1", ondblclick)
}

// Ondeactivate prop
// js:"ondeactivate"
func (*HTMLVideoElement) Ondeactivate() (ondeactivate func(window.UIEvent)) {
	js.Rewrite("$_.ondeactivate")
	return ondeactivate
}

// SetOndeactivate prop
// js:"ondeactivate"
func (*HTMLVideoElement) SetOndeactivate(ondeactivate func(window.UIEvent)) {
	js.Rewrite("$_.ondeactivate = $1", ondeactivate)
}

// Ondrag prop
// js:"ondrag"
func (*HTMLVideoElement) Ondrag() (ondrag func(*window.DragEvent)) {
	js.Rewrite("$_.ondrag")
	return ondrag
}

// SetOndrag prop
// js:"ondrag"
func (*HTMLVideoElement) SetOndrag(ondrag func(*window.DragEvent)) {
	js.Rewrite("$_.ondrag = $1", ondrag)
}

// Ondragend prop
// js:"ondragend"
func (*HTMLVideoElement) Ondragend() (ondragend func(*window.DragEvent)) {
	js.Rewrite("$_.ondragend")
	return ondragend
}

// SetOndragend prop
// js:"ondragend"
func (*HTMLVideoElement) SetOndragend(ondragend func(*window.DragEvent)) {
	js.Rewrite("$_.ondragend = $1", ondragend)
}

// Ondragenter prop
// js:"ondragenter"
func (*HTMLVideoElement) Ondragenter() (ondragenter func(*window.DragEvent)) {
	js.Rewrite("$_.ondragenter")
	return ondragenter
}

// SetOndragenter prop
// js:"ondragenter"
func (*HTMLVideoElement) SetOndragenter(ondragenter func(*window.DragEvent)) {
	js.Rewrite("$_.ondragenter = $1", ondragenter)
}

// Ondragleave prop
// js:"ondragleave"
func (*HTMLVideoElement) Ondragleave() (ondragleave func(*window.DragEvent)) {
	js.Rewrite("$_.ondragleave")
	return ondragleave
}

// SetOndragleave prop
// js:"ondragleave"
func (*HTMLVideoElement) SetOndragleave(ondragleave func(*window.DragEvent)) {
	js.Rewrite("$_.ondragleave = $1", ondragleave)
}

// Ondragover prop
// js:"ondragover"
func (*HTMLVideoElement) Ondragover() (ondragover func(*window.DragEvent)) {
	js.Rewrite("$_.ondragover")
	return ondragover
}

// SetOndragover prop
// js:"ondragover"
func (*HTMLVideoElement) SetOndragover(ondragover func(*window.DragEvent)) {
	js.Rewrite("$_.ondragover = $1", ondragover)
}

// Ondragstart prop
// js:"ondragstart"
func (*HTMLVideoElement) Ondragstart() (ondragstart func(*window.DragEvent)) {
	js.Rewrite("$_.ondragstart")
	return ondragstart
}

// SetOndragstart prop
// js:"ondragstart"
func (*HTMLVideoElement) SetOndragstart(ondragstart func(*window.DragEvent)) {
	js.Rewrite("$_.ondragstart = $1", ondragstart)
}

// Ondrop prop
// js:"ondrop"
func (*HTMLVideoElement) Ondrop() (ondrop func(*window.DragEvent)) {
	js.Rewrite("$_.ondrop")
	return ondrop
}

// SetOndrop prop
// js:"ondrop"
func (*HTMLVideoElement) SetOndrop(ondrop func(*window.DragEvent)) {
	js.Rewrite("$_.ondrop = $1", ondrop)
}

// Ondurationchange prop
// js:"ondurationchange"
func (*HTMLVideoElement) Ondurationchange() (ondurationchange func(window.Event)) {
	js.Rewrite("$_.ondurationchange")
	return ondurationchange
}

// SetOndurationchange prop
// js:"ondurationchange"
func (*HTMLVideoElement) SetOndurationchange(ondurationchange func(window.Event)) {
	js.Rewrite("$_.ondurationchange = $1", ondurationchange)
}

// Onemptied prop
// js:"onemptied"
func (*HTMLVideoElement) Onemptied() (onemptied func(window.Event)) {
	js.Rewrite("$_.onemptied")
	return onemptied
}

// SetOnemptied prop
// js:"onemptied"
func (*HTMLVideoElement) SetOnemptied(onemptied func(window.Event)) {
	js.Rewrite("$_.onemptied = $1", onemptied)
}

// Onended prop
// js:"onended"
func (*HTMLVideoElement) Onended() (onended func(window.Event)) {
	js.Rewrite("$_.onended")
	return onended
}

// SetOnended prop
// js:"onended"
func (*HTMLVideoElement) SetOnended(onended func(window.Event)) {
	js.Rewrite("$_.onended = $1", onended)
}

// Onerror prop
// js:"onerror"
func (*HTMLVideoElement) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*HTMLVideoElement) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Onfocus prop
// js:"onfocus"
func (*HTMLVideoElement) Onfocus() (onfocus func(*window.FocusEvent)) {
	js.Rewrite("$_.onfocus")
	return onfocus
}

// SetOnfocus prop
// js:"onfocus"
func (*HTMLVideoElement) SetOnfocus(onfocus func(*window.FocusEvent)) {
	js.Rewrite("$_.onfocus = $1", onfocus)
}

// Oninput prop
// js:"oninput"
func (*HTMLVideoElement) Oninput() (oninput func(window.Event)) {
	js.Rewrite("$_.oninput")
	return oninput
}

// SetOninput prop
// js:"oninput"
func (*HTMLVideoElement) SetOninput(oninput func(window.Event)) {
	js.Rewrite("$_.oninput = $1", oninput)
}

// Oninvalid prop
// js:"oninvalid"
func (*HTMLVideoElement) Oninvalid() (oninvalid func(window.Event)) {
	js.Rewrite("$_.oninvalid")
	return oninvalid
}

// SetOninvalid prop
// js:"oninvalid"
func (*HTMLVideoElement) SetOninvalid(oninvalid func(window.Event)) {
	js.Rewrite("$_.oninvalid = $1", oninvalid)
}

// Onkeydown prop
// js:"onkeydown"
func (*HTMLVideoElement) Onkeydown() (onkeydown func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeydown")
	return onkeydown
}

// SetOnkeydown prop
// js:"onkeydown"
func (*HTMLVideoElement) SetOnkeydown(onkeydown func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeydown = $1", onkeydown)
}

// Onkeypress prop
// js:"onkeypress"
func (*HTMLVideoElement) Onkeypress() (onkeypress func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeypress")
	return onkeypress
}

// SetOnkeypress prop
// js:"onkeypress"
func (*HTMLVideoElement) SetOnkeypress(onkeypress func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeypress = $1", onkeypress)
}

// Onkeyup prop
// js:"onkeyup"
func (*HTMLVideoElement) Onkeyup() (onkeyup func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeyup")
	return onkeyup
}

// SetOnkeyup prop
// js:"onkeyup"
func (*HTMLVideoElement) SetOnkeyup(onkeyup func(*window.KeyboardEvent)) {
	js.Rewrite("$_.onkeyup = $1", onkeyup)
}

// Onload prop
// js:"onload"
func (*HTMLVideoElement) Onload() (onload func(window.Event)) {
	js.Rewrite("$_.onload")
	return onload
}

// SetOnload prop
// js:"onload"
func (*HTMLVideoElement) SetOnload(onload func(window.Event)) {
	js.Rewrite("$_.onload = $1", onload)
}

// Onloadeddata prop
// js:"onloadeddata"
func (*HTMLVideoElement) Onloadeddata() (onloadeddata func(window.Event)) {
	js.Rewrite("$_.onloadeddata")
	return onloadeddata
}

// SetOnloadeddata prop
// js:"onloadeddata"
func (*HTMLVideoElement) SetOnloadeddata(onloadeddata func(window.Event)) {
	js.Rewrite("$_.onloadeddata = $1", onloadeddata)
}

// Onloadedmetadata prop
// js:"onloadedmetadata"
func (*HTMLVideoElement) Onloadedmetadata() (onloadedmetadata func(window.Event)) {
	js.Rewrite("$_.onloadedmetadata")
	return onloadedmetadata
}

// SetOnloadedmetadata prop
// js:"onloadedmetadata"
func (*HTMLVideoElement) SetOnloadedmetadata(onloadedmetadata func(window.Event)) {
	js.Rewrite("$_.onloadedmetadata = $1", onloadedmetadata)
}

// Onloadstart prop
// js:"onloadstart"
func (*HTMLVideoElement) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$_.onloadstart")
	return onloadstart
}

// SetOnloadstart prop
// js:"onloadstart"
func (*HTMLVideoElement) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$_.onloadstart = $1", onloadstart)
}

// Onmousedown prop
// js:"onmousedown"
func (*HTMLVideoElement) Onmousedown() (onmousedown func(window.MouseEvent)) {
	js.Rewrite("$_.onmousedown")
	return onmousedown
}

// SetOnmousedown prop
// js:"onmousedown"
func (*HTMLVideoElement) SetOnmousedown(onmousedown func(window.MouseEvent)) {
	js.Rewrite("$_.onmousedown = $1", onmousedown)
}

// Onmouseenter prop
// js:"onmouseenter"
func (*HTMLVideoElement) Onmouseenter() (onmouseenter func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseenter")
	return onmouseenter
}

// SetOnmouseenter prop
// js:"onmouseenter"
func (*HTMLVideoElement) SetOnmouseenter(onmouseenter func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseenter = $1", onmouseenter)
}

// Onmouseleave prop
// js:"onmouseleave"
func (*HTMLVideoElement) Onmouseleave() (onmouseleave func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseleave")
	return onmouseleave
}

// SetOnmouseleave prop
// js:"onmouseleave"
func (*HTMLVideoElement) SetOnmouseleave(onmouseleave func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseleave = $1", onmouseleave)
}

// Onmousemove prop
// js:"onmousemove"
func (*HTMLVideoElement) Onmousemove() (onmousemove func(window.MouseEvent)) {
	js.Rewrite("$_.onmousemove")
	return onmousemove
}

// SetOnmousemove prop
// js:"onmousemove"
func (*HTMLVideoElement) SetOnmousemove(onmousemove func(window.MouseEvent)) {
	js.Rewrite("$_.onmousemove = $1", onmousemove)
}

// Onmouseout prop
// js:"onmouseout"
func (*HTMLVideoElement) Onmouseout() (onmouseout func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseout")
	return onmouseout
}

// SetOnmouseout prop
// js:"onmouseout"
func (*HTMLVideoElement) SetOnmouseout(onmouseout func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseout = $1", onmouseout)
}

// Onmouseover prop
// js:"onmouseover"
func (*HTMLVideoElement) Onmouseover() (onmouseover func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseover")
	return onmouseover
}

// SetOnmouseover prop
// js:"onmouseover"
func (*HTMLVideoElement) SetOnmouseover(onmouseover func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseover = $1", onmouseover)
}

// Onmouseup prop
// js:"onmouseup"
func (*HTMLVideoElement) Onmouseup() (onmouseup func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseup")
	return onmouseup
}

// SetOnmouseup prop
// js:"onmouseup"
func (*HTMLVideoElement) SetOnmouseup(onmouseup func(window.MouseEvent)) {
	js.Rewrite("$_.onmouseup = $1", onmouseup)
}

// Onmousewheel prop
// js:"onmousewheel"
func (*HTMLVideoElement) Onmousewheel() (onmousewheel func(*window.WheelEvent)) {
	js.Rewrite("$_.onmousewheel")
	return onmousewheel
}

// SetOnmousewheel prop
// js:"onmousewheel"
func (*HTMLVideoElement) SetOnmousewheel(onmousewheel func(*window.WheelEvent)) {
	js.Rewrite("$_.onmousewheel = $1", onmousewheel)
}

// Onmscontentzoom prop
// js:"onmscontentzoom"
func (*HTMLVideoElement) Onmscontentzoom() (onmscontentzoom func(window.UIEvent)) {
	js.Rewrite("$_.onmscontentzoom")
	return onmscontentzoom
}

// SetOnmscontentzoom prop
// js:"onmscontentzoom"
func (*HTMLVideoElement) SetOnmscontentzoom(onmscontentzoom func(window.UIEvent)) {
	js.Rewrite("$_.onmscontentzoom = $1", onmscontentzoom)
}

// Onmsmanipulationstatechanged prop
// js:"onmsmanipulationstatechanged"
func (*HTMLVideoElement) Onmsmanipulationstatechanged() (onmsmanipulationstatechanged func(*window.MSManipulationEvent)) {
	js.Rewrite("$_.onmsmanipulationstatechanged")
	return onmsmanipulationstatechanged
}

// SetOnmsmanipulationstatechanged prop
// js:"onmsmanipulationstatechanged"
func (*HTMLVideoElement) SetOnmsmanipulationstatechanged(onmsmanipulationstatechanged func(*window.MSManipulationEvent)) {
	js.Rewrite("$_.onmsmanipulationstatechanged = $1", onmsmanipulationstatechanged)
}

// Onpaste prop
// js:"onpaste"
func (*HTMLVideoElement) Onpaste() (onpaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onpaste")
	return onpaste
}

// SetOnpaste prop
// js:"onpaste"
func (*HTMLVideoElement) SetOnpaste(onpaste func(*window.ClipboardEvent)) {
	js.Rewrite("$_.onpaste = $1", onpaste)
}

// Onpause prop
// js:"onpause"
func (*HTMLVideoElement) Onpause() (onpause func(window.Event)) {
	js.Rewrite("$_.onpause")
	return onpause
}

// SetOnpause prop
// js:"onpause"
func (*HTMLVideoElement) SetOnpause(onpause func(window.Event)) {
	js.Rewrite("$_.onpause = $1", onpause)
}

// Onplay prop
// js:"onplay"
func (*HTMLVideoElement) Onplay() (onplay func(window.Event)) {
	js.Rewrite("$_.onplay")
	return onplay
}

// SetOnplay prop
// js:"onplay"
func (*HTMLVideoElement) SetOnplay(onplay func(window.Event)) {
	js.Rewrite("$_.onplay = $1", onplay)
}

// Onplaying prop
// js:"onplaying"
func (*HTMLVideoElement) Onplaying() (onplaying func(window.Event)) {
	js.Rewrite("$_.onplaying")
	return onplaying
}

// SetOnplaying prop
// js:"onplaying"
func (*HTMLVideoElement) SetOnplaying(onplaying func(window.Event)) {
	js.Rewrite("$_.onplaying = $1", onplaying)
}

// Onprogress prop
// js:"onprogress"
func (*HTMLVideoElement) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop
// js:"onprogress"
func (*HTMLVideoElement) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$_.onprogress = $1", onprogress)
}

// Onratechange prop
// js:"onratechange"
func (*HTMLVideoElement) Onratechange() (onratechange func(window.Event)) {
	js.Rewrite("$_.onratechange")
	return onratechange
}

// SetOnratechange prop
// js:"onratechange"
func (*HTMLVideoElement) SetOnratechange(onratechange func(window.Event)) {
	js.Rewrite("$_.onratechange = $1", onratechange)
}

// Onreset prop
// js:"onreset"
func (*HTMLVideoElement) Onreset() (onreset func(window.Event)) {
	js.Rewrite("$_.onreset")
	return onreset
}

// SetOnreset prop
// js:"onreset"
func (*HTMLVideoElement) SetOnreset(onreset func(window.Event)) {
	js.Rewrite("$_.onreset = $1", onreset)
}

// Onscroll prop
// js:"onscroll"
func (*HTMLVideoElement) Onscroll() (onscroll func(window.UIEvent)) {
	js.Rewrite("$_.onscroll")
	return onscroll
}

// SetOnscroll prop
// js:"onscroll"
func (*HTMLVideoElement) SetOnscroll(onscroll func(window.UIEvent)) {
	js.Rewrite("$_.onscroll = $1", onscroll)
}

// Onseeked prop
// js:"onseeked"
func (*HTMLVideoElement) Onseeked() (onseeked func(window.Event)) {
	js.Rewrite("$_.onseeked")
	return onseeked
}

// SetOnseeked prop
// js:"onseeked"
func (*HTMLVideoElement) SetOnseeked(onseeked func(window.Event)) {
	js.Rewrite("$_.onseeked = $1", onseeked)
}

// Onseeking prop
// js:"onseeking"
func (*HTMLVideoElement) Onseeking() (onseeking func(window.Event)) {
	js.Rewrite("$_.onseeking")
	return onseeking
}

// SetOnseeking prop
// js:"onseeking"
func (*HTMLVideoElement) SetOnseeking(onseeking func(window.Event)) {
	js.Rewrite("$_.onseeking = $1", onseeking)
}

// Onselect prop
// js:"onselect"
func (*HTMLVideoElement) Onselect() (onselect func(window.UIEvent)) {
	js.Rewrite("$_.onselect")
	return onselect
}

// SetOnselect prop
// js:"onselect"
func (*HTMLVideoElement) SetOnselect(onselect func(window.UIEvent)) {
	js.Rewrite("$_.onselect = $1", onselect)
}

// Onselectstart prop
// js:"onselectstart"
func (*HTMLVideoElement) Onselectstart() (onselectstart func(window.Event)) {
	js.Rewrite("$_.onselectstart")
	return onselectstart
}

// SetOnselectstart prop
// js:"onselectstart"
func (*HTMLVideoElement) SetOnselectstart(onselectstart func(window.Event)) {
	js.Rewrite("$_.onselectstart = $1", onselectstart)
}

// Onstalled prop
// js:"onstalled"
func (*HTMLVideoElement) Onstalled() (onstalled func(window.Event)) {
	js.Rewrite("$_.onstalled")
	return onstalled
}

// SetOnstalled prop
// js:"onstalled"
func (*HTMLVideoElement) SetOnstalled(onstalled func(window.Event)) {
	js.Rewrite("$_.onstalled = $1", onstalled)
}

// Onsubmit prop
// js:"onsubmit"
func (*HTMLVideoElement) Onsubmit() (onsubmit func(window.Event)) {
	js.Rewrite("$_.onsubmit")
	return onsubmit
}

// SetOnsubmit prop
// js:"onsubmit"
func (*HTMLVideoElement) SetOnsubmit(onsubmit func(window.Event)) {
	js.Rewrite("$_.onsubmit = $1", onsubmit)
}

// Onsuspend prop
// js:"onsuspend"
func (*HTMLVideoElement) Onsuspend() (onsuspend func(window.Event)) {
	js.Rewrite("$_.onsuspend")
	return onsuspend
}

// SetOnsuspend prop
// js:"onsuspend"
func (*HTMLVideoElement) SetOnsuspend(onsuspend func(window.Event)) {
	js.Rewrite("$_.onsuspend = $1", onsuspend)
}

// Ontimeupdate prop
// js:"ontimeupdate"
func (*HTMLVideoElement) Ontimeupdate() (ontimeupdate func(window.Event)) {
	js.Rewrite("$_.ontimeupdate")
	return ontimeupdate
}

// SetOntimeupdate prop
// js:"ontimeupdate"
func (*HTMLVideoElement) SetOntimeupdate(ontimeupdate func(window.Event)) {
	js.Rewrite("$_.ontimeupdate = $1", ontimeupdate)
}

// Onvolumechange prop
// js:"onvolumechange"
func (*HTMLVideoElement) Onvolumechange() (onvolumechange func(window.Event)) {
	js.Rewrite("$_.onvolumechange")
	return onvolumechange
}

// SetOnvolumechange prop
// js:"onvolumechange"
func (*HTMLVideoElement) SetOnvolumechange(onvolumechange func(window.Event)) {
	js.Rewrite("$_.onvolumechange = $1", onvolumechange)
}

// Onwaiting prop
// js:"onwaiting"
func (*HTMLVideoElement) Onwaiting() (onwaiting func(window.Event)) {
	js.Rewrite("$_.onwaiting")
	return onwaiting
}

// SetOnwaiting prop
// js:"onwaiting"
func (*HTMLVideoElement) SetOnwaiting(onwaiting func(window.Event)) {
	js.Rewrite("$_.onwaiting = $1", onwaiting)
}

// OuterText prop
// js:"outerText"
func (*HTMLVideoElement) OuterText() (outerText string) {
	js.Rewrite("$_.outerText")
	return outerText
}

// SetOuterText prop
// js:"outerText"
func (*HTMLVideoElement) SetOuterText(outerText string) {
	js.Rewrite("$_.outerText = $1", outerText)
}

// Spellcheck prop
// js:"spellcheck"
func (*HTMLVideoElement) Spellcheck() (spellcheck bool) {
	js.Rewrite("$_.spellcheck")
	return spellcheck
}

// SetSpellcheck prop
// js:"spellcheck"
func (*HTMLVideoElement) SetSpellcheck(spellcheck bool) {
	js.Rewrite("$_.spellcheck = $1", spellcheck)
}

// Style prop
// js:"style"
func (*HTMLVideoElement) Style() (style *window.CSSStyleDeclaration) {
	js.Rewrite("$_.style")
	return style
}

// TabIndex prop
// js:"tabIndex"
func (*HTMLVideoElement) TabIndex() (tabIndex int8) {
	js.Rewrite("$_.tabIndex")
	return tabIndex
}

// SetTabIndex prop
// js:"tabIndex"
func (*HTMLVideoElement) SetTabIndex(tabIndex int8) {
	js.Rewrite("$_.tabIndex = $1", tabIndex)
}

// Title prop
// js:"title"
func (*HTMLVideoElement) Title() (title string) {
	js.Rewrite("$_.title")
	return title
}

// SetTitle prop
// js:"title"
func (*HTMLVideoElement) SetTitle(title string) {
	js.Rewrite("$_.title = $1", title)
}

// ClassList prop
// js:"classList"
func (*HTMLVideoElement) ClassList() (classList domtokenlist.DOMTokenList) {
	js.Rewrite("$_.classList")
	return classList
}

// ClassName prop
// js:"className"
func (*HTMLVideoElement) ClassName() (className string) {
	js.Rewrite("$_.className")
	return className
}

// SetClassName prop
// js:"className"
func (*HTMLVideoElement) SetClassName(className string) {
	js.Rewrite("$_.className = $1", className)
}

// ClientHeight prop
// js:"clientHeight"
func (*HTMLVideoElement) ClientHeight() (clientHeight int) {
	js.Rewrite("$_.clientHeight")
	return clientHeight
}

// ClientLeft prop
// js:"clientLeft"
func (*HTMLVideoElement) ClientLeft() (clientLeft int) {
	js.Rewrite("$_.clientLeft")
	return clientLeft
}

// ClientTop prop
// js:"clientTop"
func (*HTMLVideoElement) ClientTop() (clientTop int) {
	js.Rewrite("$_.clientTop")
	return clientTop
}

// ClientWidth prop
// js:"clientWidth"
func (*HTMLVideoElement) ClientWidth() (clientWidth int) {
	js.Rewrite("$_.clientWidth")
	return clientWidth
}

// ID prop
// js:"id"
func (*HTMLVideoElement) ID() (id string) {
	js.Rewrite("$_.id")
	return id
}

// SetID prop
// js:"id"
func (*HTMLVideoElement) SetID(id string) {
	js.Rewrite("$_.id = $1", id)
}

// InnerHTML prop
// js:"innerHTML"
func (*HTMLVideoElement) InnerHTML() (innerHTML string) {
	js.Rewrite("$_.innerHTML")
	return innerHTML
}

// SetInnerHTML prop
// js:"innerHTML"
func (*HTMLVideoElement) SetInnerHTML(innerHTML string) {
	js.Rewrite("$_.innerHTML = $1", innerHTML)
}

// MsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*HTMLVideoElement) MsContentZoomFactor() (msContentZoomFactor float32) {
	js.Rewrite("$_.msContentZoomFactor")
	return msContentZoomFactor
}

// SetMsContentZoomFactor prop
// js:"msContentZoomFactor"
func (*HTMLVideoElement) SetMsContentZoomFactor(msContentZoomFactor float32) {
	js.Rewrite("$_.msContentZoomFactor = $1", msContentZoomFactor)
}

// MsRegionOverflow prop
// js:"msRegionOverflow"
func (*HTMLVideoElement) MsRegionOverflow() (msRegionOverflow string) {
	js.Rewrite("$_.msRegionOverflow")
	return msRegionOverflow
}

// Onariarequest prop
// js:"onariarequest"
func (*HTMLVideoElement) Onariarequest() (onariarequest func(window.Event)) {
	js.Rewrite("$_.onariarequest")
	return onariarequest
}

// SetOnariarequest prop
// js:"onariarequest"
func (*HTMLVideoElement) SetOnariarequest(onariarequest func(window.Event)) {
	js.Rewrite("$_.onariarequest = $1", onariarequest)
}

// Oncommand prop
// js:"oncommand"
func (*HTMLVideoElement) Oncommand() (oncommand func(window.Event)) {
	js.Rewrite("$_.oncommand")
	return oncommand
}

// SetOncommand prop
// js:"oncommand"
func (*HTMLVideoElement) SetOncommand(oncommand func(window.Event)) {
	js.Rewrite("$_.oncommand = $1", oncommand)
}

// Ongotpointercapture prop
// js:"ongotpointercapture"
func (*HTMLVideoElement) Ongotpointercapture() (ongotpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.ongotpointercapture")
	return ongotpointercapture
}

// SetOngotpointercapture prop
// js:"ongotpointercapture"
func (*HTMLVideoElement) SetOngotpointercapture(ongotpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.ongotpointercapture = $1", ongotpointercapture)
}

// Onlostpointercapture prop
// js:"onlostpointercapture"
func (*HTMLVideoElement) Onlostpointercapture() (onlostpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.onlostpointercapture")
	return onlostpointercapture
}

// SetOnlostpointercapture prop
// js:"onlostpointercapture"
func (*HTMLVideoElement) SetOnlostpointercapture(onlostpointercapture func(*window.PointerEvent)) {
	js.Rewrite("$_.onlostpointercapture = $1", onlostpointercapture)
}

// Onmsgesturechange prop
// js:"onmsgesturechange"
func (*HTMLVideoElement) Onmsgesturechange() (onmsgesturechange func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturechange")
	return onmsgesturechange
}

// SetOnmsgesturechange prop
// js:"onmsgesturechange"
func (*HTMLVideoElement) SetOnmsgesturechange(onmsgesturechange func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturechange = $1", onmsgesturechange)
}

// Onmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*HTMLVideoElement) Onmsgesturedoubletap() (onmsgesturedoubletap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturedoubletap")
	return onmsgesturedoubletap
}

// SetOnmsgesturedoubletap prop
// js:"onmsgesturedoubletap"
func (*HTMLVideoElement) SetOnmsgesturedoubletap(onmsgesturedoubletap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturedoubletap = $1", onmsgesturedoubletap)
}

// Onmsgestureend prop
// js:"onmsgestureend"
func (*HTMLVideoElement) Onmsgestureend() (onmsgestureend func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgestureend")
	return onmsgestureend
}

// SetOnmsgestureend prop
// js:"onmsgestureend"
func (*HTMLVideoElement) SetOnmsgestureend(onmsgestureend func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgestureend = $1", onmsgestureend)
}

// Onmsgesturehold prop
// js:"onmsgesturehold"
func (*HTMLVideoElement) Onmsgesturehold() (onmsgesturehold func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturehold")
	return onmsgesturehold
}

// SetOnmsgesturehold prop
// js:"onmsgesturehold"
func (*HTMLVideoElement) SetOnmsgesturehold(onmsgesturehold func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturehold = $1", onmsgesturehold)
}

// Onmsgesturestart prop
// js:"onmsgesturestart"
func (*HTMLVideoElement) Onmsgesturestart() (onmsgesturestart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturestart")
	return onmsgesturestart
}

// SetOnmsgesturestart prop
// js:"onmsgesturestart"
func (*HTMLVideoElement) SetOnmsgesturestart(onmsgesturestart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturestart = $1", onmsgesturestart)
}

// Onmsgesturetap prop
// js:"onmsgesturetap"
func (*HTMLVideoElement) Onmsgesturetap() (onmsgesturetap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturetap")
	return onmsgesturetap
}

// SetOnmsgesturetap prop
// js:"onmsgesturetap"
func (*HTMLVideoElement) SetOnmsgesturetap(onmsgesturetap func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsgesturetap = $1", onmsgesturetap)
}

// Onmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*HTMLVideoElement) Onmsgotpointercapture() (onmsgotpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmsgotpointercapture")
	return onmsgotpointercapture
}

// SetOnmsgotpointercapture prop
// js:"onmsgotpointercapture"
func (*HTMLVideoElement) SetOnmsgotpointercapture(onmsgotpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmsgotpointercapture = $1", onmsgotpointercapture)
}

// Onmsinertiastart prop
// js:"onmsinertiastart"
func (*HTMLVideoElement) Onmsinertiastart() (onmsinertiastart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsinertiastart")
	return onmsinertiastart
}

// SetOnmsinertiastart prop
// js:"onmsinertiastart"
func (*HTMLVideoElement) SetOnmsinertiastart(onmsinertiastart func(*window.MSGestureEvent)) {
	js.Rewrite("$_.onmsinertiastart = $1", onmsinertiastart)
}

// Onmslostpointercapture prop
// js:"onmslostpointercapture"
func (*HTMLVideoElement) Onmslostpointercapture() (onmslostpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmslostpointercapture")
	return onmslostpointercapture
}

// SetOnmslostpointercapture prop
// js:"onmslostpointercapture"
func (*HTMLVideoElement) SetOnmslostpointercapture(onmslostpointercapture func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmslostpointercapture = $1", onmslostpointercapture)
}

// Onmspointercancel prop
// js:"onmspointercancel"
func (*HTMLVideoElement) Onmspointercancel() (onmspointercancel func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointercancel")
	return onmspointercancel
}

// SetOnmspointercancel prop
// js:"onmspointercancel"
func (*HTMLVideoElement) SetOnmspointercancel(onmspointercancel func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointercancel = $1", onmspointercancel)
}

// Onmspointerdown prop
// js:"onmspointerdown"
func (*HTMLVideoElement) Onmspointerdown() (onmspointerdown func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerdown")
	return onmspointerdown
}

// SetOnmspointerdown prop
// js:"onmspointerdown"
func (*HTMLVideoElement) SetOnmspointerdown(onmspointerdown func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerdown = $1", onmspointerdown)
}

// Onmspointerenter prop
// js:"onmspointerenter"
func (*HTMLVideoElement) Onmspointerenter() (onmspointerenter func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerenter")
	return onmspointerenter
}

// SetOnmspointerenter prop
// js:"onmspointerenter"
func (*HTMLVideoElement) SetOnmspointerenter(onmspointerenter func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerenter = $1", onmspointerenter)
}

// Onmspointerleave prop
// js:"onmspointerleave"
func (*HTMLVideoElement) Onmspointerleave() (onmspointerleave func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerleave")
	return onmspointerleave
}

// SetOnmspointerleave prop
// js:"onmspointerleave"
func (*HTMLVideoElement) SetOnmspointerleave(onmspointerleave func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerleave = $1", onmspointerleave)
}

// Onmspointermove prop
// js:"onmspointermove"
func (*HTMLVideoElement) Onmspointermove() (onmspointermove func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointermove")
	return onmspointermove
}

// SetOnmspointermove prop
// js:"onmspointermove"
func (*HTMLVideoElement) SetOnmspointermove(onmspointermove func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointermove = $1", onmspointermove)
}

// Onmspointerout prop
// js:"onmspointerout"
func (*HTMLVideoElement) Onmspointerout() (onmspointerout func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerout")
	return onmspointerout
}

// SetOnmspointerout prop
// js:"onmspointerout"
func (*HTMLVideoElement) SetOnmspointerout(onmspointerout func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerout = $1", onmspointerout)
}

// Onmspointerover prop
// js:"onmspointerover"
func (*HTMLVideoElement) Onmspointerover() (onmspointerover func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerover")
	return onmspointerover
}

// SetOnmspointerover prop
// js:"onmspointerover"
func (*HTMLVideoElement) SetOnmspointerover(onmspointerover func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerover = $1", onmspointerover)
}

// Onmspointerup prop
// js:"onmspointerup"
func (*HTMLVideoElement) Onmspointerup() (onmspointerup func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerup")
	return onmspointerup
}

// SetOnmspointerup prop
// js:"onmspointerup"
func (*HTMLVideoElement) SetOnmspointerup(onmspointerup func(*window.MSPointerEvent)) {
	js.Rewrite("$_.onmspointerup = $1", onmspointerup)
}

// Ontouchcancel prop
// js:"ontouchcancel"
func (*HTMLVideoElement) Ontouchcancel() (ontouchcancel func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchcancel")
	return ontouchcancel
}

// SetOntouchcancel prop
// js:"ontouchcancel"
func (*HTMLVideoElement) SetOntouchcancel(ontouchcancel func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchcancel = $1", ontouchcancel)
}

// Ontouchend prop
// js:"ontouchend"
func (*HTMLVideoElement) Ontouchend() (ontouchend func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchend")
	return ontouchend
}

// SetOntouchend prop
// js:"ontouchend"
func (*HTMLVideoElement) SetOntouchend(ontouchend func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchend = $1", ontouchend)
}

// Ontouchmove prop
// js:"ontouchmove"
func (*HTMLVideoElement) Ontouchmove() (ontouchmove func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchmove")
	return ontouchmove
}

// SetOntouchmove prop
// js:"ontouchmove"
func (*HTMLVideoElement) SetOntouchmove(ontouchmove func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchmove = $1", ontouchmove)
}

// Ontouchstart prop
// js:"ontouchstart"
func (*HTMLVideoElement) Ontouchstart() (ontouchstart func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchstart")
	return ontouchstart
}

// SetOntouchstart prop
// js:"ontouchstart"
func (*HTMLVideoElement) SetOntouchstart(ontouchstart func(*window.TouchEvent)) {
	js.Rewrite("$_.ontouchstart = $1", ontouchstart)
}

// Onwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*HTMLVideoElement) Onwebkitfullscreenchange() (onwebkitfullscreenchange func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenchange")
	return onwebkitfullscreenchange
}

// SetOnwebkitfullscreenchange prop
// js:"onwebkitfullscreenchange"
func (*HTMLVideoElement) SetOnwebkitfullscreenchange(onwebkitfullscreenchange func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenchange = $1", onwebkitfullscreenchange)
}

// Onwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*HTMLVideoElement) Onwebkitfullscreenerror() (onwebkitfullscreenerror func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenerror")
	return onwebkitfullscreenerror
}

// SetOnwebkitfullscreenerror prop
// js:"onwebkitfullscreenerror"
func (*HTMLVideoElement) SetOnwebkitfullscreenerror(onwebkitfullscreenerror func(window.Event)) {
	js.Rewrite("$_.onwebkitfullscreenerror = $1", onwebkitfullscreenerror)
}

// OuterHTML prop
// js:"outerHTML"
func (*HTMLVideoElement) OuterHTML() (outerHTML string) {
	js.Rewrite("$_.outerHTML")
	return outerHTML
}

// SetOuterHTML prop
// js:"outerHTML"
func (*HTMLVideoElement) SetOuterHTML(outerHTML string) {
	js.Rewrite("$_.outerHTML = $1", outerHTML)
}

// Prefix prop
// js:"prefix"
func (*HTMLVideoElement) Prefix() (prefix string) {
	js.Rewrite("$_.prefix")
	return prefix
}

// ScrollHeight prop
// js:"scrollHeight"
func (*HTMLVideoElement) ScrollHeight() (scrollHeight int) {
	js.Rewrite("$_.scrollHeight")
	return scrollHeight
}

// ScrollLeft prop
// js:"scrollLeft"
func (*HTMLVideoElement) ScrollLeft() (scrollLeft int) {
	js.Rewrite("$_.scrollLeft")
	return scrollLeft
}

// SetScrollLeft prop
// js:"scrollLeft"
func (*HTMLVideoElement) SetScrollLeft(scrollLeft int) {
	js.Rewrite("$_.scrollLeft = $1", scrollLeft)
}

// ScrollTop prop
// js:"scrollTop"
func (*HTMLVideoElement) ScrollTop() (scrollTop int) {
	js.Rewrite("$_.scrollTop")
	return scrollTop
}

// SetScrollTop prop
// js:"scrollTop"
func (*HTMLVideoElement) SetScrollTop(scrollTop int) {
	js.Rewrite("$_.scrollTop = $1", scrollTop)
}

// ScrollWidth prop
// js:"scrollWidth"
func (*HTMLVideoElement) ScrollWidth() (scrollWidth int) {
	js.Rewrite("$_.scrollWidth")
	return scrollWidth
}

// TagName prop
// js:"tagName"
func (*HTMLVideoElement) TagName() (tagName string) {
	js.Rewrite("$_.tagName")
	return tagName
}

// Onpointercancel prop
// js:"onpointercancel"
func (*HTMLVideoElement) Onpointercancel() (onpointercancel func(window.Event)) {
	js.Rewrite("$_.onpointercancel")
	return onpointercancel
}

// SetOnpointercancel prop
// js:"onpointercancel"
func (*HTMLVideoElement) SetOnpointercancel(onpointercancel func(window.Event)) {
	js.Rewrite("$_.onpointercancel = $1", onpointercancel)
}

// Onpointerdown prop
// js:"onpointerdown"
func (*HTMLVideoElement) Onpointerdown() (onpointerdown func(window.Event)) {
	js.Rewrite("$_.onpointerdown")
	return onpointerdown
}

// SetOnpointerdown prop
// js:"onpointerdown"
func (*HTMLVideoElement) SetOnpointerdown(onpointerdown func(window.Event)) {
	js.Rewrite("$_.onpointerdown = $1", onpointerdown)
}

// Onpointerenter prop
// js:"onpointerenter"
func (*HTMLVideoElement) Onpointerenter() (onpointerenter func(window.Event)) {
	js.Rewrite("$_.onpointerenter")
	return onpointerenter
}

// SetOnpointerenter prop
// js:"onpointerenter"
func (*HTMLVideoElement) SetOnpointerenter(onpointerenter func(window.Event)) {
	js.Rewrite("$_.onpointerenter = $1", onpointerenter)
}

// Onpointerleave prop
// js:"onpointerleave"
func (*HTMLVideoElement) Onpointerleave() (onpointerleave func(window.Event)) {
	js.Rewrite("$_.onpointerleave")
	return onpointerleave
}

// SetOnpointerleave prop
// js:"onpointerleave"
func (*HTMLVideoElement) SetOnpointerleave(onpointerleave func(window.Event)) {
	js.Rewrite("$_.onpointerleave = $1", onpointerleave)
}

// Onpointermove prop
// js:"onpointermove"
func (*HTMLVideoElement) Onpointermove() (onpointermove func(window.Event)) {
	js.Rewrite("$_.onpointermove")
	return onpointermove
}

// SetOnpointermove prop
// js:"onpointermove"
func (*HTMLVideoElement) SetOnpointermove(onpointermove func(window.Event)) {
	js.Rewrite("$_.onpointermove = $1", onpointermove)
}

// Onpointerout prop
// js:"onpointerout"
func (*HTMLVideoElement) Onpointerout() (onpointerout func(window.Event)) {
	js.Rewrite("$_.onpointerout")
	return onpointerout
}

// SetOnpointerout prop
// js:"onpointerout"
func (*HTMLVideoElement) SetOnpointerout(onpointerout func(window.Event)) {
	js.Rewrite("$_.onpointerout = $1", onpointerout)
}

// Onpointerover prop
// js:"onpointerover"
func (*HTMLVideoElement) Onpointerover() (onpointerover func(window.Event)) {
	js.Rewrite("$_.onpointerover")
	return onpointerover
}

// SetOnpointerover prop
// js:"onpointerover"
func (*HTMLVideoElement) SetOnpointerover(onpointerover func(window.Event)) {
	js.Rewrite("$_.onpointerover = $1", onpointerover)
}

// Onpointerup prop
// js:"onpointerup"
func (*HTMLVideoElement) Onpointerup() (onpointerup func(window.Event)) {
	js.Rewrite("$_.onpointerup")
	return onpointerup
}

// SetOnpointerup prop
// js:"onpointerup"
func (*HTMLVideoElement) SetOnpointerup(onpointerup func(window.Event)) {
	js.Rewrite("$_.onpointerup = $1", onpointerup)
}

// Onwheel prop
// js:"onwheel"
func (*HTMLVideoElement) Onwheel() (onwheel func(window.Event)) {
	js.Rewrite("$_.onwheel")
	return onwheel
}

// SetOnwheel prop
// js:"onwheel"
func (*HTMLVideoElement) SetOnwheel(onwheel func(window.Event)) {
	js.Rewrite("$_.onwheel = $1", onwheel)
}

// ChildElementCount prop
// js:"childElementCount"
func (*HTMLVideoElement) ChildElementCount() (childElementCount uint) {
	js.Rewrite("$_.childElementCount")
	return childElementCount
}

// FirstElementChild prop
// js:"firstElementChild"
func (*HTMLVideoElement) FirstElementChild() (firstElementChild window.Element) {
	js.Rewrite("$_.firstElementChild")
	return firstElementChild
}

// LastElementChild prop
// js:"lastElementChild"
func (*HTMLVideoElement) LastElementChild() (lastElementChild window.Element) {
	js.Rewrite("$_.lastElementChild")
	return lastElementChild
}

// NextElementSibling prop
// js:"nextElementSibling"
func (*HTMLVideoElement) NextElementSibling() (nextElementSibling window.Element) {
	js.Rewrite("$_.nextElementSibling")
	return nextElementSibling
}

// PreviousElementSibling prop
// js:"previousElementSibling"
func (*HTMLVideoElement) PreviousElementSibling() (previousElementSibling window.Element) {
	js.Rewrite("$_.previousElementSibling")
	return previousElementSibling
}

// Attributes prop
// js:"attributes"
func (*HTMLVideoElement) Attributes() (attributes *window.NamedNodeMap) {
	js.Rewrite("$_.attributes")
	return attributes
}

// BaseURI prop
// js:"baseURI"
func (*HTMLVideoElement) BaseURI() (baseURI string) {
	js.Rewrite("$_.baseURI")
	return baseURI
}

// ChildNodes prop
// js:"childNodes"
func (*HTMLVideoElement) ChildNodes() (childNodes *window.NodeList) {
	js.Rewrite("$_.childNodes")
	return childNodes
}

// FirstChild prop
// js:"firstChild"
func (*HTMLVideoElement) FirstChild() (firstChild window.Node) {
	js.Rewrite("$_.firstChild")
	return firstChild
}

// LastChild prop
// js:"lastChild"
func (*HTMLVideoElement) LastChild() (lastChild window.Node) {
	js.Rewrite("$_.lastChild")
	return lastChild
}

// LocalName prop
// js:"localName"
func (*HTMLVideoElement) LocalName() (localName string) {
	js.Rewrite("$_.localName")
	return localName
}

// NamespaceURI prop
// js:"namespaceURI"
func (*HTMLVideoElement) NamespaceURI() (namespaceURI string) {
	js.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// NextSibling prop
// js:"nextSibling"
func (*HTMLVideoElement) NextSibling() (nextSibling window.Node) {
	js.Rewrite("$_.nextSibling")
	return nextSibling
}

// NodeName prop
// js:"nodeName"
func (*HTMLVideoElement) NodeName() (nodeName string) {
	js.Rewrite("$_.nodeName")
	return nodeName
}

// NodeType prop
// js:"nodeType"
func (*HTMLVideoElement) NodeType() (nodeType uint8) {
	js.Rewrite("$_.nodeType")
	return nodeType
}

// NodeValue prop
// js:"nodeValue"
func (*HTMLVideoElement) NodeValue() (nodeValue string) {
	js.Rewrite("$_.nodeValue")
	return nodeValue
}

// SetNodeValue prop
// js:"nodeValue"
func (*HTMLVideoElement) SetNodeValue(nodeValue string) {
	js.Rewrite("$_.nodeValue = $1", nodeValue)
}

// OwnerDocument prop
// js:"ownerDocument"
func (*HTMLVideoElement) OwnerDocument() (ownerDocument window.Document) {
	js.Rewrite("$_.ownerDocument")
	return ownerDocument
}

// ParentElement prop
// js:"parentElement"
func (*HTMLVideoElement) ParentElement() (parentElement window.HTMLElement) {
	js.Rewrite("$_.parentElement")
	return parentElement
}

// ParentNode prop
// js:"parentNode"
func (*HTMLVideoElement) ParentNode() (parentNode window.Node) {
	js.Rewrite("$_.parentNode")
	return parentNode
}

// PreviousSibling prop
// js:"previousSibling"
func (*HTMLVideoElement) PreviousSibling() (previousSibling window.Node) {
	js.Rewrite("$_.previousSibling")
	return previousSibling
}

// TextContent prop
// js:"textContent"
func (*HTMLVideoElement) TextContent() (textContent string) {
	js.Rewrite("$_.textContent")
	return textContent
}

// SetTextContent prop
// js:"textContent"
func (*HTMLVideoElement) SetTextContent(textContent string) {
	js.Rewrite("$_.textContent = $1", textContent)
}
