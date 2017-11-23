package coordinates

import "github.com/matthewmueller/golly/js"

// js:"Coordinates,omit"
type Coordinates struct {
}

// Accuracy
func (*Coordinates) Accuracy() (accuracy float32) {
	js.Rewrite("$<.Accuracy")
	return accuracy
}

// Altitude
func (*Coordinates) Altitude() (altitude float32) {
	js.Rewrite("$<.Altitude")
	return altitude
}

// AltitudeAccuracy
func (*Coordinates) AltitudeAccuracy() (altitudeAccuracy float32) {
	js.Rewrite("$<.AltitudeAccuracy")
	return altitudeAccuracy
}

// Heading
func (*Coordinates) Heading() (heading float32) {
	js.Rewrite("$<.Heading")
	return heading
}

// Latitude
func (*Coordinates) Latitude() (latitude float32) {
	js.Rewrite("$<.Latitude")
	return latitude
}

// Longitude
func (*Coordinates) Longitude() (longitude float32) {
	js.Rewrite("$<.Longitude")
	return longitude
}

// Speed
func (*Coordinates) Speed() (speed float32) {
	js.Rewrite("$<.Speed")
	return speed
}
