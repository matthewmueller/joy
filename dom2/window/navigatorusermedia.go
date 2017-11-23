package window

import (
	"github.com/matthewmueller/golly/dom2/mediastreamconstraints"
	"github.com/matthewmueller/golly/dom2/mediastreamerror"
)

// js:"NavigatorUserMedia,omit"
type NavigatorUserMedia interface {

	// GetUserMedia
	GetUserMedia(constraints *mediastreamconstraints.MediaStreamConstraints, successCallback func(stream *MediaStream), errorCallback func(err *mediastreamerror.MediaStreamError))

	// MediaDevices
	MediaDevices() (mediaDevices *MediaDevices)
}
