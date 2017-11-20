package htmlscriptelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLScriptElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLScriptElement) GetAsync() (async bool) {
	js.Rewrite("$<.async")
	return async
}

func (*HTMLScriptElement) SetAsync(async bool) {
	js.Rewrite("$<.async = $1", async)
}

func (*HTMLScriptElement) GetCharset() (charset string) {
	js.Rewrite("$<.charset")
	return charset
}

func (*HTMLScriptElement) SetCharset(charset string) {
	js.Rewrite("$<.charset = $1", charset)
}

func (*HTMLScriptElement) GetCrossOrigin() (crossOrigin string) {
	js.Rewrite("$<.crossOrigin")
	return crossOrigin
}

func (*HTMLScriptElement) SetCrossOrigin(crossOrigin string) {
	js.Rewrite("$<.crossOrigin = $1", crossOrigin)
}

func (*HTMLScriptElement) GetDefer() (dfr bool) {
	js.Rewrite("$<.defer")
	return dfr
}

func (*HTMLScriptElement) SetDefer(dfr bool) {
	js.Rewrite("$<.defer = $1", dfr)
}

func (*HTMLScriptElement) GetEvent() (event string) {
	js.Rewrite("$<.event")
	return event
}

func (*HTMLScriptElement) SetEvent(event string) {
	js.Rewrite("$<.event = $1", event)
}

func (*HTMLScriptElement) GetHTMLFor() (htmlFor string) {
	js.Rewrite("$<.htmlFor")
	return htmlFor
}

func (*HTMLScriptElement) SetHTMLFor(htmlFor string) {
	js.Rewrite("$<.htmlFor = $1", htmlFor)
}

func (*HTMLScriptElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLScriptElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLScriptElement) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*HTMLScriptElement) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

func (*HTMLScriptElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLScriptElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
