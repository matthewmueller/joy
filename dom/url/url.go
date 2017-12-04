package url

import (
	"github.com/matthewmueller/joy/dom/objecturloptions"
	"github.com/matthewmueller/joy/macro"
)

// New fn
func New(url string, base *string) *URL {
	macro.Rewrite("new URL($1, $2)", url, base)
	return &URL{}
}

// URL struct
// js:"URL,omit"
type URL struct {
}

// CreateObjectURL fn
// js:"createObjectURL"
func (*URL) CreateObjectURL(object interface{}, options *objecturloptions.ObjectURLOptions) (s string) {
	macro.Rewrite("$_.createObjectURL($1, $2)", object, options)
	return s
}

// RevokeObjectURL fn
// js:"revokeObjectURL"
func (*URL) RevokeObjectURL(url string) {
	macro.Rewrite("$_.revokeObjectURL($1)", url)
}

// ToString fn
// js:"toString"
func (*URL) ToString() (s string) {
	macro.Rewrite("$_.toString()")
	return s
}

// Hash prop
// js:"hash"
func (*URL) Hash() (hash string) {
	macro.Rewrite("$_.hash")
	return hash
}

// SetHash prop
// js:"hash"
func (*URL) SetHash(hash string) {
	macro.Rewrite("$_.hash = $1", hash)
}

// Host prop
// js:"host"
func (*URL) Host() (host string) {
	macro.Rewrite("$_.host")
	return host
}

// SetHost prop
// js:"host"
func (*URL) SetHost(host string) {
	macro.Rewrite("$_.host = $1", host)
}

// Hostname prop
// js:"hostname"
func (*URL) Hostname() (hostname string) {
	macro.Rewrite("$_.hostname")
	return hostname
}

// SetHostname prop
// js:"hostname"
func (*URL) SetHostname(hostname string) {
	macro.Rewrite("$_.hostname = $1", hostname)
}

// Href prop
// js:"href"
func (*URL) Href() (href string) {
	macro.Rewrite("$_.href")
	return href
}

// SetHref prop
// js:"href"
func (*URL) SetHref(href string) {
	macro.Rewrite("$_.href = $1", href)
}

// Origin prop
// js:"origin"
func (*URL) Origin() (origin string) {
	macro.Rewrite("$_.origin")
	return origin
}

// Password prop
// js:"password"
func (*URL) Password() (password string) {
	macro.Rewrite("$_.password")
	return password
}

// SetPassword prop
// js:"password"
func (*URL) SetPassword(password string) {
	macro.Rewrite("$_.password = $1", password)
}

// Pathname prop
// js:"pathname"
func (*URL) Pathname() (pathname string) {
	macro.Rewrite("$_.pathname")
	return pathname
}

// SetPathname prop
// js:"pathname"
func (*URL) SetPathname(pathname string) {
	macro.Rewrite("$_.pathname = $1", pathname)
}

// Port prop
// js:"port"
func (*URL) Port() (port string) {
	macro.Rewrite("$_.port")
	return port
}

// SetPort prop
// js:"port"
func (*URL) SetPort(port string) {
	macro.Rewrite("$_.port = $1", port)
}

// Protocol prop
// js:"protocol"
func (*URL) Protocol() (protocol string) {
	macro.Rewrite("$_.protocol")
	return protocol
}

// SetProtocol prop
// js:"protocol"
func (*URL) SetProtocol(protocol string) {
	macro.Rewrite("$_.protocol = $1", protocol)
}

// Search prop
// js:"search"
func (*URL) Search() (search string) {
	macro.Rewrite("$_.search")
	return search
}

// SetSearch prop
// js:"search"
func (*URL) SetSearch(search string) {
	macro.Rewrite("$_.search = $1", search)
}

// Username prop
// js:"username"
func (*URL) Username() (username string) {
	macro.Rewrite("$_.username")
	return username
}

// SetUsername prop
// js:"username"
func (*URL) SetUsername(username string) {
	macro.Rewrite("$_.username = $1", username)
}
