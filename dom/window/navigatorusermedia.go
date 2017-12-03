package window

import (
	"github.com/matthewmueller/joy/dom/mediastreamconstraints"
	"github.com/matthewmueller/joy/dom/mediastreamerror"
)

// NavigatorUserMedia interface
// js:"NavigatorUserMedia"
type NavigatorUserMedia interface {

	// GetUserMedia
	// js:"getUserMedia"
	// jsrewrite:"$_.getUserMedia($1, $2, $3)"
	GetUserMedia(constraints *mediastreamconstraints.MediaStreamConstraints, successCallback func(stream *MediaStream), errorCallback func(err *mediastreamerror.MediaStreamError))

	// MediaDevices prop
	// js:"mediaDevices"
	// jsrewrite:"$_.mediaDevices"
	MediaDevices() (mediaDevices *MediaDevices)
}
