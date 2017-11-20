package cachestorage

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cache"
	"github.com/matthewmueller/golly/internal/gendom/dom/cachequeryoptions"
	"github.com/matthewmueller/golly/internal/gendom/dom/request"
	"github.com/matthewmueller/golly/js"
)

type CacheStorage struct {
}

func (*CacheStorage) Delete(cacheName string) (b bool) {
	js.Rewrite("await $<.delete($1)", cacheName)
	return b
}

func (*CacheStorage) Has(cacheName string) (b bool) {
	js.Rewrite("await $<.has($1)", cacheName)
	return b
}

func (*CacheStorage) Keys() (s []string) {
	js.Rewrite("await $<.keys()")
	return s
}

func (*CacheStorage) Match(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (i interface{}) {
	js.Rewrite("await $<.match($1, $2)", request, options)
	return i
}

func (*CacheStorage) Open(cacheName string) (c *cache.Cache) {
	js.Rewrite("await $<.open($1)", cacheName)
	return c
}
