package requestinit

import (
	"github.com/matthewmueller/golly/dom2/requestcache"
	"github.com/matthewmueller/golly/dom2/requestcredentials"
	"github.com/matthewmueller/golly/dom2/requestredirect"
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
