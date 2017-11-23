package url

import "github.com/matthewmueller/golly/js"

// js:"URL,omit"
type URL struct {
}

// CreateObjectURL
func (*URL) CreateObjectURL(object interface{}, options *objecturloptions.ObjectURLOptions) (s string) {
	js.Rewrite("$<.CreateObjectURL($1, $2)", object, options)
	return s
}

// RevokeObjectURL
func (*URL) RevokeObjectURL(url string) {
	js.Rewrite("$<.RevokeObjectURL($1)", url)
}

// ToString
func (*URL) ToString() (s string) {
	js.Rewrite("$<.ToString()")
	return s
}

// Hash
func (*URL) Hash() (hash string) {
	js.Rewrite("$<.Hash")
	return hash
}

// Hash
func (*URL) SetHash(hash string) {
	js.Rewrite("$<.Hash = $1", hash)
}

// Host
func (*URL) Host() (host string) {
	js.Rewrite("$<.Host")
	return host
}

// Host
func (*URL) SetHost(host string) {
	js.Rewrite("$<.Host = $1", host)
}

// Hostname
func (*URL) Hostname() (hostname string) {
	js.Rewrite("$<.Hostname")
	return hostname
}

// Hostname
func (*URL) SetHostname(hostname string) {
	js.Rewrite("$<.Hostname = $1", hostname)
}

// Href
func (*URL) Href() (href string) {
	js.Rewrite("$<.Href")
	return href
}

// Href
func (*URL) SetHref(href string) {
	js.Rewrite("$<.Href = $1", href)
}

// Origin
func (*URL) Origin() (origin string) {
	js.Rewrite("$<.Origin")
	return origin
}

// Password
func (*URL) Password() (password string) {
	js.Rewrite("$<.Password")
	return password
}

// Password
func (*URL) SetPassword(password string) {
	js.Rewrite("$<.Password = $1", password)
}

// Pathname
func (*URL) Pathname() (pathname string) {
	js.Rewrite("$<.Pathname")
	return pathname
}

// Pathname
func (*URL) SetPathname(pathname string) {
	js.Rewrite("$<.Pathname = $1", pathname)
}

// Port
func (*URL) Port() (port string) {
	js.Rewrite("$<.Port")
	return port
}

// Port
func (*URL) SetPort(port string) {
	js.Rewrite("$<.Port = $1", port)
}

// Protocol
func (*URL) Protocol() (protocol string) {
	js.Rewrite("$<.Protocol")
	return protocol
}

// Protocol
func (*URL) SetProtocol(protocol string) {
	js.Rewrite("$<.Protocol = $1", protocol)
}

// Search
func (*URL) Search() (search string) {
	js.Rewrite("$<.Search")
	return search
}

// Search
func (*URL) SetSearch(search string) {
	js.Rewrite("$<.Search = $1", search)
}

// Username
func (*URL) Username() (username string) {
	js.Rewrite("$<.Username")
	return username
}

// Username
func (*URL) SetUsername(username string) {
	js.Rewrite("$<.Username = $1", username)
}
