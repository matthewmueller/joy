package window

import (
	"github.com/matthewmueller/golly/dom/mediastreamconstraints"
	"github.com/matthewmueller/golly/dom/mediastreamerror"
)

// NavigatorUserMedia interface
// js:"NavigatorUserMedia"
type NavigatorUserMedia interface {

	// GetUserMedia
	// js:"getUserMedia"
	GetUserMedia(constraints *mediastreamconstraints.MediaStreamConstraints, successCallback func(stream *MediaStream), errorCallback func(err *mediastreamerror.MediaStreamError))

	// MediaDevices prop
	// js:"mediaDevices"
	MediaDevices() (mediaDevices *MediaDevices)
}
