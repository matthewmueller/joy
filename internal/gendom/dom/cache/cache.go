package cache

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cachequeryoptions"
	"github.com/matthewmueller/golly/internal/gendom/dom/request"
	"github.com/matthewmueller/golly/internal/gendom/dom/response"
	"github.com/matthewmueller/golly/js"
)

type Cache struct {
}

func (*Cache) Add(request *request.Request) {
	js.Rewrite("await $<.add($1)", request)
}

func (*Cache) AddAll(requests []*request.Request) {
	js.Rewrite("await $<.addAll($1)", requests)
}

func (*Cache) Delete(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (b bool) {
	js.Rewrite("await $<.delete($1, $2)", request, options)
	return b
}

func (*Cache) Keys(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (r []*request.Request) {
	js.Rewrite("await $<.keys($1, $2)", request, options)
	return r
}

func (*Cache) Match(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (r *response.Response) {
	js.Rewrite("await $<.match($1, $2)", request, options)
	return r
}

func (*Cache) MatchAll(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (r []*response.Response) {
	js.Rewrite("await $<.matchAll($1, $2)", request, options)
	return r
}

func (*Cache) Put(request *request.Request, response *response.Response) {
	js.Rewrite("await $<.put($1, $2)", request, response)
}
