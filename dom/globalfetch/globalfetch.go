package globalfetch

import (
	"github.com/matthewmueller/joy/dom/request"
	"github.com/matthewmueller/joy/dom/requestinit"
	"github.com/matthewmueller/joy/dom/response"
)

// GlobalFetch interface
// js:"GlobalFetch"
type GlobalFetch interface {

	// Fetch
	// js:"fetch"
	// jsrewrite:"await $_.fetch($1, $2)"
	Fetch(input *request.Request, init *requestinit.RequestInit) (r *response.Response)
}
