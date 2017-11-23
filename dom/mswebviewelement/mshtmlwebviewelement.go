package mswebviewelement

import (
	"github.com/matthewmueller/golly/dom2/deferredpermissionrequest"
	"github.com/matthewmueller/golly/dom2/focusnavigationorigin"
	"github.com/matthewmueller/golly/dom2/mswebviewsettings"
	"github.com/matthewmueller/golly/dom2/navigationreason"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// MSHTMLWebViewElement struct
// js:"MSHTMLWebViewElement,omit"
type MSHTMLWebViewElement struct {
	window.HTMLElement
}

// AddWebAllowedObject fn
func (*MSHTMLWebViewElement) AddWebAllowedObject(name string, applicationObject interface{}) {
	js.Rewrite("$<.addWebAllowedObject($1, $2)", name, applicationObject)
}

// BuildLocalStreamURI fn
func (*MSHTMLWebViewElement) BuildLocalStreamURI(contentIdentifier string, relativePath string) (s string) {
	js.Rewrite("$<.buildLocalStreamUri($1, $2)", contentIdentifier, relativePath)
	return s
}

// CapturePreviewToBlobAsync fn
func (*MSHTMLWebViewElement) CapturePreviewToBlobAsync() (m *MSWebViewAsyncOperation) {
	js.Rewrite("$<.capturePreviewToBlobAsync()")
	return m
}

// CaptureSelectedContentToDataPackageAsync fn
func (*MSHTMLWebViewElement) CaptureSelectedContentToDataPackageAsync() (m *MSWebViewAsyncOperation) {
	js.Rewrite("$<.captureSelectedContentToDataPackageAsync()")
	return m
}

// GetDeferredPermissionRequestByID fn
func (*MSHTMLWebViewElement) GetDeferredPermissionRequestByID(id uint) (d deferredpermissionrequest.DeferredPermissionRequest) {
	js.Rewrite("$<.getDeferredPermissionRequestById($1)", id)
	return d
}

// GetDeferredPermissionRequests fn
func (*MSHTMLWebViewElement) GetDeferredPermissionRequests() (d []deferredpermissionrequest.DeferredPermissionRequest) {
	js.Rewrite("$<.getDeferredPermissionRequests()")
	return d
}

// GoBack fn
func (*MSHTMLWebViewElement) GoBack() {
	js.Rewrite("$<.goBack()")
}

// GoForward fn
func (*MSHTMLWebViewElement) GoForward() {
	js.Rewrite("$<.goForward()")
}

// InvokeScriptAsync fn
func (*MSHTMLWebViewElement) InvokeScriptAsync(scriptName string, args interface{}) (m *MSWebViewAsyncOperation) {
	js.Rewrite("$<.invokeScriptAsync($1, $2)", scriptName, args)
	return m
}

// Navigate fn
func (*MSHTMLWebViewElement) Navigate(uri string) {
	js.Rewrite("$<.navigate($1)", uri)
}

// NavigateFocus fn
func (*MSHTMLWebViewElement) NavigateFocus(navigationReason *navigationreason.NavigationReason, origin *focusnavigationorigin.FocusNavigationOrigin) {
	js.Rewrite("$<.navigateFocus($1, $2)", navigationReason, origin)
}

// NavigateToLocalStreamURI fn
func (*MSHTMLWebViewElement) NavigateToLocalStreamURI(source string, streamResolver interface{}) {
	js.Rewrite("$<.navigateToLocalStreamUri($1, $2)", source, streamResolver)
}

// NavigateToString fn
func (*MSHTMLWebViewElement) NavigateToString(contents string) {
	js.Rewrite("$<.navigateToString($1)", contents)
}

// NavigateWithHTTPRequestMessage fn
func (*MSHTMLWebViewElement) NavigateWithHTTPRequestMessage(requestMessage interface{}) {
	js.Rewrite("$<.navigateWithHttpRequestMessage($1)", requestMessage)
}

// Refresh fn
func (*MSHTMLWebViewElement) Refresh() {
	js.Rewrite("$<.refresh()")
}

// Stop fn
func (*MSHTMLWebViewElement) Stop() {
	js.Rewrite("$<.stop()")
}

// CanGoBack prop
func (*MSHTMLWebViewElement) CanGoBack() (canGoBack bool) {
	js.Rewrite("$<.canGoBack")
	return canGoBack
}

// CanGoForward prop
func (*MSHTMLWebViewElement) CanGoForward() (canGoForward bool) {
	js.Rewrite("$<.canGoForward")
	return canGoForward
}

// ContainsFullScreenElement prop
func (*MSHTMLWebViewElement) ContainsFullScreenElement() (containsFullScreenElement bool) {
	js.Rewrite("$<.containsFullScreenElement")
	return containsFullScreenElement
}

// DocumentTitle prop
func (*MSHTMLWebViewElement) DocumentTitle() (documentTitle string) {
	js.Rewrite("$<.documentTitle")
	return documentTitle
}

// Height prop
func (*MSHTMLWebViewElement) Height() (height uint) {
	js.Rewrite("$<.height")
	return height
}

// Height prop
func (*MSHTMLWebViewElement) SetHeight(height uint) {
	js.Rewrite("$<.height = $1", height)
}

// Settings prop
func (*MSHTMLWebViewElement) Settings() (settings *mswebviewsettings.MSWebViewSettings) {
	js.Rewrite("$<.settings")
	return settings
}

// Src prop
func (*MSHTMLWebViewElement) Src() (src string) {
	js.Rewrite("$<.src")
	return src
}

// Src prop
func (*MSHTMLWebViewElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

// Width prop
func (*MSHTMLWebViewElement) Width() (width uint) {
	js.Rewrite("$<.width")
	return width
}

// Width prop
func (*MSHTMLWebViewElement) SetWidth(width uint) {
	js.Rewrite("$<.width = $1", width)
}
