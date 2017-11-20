package request

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/body"
	"github.com/matthewmueller/golly/internal/gendom/dom/headers"
	"github.com/matthewmueller/golly/internal/gendom/dom/referrerpolicy"
	"github.com/matthewmueller/golly/internal/gendom/dom/requestcache"
	"github.com/matthewmueller/golly/internal/gendom/dom/requestcredentials"
	"github.com/matthewmueller/golly/internal/gendom/dom/requestdestination"
	"github.com/matthewmueller/golly/internal/gendom/dom/requestmode"
	"github.com/matthewmueller/golly/internal/gendom/dom/requestredirect"
	"github.com/matthewmueller/golly/internal/gendom/dom/requesttype"
	"github.com/matthewmueller/golly/js"
)

type Request struct {
	*body.Body
}

func (*Request) Clone() (r *Request) {
	js.Rewrite("$<.clone()")
	return r
}

func (*Request) GetCache() (cache *requestcache.RequestCache) {
	js.Rewrite("$<.cache")
	return cache
}

func (*Request) GetCredentials() (credentials *requestcredentials.RequestCredentials) {
	js.Rewrite("$<.credentials")
	return credentials
}

func (*Request) GetDestination() (destination *requestdestination.RequestDestination) {
	js.Rewrite("$<.destination")
	return destination
}

func (*Request) GetHeaders() (headers *headers.Headers) {
	js.Rewrite("$<.headers")
	return headers
}

func (*Request) GetIntegrity() (integrity string) {
	js.Rewrite("$<.integrity")
	return integrity
}

func (*Request) GetKeepalive() (keepalive bool) {
	js.Rewrite("$<.keepalive")
	return keepalive
}

func (*Request) GetMethod() (method string) {
	js.Rewrite("$<.method")
	return method
}

func (*Request) GetMode() (mode *requestmode.RequestMode) {
	js.Rewrite("$<.mode")
	return mode
}

func (*Request) GetRedirect() (redirect *requestredirect.RequestRedirect) {
	js.Rewrite("$<.redirect")
	return redirect
}

func (*Request) GetReferrer() (referrer string) {
	js.Rewrite("$<.referrer")
	return referrer
}

func (*Request) GetReferrerPolicy() (referrerPolicy *referrerpolicy.ReferrerPolicy) {
	js.Rewrite("$<.referrerPolicy")
	return referrerPolicy
}

func (*Request) GetType() (kind *requesttype.RequestType) {
	js.Rewrite("$<.type")
	return kind
}

func (*Request) GetURL() (url string) {
	js.Rewrite("$<.url")
	return url
}
