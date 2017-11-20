package videoplaybackquality

import "github.com/matthewmueller/golly/js"

type VideoPlaybackQuality struct {
}

func (*VideoPlaybackQuality) GetCorruptedVideoFrames() (corruptedVideoFrames uint) {
	js.Rewrite("$<.corruptedVideoFrames")
	return corruptedVideoFrames
}

func (*VideoPlaybackQuality) GetCreationTime() (creationTime int) {
	js.Rewrite("$<.creationTime")
	return creationTime
}

func (*VideoPlaybackQuality) GetDroppedVideoFrames() (droppedVideoFrames uint) {
	js.Rewrite("$<.droppedVideoFrames")
	return droppedVideoFrames
}

func (*VideoPlaybackQuality) GetTotalFrameDelay() (totalFrameDelay float32) {
	js.Rewrite("$<.totalFrameDelay")
	return totalFrameDelay
}

func (*VideoPlaybackQuality) GetTotalVideoFrames() (totalVideoFrames uint) {
	js.Rewrite("$<.totalVideoFrames")
	return totalVideoFrames
}
