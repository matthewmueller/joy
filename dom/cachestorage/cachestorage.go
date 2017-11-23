package cachestorage

import (
	"github.com/matthewmueller/golly/dom/cache"
	"github.com/matthewmueller/golly/dom/cachequeryoptions"
	"github.com/matthewmueller/golly/dom/request"
	"github.com/matthewmueller/golly/js"
)

// CacheStorage struct
// js:"CacheStorage,omit"
type CacheStorage struct {
}

// Delete fn
func (*CacheStorage) Delete(cacheName string) (b bool) {
	js.Rewrite("await $<.delete($1)", cacheName)
	return b
}

// Has fn
func (*CacheStorage) Has(cacheName string) (b bool) {
	js.Rewrite("await $<.has($1)", cacheName)
	return b
}

// Keys fn
func (*CacheStorage) Keys() (s []string) {
	js.Rewrite("await $<.keys()")
	return s
}

// Match fn
func (*CacheStorage) Match(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (i interface{}) {
	js.Rewrite("await $<.match($1, $2)", request, options)
	return i
}

// Open fn
func (*CacheStorage) Open(cacheName string) (c *cache.Cache) {
	js.Rewrite("await $<.open($1)", cacheName)
	return c
}
