package htmlscriptelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLScriptElement struct
// js:"HTMLScriptElement,omit"
type HTMLScriptElement struct {
	window.HTMLElement
}

// Async prop
func (*HTMLScriptElement) Async() (async bool) {
	js.Rewrite("$<.async")
	return async
}

// Async prop
func (*HTMLScriptElement) SetAsync(async bool) {
	js.Rewrite("$<.async = $1", async)
}

// Charset prop Sets or retrieves the character set used to encode the object.
func (*HTMLScriptElement) Charset() (charset string) {
	js.Rewrite("$<.charset")
	return charset
}

// Charset prop Sets or retrieves the character set used to encode the object.
func (*HTMLScriptElement) SetCharset(charset string) {
	js.Rewrite("$<.charset = $1", charset)
}

// CrossOrigin prop
func (*HTMLScriptElement) CrossOrigin() (crossOrigin string) {
	js.Rewrite("$<.crossOrigin")
	return crossOrigin
}

// CrossOrigin prop
func (*HTMLScriptElement) SetCrossOrigin(crossOrigin string) {
	js.Rewrite("$<.crossOrigin = $1", crossOrigin)
}

// Defer prop Sets or retrieves the status of the script.
func (*HTMLScriptElement) Defer() (dfr bool) {
	js.Rewrite("$<.defer")
	return dfr
}

// Defer prop Sets or retrieves the status of the script.
func (*HTMLScriptElement) SetDefer(dfr bool) {
	js.Rewrite("$<.defer = $1", dfr)
}

// Event prop Sets or retrieves the event for which the script is written.
func (*HTMLScriptElement) Event() (event string) {
	js.Rewrite("$<.event")
	return event
}

// Event prop Sets or retrieves the event for which the script is written.
func (*HTMLScriptElement) SetEvent(event string) {
	js.Rewrite("$<.event = $1", event)
}

// HTMLFor prop Sets or retrieves the object that is bound to the event script.
func (*HTMLScriptElement) HTMLFor() (htmlFor string) {
	js.Rewrite("$<.htmlFor")
	return htmlFor
}

// HTMLFor prop Sets or retrieves the object that is bound to the event script.
func (*HTMLScriptElement) SetHTMLFor(htmlFor string) {
	js.Rewrite("$<.htmlFor = $1", htmlFor)
}

// Src prop Retrieves the URL to an external file that contains the source code or data.
func (*HTMLScriptElement) Src() (src string) {
	js.Rewrite("$<.src")
	return src
}

// Src prop Retrieves the URL to an external file that contains the source code or data.
func (*HTMLScriptElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

// Text prop Retrieves or sets the text of the object as a string.
func (*HTMLScriptElement) Text() (text string) {
	js.Rewrite("$<.text")
	return text
}

// Text prop Retrieves or sets the text of the object as a string.
func (*HTMLScriptElement) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

// Type prop Sets or retrieves the MIME type for the associated scripting engine.
func (*HTMLScriptElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop Sets or retrieves the MIME type for the associated scripting engine.
func (*HTMLScriptElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
