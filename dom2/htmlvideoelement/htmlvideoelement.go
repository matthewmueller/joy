package htmlvideoelement

import (
	"github.com/matthewmueller/golly/dom2/htmlmediaelement"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLVideoElement,omit"
type HTMLVideoElement struct {
	htmlmediaelement.HTMLMediaElement
}

// GetVideoPlaybackQuality
func (*HTMLVideoElement) GetVideoPlaybackQuality() (v *videoplaybackquality.VideoPlaybackQuality) {
	js.Rewrite("$<.GetVideoPlaybackQuality()")
	return v
}

// MsFrameStep
func (*HTMLVideoElement) MsFrameStep(forward bool) {
	js.Rewrite("$<.MsFrameStep($1)", forward)
}

// MsInsertVideoEffect
func (*HTMLVideoElement) MsInsertVideoEffect(activatableClassId string, effectRequired bool, config *interface{}) {
	js.Rewrite("$<.MsInsertVideoEffect($1, $2, $3)", activatableClassId, effectRequired, config)
}

// MsSetVideoRectangle
func (*HTMLVideoElement) MsSetVideoRectangle(left float32, top float32, right float32, bottom float32) {
	js.Rewrite("$<.MsSetVideoRectangle($1, $2, $3, $4)", left, top, right, bottom)
}

// WebkitEnterFullscreen
func (*HTMLVideoElement) WebkitEnterFullscreen() {
	js.Rewrite("$<.WebkitEnterFullscreen()")
}

// WebkitEnterFullScreen
func (*HTMLVideoElement) WebkitEnterFullScreen() {
	js.Rewrite("$<.WebkitEnterFullScreen()")
}

// WebkitExitFullscreen
func (*HTMLVideoElement) WebkitExitFullscreen() {
	js.Rewrite("$<.WebkitExitFullscreen()")
}

// WebkitExitFullScreen
func (*HTMLVideoElement) WebkitExitFullScreen() {
	js.Rewrite("$<.WebkitExitFullScreen()")
}

// Height Gets or sets the height of the video element.
func (*HTMLVideoElement) Height() (height uint) {
	js.Rewrite("$<.Height")
	return height
}

// Height Gets or sets the height of the video element.
func (*HTMLVideoElement) SetHeight(height uint) {
	js.Rewrite("$<.Height = $1", height)
}

// MsHorizontalMirror
func (*HTMLVideoElement) MsHorizontalMirror() (msHorizontalMirror bool) {
	js.Rewrite("$<.MsHorizontalMirror")
	return msHorizontalMirror
}

// MsHorizontalMirror
func (*HTMLVideoElement) SetMsHorizontalMirror(msHorizontalMirror bool) {
	js.Rewrite("$<.MsHorizontalMirror = $1", msHorizontalMirror)
}

// MsIsLayoutOptimalForPlayback
func (*HTMLVideoElement) MsIsLayoutOptimalForPlayback() (msIsLayoutOptimalForPlayback bool) {
	js.Rewrite("$<.MsIsLayoutOptimalForPlayback")
	return msIsLayoutOptimalForPlayback
}

// MsIsStereo3d
func (*HTMLVideoElement) MsIsStereo3d() (msIsStereo3D bool) {
	js.Rewrite("$<.MsIsStereo3d")
	return msIsStereo3D
}

// MsStereo3dPackingMode
func (*HTMLVideoElement) MsStereo3dPackingMode() (msStereo3DPackingMode string) {
	js.Rewrite("$<.MsStereo3dPackingMode")
	return msStereo3DPackingMode
}

// MsStereo3dPackingMode
func (*HTMLVideoElement) SetMsStereo3dPackingMode(msStereo3DPackingMode string) {
	js.Rewrite("$<.MsStereo3dPackingMode = $1", msStereo3DPackingMode)
}

// MsStereo3dRenderMode
func (*HTMLVideoElement) MsStereo3dRenderMode() (msStereo3DRenderMode string) {
	js.Rewrite("$<.MsStereo3dRenderMode")
	return msStereo3DRenderMode
}

// MsStereo3dRenderMode
func (*HTMLVideoElement) SetMsStereo3dRenderMode(msStereo3DRenderMode string) {
	js.Rewrite("$<.MsStereo3dRenderMode = $1", msStereo3DRenderMode)
}

// MsZoom
func (*HTMLVideoElement) MsZoom() (msZoom bool) {
	js.Rewrite("$<.MsZoom")
	return msZoom
}

// MsZoom
func (*HTMLVideoElement) SetMsZoom(msZoom bool) {
	js.Rewrite("$<.MsZoom = $1", msZoom)
}

// OnMSVideoFormatChanged
func (*HTMLVideoElement) OnMSVideoFormatChanged() (onmsvideoformatchanged func(window.Event)) {
	js.Rewrite("$<.OnMSVideoFormatChanged")
	return onmsvideoformatchanged
}

// OnMSVideoFormatChanged
func (*HTMLVideoElement) SetOnMSVideoFormatChanged(onmsvideoformatchanged func(window.Event)) {
	js.Rewrite("$<.OnMSVideoFormatChanged = $1", onmsvideoformatchanged)
}

// OnMSVideoFrameStepCompleted
func (*HTMLVideoElement) OnMSVideoFrameStepCompleted() (onmsvideoframestepcompleted func(window.Event)) {
	js.Rewrite("$<.OnMSVideoFrameStepCompleted")
	return onmsvideoframestepcompleted
}

// OnMSVideoFrameStepCompleted
func (*HTMLVideoElement) SetOnMSVideoFrameStepCompleted(onmsvideoframestepcompleted func(window.Event)) {
	js.Rewrite("$<.OnMSVideoFrameStepCompleted = $1", onmsvideoframestepcompleted)
}

// OnMSVideoOptimalLayoutChanged
func (*HTMLVideoElement) OnMSVideoOptimalLayoutChanged() (onmsvideooptimallayoutchanged func(window.Event)) {
	js.Rewrite("$<.OnMSVideoOptimalLayoutChanged")
	return onmsvideooptimallayoutchanged
}

// OnMSVideoOptimalLayoutChanged
func (*HTMLVideoElement) SetOnMSVideoOptimalLayoutChanged(onmsvideooptimallayoutchanged func(window.Event)) {
	js.Rewrite("$<.OnMSVideoOptimalLayoutChanged = $1", onmsvideooptimallayoutchanged)
}

// Poster Gets or sets a URL of an image to display, for example, like a movie poster. This can be a still frame from the video, or another image if no video data is available.
func (*HTMLVideoElement) Poster() (poster string) {
	js.Rewrite("$<.Poster")
	return poster
}

// Poster Gets or sets a URL of an image to display, for example, like a movie poster. This can be a still frame from the video, or another image if no video data is available.
func (*HTMLVideoElement) SetPoster(poster string) {
	js.Rewrite("$<.Poster = $1", poster)
}

// VideoHeight Gets the intrinsic height of a video in CSS pixels, or zero if the dimensions are not known.
func (*HTMLVideoElement) VideoHeight() (videoHeight uint) {
	js.Rewrite("$<.VideoHeight")
	return videoHeight
}

// VideoWidth Gets the intrinsic width of a video in CSS pixels, or zero if the dimensions are not known.
func (*HTMLVideoElement) VideoWidth() (videoWidth uint) {
	js.Rewrite("$<.VideoWidth")
	return videoWidth
}

// WebkitDisplayingFullscreen
func (*HTMLVideoElement) WebkitDisplayingFullscreen() (webkitDisplayingFullscreen bool) {
	js.Rewrite("$<.WebkitDisplayingFullscreen")
	return webkitDisplayingFullscreen
}

// WebkitSupportsFullscreen
func (*HTMLVideoElement) WebkitSupportsFullscreen() (webkitSupportsFullscreen bool) {
	js.Rewrite("$<.WebkitSupportsFullscreen")
	return webkitSupportsFullscreen
}

// Width Gets or sets the width of the video element.
func (*HTMLVideoElement) Width() (width uint) {
	js.Rewrite("$<.Width")
	return width
}

// Width Gets or sets the width of the video element.
func (*HTMLVideoElement) SetWidth(width uint) {
	js.Rewrite("$<.Width = $1", width)
}
