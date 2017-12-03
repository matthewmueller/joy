package videoplaybackquality

import "github.com/matthewmueller/joy/macro"

// VideoPlaybackQuality struct
// js:"VideoPlaybackQuality,omit"
type VideoPlaybackQuality struct {
}

// CorruptedVideoFrames prop
// js:"corruptedVideoFrames"
func (*VideoPlaybackQuality) CorruptedVideoFrames() (corruptedVideoFrames uint) {
	macro.Rewrite("$_.corruptedVideoFrames")
	return corruptedVideoFrames
}

// CreationTime prop
// js:"creationTime"
func (*VideoPlaybackQuality) CreationTime() (creationTime int) {
	macro.Rewrite("$_.creationTime")
	return creationTime
}

// DroppedVideoFrames prop
// js:"droppedVideoFrames"
func (*VideoPlaybackQuality) DroppedVideoFrames() (droppedVideoFrames uint) {
	macro.Rewrite("$_.droppedVideoFrames")
	return droppedVideoFrames
}

// TotalFrameDelay prop
// js:"totalFrameDelay"
func (*VideoPlaybackQuality) TotalFrameDelay() (totalFrameDelay float32) {
	macro.Rewrite("$_.totalFrameDelay")
	return totalFrameDelay
}

// TotalVideoFrames prop
// js:"totalVideoFrames"
func (*VideoPlaybackQuality) TotalVideoFrames() (totalVideoFrames uint) {
	macro.Rewrite("$_.totalVideoFrames")
	return totalVideoFrames
}
