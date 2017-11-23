package cache

import (
	"github.com/matthewmueller/golly/dom2/cachequeryoptions"
	"github.com/matthewmueller/golly/dom2/request"
	"github.com/matthewmueller/golly/js"
)

// js:"Cache,omit"
type Cache struct {
}

// Add
func (*Cache) Add(request *Request) {
	js.Rewrite("await $<.Add($1)", request)
}

// AddAll
func (*Cache) AddAll(requests []*Request) {
	js.Rewrite("await $<.AddAll($1)", requests)
}

// Delete
func (*Cache) Delete(request *Request, options *cachequeryoptions.CacheQueryOptions) (b bool) {
	js.Rewrite("await $<.Delete($1, $2)", request, options)
	return b
}

// Keys
func (*Cache) Keys(request *Request, options *cachequeryoptions.CacheQueryOptions) (r []*request.Request) {
	js.Rewrite("await $<.Keys($1, $2)", request, options)
	return r
}

// Match
func (*Cache) Match(request *Request, options *cachequeryoptions.CacheQueryOptions) (r *response.Response) {
	js.Rewrite("await $<.Match($1, $2)", request, options)
	return r
}

// MatchAll
func (*Cache) MatchAll(request *Request, options *cachequeryoptions.CacheQueryOptions) (r []*response.Response) {
	js.Rewrite("await $<.MatchAll($1, $2)", request, options)
	return r
}

// Put
func (*Cache) Put(request *Request, response *response.Response) {
	js.Rewrite("await $<.Put($1, $2)", request, response)
}
