package cache

import (
	"github.com/matthewmueller/golly/dom2/cachequeryoptions"
	"github.com/matthewmueller/golly/dom2/request"
	"github.com/matthewmueller/golly/dom2/response"
	"github.com/matthewmueller/golly/js"
)

// Cache struct
// js:"Cache,omit"
type Cache struct {
}

// Add fn
func (*Cache) Add(request *request.Request) {
	js.Rewrite("await $<.add($1)", request)
}

// AddAll fn
func (*Cache) AddAll(requests []*request.Request) {
	js.Rewrite("await $<.addAll($1)", requests)
}

// Delete fn
func (*Cache) Delete(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (b bool) {
	js.Rewrite("await $<.delete($1, $2)", request, options)
	return b
}

// Keys fn
func (*Cache) Keys(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (r []*request.Request) {
	js.Rewrite("await $<.keys($1, $2)", request, options)
	return r
}

// Match fn
func (*Cache) Match(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (r *response.Response) {
	js.Rewrite("await $<.match($1, $2)", request, options)
	return r
}

// MatchAll fn
func (*Cache) MatchAll(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (r []*response.Response) {
	js.Rewrite("await $<.matchAll($1, $2)", request, options)
	return r
}

// Put fn
func (*Cache) Put(request *request.Request, response *response.Response) {
	js.Rewrite("await $<.put($1, $2)", request, response)
}
