package navigatorusermedia

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/mediadevices"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastream"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastreamconstraints"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastreamerror"
	"github.com/matthewmueller/golly/js"
)

type NavigatorUserMedia struct {
}

func (*NavigatorUserMedia) GetUserMedia(constraints *mediastreamconstraints.MediaStreamConstraints, successCallback func(stream *mediastream.MediaStream), errorCallback func(err *mediastreamerror.MediaStreamError)) {
	js.Rewrite("$<.getUserMedia($1, $2, $3)", constraints, successCallback, errorCallback)
}

func (*NavigatorUserMedia) GetMediaDevices() (mediaDevices *mediadevices.MediaDevices) {
	js.Rewrite("$<.mediaDevices")
	return mediaDevices
}
