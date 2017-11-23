package location

import "github.com/matthewmueller/golly/js"

// Location struct
// js:"Location,omit"
type Location struct {
}

// Assign fn
func (*Location) Assign(url string) {
	js.Rewrite("$<.assign($1)", url)
}

// Reload fn
func (*Location) Reload(forcedReload *bool) {
	js.Rewrite("$<.reload($1)", forcedReload)
}

// Replace fn
func (*Location) Replace(url string) {
	js.Rewrite("$<.replace($1)", url)
}

// ToString fn
func (*Location) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

// Hash prop
func (*Location) Hash() (hash string) {
	js.Rewrite("$<.hash")
	return hash
}

// Hash prop
func (*Location) SetHash(hash string) {
	js.Rewrite("$<.hash = $1", hash)
}

// Host prop
func (*Location) Host() (host string) {
	js.Rewrite("$<.host")
	return host
}

// Host prop
func (*Location) SetHost(host string) {
	js.Rewrite("$<.host = $1", host)
}

// Hostname prop
func (*Location) Hostname() (hostname string) {
	js.Rewrite("$<.hostname")
	return hostname
}

// Hostname prop
func (*Location) SetHostname(hostname string) {
	js.Rewrite("$<.hostname = $1", hostname)
}

// Href prop
func (*Location) Href() (href string) {
	js.Rewrite("$<.href")
	return href
}

// Href prop
func (*Location) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

// Origin prop
func (*Location) Origin() (origin string) {
	js.Rewrite("$<.origin")
	return origin
}

// Pathname prop
func (*Location) Pathname() (pathname string) {
	js.Rewrite("$<.pathname")
	return pathname
}

// Pathname prop
func (*Location) SetPathname(pathname string) {
	js.Rewrite("$<.pathname = $1", pathname)
}

// Port prop
func (*Location) Port() (port string) {
	js.Rewrite("$<.port")
	return port
}

// Port prop
func (*Location) SetPort(port string) {
	js.Rewrite("$<.port = $1", port)
}

// Protocol prop
func (*Location) Protocol() (protocol string) {
	js.Rewrite("$<.protocol")
	return protocol
}

// Protocol prop
func (*Location) SetProtocol(protocol string) {
	js.Rewrite("$<.protocol = $1", protocol)
}

// Search prop
func (*Location) Search() (search string) {
	js.Rewrite("$<.search")
	return search
}

// Search prop
func (*Location) SetSearch(search string) {
	js.Rewrite("$<.search = $1", search)
}
