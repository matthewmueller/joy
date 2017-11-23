package url

import (
	"github.com/matthewmueller/golly/dom2/objecturloptions"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(url string, base *string) *URL {
	js.Rewrite("URL")
	return &URL{}
}

// URL struct
// js:"URL,omit"
type URL struct {
}

// CreateObjectURL fn
func (*URL) CreateObjectURL(object interface{}, options *objecturloptions.ObjectURLOptions) (s string) {
	js.Rewrite("$<.createObjectURL($1, $2)", object, options)
	return s
}

// RevokeObjectURL fn
func (*URL) RevokeObjectURL(url string) {
	js.Rewrite("$<.revokeObjectURL($1)", url)
}

// ToString fn
func (*URL) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

// Hash prop
func (*URL) Hash() (hash string) {
	js.Rewrite("$<.hash")
	return hash
}

// Hash prop
func (*URL) SetHash(hash string) {
	js.Rewrite("$<.hash = $1", hash)
}

// Host prop
func (*URL) Host() (host string) {
	js.Rewrite("$<.host")
	return host
}

// Host prop
func (*URL) SetHost(host string) {
	js.Rewrite("$<.host = $1", host)
}

// Hostname prop
func (*URL) Hostname() (hostname string) {
	js.Rewrite("$<.hostname")
	return hostname
}

// Hostname prop
func (*URL) SetHostname(hostname string) {
	js.Rewrite("$<.hostname = $1", hostname)
}

// Href prop
func (*URL) Href() (href string) {
	js.Rewrite("$<.href")
	return href
}

// Href prop
func (*URL) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

// Origin prop
func (*URL) Origin() (origin string) {
	js.Rewrite("$<.origin")
	return origin
}

// Password prop
func (*URL) Password() (password string) {
	js.Rewrite("$<.password")
	return password
}

// Password prop
func (*URL) SetPassword(password string) {
	js.Rewrite("$<.password = $1", password)
}

// Pathname prop
func (*URL) Pathname() (pathname string) {
	js.Rewrite("$<.pathname")
	return pathname
}

// Pathname prop
func (*URL) SetPathname(pathname string) {
	js.Rewrite("$<.pathname = $1", pathname)
}

// Port prop
func (*URL) Port() (port string) {
	js.Rewrite("$<.port")
	return port
}

// Port prop
func (*URL) SetPort(port string) {
	js.Rewrite("$<.port = $1", port)
}

// Protocol prop
func (*URL) Protocol() (protocol string) {
	js.Rewrite("$<.protocol")
	return protocol
}

// Protocol prop
func (*URL) SetProtocol(protocol string) {
	js.Rewrite("$<.protocol = $1", protocol)
}

// Search prop
func (*URL) Search() (search string) {
	js.Rewrite("$<.search")
	return search
}

// Search prop
func (*URL) SetSearch(search string) {
	js.Rewrite("$<.search = $1", search)
}

// Username prop
func (*URL) Username() (username string) {
	js.Rewrite("$<.username")
	return username
}

// Username prop
func (*URL) SetUsername(username string) {
	js.Rewrite("$<.username = $1", username)
}
