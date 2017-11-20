package htmlareaelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLAreaElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLAreaElement) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*HTMLAreaElement) GetAlt() (alt string) {
	js.Rewrite("$<.alt")
	return alt
}

func (*HTMLAreaElement) SetAlt(alt string) {
	js.Rewrite("$<.alt = $1", alt)
}

func (*HTMLAreaElement) GetCoords() (coords string) {
	js.Rewrite("$<.coords")
	return coords
}

func (*HTMLAreaElement) SetCoords(coords string) {
	js.Rewrite("$<.coords = $1", coords)
}

func (*HTMLAreaElement) GetDownload() (download string) {
	js.Rewrite("$<.download")
	return download
}

func (*HTMLAreaElement) SetDownload(download string) {
	js.Rewrite("$<.download = $1", download)
}

func (*HTMLAreaElement) GetHash() (hash string) {
	js.Rewrite("$<.hash")
	return hash
}

func (*HTMLAreaElement) SetHash(hash string) {
	js.Rewrite("$<.hash = $1", hash)
}

func (*HTMLAreaElement) GetHost() (host string) {
	js.Rewrite("$<.host")
	return host
}

func (*HTMLAreaElement) SetHost(host string) {
	js.Rewrite("$<.host = $1", host)
}

func (*HTMLAreaElement) GetHostname() (hostname string) {
	js.Rewrite("$<.hostname")
	return hostname
}

func (*HTMLAreaElement) SetHostname(hostname string) {
	js.Rewrite("$<.hostname = $1", hostname)
}

func (*HTMLAreaElement) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*HTMLAreaElement) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

func (*HTMLAreaElement) GetNoHref() (noHref bool) {
	js.Rewrite("$<.noHref")
	return noHref
}

func (*HTMLAreaElement) SetNoHref(noHref bool) {
	js.Rewrite("$<.noHref = $1", noHref)
}

func (*HTMLAreaElement) GetPathname() (pathname string) {
	js.Rewrite("$<.pathname")
	return pathname
}

func (*HTMLAreaElement) SetPathname(pathname string) {
	js.Rewrite("$<.pathname = $1", pathname)
}

func (*HTMLAreaElement) GetPort() (port string) {
	js.Rewrite("$<.port")
	return port
}

func (*HTMLAreaElement) SetPort(port string) {
	js.Rewrite("$<.port = $1", port)
}

func (*HTMLAreaElement) GetProtocol() (protocol string) {
	js.Rewrite("$<.protocol")
	return protocol
}

func (*HTMLAreaElement) SetProtocol(protocol string) {
	js.Rewrite("$<.protocol = $1", protocol)
}

func (*HTMLAreaElement) GetRel() (rel string) {
	js.Rewrite("$<.rel")
	return rel
}

func (*HTMLAreaElement) SetRel(rel string) {
	js.Rewrite("$<.rel = $1", rel)
}

func (*HTMLAreaElement) GetSearch() (search string) {
	js.Rewrite("$<.search")
	return search
}

func (*HTMLAreaElement) SetSearch(search string) {
	js.Rewrite("$<.search = $1", search)
}

func (*HTMLAreaElement) GetShape() (shape string) {
	js.Rewrite("$<.shape")
	return shape
}

func (*HTMLAreaElement) SetShape(shape string) {
	js.Rewrite("$<.shape = $1", shape)
}

func (*HTMLAreaElement) GetTarget() (target string) {
	js.Rewrite("$<.target")
	return target
}

func (*HTMLAreaElement) SetTarget(target string) {
	js.Rewrite("$<.target = $1", target)
}
