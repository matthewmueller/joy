package geolocation

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/position"
	"github.com/matthewmueller/golly/internal/gendom/dom/positionerror"
	"github.com/matthewmueller/golly/internal/gendom/dom/positionoptions"
	"github.com/matthewmueller/golly/js"
)

type Geolocation struct {
}

func (*Geolocation) ClearWatch(watchId int) {
	js.Rewrite("$<.clearWatch($1)", watchId)
}

func (*Geolocation) GetCurrentPosition(successCallback func(position *position.Position), errorCallback func(err *positionerror.PositionError), options *positionoptions.PositionOptions) {
	js.Rewrite("$<.getCurrentPosition($1, $2, $3)", successCallback, errorCallback, options)
}

func (*Geolocation) WatchPosition(successCallback func(position *position.Position), errorCallback func(err *positionerror.PositionError), options *positionoptions.PositionOptions) (i int) {
	js.Rewrite("$<.watchPosition($1, $2, $3)", successCallback, errorCallback, options)
	return i
}
