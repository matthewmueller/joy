package coordinates

import "github.com/matthewmueller/golly/js"

type Coordinates struct {
}

func (*Coordinates) GetAccuracy() (accuracy float32) {
	js.Rewrite("$<.accuracy")
	return accuracy
}

func (*Coordinates) GetAltitude() (altitude float32) {
	js.Rewrite("$<.altitude")
	return altitude
}

func (*Coordinates) GetAltitudeAccuracy() (altitudeAccuracy float32) {
	js.Rewrite("$<.altitudeAccuracy")
	return altitudeAccuracy
}

func (*Coordinates) GetHeading() (heading float32) {
	js.Rewrite("$<.heading")
	return heading
}

func (*Coordinates) GetLatitude() (latitude float32) {
	js.Rewrite("$<.latitude")
	return latitude
}

func (*Coordinates) GetLongitude() (longitude float32) {
	js.Rewrite("$<.longitude")
	return longitude
}

func (*Coordinates) GetSpeed() (speed float32) {
	js.Rewrite("$<.speed")
	return speed
}
