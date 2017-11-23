package htmllinkelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLLinkElement struct
// js:"HTMLLinkElement,omit"
type HTMLLinkElement struct {
	window.HTMLElement
}

// Sheet prop
func (*HTMLLinkElement) Sheet() (sheet window.StyleSheet) {
	js.Rewrite("$<.sheet")
	return sheet
}

// Charset prop Sets or retrieves the character set used to encode the object.
func (*HTMLLinkElement) Charset() (charset string) {
	js.Rewrite("$<.charset")
	return charset
}

// Charset prop Sets or retrieves the character set used to encode the object.
func (*HTMLLinkElement) SetCharset(charset string) {
	js.Rewrite("$<.charset = $1", charset)
}

// Disabled prop
func (*HTMLLinkElement) Disabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

// Disabled prop
func (*HTMLLinkElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

// Href prop Sets or retrieves a destination URL or an anchor point.
func (*HTMLLinkElement) Href() (href string) {
	js.Rewrite("$<.href")
	return href
}

// Href prop Sets or retrieves a destination URL or an anchor point.
func (*HTMLLinkElement) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

// Hreflang prop Sets or retrieves the language code of the object.
func (*HTMLLinkElement) Hreflang() (hreflang string) {
	js.Rewrite("$<.hreflang")
	return hreflang
}

// Hreflang prop Sets or retrieves the language code of the object.
func (*HTMLLinkElement) SetHreflang(hreflang string) {
	js.Rewrite("$<.hreflang = $1", hreflang)
}

// Media prop Sets or retrieves the media type.
func (*HTMLLinkElement) Media() (media string) {
	js.Rewrite("$<.media")
	return media
}

// Media prop Sets or retrieves the media type.
func (*HTMLLinkElement) SetMedia(media string) {
	js.Rewrite("$<.media = $1", media)
}

// Rel prop Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLLinkElement) Rel() (rel string) {
	js.Rewrite("$<.rel")
	return rel
}

// Rel prop Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLLinkElement) SetRel(rel string) {
	js.Rewrite("$<.rel = $1", rel)
}

// Rev prop Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLLinkElement) Rev() (rev string) {
	js.Rewrite("$<.rev")
	return rev
}

// Rev prop Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLLinkElement) SetRev(rev string) {
	js.Rewrite("$<.rev = $1", rev)
}

// Target prop Sets or retrieves the window or frame at which to target content.
func (*HTMLLinkElement) Target() (target string) {
	js.Rewrite("$<.target")
	return target
}

// Target prop Sets or retrieves the window or frame at which to target content.
func (*HTMLLinkElement) SetTarget(target string) {
	js.Rewrite("$<.target = $1", target)
}

// Type prop Sets or retrieves the MIME type of the object.
func (*HTMLLinkElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop Sets or retrieves the MIME type of the object.
func (*HTMLLinkElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
