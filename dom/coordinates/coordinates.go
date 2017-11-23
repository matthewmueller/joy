package coordinates

import "github.com/matthewmueller/golly/js"

// Coordinates struct
// js:"Coordinates,omit"
type Coordinates struct {
}

// Accuracy prop
func (*Coordinates) Accuracy() (accuracy float32) {
	js.Rewrite("$<.accuracy")
	return accuracy
}

// Altitude prop
func (*Coordinates) Altitude() (altitude float32) {
	js.Rewrite("$<.altitude")
	return altitude
}

// AltitudeAccuracy prop
func (*Coordinates) AltitudeAccuracy() (altitudeAccuracy float32) {
	js.Rewrite("$<.altitudeAccuracy")
	return altitudeAccuracy
}

// Heading prop
func (*Coordinates) Heading() (heading float32) {
	js.Rewrite("$<.heading")
	return heading
}

// Latitude prop
func (*Coordinates) Latitude() (latitude float32) {
	js.Rewrite("$<.latitude")
	return latitude
}

// Longitude prop
func (*Coordinates) Longitude() (longitude float32) {
	js.Rewrite("$<.longitude")
	return longitude
}

// Speed prop
func (*Coordinates) Speed() (speed float32) {
	js.Rewrite("$<.speed")
	return speed
}
