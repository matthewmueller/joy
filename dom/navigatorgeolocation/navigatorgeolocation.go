package navigatorgeolocation

import "github.com/matthewmueller/joy/dom/geolocation"

// NavigatorGeolocation interface
// js:"NavigatorGeolocation"
type NavigatorGeolocation interface {

	// Geolocation prop
	// js:"geolocation"
	// jsrewrite:"$_.geolocation"
	Geolocation() (geolocation *geolocation.Geolocation)
}
