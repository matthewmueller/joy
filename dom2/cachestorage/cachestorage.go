package cachestorage

import (
	"github.com/matthewmueller/golly/dom2/cache"
	"github.com/matthewmueller/golly/dom2/cachequeryoptions"
	"github.com/matthewmueller/golly/js"
)

// js:"CacheStorage,omit"
type CacheStorage struct {
}

// Delete
func (*CacheStorage) Delete(cacheName string) (b bool) {
	js.Rewrite("await $<.Delete($1)", cacheName)
	return b
}

// Has
func (*CacheStorage) Has(cacheName string) (b bool) {
	js.Rewrite("await $<.Has($1)", cacheName)
	return b
}

// Keys
func (*CacheStorage) Keys() (s []string) {
	js.Rewrite("await $<.Keys()")
	return s
}

// Match
func (*CacheStorage) Match(request *Request, options *cachequeryoptions.CacheQueryOptions) (i interface{}) {
	js.Rewrite("await $<.Match($1, $2)", request, options)
	return i
}

// Open
func (*CacheStorage) Open(cacheName string) (c *cache.Cache) {
	js.Rewrite("await $<.Open($1)", cacheName)
	return c
}
