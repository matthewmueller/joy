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

// Charset Sets or retrieves the character set used to encode the object.
func (*HTMLMetaElement) Charset() (charset string) {
	js.Rewrite("$<.Charset")
	return charset
}

// Charset Sets or retrieves the character set used to encode the object.
func (*HTMLMetaElement) SetCharset(charset string) {
	js.Rewrite("$<.Charset = $1", charset)
}

// Content Gets or sets meta-information to associate with httpEquiv or name.
func (*HTMLMetaElement) Content() (content string) {
	js.Rewrite("$<.Content")
	return content
}

// Content Gets or sets meta-information to associate with httpEquiv or name.
func (*HTMLMetaElement) SetContent(content string) {
	js.Rewrite("$<.Content = $1", content)
}

// HTTPEquiv Gets or sets information used to bind the value of a content attribute of a meta element to an HTTP response header.
func (*HTMLMetaElement) HTTPEquiv() (httpEquiv string) {
	js.Rewrite("$<.HTTPEquiv")
	return httpEquiv
}

// HTTPEquiv Gets or sets information used to bind the value of a content attribute of a meta element to an HTTP response header.
func (*HTMLMetaElement) SetHTTPEquiv(httpEquiv string) {
	js.Rewrite("$<.HTTPEquiv = $1", httpEquiv)
}

// Name Sets or retrieves the value specified in the content attribute of the meta object.
func (*HTMLMetaElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name Sets or retrieves the value specified in the content attribute of the meta object.
func (*HTMLMetaElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// Scheme Sets or retrieves a scheme to be used in interpreting the value of a property specified for the object.
func (*HTMLMetaElement) Scheme() (scheme string) {
	js.Rewrite("$<.Scheme")
	return scheme
}

// Scheme Sets or retrieves a scheme to be used in interpreting the value of a property specified for the object.
func (*HTMLMetaElement) SetScheme(scheme string) {
	js.Rewrite("$<.Scheme = $1", scheme)
}

// URL Sets or retrieves the URL property that will be loaded after the specified time has elapsed.
func (*HTMLMetaElement) URL() (url string) {
	js.Rewrite("$<.URL")
	return url
}

// URL Sets or retrieves the URL property that will be loaded after the specified time has elapsed.
func (*HTMLMetaElement) SetURL(url string) {
	js.Rewrite("$<.URL = $1", url)
}
