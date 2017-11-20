package url

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/objecturloptions"
	"github.com/matthewmueller/golly/js"
)

type URL struct {
}

func (*URL) CreateObjectURL(object interface{}, options *objecturloptions.ObjectURLOptions) (s string) {
	js.Rewrite("$<.createObjectURL($1, $2)", object, options)
	return s
}

func (*URL) RevokeObjectURL(url string) {
	js.Rewrite("$<.revokeObjectURL($1)", url)
}

func (*URL) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*URL) GetHash() (hash string) {
	js.Rewrite("$<.hash")
	return hash
}

func (*URL) SetHash(hash string) {
	js.Rewrite("$<.hash = $1", hash)
}

func (*URL) GetHost() (host string) {
	js.Rewrite("$<.host")
	return host
}

func (*URL) SetHost(host string) {
	js.Rewrite("$<.host = $1", host)
}

func (*URL) GetHostname() (hostname string) {
	js.Rewrite("$<.hostname")
	return hostname
}

func (*URL) SetHostname(hostname string) {
	js.Rewrite("$<.hostname = $1", hostname)
}

func (*URL) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*URL) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

func (*URL) GetOrigin() (origin string) {
	js.Rewrite("$<.origin")
	return origin
}

func (*URL) GetPassword() (password string) {
	js.Rewrite("$<.password")
	return password
}

func (*URL) SetPassword(password string) {
	js.Rewrite("$<.password = $1", password)
}

func (*URL) GetPathname() (pathname string) {
	js.Rewrite("$<.pathname")
	return pathname
}

func (*URL) SetPathname(pathname string) {
	js.Rewrite("$<.pathname = $1", pathname)
}

func (*URL) GetPort() (port string) {
	js.Rewrite("$<.port")
	return port
}

func (*URL) SetPort(port string) {
	js.Rewrite("$<.port = $1", port)
}

func (*URL) GetProtocol() (protocol string) {
	js.Rewrite("$<.protocol")
	return protocol
}

func (*URL) SetProtocol(protocol string) {
	js.Rewrite("$<.protocol = $1", protocol)
}

func (*URL) GetSearch() (search string) {
	js.Rewrite("$<.search")
	return search
}

func (*URL) SetSearch(search string) {
	js.Rewrite("$<.search = $1", search)
}

func (*URL) GetUsername() (username string) {
	js.Rewrite("$<.username")
	return username
}

func (*URL) SetUsername(username string) {
	js.Rewrite("$<.username = $1", username)
}
