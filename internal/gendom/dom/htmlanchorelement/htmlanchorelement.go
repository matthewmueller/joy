package htmlanchorelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLAnchorElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLAnchorElement) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*HTMLAnchorElement) GetCharset() (charset string) {
	js.Rewrite("$<.charset")
	return charset
}

func (*HTMLAnchorElement) SetCharset(charset string) {
	js.Rewrite("$<.charset = $1", charset)
}

func (*HTMLAnchorElement) GetCoords() (coords string) {
	js.Rewrite("$<.coords")
	return coords
}

func (*HTMLAnchorElement) SetCoords(coords string) {
	js.Rewrite("$<.coords = $1", coords)
}

func (*HTMLAnchorElement) GetDownload() (download string) {
	js.Rewrite("$<.download")
	return download
}

func (*HTMLAnchorElement) SetDownload(download string) {
	js.Rewrite("$<.download = $1", download)
}

func (*HTMLAnchorElement) GetHash() (hash string) {
	js.Rewrite("$<.hash")
	return hash
}

func (*HTMLAnchorElement) SetHash(hash string) {
	js.Rewrite("$<.hash = $1", hash)
}

func (*HTMLAnchorElement) GetHost() (host string) {
	js.Rewrite("$<.host")
	return host
}

func (*HTMLAnchorElement) SetHost(host string) {
	js.Rewrite("$<.host = $1", host)
}

func (*HTMLAnchorElement) GetHostname() (hostname string) {
	js.Rewrite("$<.hostname")
	return hostname
}

func (*HTMLAnchorElement) SetHostname(hostname string) {
	js.Rewrite("$<.hostname = $1", hostname)
}

func (*HTMLAnchorElement) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*HTMLAnchorElement) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

func (*HTMLAnchorElement) GetHreflang() (hreflang string) {
	js.Rewrite("$<.hreflang")
	return hreflang
}

func (*HTMLAnchorElement) SetHreflang(hreflang string) {
	js.Rewrite("$<.hreflang = $1", hreflang)
}

func (*HTMLAnchorElement) GetMethods() (Methods string) {
	js.Rewrite("$<.Methods")
	return Methods
}

func (*HTMLAnchorElement) SetMethods(Methods string) {
	js.Rewrite("$<.Methods = $1", Methods)
}

func (*HTMLAnchorElement) GetMimeType() (mimeType string) {
	js.Rewrite("$<.mimeType")
	return mimeType
}

func (*HTMLAnchorElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLAnchorElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLAnchorElement) GetNameProp() (nameProp string) {
	js.Rewrite("$<.nameProp")
	return nameProp
}

func (*HTMLAnchorElement) GetPathname() (pathname string) {
	js.Rewrite("$<.pathname")
	return pathname
}

func (*HTMLAnchorElement) SetPathname(pathname string) {
	js.Rewrite("$<.pathname = $1", pathname)
}

func (*HTMLAnchorElement) GetPort() (port string) {
	js.Rewrite("$<.port")
	return port
}

func (*HTMLAnchorElement) SetPort(port string) {
	js.Rewrite("$<.port = $1", port)
}

func (*HTMLAnchorElement) GetProtocol() (protocol string) {
	js.Rewrite("$<.protocol")
	return protocol
}

func (*HTMLAnchorElement) SetProtocol(protocol string) {
	js.Rewrite("$<.protocol = $1", protocol)
}

func (*HTMLAnchorElement) GetProtocolLong() (protocolLong string) {
	js.Rewrite("$<.protocolLong")
	return protocolLong
}

func (*HTMLAnchorElement) GetRel() (rel string) {
	js.Rewrite("$<.rel")
	return rel
}

func (*HTMLAnchorElement) SetRel(rel string) {
	js.Rewrite("$<.rel = $1", rel)
}

func (*HTMLAnchorElement) GetRev() (rev string) {
	js.Rewrite("$<.rev")
	return rev
}

func (*HTMLAnchorElement) SetRev(rev string) {
	js.Rewrite("$<.rev = $1", rev)
}

func (*HTMLAnchorElement) GetSearch() (search string) {
	js.Rewrite("$<.search")
	return search
}

func (*HTMLAnchorElement) SetSearch(search string) {
	js.Rewrite("$<.search = $1", search)
}

func (*HTMLAnchorElement) GetShape() (shape string) {
	js.Rewrite("$<.shape")
	return shape
}

func (*HTMLAnchorElement) SetShape(shape string) {
	js.Rewrite("$<.shape = $1", shape)
}

func (*HTMLAnchorElement) GetTarget() (target string) {
	js.Rewrite("$<.target")
	return target
}

func (*HTMLAnchorElement) SetTarget(target string) {
	js.Rewrite("$<.target = $1", target)
}

func (*HTMLAnchorElement) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*HTMLAnchorElement) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

func (*HTMLAnchorElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLAnchorElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

func (*HTMLAnchorElement) GetUrn() (urn string) {
	js.Rewrite("$<.urn")
	return urn
}

func (*HTMLAnchorElement) SetUrn(urn string) {
	js.Rewrite("$<.urn = $1", urn)
}
