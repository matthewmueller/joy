package globalfetch

import (
	"github.com/matthewmueller/golly/dom/request"
	"github.com/matthewmueller/golly/dom/requestinit"
	"github.com/matthewmueller/golly/dom/response"
	"github.com/matthewmueller/golly/js"
)

// GlobalFetch struct
// js:"GlobalFetch,omit"
type GlobalFetch struct {
}

// Fetch fn
func (*GlobalFetch) Fetch(input *request.Request, init *requestinit.RequestInit) (r *response.Response) {
	js.Rewrite("await $<.fetch($1, $2)", input, init)
	return r
}
