package htmlvideoelement

import (
	"github.com/matthewmueller/golly/dom2/htmlmediaelement"
	"github.com/matthewmueller/golly/dom2/videoplaybackquality"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLVideoElement struct
// js:"HTMLVideoElement,omit"
type HTMLVideoElement struct {
	htmlmediaelement.HTMLMediaElement
}

// GetVideoPlaybackQuality fn
func (*HTMLVideoElement) GetVideoPlaybackQuality() (v *videoplaybackquality.VideoPlaybackQuality) {
	js.Rewrite("$<.getVideoPlaybackQuality()")
	return v
}

// MsFrameStep fn
func (*HTMLVideoElement) MsFrameStep(forward bool) {
	js.Rewrite("$<.msFrameStep($1)", forward)
}

// MsInsertVideoEffect fn
func (*HTMLVideoElement) MsInsertVideoEffect(activatableClassId string, effectRequired bool, config *interface{}) {
	js.Rewrite("$<.msInsertVideoEffect($1, $2, $3)", activatableClassId, effectRequired, config)
}

// MsSetVideoRectangle fn
func (*HTMLVideoElement) MsSetVideoRectangle(left float32, top float32, right float32, bottom float32) {
	js.Rewrite("$<.msSetVideoRectangle($1, $2, $3, $4)", left, top, right, bottom)
}

// WebkitEnterFullscreen fn
func (*HTMLVideoElement) WebkitEnterFullscreen() {
	js.Rewrite("$<.webkitEnterFullscreen()")
}

// WebkitEnterFullScreen fn
func (*HTMLVideoElement) WebkitEnterFullScreen() {
	js.Rewrite("$<.webkitEnterFullScreen()")
}

// WebkitExitFullscreen fn
func (*HTMLVideoElement) WebkitExitFullscreen() {
	js.Rewrite("$<.webkitExitFullscreen()")
}

// WebkitExitFullScreen fn
func (*HTMLVideoElement) WebkitExitFullScreen() {
	js.Rewrite("$<.webkitExitFullScreen()")
}

// Height prop Gets or sets the height of the video element.
func (*HTMLVideoElement) Height() (height uint) {
	js.Rewrite("$<.height")
	return height
}

// Height prop Gets or sets the height of the video element.
func (*HTMLVideoElement) SetHeight(height uint) {
	js.Rewrite("$<.height = $1", height)
}

// MsHorizontalMirror prop
func (*HTMLVideoElement) MsHorizontalMirror() (msHorizontalMirror bool) {
	js.Rewrite("$<.msHorizontalMirror")
	return msHorizontalMirror
}

// MsHorizontalMirror prop
func (*HTMLVideoElement) SetMsHorizontalMirror(msHorizontalMirror bool) {
	js.Rewrite("$<.msHorizontalMirror = $1", msHorizontalMirror)
}

// MsIsLayoutOptimalForPlayback prop
func (*HTMLVideoElement) MsIsLayoutOptimalForPlayback() (msIsLayoutOptimalForPlayback bool) {
	js.Rewrite("$<.msIsLayoutOptimalForPlayback")
	return msIsLayoutOptimalForPlayback
}

// MsIsStereo3d prop
func (*HTMLVideoElement) MsIsStereo3d() (msIsStereo3D bool) {
	js.Rewrite("$<.msIsStereo3D")
	return msIsStereo3D
}

// MsStereo3dPackingMode prop
func (*HTMLVideoElement) MsStereo3dPackingMode() (msStereo3DPackingMode string) {
	js.Rewrite("$<.msStereo3DPackingMode")
	return msStereo3DPackingMode
}

// MsStereo3dPackingMode prop
func (*HTMLVideoElement) SetMsStereo3dPackingMode(msStereo3DPackingMode string) {
	js.Rewrite("$<.msStereo3DPackingMode = $1", msStereo3DPackingMode)
}

// MsStereo3dRenderMode prop
func (*HTMLVideoElement) MsStereo3dRenderMode() (msStereo3DRenderMode string) {
	js.Rewrite("$<.msStereo3DRenderMode")
	return msStereo3DRenderMode
}

// MsStereo3dRenderMode prop
func (*HTMLVideoElement) SetMsStereo3dRenderMode(msStereo3DRenderMode string) {
	js.Rewrite("$<.msStereo3DRenderMode = $1", msStereo3DRenderMode)
}

// MsZoom prop
func (*HTMLVideoElement) MsZoom() (msZoom bool) {
	js.Rewrite("$<.msZoom")
	return msZoom
}

// MsZoom prop
func (*HTMLVideoElement) SetMsZoom(msZoom bool) {
	js.Rewrite("$<.msZoom = $1", msZoom)
}

// OnMSVideoFormatChanged prop
func (*HTMLVideoElement) OnMSVideoFormatChanged() (onmsvideoformatchanged func(window.Event)) {
	js.Rewrite("$<.onMSVideoFormatChanged")
	return onmsvideoformatchanged
}

// OnMSVideoFormatChanged prop
func (*HTMLVideoElement) SetOnMSVideoFormatChanged(onmsvideoformatchanged func(window.Event)) {
	js.Rewrite("$<.onMSVideoFormatChanged = $1", onmsvideoformatchanged)
}

// OnMSVideoFrameStepCompleted prop
func (*HTMLVideoElement) OnMSVideoFrameStepCompleted() (onmsvideoframestepcompleted func(window.Event)) {
	js.Rewrite("$<.onMSVideoFrameStepCompleted")
	return onmsvideoframestepcompleted
}

// OnMSVideoFrameStepCompleted prop
func (*HTMLVideoElement) SetOnMSVideoFrameStepCompleted(onmsvideoframestepcompleted func(window.Event)) {
	js.Rewrite("$<.onMSVideoFrameStepCompleted = $1", onmsvideoframestepcompleted)
}

// OnMSVideoOptimalLayoutChanged prop
func (*HTMLVideoElement) OnMSVideoOptimalLayoutChanged() (onmsvideooptimallayoutchanged func(window.Event)) {
	js.Rewrite("$<.onMSVideoOptimalLayoutChanged")
	return onmsvideooptimallayoutchanged
}

// OnMSVideoOptimalLayoutChanged prop
func (*HTMLVideoElement) SetOnMSVideoOptimalLayoutChanged(onmsvideooptimallayoutchanged func(window.Event)) {
	js.Rewrite("$<.onMSVideoOptimalLayoutChanged = $1", onmsvideooptimallayoutchanged)
}

// Poster prop Gets or sets a URL of an image to display, for example, like a movie poster. This can be a still frame from the video, or another image if no video data is available.
func (*HTMLVideoElement) Poster() (poster string) {
	js.Rewrite("$<.poster")
	return poster
}

// Poster prop Gets or sets a URL of an image to display, for example, like a movie poster. This can be a still frame from the video, or another image if no video data is available.
func (*HTMLVideoElement) SetPoster(poster string) {
	js.Rewrite("$<.poster = $1", poster)
}

// VideoHeight prop Gets the intrinsic height of a video in CSS pixels, or zero if the dimensions are not known.
func (*HTMLVideoElement) VideoHeight() (videoHeight uint) {
	js.Rewrite("$<.videoHeight")
	return videoHeight
}

// VideoWidth prop Gets the intrinsic width of a video in CSS pixels, or zero if the dimensions are not known.
func (*HTMLVideoElement) VideoWidth() (videoWidth uint) {
	js.Rewrite("$<.videoWidth")
	return videoWidth
}

// WebkitDisplayingFullscreen prop
func (*HTMLVideoElement) WebkitDisplayingFullscreen() (webkitDisplayingFullscreen bool) {
	js.Rewrite("$<.webkitDisplayingFullscreen")
	return webkitDisplayingFullscreen
}

// WebkitSupportsFullscreen prop
func (*HTMLVideoElement) WebkitSupportsFullscreen() (webkitSupportsFullscreen bool) {
	js.Rewrite("$<.webkitSupportsFullscreen")
	return webkitSupportsFullscreen
}

// Width prop Gets or sets the width of the video element.
func (*HTMLVideoElement) Width() (width uint) {
	js.Rewrite("$<.width")
	return width
}

// Width prop Gets or sets the width of the video element.
func (*HTMLVideoElement) SetWidth(width uint) {
	js.Rewrite("$<.width = $1", width)
}
