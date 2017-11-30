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
	Fetch(input *request.Request, init *requestinit.RequestInit) (r *response.Response)
}
