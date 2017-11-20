package navigatorgeolocation

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/geolocation"
	"github.com/matthewmueller/golly/js"
)

type NavigatorGeolocation struct {
}

func (*NavigatorGeolocation) GetGeolocation() (geolocation *geolocation.Geolocation) {
	js.Rewrite("$<.geolocation")
	return geolocation
}
