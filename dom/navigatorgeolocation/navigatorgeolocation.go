package navigatorgeolocation

import "github.com/matthewmueller/golly/dom/geolocation"

// NavigatorGeolocation interface
// js:"NavigatorGeolocation"
type NavigatorGeolocation interface {

	// Geolocation prop
	// js:"geolocation"
	Geolocation() (geolocation *geolocation.Geolocation)
}
