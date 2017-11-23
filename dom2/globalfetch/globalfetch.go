package globalfetch

import (
	"github.com/matthewmueller/golly/dom2/request"
	"github.com/matthewmueller/golly/dom2/requestinit"
	"github.com/matthewmueller/golly/dom2/response"
	"github.com/matthewmueller/golly/js"
)

// GlobalFetch struct
// js:"GlobalFetch,omit"
type GlobalFetch struct {
}

// Fetch
func (*GlobalFetch) Fetch(input *request.Request, init *requestinit.RequestInit) (r *response.Response) {
	js.Rewrite("await $<.Fetch($1, $2)", input, init)
	return r
}
