package location

import "github.com/matthewmueller/golly/js"

// js:"Location,omit"
type Location struct {
}

// Assign
func (*Location) Assign(url string) {
	js.Rewrite("$<.Assign($1)", url)
}

// Reload
func (*Location) Reload(forcedReload *bool) {
	js.Rewrite("$<.Reload($1)", forcedReload)
}

// Replace
func (*Location) Replace(url string) {
	js.Rewrite("$<.Replace($1)", url)
}

// ToString
func (*Location) ToString() (s string) {
	js.Rewrite("$<.ToString()")
	return s
}

// Hash
func (*Location) Hash() (hash string) {
	js.Rewrite("$<.Hash")
	return hash
}

// Hash
func (*Location) SetHash(hash string) {
	js.Rewrite("$<.Hash = $1", hash)
}

// Host
func (*Location) Host() (host string) {
	js.Rewrite("$<.Host")
	return host
}

// Host
func (*Location) SetHost(host string) {
	js.Rewrite("$<.Host = $1", host)
}

// Hostname
func (*Location) Hostname() (hostname string) {
	js.Rewrite("$<.Hostname")
	return hostname
}

// Hostname
func (*Location) SetHostname(hostname string) {
	js.Rewrite("$<.Hostname = $1", hostname)
}

// Href
func (*Location) Href() (href string) {
	js.Rewrite("$<.Href")
	return href
}

// Href
func (*Location) SetHref(href string) {
	js.Rewrite("$<.Href = $1", href)
}

// Origin
func (*Location) Origin() (origin string) {
	js.Rewrite("$<.Origin")
	return origin
}

// Pathname
func (*Location) Pathname() (pathname string) {
	js.Rewrite("$<.Pathname")
	return pathname
}

// Pathname
func (*Location) SetPathname(pathname string) {
	js.Rewrite("$<.Pathname = $1", pathname)
}

// Port
func (*Location) Port() (port string) {
	js.Rewrite("$<.Port")
	return port
}

// Port
func (*Location) SetPort(port string) {
	js.Rewrite("$<.Port = $1", port)
}

// Protocol
func (*Location) Protocol() (protocol string) {
	js.Rewrite("$<.Protocol")
	return protocol
}

// Protocol
func (*Location) SetProtocol(protocol string) {
	js.Rewrite("$<.Protocol = $1", protocol)
}

// Search
func (*Location) Search() (search string) {
	js.Rewrite("$<.Search")
	return search
}

// Search
func (*Location) SetSearch(search string) {
	js.Rewrite("$<.Search = $1", search)
}
