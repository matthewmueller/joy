package globalfetch

import "github.com/matthewmueller/golly/dom2/requestinit"

// js:"GlobalFetch,omit"
type GlobalFetch interface {

	// Fetch
	Fetch(input *Request, init *requestinit.RequestInit) (r *response.Response)
}
