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

// Async
func (*HTMLScriptElement) Async() (async bool) {
	js.Rewrite("$<.Async")
	return async
}

// Async
func (*HTMLScriptElement) SetAsync(async bool) {
	js.Rewrite("$<.Async = $1", async)
}

// Charset Sets or retrieves the character set used to encode the object.
func (*HTMLScriptElement) Charset() (charset string) {
	js.Rewrite("$<.Charset")
	return charset
}

// Charset Sets or retrieves the character set used to encode the object.
func (*HTMLScriptElement) SetCharset(charset string) {
	js.Rewrite("$<.Charset = $1", charset)
}

// CrossOrigin
func (*HTMLScriptElement) CrossOrigin() (crossOrigin string) {
	js.Rewrite("$<.CrossOrigin")
	return crossOrigin
}

// CrossOrigin
func (*HTMLScriptElement) SetCrossOrigin(crossOrigin string) {
	js.Rewrite("$<.CrossOrigin = $1", crossOrigin)
}

// Defer Sets or retrieves the status of the script.
func (*HTMLScriptElement) Defer() (dfr bool) {
	js.Rewrite("$<.Defer")
	return dfr
}

// Defer Sets or retrieves the status of the script.
func (*HTMLScriptElement) SetDefer(dfr bool) {
	js.Rewrite("$<.Defer = $1", dfr)
}

// Event Sets or retrieves the event for which the script is written.
func (*HTMLScriptElement) Event() (event string) {
	js.Rewrite("$<.Event")
	return event
}

// Event Sets or retrieves the event for which the script is written.
func (*HTMLScriptElement) SetEvent(event string) {
	js.Rewrite("$<.Event = $1", event)
}

// HTMLFor Sets or retrieves the object that is bound to the event script.
func (*HTMLScriptElement) HTMLFor() (htmlFor string) {
	js.Rewrite("$<.HTMLFor")
	return htmlFor
}

// HTMLFor Sets or retrieves the object that is bound to the event script.
func (*HTMLScriptElement) SetHTMLFor(htmlFor string) {
	js.Rewrite("$<.HTMLFor = $1", htmlFor)
}

// Src Retrieves the URL to an external file that contains the source code or data.
func (*HTMLScriptElement) Src() (src string) {
	js.Rewrite("$<.Src")
	return src
}

// Src Retrieves the URL to an external file that contains the source code or data.
func (*HTMLScriptElement) SetSrc(src string) {
	js.Rewrite("$<.Src = $1", src)
}

// Text Retrieves or sets the text of the object as a string.
func (*HTMLScriptElement) Text() (text string) {
	js.Rewrite("$<.Text")
	return text
}

// Text Retrieves or sets the text of the object as a string.
func (*HTMLScriptElement) SetText(text string) {
	js.Rewrite("$<.Text = $1", text)
}

// Type Sets or retrieves the MIME type for the associated scripting engine.
func (*HTMLScriptElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type Sets or retrieves the MIME type for the associated scripting engine.
func (*HTMLScriptElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}
