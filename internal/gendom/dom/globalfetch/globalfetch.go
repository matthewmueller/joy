package globalfetch

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/request"
	"github.com/matthewmueller/golly/internal/gendom/dom/requestinit"
	"github.com/matthewmueller/golly/internal/gendom/dom/response"
	"github.com/matthewmueller/golly/js"
)

type GlobalFetch struct {
}

func (*GlobalFetch) Fetch(input *request.Request, init *requestinit.RequestInit) (r *response.Response) {
	js.Rewrite("await $<.fetch($1, $2)", input, init)
	return r
}
