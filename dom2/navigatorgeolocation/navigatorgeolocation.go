package navigatorgeolocation

import "github.com/matthewmueller/golly/dom2/geolocation"

// js:"NavigatorGeolocation,omit"
type NavigatorGeolocation interface {

	// Geolocation
	Geolocation() (geolocation *geolocation.Geolocation)
}
