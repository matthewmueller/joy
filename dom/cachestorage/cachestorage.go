package cachestorage

import (
	"github.com/matthewmueller/joy/dom/cache"
	"github.com/matthewmueller/joy/dom/cachequeryoptions"
	"github.com/matthewmueller/joy/dom/request"
	"github.com/matthewmueller/joy/macro"
)

// CacheStorage struct
// js:"CacheStorage,omit"
type CacheStorage struct {
}

// Delete fn
// js:"delete"
func (*CacheStorage) Delete(cacheName string) (b bool) {
	macro.Rewrite("await $_.delete($1)", cacheName)
	return b
}

// Has fn
// js:"has"
func (*CacheStorage) Has(cacheName string) (b bool) {
	macro.Rewrite("await $_.has($1)", cacheName)
	return b
}

// Keys fn
// js:"keys"
func (*CacheStorage) Keys() (s []string) {
	macro.Rewrite("await $_.keys()")
	return s
}

// Match fn
// js:"match"
func (*CacheStorage) Match(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (i interface{}) {
	macro.Rewrite("await $_.match($1, $2)", request, options)
	return i
}

// Open fn
// js:"open"
func (*CacheStorage) Open(cacheName string) (c *cache.Cache) {
	macro.Rewrite("await $_.open($1)", cacheName)
	return c
}
