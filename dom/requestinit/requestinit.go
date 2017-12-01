package requestinit

import (
	"github.com/matthewmueller/golly/dom/referrerpolicy"
	"github.com/matthewmueller/golly/dom/requestcache"
	"github.com/matthewmueller/golly/dom/requestcredentials"
	"github.com/matthewmueller/golly/dom/requestmode"
	"github.com/matthewmueller/golly/dom/requestredirect"
)

type RequestInit struct {
	body           *interface{}
	cache          *requestcache.RequestCache
	credentials    *requestcredentials.RequestCredentials
	headers        *interface{}
	integrity      *string
	keepalive      *bool
	method         *string
	mode           *requestmode.RequestMode
	redirect       *requestredirect.RequestRedirect
	referrer       *string
	referrerPolicy *referrerpolicy.ReferrerPolicy
	window         *interface{}
}
