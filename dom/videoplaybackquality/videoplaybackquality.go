package videoplaybackquality

import "github.com/matthewmueller/golly/js"

// VideoPlaybackQuality struct
// js:"VideoPlaybackQuality,omit"
type VideoPlaybackQuality struct {
}

// CorruptedVideoFrames prop
func (*VideoPlaybackQuality) CorruptedVideoFrames() (corruptedVideoFrames uint) {
	js.Rewrite("$<.corruptedVideoFrames")
	return corruptedVideoFrames
}

// CreationTime prop
func (*VideoPlaybackQuality) CreationTime() (creationTime int) {
	js.Rewrite("$<.creationTime")
	return creationTime
}

// DroppedVideoFrames prop
func (*VideoPlaybackQuality) DroppedVideoFrames() (droppedVideoFrames uint) {
	js.Rewrite("$<.droppedVideoFrames")
	return droppedVideoFrames
}

// TotalFrameDelay prop
func (*VideoPlaybackQuality) TotalFrameDelay() (totalFrameDelay float32) {
	js.Rewrite("$<.totalFrameDelay")
	return totalFrameDelay
}

// TotalVideoFrames prop
func (*VideoPlaybackQuality) TotalVideoFrames() (totalVideoFrames uint) {
	js.Rewrite("$<.totalVideoFrames")
	return totalVideoFrames
}
