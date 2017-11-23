package navigatorgeolocation

import (
	"github.com/matthewmueller/golly/dom/geolocation"
	"github.com/matthewmueller/golly/js"
)

// NavigatorGeolocation struct
// js:"NavigatorGeolocation,omit"
type NavigatorGeolocation struct {
}

// Geolocation prop
func (*NavigatorGeolocation) Geolocation() (geolocation *geolocation.Geolocation) {
	js.Rewrite("$<.geolocation")
	return geolocation
}
