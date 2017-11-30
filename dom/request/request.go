package request

import (
	"github.com/matthewmueller/golly/dom/blob"
	"github.com/matthewmueller/golly/dom/headers"
	"github.com/matthewmueller/golly/dom/referrerpolicy"
	"github.com/matthewmueller/golly/dom/requestcache"
	"github.com/matthewmueller/golly/dom/requestcredentials"
	"github.com/matthewmueller/golly/dom/requestdestination"
	"github.com/matthewmueller/golly/dom/requestmode"
	"github.com/matthewmueller/golly/dom/requestredirect"
	"github.com/matthewmueller/golly/dom/requesttype"
	"github.com/matthewmueller/golly/js"
)

// Request struct
// js:"Request,omit"
type Request struct {
}

// Clone fn
// js:"clone"
func (*Request) Clone() (r *Request) {
	js.Rewrite("$_.clone()")
	return r
}

// ArrayBuffer fn
// js:"arrayBuffer"
func (*Request) ArrayBuffer() (b []byte) {
	js.Rewrite("await $_.arrayBuffer()")
	return b
}

// Blob fn
// js:"blob"
func (*Request) Blob() (b blob.Blob) {
	js.Rewrite("await $_.blob()")
	return b
}

// JSON fn
// js:"json"
func (*Request) JSON() (i interface{}) {
	js.Rewrite("await $_.json()")
	return i
}

// Text fn
// js:"text"
func (*Request) Text() (s string) {
	js.Rewrite("await $_.text()")
	return s
}

// Cache prop
// js:"cache"
func (*Request) Cache() (cache *requestcache.RequestCache) {
	js.Rewrite("$_.cache")
	return cache
}

// Credentials prop
// js:"credentials"
func (*Request) Credentials() (credentials *requestcredentials.RequestCredentials) {
	js.Rewrite("$_.credentials")
	return credentials
}

// Destination prop
// js:"destination"
func (*Request) Destination() (destination *requestdestination.RequestDestination) {
	js.Rewrite("$_.destination")
	return destination
}

// Headers prop
// js:"headers"
func (*Request) Headers() (headers *headers.Headers) {
	js.Rewrite("$_.headers")
	return headers
}

// Integrity prop
// js:"integrity"
func (*Request) Integrity() (integrity string) {
	js.Rewrite("$_.integrity")
	return integrity
}

// Keepalive prop
// js:"keepalive"
func (*Request) Keepalive() (keepalive bool) {
	js.Rewrite("$_.keepalive")
	return keepalive
}

// Method prop
// js:"method"
func (*Request) Method() (method string) {
	js.Rewrite("$_.method")
	return method
}

// Mode prop
// js:"mode"
func (*Request) Mode() (mode *requestmode.RequestMode) {
	js.Rewrite("$_.mode")
	return mode
}

// Redirect prop
// js:"redirect"
func (*Request) Redirect() (redirect *requestredirect.RequestRedirect) {
	js.Rewrite("$_.redirect")
	return redirect
}

// Referrer prop
// js:"referrer"
func (*Request) Referrer() (referrer string) {
	js.Rewrite("$_.referrer")
	return referrer
}

// ReferrerPolicy prop
// js:"referrerPolicy"
func (*Request) ReferrerPolicy() (referrerPolicy *referrerpolicy.ReferrerPolicy) {
	js.Rewrite("$_.referrerPolicy")
	return referrerPolicy
}

// Type prop
// js:"type"
func (*Request) Type() (kind *requesttype.RequestType) {
	js.Rewrite("$_.type")
	return kind
}

// URL prop
// js:"url"
func (*Request) URL() (url string) {
	js.Rewrite("$_.url")
	return url
}

// BodyUsed prop
// js:"bodyUsed"
func (*Request) BodyUsed() (bodyUsed bool) {
	js.Rewrite("$_.bodyUsed")
	return bodyUsed
}
