package htmlareaelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLAreaElement,omit"
type HTMLAreaElement struct {
	window.HTMLElement
}

// ToString Returns a string representation of an object.
func (*HTMLAreaElement) ToString() (s string) {
	js.Rewrite("$<.ToString()")
	return s
}

// Alt Sets or retrieves a text alternative to the graphic.
func (*HTMLAreaElement) Alt() (alt string) {
	js.Rewrite("$<.Alt")
	return alt
}

// Alt Sets or retrieves a text alternative to the graphic.
func (*HTMLAreaElement) SetAlt(alt string) {
	js.Rewrite("$<.Alt = $1", alt)
}

// Coords Sets or retrieves the coordinates of the object.
func (*HTMLAreaElement) Coords() (coords string) {
	js.Rewrite("$<.Coords")
	return coords
}

// Coords Sets or retrieves the coordinates of the object.
func (*HTMLAreaElement) SetCoords(coords string) {
	js.Rewrite("$<.Coords = $1", coords)
}

// Download
func (*HTMLAreaElement) Download() (download string) {
	js.Rewrite("$<.Download")
	return download
}

// Download
func (*HTMLAreaElement) SetDownload(download string) {
	js.Rewrite("$<.Download = $1", download)
}

// Hash Sets or retrieves the subsection of the href property that follows the number sign (#).
func (*HTMLAreaElement) Hash() (hash string) {
	js.Rewrite("$<.Hash")
	return hash
}

// Hash Sets or retrieves the subsection of the href property that follows the number sign (#).
func (*HTMLAreaElement) SetHash(hash string) {
	js.Rewrite("$<.Hash = $1", hash)
}

// Host Sets or retrieves the hostname and port number of the location or URL.
func (*HTMLAreaElement) Host() (host string) {
	js.Rewrite("$<.Host")
	return host
}

// Host Sets or retrieves the hostname and port number of the location or URL.
func (*HTMLAreaElement) SetHost(host string) {
	js.Rewrite("$<.Host = $1", host)
}

// Hostname Sets or retrieves the host name part of the location or URL.
func (*HTMLAreaElement) Hostname() (hostname string) {
	js.Rewrite("$<.Hostname")
	return hostname
}

// Hostname Sets or retrieves the host name part of the location or URL.
func (*HTMLAreaElement) SetHostname(hostname string) {
	js.Rewrite("$<.Hostname = $1", hostname)
}

// Href Sets or retrieves a destination URL or an anchor point.
func (*HTMLAreaElement) Href() (href string) {
	js.Rewrite("$<.Href")
	return href
}

// Href Sets or retrieves a destination URL or an anchor point.
func (*HTMLAreaElement) SetHref(href string) {
	js.Rewrite("$<.Href = $1", href)
}

// NoHref Sets or gets whether clicks in this region cause action.
func (*HTMLAreaElement) NoHref() (noHref bool) {
	js.Rewrite("$<.NoHref")
	return noHref
}

// NoHref Sets or gets whether clicks in this region cause action.
func (*HTMLAreaElement) SetNoHref(noHref bool) {
	js.Rewrite("$<.NoHref = $1", noHref)
}

// Pathname Sets or retrieves the file name or path specified by the object.
func (*HTMLAreaElement) Pathname() (pathname string) {
	js.Rewrite("$<.Pathname")
	return pathname
}

// Pathname Sets or retrieves the file name or path specified by the object.
func (*HTMLAreaElement) SetPathname(pathname string) {
	js.Rewrite("$<.Pathname = $1", pathname)
}

// Port Sets or retrieves the port number associated with a URL.
func (*HTMLAreaElement) Port() (port string) {
	js.Rewrite("$<.Port")
	return port
}

// Port Sets or retrieves the port number associated with a URL.
func (*HTMLAreaElement) SetPort(port string) {
	js.Rewrite("$<.Port = $1", port)
}

// Protocol Sets or retrieves the protocol portion of a URL.
func (*HTMLAreaElement) Protocol() (protocol string) {
	js.Rewrite("$<.Protocol")
	return protocol
}

// Protocol Sets or retrieves the protocol portion of a URL.
func (*HTMLAreaElement) SetProtocol(protocol string) {
	js.Rewrite("$<.Protocol = $1", protocol)
}

// Rel
func (*HTMLAreaElement) Rel() (rel string) {
	js.Rewrite("$<.Rel")
	return rel
}

// Rel
func (*HTMLAreaElement) SetRel(rel string) {
	js.Rewrite("$<.Rel = $1", rel)
}

// Search Sets or retrieves the substring of the href property that follows the question mark.
func (*HTMLAreaElement) Search() (search string) {
	js.Rewrite("$<.Search")
	return search
}

// Search Sets or retrieves the substring of the href property that follows the question mark.
func (*HTMLAreaElement) SetSearch(search string) {
	js.Rewrite("$<.Search = $1", search)
}

// Shape Sets or retrieves the shape of the object.
func (*HTMLAreaElement) Shape() (shape string) {
	js.Rewrite("$<.Shape")
	return shape
}

// Shape Sets or retrieves the shape of the object.
func (*HTMLAreaElement) SetShape(shape string) {
	js.Rewrite("$<.Shape = $1", shape)
}

// Target Sets or retrieves the window or frame at which to target content.
func (*HTMLAreaElement) Target() (target string) {
	js.Rewrite("$<.Target")
	return target
}

// Target Sets or retrieves the window or frame at which to target content.
func (*HTMLAreaElement) SetTarget(target string) {
	js.Rewrite("$<.Target = $1", target)
}
