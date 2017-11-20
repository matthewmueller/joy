package location

import "github.com/matthewmueller/golly/js"

type Location struct {
}

func (*Location) Assign(url string) {
	js.Rewrite("$<.assign($1)", url)
}

func (*Location) Reload(forcedReload bool) {
	js.Rewrite("$<.reload($1)", forcedReload)
}

func (*Location) Replace(url string) {
	js.Rewrite("$<.replace($1)", url)
}

func (*Location) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*Location) GetHash() (hash string) {
	js.Rewrite("$<.hash")
	return hash
}

func (*Location) SetHash(hash string) {
	js.Rewrite("$<.hash = $1", hash)
}

func (*Location) GetHost() (host string) {
	js.Rewrite("$<.host")
	return host
}

func (*Location) SetHost(host string) {
	js.Rewrite("$<.host = $1", host)
}

func (*Location) GetHostname() (hostname string) {
	js.Rewrite("$<.hostname")
	return hostname
}

func (*Location) SetHostname(hostname string) {
	js.Rewrite("$<.hostname = $1", hostname)
}

func (*Location) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*Location) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

func (*Location) GetOrigin() (origin string) {
	js.Rewrite("$<.origin")
	return origin
}

func (*Location) GetPathname() (pathname string) {
	js.Rewrite("$<.pathname")
	return pathname
}

func (*Location) SetPathname(pathname string) {
	js.Rewrite("$<.pathname = $1", pathname)
}

func (*Location) GetPort() (port string) {
	js.Rewrite("$<.port")
	return port
}

func (*Location) SetPort(port string) {
	js.Rewrite("$<.port = $1", port)
}

func (*Location) GetProtocol() (protocol string) {
	js.Rewrite("$<.protocol")
	return protocol
}

func (*Location) SetProtocol(protocol string) {
	js.Rewrite("$<.protocol = $1", protocol)
}

func (*Location) GetSearch() (search string) {
	js.Rewrite("$<.search")
	return search
}

func (*Location) SetSearch(search string) {
	js.Rewrite("$<.search = $1", search)
}
