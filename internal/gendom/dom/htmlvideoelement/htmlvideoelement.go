package htmlvideoelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlmediaelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/videoplaybackquality"
	"github.com/matthewmueller/golly/js"
)

type HTMLVideoElement struct {
	*htmlmediaelement.HTMLMediaElement
}

func (*HTMLVideoElement) GetVideoPlaybackQuality() (v *videoplaybackquality.VideoPlaybackQuality) {
	js.Rewrite("$<.getVideoPlaybackQuality()")
	return v
}

func (*HTMLVideoElement) MsFrameStep(forward bool) {
	js.Rewrite("$<.msFrameStep($1)", forward)
}

func (*HTMLVideoElement) MsInsertVideoEffect(activatableClassId string, effectRequired bool, config interface{}) {
	js.Rewrite("$<.msInsertVideoEffect($1, $2, $3)", activatableClassId, effectRequired, config)
}

func (*HTMLVideoElement) MsSetVideoRectangle(left float32, top float32, right float32, bottom float32) {
	js.Rewrite("$<.msSetVideoRectangle($1, $2, $3, $4)", left, top, right, bottom)
}

func (*HTMLVideoElement) WebkitEnterFullscreen() {
	js.Rewrite("$<.webkitEnterFullscreen()")
}

func (*HTMLVideoElement) WebkitEnterFullScreen() {
	js.Rewrite("$<.webkitEnterFullScreen()")
}

func (*HTMLVideoElement) WebkitExitFullscreen() {
	js.Rewrite("$<.webkitExitFullscreen()")
}

func (*HTMLVideoElement) WebkitExitFullScreen() {
	js.Rewrite("$<.webkitExitFullScreen()")
}

func (*HTMLVideoElement) GetHeight() (height uint) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLVideoElement) SetHeight(height uint) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLVideoElement) GetMsHorizontalMirror() (msHorizontalMirror bool) {
	js.Rewrite("$<.msHorizontalMirror")
	return msHorizontalMirror
}

func (*HTMLVideoElement) SetMsHorizontalMirror(msHorizontalMirror bool) {
	js.Rewrite("$<.msHorizontalMirror = $1", msHorizontalMirror)
}

func (*HTMLVideoElement) GetMsIsLayoutOptimalForPlayback() (msIsLayoutOptimalForPlayback bool) {
	js.Rewrite("$<.msIsLayoutOptimalForPlayback")
	return msIsLayoutOptimalForPlayback
}

func (*HTMLVideoElement) GetMsIsStereo3d() (msIsStereo3D bool) {
	js.Rewrite("$<.msIsStereo3D")
	return msIsStereo3D
}

func (*HTMLVideoElement) GetMsStereo3dPackingMode() (msStereo3DPackingMode string) {
	js.Rewrite("$<.msStereo3DPackingMode")
	return msStereo3DPackingMode
}

func (*HTMLVideoElement) SetMsStereo3dPackingMode(msStereo3DPackingMode string) {
	js.Rewrite("$<.msStereo3DPackingMode = $1", msStereo3DPackingMode)
}

func (*HTMLVideoElement) GetMsStereo3dRenderMode() (msStereo3DRenderMode string) {
	js.Rewrite("$<.msStereo3DRenderMode")
	return msStereo3DRenderMode
}

func (*HTMLVideoElement) SetMsStereo3dRenderMode(msStereo3DRenderMode string) {
	js.Rewrite("$<.msStereo3DRenderMode = $1", msStereo3DRenderMode)
}

func (*HTMLVideoElement) GetMsZoom() (msZoom bool) {
	js.Rewrite("$<.msZoom")
	return msZoom
}

func (*HTMLVideoElement) SetMsZoom(msZoom bool) {
	js.Rewrite("$<.msZoom = $1", msZoom)
}

func (*HTMLVideoElement) GetOnMSVideoFormatChanged() (MSVideoFormatChanged *event.Event) {
	js.Rewrite("$<.onMSVideoFormatChanged")
	return MSVideoFormatChanged
}

func (*HTMLVideoElement) SetOnMSVideoFormatChanged(MSVideoFormatChanged *event.Event) {
	js.Rewrite("$<.onMSVideoFormatChanged = $1", MSVideoFormatChanged)
}

func (*HTMLVideoElement) GetOnMSVideoFrameStepCompleted() (MSVideoFrameStepCompleted *event.Event) {
	js.Rewrite("$<.onMSVideoFrameStepCompleted")
	return MSVideoFrameStepCompleted
}

func (*HTMLVideoElement) SetOnMSVideoFrameStepCompleted(MSVideoFrameStepCompleted *event.Event) {
	js.Rewrite("$<.onMSVideoFrameStepCompleted = $1", MSVideoFrameStepCompleted)
}

func (*HTMLVideoElement) GetOnMSVideoOptimalLayoutChanged() (MSVideoOptimalLayoutChanged *event.Event) {
	js.Rewrite("$<.onMSVideoOptimalLayoutChanged")
	return MSVideoOptimalLayoutChanged
}

func (*HTMLVideoElement) SetOnMSVideoOptimalLayoutChanged(MSVideoOptimalLayoutChanged *event.Event) {
	js.Rewrite("$<.onMSVideoOptimalLayoutChanged = $1", MSVideoOptimalLayoutChanged)
}

func (*HTMLVideoElement) GetPoster() (poster string) {
	js.Rewrite("$<.poster")
	return poster
}

func (*HTMLVideoElement) SetPoster(poster string) {
	js.Rewrite("$<.poster = $1", poster)
}

func (*HTMLVideoElement) GetVideoHeight() (videoHeight uint) {
	js.Rewrite("$<.videoHeight")
	return videoHeight
}

func (*HTMLVideoElement) GetVideoWidth() (videoWidth uint) {
	js.Rewrite("$<.videoWidth")
	return videoWidth
}

func (*HTMLVideoElement) GetWebkitDisplayingFullscreen() (webkitDisplayingFullscreen bool) {
	js.Rewrite("$<.webkitDisplayingFullscreen")
	return webkitDisplayingFullscreen
}

func (*HTMLVideoElement) GetWebkitSupportsFullscreen() (webkitSupportsFullscreen bool) {
	js.Rewrite("$<.webkitSupportsFullscreen")
	return webkitSupportsFullscreen
}

func (*HTMLVideoElement) GetWidth() (width uint) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLVideoElement) SetWidth(width uint) {
	js.Rewrite("$<.width = $1", width)
}
