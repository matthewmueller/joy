package geolocation

import (
	"github.com/matthewmueller/golly/dom/position"
	"github.com/matthewmueller/golly/dom/positionerror"
	"github.com/matthewmueller/golly/dom/positionoptions"
	"github.com/matthewmueller/golly/js"
)

// Geolocation struct
// js:"Geolocation,omit"
type Geolocation struct {
}

// ClearWatch fn
func (*Geolocation) ClearWatch(watchId int) {
	js.Rewrite("$<.clearWatch($1)", watchId)
}

// GetCurrentPosition fn
func (*Geolocation) GetCurrentPosition(successCallback func(position *position.Position), errorCallback *func(err *positionerror.PositionError), options *positionoptions.PositionOptions) {
	js.Rewrite("$<.getCurrentPosition($1, $2, $3)", successCallback, errorCallback, options)
}

// WatchPosition fn
func (*Geolocation) WatchPosition(successCallback func(position *position.Position), errorCallback *func(err *positionerror.PositionError), options *positionoptions.PositionOptions) (i int) {
	js.Rewrite("$<.watchPosition($1, $2, $3)", successCallback, errorCallback, options)
	return i
}
