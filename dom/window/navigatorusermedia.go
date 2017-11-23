package window

import (
	"github.com/matthewmueller/golly/dom/mediastreamconstraints"
	"github.com/matthewmueller/golly/dom/mediastreamerror"
	"github.com/matthewmueller/golly/js"
)

// NavigatorUserMedia struct
// js:"NavigatorUserMedia,omit"
type NavigatorUserMedia struct {
}

// GetUserMedia fn
func (*NavigatorUserMedia) GetUserMedia(constraints *mediastreamconstraints.MediaStreamConstraints, successCallback func(stream *MediaStream), errorCallback func(err *mediastreamerror.MediaStreamError)) {
	js.Rewrite("$<.getUserMedia($1, $2, $3)", constraints, successCallback, errorCallback)
}

// MediaDevices prop
func (*NavigatorUserMedia) MediaDevices() (mediaDevices *MediaDevices) {
	js.Rewrite("$<.mediaDevices")
	return mediaDevices
}
