package videoplaybackquality

import "github.com/matthewmueller/joy/js"

// VideoPlaybackQuality struct
// js:"VideoPlaybackQuality,omit"
type VideoPlaybackQuality struct {
}

// CorruptedVideoFrames prop
// js:"corruptedVideoFrames"
func (*VideoPlaybackQuality) CorruptedVideoFrames() (corruptedVideoFrames uint) {
	js.Rewrite("$_.corruptedVideoFrames")
	return corruptedVideoFrames
}

// CreationTime prop
// js:"creationTime"
func (*VideoPlaybackQuality) CreationTime() (creationTime int) {
	js.Rewrite("$_.creationTime")
	return creationTime
}

// DroppedVideoFrames prop
// js:"droppedVideoFrames"
func (*VideoPlaybackQuality) DroppedVideoFrames() (droppedVideoFrames uint) {
	js.Rewrite("$_.droppedVideoFrames")
	return droppedVideoFrames
}

// TotalFrameDelay prop
// js:"totalFrameDelay"
func (*VideoPlaybackQuality) TotalFrameDelay() (totalFrameDelay float32) {
	js.Rewrite("$_.totalFrameDelay")
	return totalFrameDelay
}

// TotalVideoFrames prop
// js:"totalVideoFrames"
func (*VideoPlaybackQuality) TotalVideoFrames() (totalVideoFrames uint) {
	js.Rewrite("$_.totalVideoFrames")
	return totalVideoFrames
}
