package coordinates

import "github.com/matthewmueller/joy/js"

// Coordinates struct
// js:"Coordinates,omit"
type Coordinates struct {
}

// Accuracy prop
// js:"accuracy"
func (*Coordinates) Accuracy() (accuracy float32) {
	js.Rewrite("$_.accuracy")
	return accuracy
}

// Altitude prop
// js:"altitude"
func (*Coordinates) Altitude() (altitude float32) {
	js.Rewrite("$_.altitude")
	return altitude
}

// AltitudeAccuracy prop
// js:"altitudeAccuracy"
func (*Coordinates) AltitudeAccuracy() (altitudeAccuracy float32) {
	js.Rewrite("$_.altitudeAccuracy")
	return altitudeAccuracy
}

// Heading prop
// js:"heading"
func (*Coordinates) Heading() (heading float32) {
	js.Rewrite("$_.heading")
	return heading
}

// Latitude prop
// js:"latitude"
func (*Coordinates) Latitude() (latitude float32) {
	js.Rewrite("$_.latitude")
	return latitude
}

// Longitude prop
// js:"longitude"
func (*Coordinates) Longitude() (longitude float32) {
	js.Rewrite("$_.longitude")
	return longitude
}

// Speed prop
// js:"speed"
func (*Coordinates) Speed() (speed float32) {
	js.Rewrite("$_.speed")
	return speed
}
