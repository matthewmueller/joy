package requestinit

import (
	"github.com/matthewmueller/joy/dom/referrerpolicy"
	"github.com/matthewmueller/joy/dom/requestcache"
	"github.com/matthewmueller/joy/dom/requestcredentials"
	"github.com/matthewmueller/joy/dom/requestmode"
	"github.com/matthewmueller/joy/dom/requestredirect"
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
