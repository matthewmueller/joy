package request

import (
	"github.com/matthewmueller/golly/dom/blob"
	"github.com/matthewmueller/golly/dom/headers"
	"github.com/matthewmueller/golly/dom/referrerpolicy"
	"github.com/matthewmueller/golly/dom/requestcache"
	"github.com/matthewmueller/golly/dom/requestcredentials"
	"github.com/matthewmueller/golly/dom/requestdestination"
	"github.com/matthewmueller/golly/dom/requestinit"
	"github.com/matthewmueller/golly/dom/requestmode"
	"github.com/matthewmueller/golly/dom/requestredirect"
	"github.com/matthewmueller/golly/dom/requesttype"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(input interface{}, init *requestinit.RequestInit) *Request {
	js.Rewrite("Request")
	return &Request{}
}

// Request struct
// js:"Request,omit"
type Request struct {
}

// ArrayBuffer fn
func (*Request) ArrayBuffer() (b []byte) {
	js.Rewrite("await $<.arrayBuffer()")
	return b
}

// Blob fn
func (*Request) Blob() (b blob.Blob) {
	js.Rewrite("await $<.blob()")
	return b
}

// JSON fn
func (*Request) JSON() (i interface{}) {
	js.Rewrite("await $<.json()")
	return i
}

// Text fn
func (*Request) Text() (s string) {
	js.Rewrite("await $<.text()")
	return s
}

// Clone fn
func (*Request) Clone() (r *Request) {
	js.Rewrite("$<.clone()")
	return r
}

// BodyUsed prop
func (*Request) BodyUsed() (bodyUsed bool) {
	js.Rewrite("$<.bodyUsed")
	return bodyUsed
}

// Cache prop
func (*Request) Cache() (cache *requestcache.RequestCache) {
	js.Rewrite("$<.cache")
	return cache
}

// Credentials prop
func (*Request) Credentials() (credentials *requestcredentials.RequestCredentials) {
	js.Rewrite("$<.credentials")
	return credentials
}

// Destination prop
func (*Request) Destination() (destination *requestdestination.RequestDestination) {
	js.Rewrite("$<.destination")
	return destination
}

// Headers prop
func (*Request) Headers() (headers *headers.Headers) {
	js.Rewrite("$<.headers")
	return headers
}

// Integrity prop
func (*Request) Integrity() (integrity string) {
	js.Rewrite("$<.integrity")
	return integrity
}

// Keepalive prop
func (*Request) Keepalive() (keepalive bool) {
	js.Rewrite("$<.keepalive")
	return keepalive
}

// Method prop
func (*Request) Method() (method string) {
	js.Rewrite("$<.method")
	return method
}

// Mode prop
func (*Request) Mode() (mode *requestmode.RequestMode) {
	js.Rewrite("$<.mode")
	return mode
}

// Redirect prop
func (*Request) Redirect() (redirect *requestredirect.RequestRedirect) {
	js.Rewrite("$<.redirect")
	return redirect
}

// Referrer prop
func (*Request) Referrer() (referrer string) {
	js.Rewrite("$<.referrer")
	return referrer
}

// ReferrerPolicy prop
func (*Request) ReferrerPolicy() (referrerPolicy *referrerpolicy.ReferrerPolicy) {
	js.Rewrite("$<.referrerPolicy")
	return referrerPolicy
}

// Type prop
func (*Request) Type() (kind *requesttype.RequestType) {
	js.Rewrite("$<.type")
	return kind
}

// URL prop
func (*Request) URL() (url string) {
	js.Rewrite("$<.url")
	return url
}
