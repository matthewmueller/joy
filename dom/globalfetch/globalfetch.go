package globalfetch

import (
	"github.com/matthewmueller/golly/dom/request"
	"github.com/matthewmueller/golly/dom/requestinit"
	"github.com/matthewmueller/golly/dom/response"
)

// GlobalFetch interface
// js:"GlobalFetch"
type GlobalFetch interface {

	// Fetch
	// js:"fetch"
	// jsrewrite:"await $_.fetch($1, $2)"
	Fetch(input *request.Request, init *requestinit.RequestInit) (r *response.Response)
}
