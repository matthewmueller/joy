package request

import (
	"github.com/matthewmueller/golly/dom2/headers"
	"github.com/matthewmueller/golly/dom2/requestredirect"
	"github.com/matthewmueller/golly/js"
	"github.com/matthewmueller/golly/vdom/h/body"
)

// js:"Request,omit"
type Request struct {
	body.Body
}

// Clone
func (*Request) Clone() (r *Request) {
	js.Rewrite("$<.Clone()")
	return r
}

// Cache
func (*Request) Cache() (cache *requestcache.RequestCache) {
	js.Rewrite("$<.Cache")
	return cache
}

// Credentials
func (*Request) Credentials() (credentials *requestcredentials.RequestCredentials) {
	js.Rewrite("$<.Credentials")
	return credentials
}

// Destination
func (*Request) Destination() (destination *requestdestination.RequestDestination) {
	js.Rewrite("$<.Destination")
	return destination
}

// Headers
func (*Request) Headers() (headers *headers.Headers) {
	js.Rewrite("$<.Headers")
	return headers
}

// Integrity
func (*Request) Integrity() (integrity string) {
	js.Rewrite("$<.Integrity")
	return integrity
}

// Keepalive
func (*Request) Keepalive() (keepalive bool) {
	js.Rewrite("$<.Keepalive")
	return keepalive
}

// Method
func (*Request) Method() (method string) {
	js.Rewrite("$<.Method")
	return method
}

// Mode
func (*Request) Mode() (mode *requestmode.RequestMode) {
	js.Rewrite("$<.Mode")
	return mode
}

// Redirect
func (*Request) Redirect() (redirect *requestredirect.RequestRedirect) {
	js.Rewrite("$<.Redirect")
	return redirect
}

// Referrer
func (*Request) Referrer() (referrer string) {
	js.Rewrite("$<.Referrer")
	return referrer
}

// ReferrerPolicy
func (*Request) ReferrerPolicy() (referrerPolicy *referrerpolicy.ReferrerPolicy) {
	js.Rewrite("$<.ReferrerPolicy")
	return referrerPolicy
}

// Type
func (*Request) Type() (kind *requesttype.RequestType) {
	js.Rewrite("$<.Type")
	return kind
}

// URL
func (*Request) URL() (url string) {
	js.Rewrite("$<.URL")
	return url
}
