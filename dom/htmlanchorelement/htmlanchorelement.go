package htmlanchorelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLAnchorElement struct
// js:"HTMLAnchorElement,omit"
type HTMLAnchorElement struct {
	window.HTMLElement
}

// ToString fn Returns a string representation of an object.
func (*HTMLAnchorElement) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

// Charset prop Sets or retrieves the character set used to encode the object.
func (*HTMLAnchorElement) Charset() (charset string) {
	js.Rewrite("$<.charset")
	return charset
}

// Charset prop Sets or retrieves the character set used to encode the object.
func (*HTMLAnchorElement) SetCharset(charset string) {
	js.Rewrite("$<.charset = $1", charset)
}

// Coords prop Sets or retrieves the coordinates of the object.
func (*HTMLAnchorElement) Coords() (coords string) {
	js.Rewrite("$<.coords")
	return coords
}

// Coords prop Sets or retrieves the coordinates of the object.
func (*HTMLAnchorElement) SetCoords(coords string) {
	js.Rewrite("$<.coords = $1", coords)
}

// Download prop
func (*HTMLAnchorElement) Download() (download string) {
	js.Rewrite("$<.download")
	return download
}

// Download prop
func (*HTMLAnchorElement) SetDownload(download string) {
	js.Rewrite("$<.download = $1", download)
}

// Hash prop Contains the anchor portion of the URL including the hash sign (#).
func (*HTMLAnchorElement) Hash() (hash string) {
	js.Rewrite("$<.hash")
	return hash
}

// Hash prop Contains the anchor portion of the URL including the hash sign (#).
func (*HTMLAnchorElement) SetHash(hash string) {
	js.Rewrite("$<.hash = $1", hash)
}

// Host prop Contains the hostname and port values of the URL.
func (*HTMLAnchorElement) Host() (host string) {
	js.Rewrite("$<.host")
	return host
}

// Host prop Contains the hostname and port values of the URL.
func (*HTMLAnchorElement) SetHost(host string) {
	js.Rewrite("$<.host = $1", host)
}

// Hostname prop Contains the hostname of a URL.
func (*HTMLAnchorElement) Hostname() (hostname string) {
	js.Rewrite("$<.hostname")
	return hostname
}

// Hostname prop Contains the hostname of a URL.
func (*HTMLAnchorElement) SetHostname(hostname string) {
	js.Rewrite("$<.hostname = $1", hostname)
}

// Href prop Sets or retrieves a destination URL or an anchor point.
func (*HTMLAnchorElement) Href() (href string) {
	js.Rewrite("$<.href")
	return href
}

// Href prop Sets or retrieves a destination URL or an anchor point.
func (*HTMLAnchorElement) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

// Hreflang prop Sets or retrieves the language code of the object.
func (*HTMLAnchorElement) Hreflang() (hreflang string) {
	js.Rewrite("$<.hreflang")
	return hreflang
}

// Hreflang prop Sets or retrieves the language code of the object.
func (*HTMLAnchorElement) SetHreflang(hreflang string) {
	js.Rewrite("$<.hreflang = $1", hreflang)
}

// Methods prop
func (*HTMLAnchorElement) Methods() (Methods string) {
	js.Rewrite("$<.Methods")
	return Methods
}

// Methods prop
func (*HTMLAnchorElement) SetMethods(Methods string) {
	js.Rewrite("$<.Methods = $1", Methods)
}

// MimeType prop
func (*HTMLAnchorElement) MimeType() (mimeType string) {
	js.Rewrite("$<.mimeType")
	return mimeType
}

// Name prop Sets or retrieves the shape of the object.
func (*HTMLAnchorElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop Sets or retrieves the shape of the object.
func (*HTMLAnchorElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// NameProp prop
func (*HTMLAnchorElement) NameProp() (nameProp string) {
	js.Rewrite("$<.nameProp")
	return nameProp
}

// Pathname prop Contains the pathname of the URL.
func (*HTMLAnchorElement) Pathname() (pathname string) {
	js.Rewrite("$<.pathname")
	return pathname
}

// Pathname prop Contains the pathname of the URL.
func (*HTMLAnchorElement) SetPathname(pathname string) {
	js.Rewrite("$<.pathname = $1", pathname)
}

// Port prop Sets or retrieves the port number associated with a URL.
func (*HTMLAnchorElement) Port() (port string) {
	js.Rewrite("$<.port")
	return port
}

// Port prop Sets or retrieves the port number associated with a URL.
func (*HTMLAnchorElement) SetPort(port string) {
	js.Rewrite("$<.port = $1", port)
}

// Protocol prop Contains the protocol of the URL.
func (*HTMLAnchorElement) Protocol() (protocol string) {
	js.Rewrite("$<.protocol")
	return protocol
}

// Protocol prop Contains the protocol of the URL.
func (*HTMLAnchorElement) SetProtocol(protocol string) {
	js.Rewrite("$<.protocol = $1", protocol)
}

// ProtocolLong prop
func (*HTMLAnchorElement) ProtocolLong() (protocolLong string) {
	js.Rewrite("$<.protocolLong")
	return protocolLong
}

// Rel prop Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLAnchorElement) Rel() (rel string) {
	js.Rewrite("$<.rel")
	return rel
}

// Rel prop Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLAnchorElement) SetRel(rel string) {
	js.Rewrite("$<.rel = $1", rel)
}

// Rev prop Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLAnchorElement) Rev() (rev string) {
	js.Rewrite("$<.rev")
	return rev
}

// Rev prop Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLAnchorElement) SetRev(rev string) {
	js.Rewrite("$<.rev = $1", rev)
}

// Search prop Sets or retrieves the substring of the href property that follows the question mark.
func (*HTMLAnchorElement) Search() (search string) {
	js.Rewrite("$<.search")
	return search
}

// Search prop Sets or retrieves the substring of the href property that follows the question mark.
func (*HTMLAnchorElement) SetSearch(search string) {
	js.Rewrite("$<.search = $1", search)
}

// Shape prop Sets or retrieves the shape of the object.
func (*HTMLAnchorElement) Shape() (shape string) {
	js.Rewrite("$<.shape")
	return shape
}

// Shape prop Sets or retrieves the shape of the object.
func (*HTMLAnchorElement) SetShape(shape string) {
	js.Rewrite("$<.shape = $1", shape)
}

// Target prop Sets or retrieves the window or frame at which to target content.
func (*HTMLAnchorElement) Target() (target string) {
	js.Rewrite("$<.target")
	return target
}

// Target prop Sets or retrieves the window or frame at which to target content.
func (*HTMLAnchorElement) SetTarget(target string) {
	js.Rewrite("$<.target = $1", target)
}

// Text prop Retrieves or sets the text of the object as a string.
func (*HTMLAnchorElement) Text() (text string) {
	js.Rewrite("$<.text")
	return text
}

// Text prop Retrieves or sets the text of the object as a string.
func (*HTMLAnchorElement) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

// Type prop
func (*HTMLAnchorElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop
func (*HTMLAnchorElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

// Urn prop
func (*HTMLAnchorElement) Urn() (urn string) {
	js.Rewrite("$<.urn")
	return urn
}

// Urn prop
func (*HTMLAnchorElement) SetUrn(urn string) {
	js.Rewrite("$<.urn = $1", urn)
}
