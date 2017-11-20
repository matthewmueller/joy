package htmlmetaelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLMetaElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLMetaElement) GetCharset() (charset string) {
	js.Rewrite("$<.charset")
	return charset
}

func (*HTMLMetaElement) SetCharset(charset string) {
	js.Rewrite("$<.charset = $1", charset)
}

func (*HTMLMetaElement) GetContent() (content string) {
	js.Rewrite("$<.content")
	return content
}

func (*HTMLMetaElement) SetContent(content string) {
	js.Rewrite("$<.content = $1", content)
}

func (*HTMLMetaElement) GetHTTPEquiv() (httpEquiv string) {
	js.Rewrite("$<.httpEquiv")
	return httpEquiv
}

func (*HTMLMetaElement) SetHTTPEquiv(httpEquiv string) {
	js.Rewrite("$<.httpEquiv = $1", httpEquiv)
}

func (*HTMLMetaElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLMetaElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLMetaElement) GetScheme() (scheme string) {
	js.Rewrite("$<.scheme")
	return scheme
}

func (*HTMLMetaElement) SetScheme(scheme string) {
	js.Rewrite("$<.scheme = $1", scheme)
}

func (*HTMLMetaElement) GetURL() (url string) {
	js.Rewrite("$<.url")
	return url
}

func (*HTMLMetaElement) SetURL(url string) {
	js.Rewrite("$<.url = $1", url)
}
