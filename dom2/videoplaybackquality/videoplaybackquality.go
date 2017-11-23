package videoplaybackquality

import "github.com/matthewmueller/golly/js"

// VideoPlaybackQuality struct
// js:"VideoPlaybackQuality,omit"
type VideoPlaybackQuality struct {
}

// CorruptedVideoFrames
func (*VideoPlaybackQuality) CorruptedVideoFrames() (corruptedVideoFrames uint) {
	js.Rewrite("$<.CorruptedVideoFrames")
	return corruptedVideoFrames
}

// CreationTime
func (*VideoPlaybackQuality) CreationTime() (creationTime int) {
	js.Rewrite("$<.CreationTime")
	return creationTime
}

// DroppedVideoFrames
func (*VideoPlaybackQuality) DroppedVideoFrames() (droppedVideoFrames uint) {
	js.Rewrite("$<.DroppedVideoFrames")
	return droppedVideoFrames
}

// TotalFrameDelay
func (*VideoPlaybackQuality) TotalFrameDelay() (totalFrameDelay float32) {
	js.Rewrite("$<.TotalFrameDelay")
	return totalFrameDelay
}

// TotalVideoFrames
func (*VideoPlaybackQuality) TotalVideoFrames() (totalVideoFrames uint) {
	js.Rewrite("$<.TotalVideoFrames")
	return totalVideoFrames
}
