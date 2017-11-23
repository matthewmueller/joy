package htmlmetaelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLMetaElement struct
// js:"HTMLMetaElement,omit"
type HTMLMetaElement struct {
	window.HTMLElement
}

// Charset prop Sets or retrieves the character set used to encode the object.
func (*HTMLMetaElement) Charset() (charset string) {
	js.Rewrite("$<.charset")
	return charset
}

// Charset prop Sets or retrieves the character set used to encode the object.
func (*HTMLMetaElement) SetCharset(charset string) {
	js.Rewrite("$<.charset = $1", charset)
}

// Content prop Gets or sets meta-information to associate with httpEquiv or name.
func (*HTMLMetaElement) Content() (content string) {
	js.Rewrite("$<.content")
	return content
}

// Content prop Gets or sets meta-information to associate with httpEquiv or name.
func (*HTMLMetaElement) SetContent(content string) {
	js.Rewrite("$<.content = $1", content)
}

// HTTPEquiv prop Gets or sets information used to bind the value of a content attribute of a meta element to an HTTP response header.
func (*HTMLMetaElement) HTTPEquiv() (httpEquiv string) {
	js.Rewrite("$<.httpEquiv")
	return httpEquiv
}

// HTTPEquiv prop Gets or sets information used to bind the value of a content attribute of a meta element to an HTTP response header.
func (*HTMLMetaElement) SetHTTPEquiv(httpEquiv string) {
	js.Rewrite("$<.httpEquiv = $1", httpEquiv)
}

// Name prop Sets or retrieves the value specified in the content attribute of the meta object.
func (*HTMLMetaElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop Sets or retrieves the value specified in the content attribute of the meta object.
func (*HTMLMetaElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// Scheme prop Sets or retrieves a scheme to be used in interpreting the value of a property specified for the object.
func (*HTMLMetaElement) Scheme() (scheme string) {
	js.Rewrite("$<.scheme")
	return scheme
}

// Scheme prop Sets or retrieves a scheme to be used in interpreting the value of a property specified for the object.
func (*HTMLMetaElement) SetScheme(scheme string) {
	js.Rewrite("$<.scheme = $1", scheme)
}

// URL prop Sets or retrieves the URL property that will be loaded after the specified time has elapsed.
func (*HTMLMetaElement) URL() (url string) {
	js.Rewrite("$<.url")
	return url
}

// URL prop Sets or retrieves the URL property that will be loaded after the specified time has elapsed.
func (*HTMLMetaElement) SetURL(url string) {
	js.Rewrite("$<.url = $1", url)
}
