package htmlanchorelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLAnchorElement struct
// js:"HTMLAnchorElement,omit"
type HTMLAnchorElement struct {
	window.HTMLElement
}

// ToString Returns a string representation of an object.
func (*HTMLAnchorElement) ToString() (s string) {
	js.Rewrite("$<.ToString()")
	return s
}

// Charset Sets or retrieves the character set used to encode the object.
func (*HTMLAnchorElement) Charset() (charset string) {
	js.Rewrite("$<.Charset")
	return charset
}

// Charset Sets or retrieves the character set used to encode the object.
func (*HTMLAnchorElement) SetCharset(charset string) {
	js.Rewrite("$<.Charset = $1", charset)
}

// Coords Sets or retrieves the coordinates of the object.
func (*HTMLAnchorElement) Coords() (coords string) {
	js.Rewrite("$<.Coords")
	return coords
}

// Coords Sets or retrieves the coordinates of the object.
func (*HTMLAnchorElement) SetCoords(coords string) {
	js.Rewrite("$<.Coords = $1", coords)
}

// Download
func (*HTMLAnchorElement) Download() (download string) {
	js.Rewrite("$<.Download")
	return download
}

// Download
func (*HTMLAnchorElement) SetDownload(download string) {
	js.Rewrite("$<.Download = $1", download)
}

// Hash Contains the anchor portion of the URL including the hash sign (#).
func (*HTMLAnchorElement) Hash() (hash string) {
	js.Rewrite("$<.Hash")
	return hash
}

// Hash Contains the anchor portion of the URL including the hash sign (#).
func (*HTMLAnchorElement) SetHash(hash string) {
	js.Rewrite("$<.Hash = $1", hash)
}

// Host Contains the hostname and port values of the URL.
func (*HTMLAnchorElement) Host() (host string) {
	js.Rewrite("$<.Host")
	return host
}

// Host Contains the hostname and port values of the URL.
func (*HTMLAnchorElement) SetHost(host string) {
	js.Rewrite("$<.Host = $1", host)
}

// Hostname Contains the hostname of a URL.
func (*HTMLAnchorElement) Hostname() (hostname string) {
	js.Rewrite("$<.Hostname")
	return hostname
}

// Hostname Contains the hostname of a URL.
func (*HTMLAnchorElement) SetHostname(hostname string) {
	js.Rewrite("$<.Hostname = $1", hostname)
}

// Href Sets or retrieves a destination URL or an anchor point.
func (*HTMLAnchorElement) Href() (href string) {
	js.Rewrite("$<.Href")
	return href
}

// Href Sets or retrieves a destination URL or an anchor point.
func (*HTMLAnchorElement) SetHref(href string) {
	js.Rewrite("$<.Href = $1", href)
}

// Hreflang Sets or retrieves the language code of the object.
func (*HTMLAnchorElement) Hreflang() (hreflang string) {
	js.Rewrite("$<.Hreflang")
	return hreflang
}

// Hreflang Sets or retrieves the language code of the object.
func (*HTMLAnchorElement) SetHreflang(hreflang string) {
	js.Rewrite("$<.Hreflang = $1", hreflang)
}

// Methods
func (*HTMLAnchorElement) Methods() (Methods string) {
	js.Rewrite("$<.Methods")
	return Methods
}

// Methods
func (*HTMLAnchorElement) SetMethods(Methods string) {
	js.Rewrite("$<.Methods = $1", Methods)
}

// MimeType
func (*HTMLAnchorElement) MimeType() (mimeType string) {
	js.Rewrite("$<.MimeType")
	return mimeType
}

// Name Sets or retrieves the shape of the object.
func (*HTMLAnchorElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name Sets or retrieves the shape of the object.
func (*HTMLAnchorElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// NameProp
func (*HTMLAnchorElement) NameProp() (nameProp string) {
	js.Rewrite("$<.NameProp")
	return nameProp
}

// Pathname Contains the pathname of the URL.
func (*HTMLAnchorElement) Pathname() (pathname string) {
	js.Rewrite("$<.Pathname")
	return pathname
}

// Pathname Contains the pathname of the URL.
func (*HTMLAnchorElement) SetPathname(pathname string) {
	js.Rewrite("$<.Pathname = $1", pathname)
}

// Port Sets or retrieves the port number associated with a URL.
func (*HTMLAnchorElement) Port() (port string) {
	js.Rewrite("$<.Port")
	return port
}

// Port Sets or retrieves the port number associated with a URL.
func (*HTMLAnchorElement) SetPort(port string) {
	js.Rewrite("$<.Port = $1", port)
}

// Protocol Contains the protocol of the URL.
func (*HTMLAnchorElement) Protocol() (protocol string) {
	js.Rewrite("$<.Protocol")
	return protocol
}

// Protocol Contains the protocol of the URL.
func (*HTMLAnchorElement) SetProtocol(protocol string) {
	js.Rewrite("$<.Protocol = $1", protocol)
}

// ProtocolLong
func (*HTMLAnchorElement) ProtocolLong() (protocolLong string) {
	js.Rewrite("$<.ProtocolLong")
	return protocolLong
}

// Rel Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLAnchorElement) Rel() (rel string) {
	js.Rewrite("$<.Rel")
	return rel
}

// Rel Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLAnchorElement) SetRel(rel string) {
	js.Rewrite("$<.Rel = $1", rel)
}

// Rev Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLAnchorElement) Rev() (rev string) {
	js.Rewrite("$<.Rev")
	return rev
}

// Rev Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLAnchorElement) SetRev(rev string) {
	js.Rewrite("$<.Rev = $1", rev)
}

// Search Sets or retrieves the substring of the href property that follows the question mark.
func (*HTMLAnchorElement) Search() (search string) {
	js.Rewrite("$<.Search")
	return search
}

// Search Sets or retrieves the substring of the href property that follows the question mark.
func (*HTMLAnchorElement) SetSearch(search string) {
	js.Rewrite("$<.Search = $1", search)
}

// Shape Sets or retrieves the shape of the object.
func (*HTMLAnchorElement) Shape() (shape string) {
	js.Rewrite("$<.Shape")
	return shape
}

// Shape Sets or retrieves the shape of the object.
func (*HTMLAnchorElement) SetShape(shape string) {
	js.Rewrite("$<.Shape = $1", shape)
}

// Target Sets or retrieves the window or frame at which to target content.
func (*HTMLAnchorElement) Target() (target string) {
	js.Rewrite("$<.Target")
	return target
}

// Target Sets or retrieves the window or frame at which to target content.
func (*HTMLAnchorElement) SetTarget(target string) {
	js.Rewrite("$<.Target = $1", target)
}

// Text Retrieves or sets the text of the object as a string.
func (*HTMLAnchorElement) Text() (text string) {
	js.Rewrite("$<.Text")
	return text
}

// Text Retrieves or sets the text of the object as a string.
func (*HTMLAnchorElement) SetText(text string) {
	js.Rewrite("$<.Text = $1", text)
}

// Type
func (*HTMLAnchorElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type
func (*HTMLAnchorElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}

// Urn
func (*HTMLAnchorElement) Urn() (urn string) {
	js.Rewrite("$<.Urn")
	return urn
}

// Urn
func (*HTMLAnchorElement) SetUrn(urn string) {
	js.Rewrite("$<.Urn = $1", urn)
}
