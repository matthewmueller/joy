package htmllinkelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLLinkElement,omit"
type HTMLLinkElement struct {
	window.HTMLElement
	linkstyle.LinkStyle
}

// Charset Sets or retrieves the character set used to encode the object.
func (*HTMLLinkElement) Charset() (charset string) {
	js.Rewrite("$<.Charset")
	return charset
}

// Charset Sets or retrieves the character set used to encode the object.
func (*HTMLLinkElement) SetCharset(charset string) {
	js.Rewrite("$<.Charset = $1", charset)
}

// Disabled
func (*HTMLLinkElement) Disabled() (disabled bool) {
	js.Rewrite("$<.Disabled")
	return disabled
}

// Disabled
func (*HTMLLinkElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.Disabled = $1", disabled)
}

// Href Sets or retrieves a destination URL or an anchor point.
func (*HTMLLinkElement) Href() (href string) {
	js.Rewrite("$<.Href")
	return href
}

// Href Sets or retrieves a destination URL or an anchor point.
func (*HTMLLinkElement) SetHref(href string) {
	js.Rewrite("$<.Href = $1", href)
}

// Hreflang Sets or retrieves the language code of the object.
func (*HTMLLinkElement) Hreflang() (hreflang string) {
	js.Rewrite("$<.Hreflang")
	return hreflang
}

// Hreflang Sets or retrieves the language code of the object.
func (*HTMLLinkElement) SetHreflang(hreflang string) {
	js.Rewrite("$<.Hreflang = $1", hreflang)
}

// Media Sets or retrieves the media type.
func (*HTMLLinkElement) Media() (media string) {
	js.Rewrite("$<.Media")
	return media
}

// Media Sets or retrieves the media type.
func (*HTMLLinkElement) SetMedia(media string) {
	js.Rewrite("$<.Media = $1", media)
}

// Rel Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLLinkElement) Rel() (rel string) {
	js.Rewrite("$<.Rel")
	return rel
}

// Rel Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLLinkElement) SetRel(rel string) {
	js.Rewrite("$<.Rel = $1", rel)
}

// Rev Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLLinkElement) Rev() (rev string) {
	js.Rewrite("$<.Rev")
	return rev
}

// Rev Sets or retrieves the relationship between the object and the destination of the link.
func (*HTMLLinkElement) SetRev(rev string) {
	js.Rewrite("$<.Rev = $1", rev)
}

// Target Sets or retrieves the window or frame at which to target content.
func (*HTMLLinkElement) Target() (target string) {
	js.Rewrite("$<.Target")
	return target
}

// Target Sets or retrieves the window or frame at which to target content.
func (*HTMLLinkElement) SetTarget(target string) {
	js.Rewrite("$<.Target = $1", target)
}

// Type Sets or retrieves the MIME type of the object.
func (*HTMLLinkElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type Sets or retrieves the MIME type of the object.
func (*HTMLLinkElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}
