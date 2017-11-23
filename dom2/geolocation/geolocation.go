package geolocation

import "github.com/matthewmueller/golly/js"

// js:"Geolocation,omit"
type Geolocation struct {
}

// ClearWatch
func (*Geolocation) ClearWatch(watchId int) {
	js.Rewrite("$<.ClearWatch($1)", watchId)
}

// GetCurrentPosition
func (*Geolocation) GetCurrentPosition(successCallback func(position *position.Position), errorCallback *func(err *positionerror.PositionError), options *positionoptions.PositionOptions) {
	js.Rewrite("$<.GetCurrentPosition($1, $2, $3)", successCallback, errorCallback, options)
}

// WatchPosition
func (*Geolocation) WatchPosition(successCallback func(position *position.Position), errorCallback *func(err *positionerror.PositionError), options *positionoptions.PositionOptions) (i int) {
	js.Rewrite("$<.WatchPosition($1, $2, $3)", successCallback, errorCallback, options)
	return i
}
