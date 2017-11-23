package mswebviewelement

import (
	"github.com/matthewmueller/golly/dom2/deferredpermissionrequest"
	"github.com/matthewmueller/golly/dom2/focusnavigationorigin"
	"github.com/matthewmueller/golly/dom2/mswebviewsettings"
	"github.com/matthewmueller/golly/dom2/navigationreason"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"MSHTMLWebViewElement,omit"
type MSHTMLWebViewElement struct {
	window.HTMLElement
}

// AddWebAllowedObject
func (*MSHTMLWebViewElement) AddWebAllowedObject(name string, applicationObject interface{}) {
	js.Rewrite("$<.AddWebAllowedObject($1, $2)", name, applicationObject)
}

// BuildLocalStreamURI
func (*MSHTMLWebViewElement) BuildLocalStreamURI(contentIdentifier string, relativePath string) (s string) {
	js.Rewrite("$<.BuildLocalStreamURI($1, $2)", contentIdentifier, relativePath)
	return s
}

// CapturePreviewToBlobAsync
func (*MSHTMLWebViewElement) CapturePreviewToBlobAsync() (m *MSWebViewAsyncOperation) {
	js.Rewrite("$<.CapturePreviewToBlobAsync()")
	return m
}

// CaptureSelectedContentToDataPackageAsync
func (*MSHTMLWebViewElement) CaptureSelectedContentToDataPackageAsync() (m *MSWebViewAsyncOperation) {
	js.Rewrite("$<.CaptureSelectedContentToDataPackageAsync()")
	return m
}

// GetDeferredPermissionRequestByID
func (*MSHTMLWebViewElement) GetDeferredPermissionRequestByID(id uint) (d deferredpermissionrequest.DeferredPermissionRequest) {
	js.Rewrite("$<.GetDeferredPermissionRequestByID($1)", id)
	return d
}

// GetDeferredPermissionRequests
func (*MSHTMLWebViewElement) GetDeferredPermissionRequests() (d []deferredpermissionrequest.DeferredPermissionRequest) {
	js.Rewrite("$<.GetDeferredPermissionRequests()")
	return d
}

// GoBack
func (*MSHTMLWebViewElement) GoBack() {
	js.Rewrite("$<.GoBack()")
}

// GoForward
func (*MSHTMLWebViewElement) GoForward() {
	js.Rewrite("$<.GoForward()")
}

// InvokeScriptAsync
func (*MSHTMLWebViewElement) InvokeScriptAsync(scriptName string, args interface{}) (m *MSWebViewAsyncOperation) {
	js.Rewrite("$<.InvokeScriptAsync($1, $2)", scriptName, args)
	return m
}

// Navigate
func (*MSHTMLWebViewElement) Navigate(uri string) {
	js.Rewrite("$<.Navigate($1)", uri)
}

// NavigateFocus
func (*MSHTMLWebViewElement) NavigateFocus(navigationReason *navigationreason.NavigationReason, origin *focusnavigationorigin.FocusNavigationOrigin) {
	js.Rewrite("$<.NavigateFocus($1, $2)", navigationReason, origin)
}

// NavigateToLocalStreamURI
func (*MSHTMLWebViewElement) NavigateToLocalStreamURI(source string, streamResolver interface{}) {
	js.Rewrite("$<.NavigateToLocalStreamURI($1, $2)", source, streamResolver)
}

// NavigateToString
func (*MSHTMLWebViewElement) NavigateToString(contents string) {
	js.Rewrite("$<.NavigateToString($1)", contents)
}

// NavigateWithHTTPRequestMessage
func (*MSHTMLWebViewElement) NavigateWithHTTPRequestMessage(requestMessage interface{}) {
	js.Rewrite("$<.NavigateWithHTTPRequestMessage($1)", requestMessage)
}

// Refresh
func (*MSHTMLWebViewElement) Refresh() {
	js.Rewrite("$<.Refresh()")
}

// Stop
func (*MSHTMLWebViewElement) Stop() {
	js.Rewrite("$<.Stop()")
}

// CanGoBack
func (*MSHTMLWebViewElement) CanGoBack() (canGoBack bool) {
	js.Rewrite("$<.CanGoBack")
	return canGoBack
}

// CanGoForward
func (*MSHTMLWebViewElement) CanGoForward() (canGoForward bool) {
	js.Rewrite("$<.CanGoForward")
	return canGoForward
}

// ContainsFullScreenElement
func (*MSHTMLWebViewElement) ContainsFullScreenElement() (containsFullScreenElement bool) {
	js.Rewrite("$<.ContainsFullScreenElement")
	return containsFullScreenElement
}

// DocumentTitle
func (*MSHTMLWebViewElement) DocumentTitle() (documentTitle string) {
	js.Rewrite("$<.DocumentTitle")
	return documentTitle
}

// Height
func (*MSHTMLWebViewElement) Height() (height uint) {
	js.Rewrite("$<.Height")
	return height
}

// Height
func (*MSHTMLWebViewElement) SetHeight(height uint) {
	js.Rewrite("$<.Height = $1", height)
}

// Settings
func (*MSHTMLWebViewElement) Settings() (settings *mswebviewsettings.MSWebViewSettings) {
	js.Rewrite("$<.Settings")
	return settings
}

// Src
func (*MSHTMLWebViewElement) Src() (src string) {
	js.Rewrite("$<.Src")
	return src
}

// Src
func (*MSHTMLWebViewElement) SetSrc(src string) {
	js.Rewrite("$<.Src = $1", src)
}

// Width
func (*MSHTMLWebViewElement) Width() (width uint) {
	js.Rewrite("$<.Width")
	return width
}

// Width
func (*MSHTMLWebViewElement) SetWidth(width uint) {
	js.Rewrite("$<.Width = $1", width)
}
