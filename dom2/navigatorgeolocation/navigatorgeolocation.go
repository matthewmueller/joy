package navigatorgeolocation

import (
	"github.com/matthewmueller/golly/dom2/geolocation"
	"github.com/matthewmueller/golly/js"
)

// NavigatorGeolocation struct
// js:"NavigatorGeolocation,omit"
type NavigatorGeolocation struct {
}

// Geolocation
func (*NavigatorGeolocation) Geolocation() (geolocation *geolocation.Geolocation) {
	js.Rewrite("$<.Geolocation")
	return geolocation
}
