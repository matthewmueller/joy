package cache

import (
	"github.com/matthewmueller/joy/dom/cachequeryoptions"
	"github.com/matthewmueller/joy/dom/request"
	"github.com/matthewmueller/joy/dom/response"
	"github.com/matthewmueller/joy/js"
)

// Cache struct
// js:"Cache,omit"
type Cache struct {
}

// Add fn
// js:"add"
func (*Cache) Add(request *request.Request) {
	js.Rewrite("await $_.add($1)", request)
}

// AddAll fn
// js:"addAll"
func (*Cache) AddAll(requests []*request.Request) {
	js.Rewrite("await $_.addAll($1)", requests)
}

// Delete fn
// js:"delete"
func (*Cache) Delete(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (b bool) {
	js.Rewrite("await $_.delete($1, $2)", request, options)
	return b
}

// Keys fn
// js:"keys"
func (*Cache) Keys(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (r []*request.Request) {
	js.Rewrite("await $_.keys($1, $2)", request, options)
	return r
}

// Match fn
// js:"match"
func (*Cache) Match(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (r *response.Response) {
	js.Rewrite("await $_.match($1, $2)", request, options)
	return r
}

// MatchAll fn
// js:"matchAll"
func (*Cache) MatchAll(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (r []*response.Response) {
	js.Rewrite("await $_.matchAll($1, $2)", request, options)
	return r
}

// Put fn
// js:"put"
func (*Cache) Put(request *request.Request, response *response.Response) {
	js.Rewrite("await $_.put($1, $2)", request, response)
}
