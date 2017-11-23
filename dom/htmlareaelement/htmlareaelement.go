package htmlareaelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLAreaElement struct
// js:"HTMLAreaElement,omit"
type HTMLAreaElement struct {
	window.HTMLElement
}

// ToString fn Returns a string representation of an object.
func (*HTMLAreaElement) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

// Alt prop Sets or retrieves a text alternative to the graphic.
func (*HTMLAreaElement) Alt() (alt string) {
	js.Rewrite("$<.alt")
	return alt
}

// Alt prop Sets or retrieves a text alternative to the graphic.
func (*HTMLAreaElement) SetAlt(alt string) {
	js.Rewrite("$<.alt = $1", alt)
}

// Coords prop Sets or retrieves the coordinates of the object.
func (*HTMLAreaElement) Coords() (coords string) {
	js.Rewrite("$<.coords")
	return coords
}

// Coords prop Sets or retrieves the coordinates of the object.
func (*HTMLAreaElement) SetCoords(coords string) {
	js.Rewrite("$<.coords = $1", coords)
}

// Download prop
func (*HTMLAreaElement) Download() (download string) {
	js.Rewrite("$<.download")
	return download
}

// Download prop
func (*HTMLAreaElement) SetDownload(download string) {
	js.Rewrite("$<.download = $1", download)
}

// Hash prop Sets or retrieves the subsection of the href property that follows the number sign (#).
func (*HTMLAreaElement) Hash() (hash string) {
	js.Rewrite("$<.hash")
	return hash
}

// Hash prop Sets or retrieves the subsection of the href property that follows the number sign (#).
func (*HTMLAreaElement) SetHash(hash string) {
	js.Rewrite("$<.hash = $1", hash)
}

// Host prop Sets or retrieves the hostname and port number of the location or URL.
func (*HTMLAreaElement) Host() (host string) {
	js.Rewrite("$<.host")
	return host
}

// Host prop Sets or retrieves the hostname and port number of the location or URL.
func (*HTMLAreaElement) SetHost(host string) {
	js.Rewrite("$<.host = $1", host)
}

// Hostname prop Sets or retrieves the host name part of the location or URL.
func (*HTMLAreaElement) Hostname() (hostname string) {
	js.Rewrite("$<.hostname")
	return hostname
}

// Hostname prop Sets or retrieves the host name part of the location or URL.
func (*HTMLAreaElement) SetHostname(hostname string) {
	js.Rewrite("$<.hostname = $1", hostname)
}

// Href prop Sets or retrieves a destination URL or an anchor point.
func (*HTMLAreaElement) Href() (href string) {
	js.Rewrite("$<.href")
	return href
}

// Href prop Sets or retrieves a destination URL or an anchor point.
func (*HTMLAreaElement) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

// NoHref prop Sets or gets whether clicks in this region cause action.
func (*HTMLAreaElement) NoHref() (noHref bool) {
	js.Rewrite("$<.noHref")
	return noHref
}

// NoHref prop Sets or gets whether clicks in this region cause action.
func (*HTMLAreaElement) SetNoHref(noHref bool) {
	js.Rewrite("$<.noHref = $1", noHref)
}

// Pathname prop Sets or retrieves the file name or path specified by the object.
func (*HTMLAreaElement) Pathname() (pathname string) {
	js.Rewrite("$<.pathname")
	return pathname
}

// Pathname prop Sets or retrieves the file name or path specified by the object.
func (*HTMLAreaElement) SetPathname(pathname string) {
	js.Rewrite("$<.pathname = $1", pathname)
}

// Port prop Sets or retrieves the port number associated with a URL.
func (*HTMLAreaElement) Port() (port string) {
	js.Rewrite("$<.port")
	return port
}

// Port prop Sets or retrieves the port number associated with a URL.
func (*HTMLAreaElement) SetPort(port string) {
	js.Rewrite("$<.port = $1", port)
}

// Protocol prop Sets or retrieves the protocol portion of a URL.
func (*HTMLAreaElement) Protocol() (protocol string) {
	js.Rewrite("$<.protocol")
	return protocol
}

// Protocol prop Sets or retrieves the protocol portion of a URL.
func (*HTMLAreaElement) SetProtocol(protocol string) {
	js.Rewrite("$<.protocol = $1", protocol)
}

// Rel prop
func (*HTMLAreaElement) Rel() (rel string) {
	js.Rewrite("$<.rel")
	return rel
}

// Rel prop
func (*HTMLAreaElement) SetRel(rel string) {
	js.Rewrite("$<.rel = $1", rel)
}

// Search prop Sets or retrieves the substring of the href property that follows the question mark.
func (*HTMLAreaElement) Search() (search string) {
	js.Rewrite("$<.search")
	return search
}

// Search prop Sets or retrieves the substring of the href property that follows the question mark.
func (*HTMLAreaElement) SetSearch(search string) {
	js.Rewrite("$<.search = $1", search)
}

// Shape prop Sets or retrieves the shape of the object.
func (*HTMLAreaElement) Shape() (shape string) {
	js.Rewrite("$<.shape")
	return shape
}

// Shape prop Sets or retrieves the shape of the object.
func (*HTMLAreaElement) SetShape(shape string) {
	js.Rewrite("$<.shape = $1", shape)
}

// Target prop Sets or retrieves the window or frame at which to target content.
func (*HTMLAreaElement) Target() (target string) {
	js.Rewrite("$<.target")
	return target
}

// Target prop Sets or retrieves the window or frame at which to target content.
func (*HTMLAreaElement) SetTarget(target string) {
	js.Rewrite("$<.target = $1", target)
}
