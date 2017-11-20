package mshtmlwebviewelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/deferredpermissionrequest"
	"github.com/matthewmueller/golly/internal/gendom/dom/focusnavigationorigin"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/mswebviewasyncoperation"
	"github.com/matthewmueller/golly/internal/gendom/dom/mswebviewsettings"
	"github.com/matthewmueller/golly/internal/gendom/dom/navigationreason"
	"github.com/matthewmueller/golly/js"
)

type MSHTMLWebViewElement struct {
	*htmlelement.HTMLElement
}

func (*MSHTMLWebViewElement) AddWebAllowedObject(name string, applicationObject interface{}) {
	js.Rewrite("$<.addWebAllowedObject($1, $2)", name, applicationObject)
}

func (*MSHTMLWebViewElement) BuildLocalStreamURI(contentIdentifier string, relativePath string) (s string) {
	js.Rewrite("$<.buildLocalStreamUri($1, $2)", contentIdentifier, relativePath)
	return s
}

func (*MSHTMLWebViewElement) CapturePreviewToBlobAsync() (m *mswebviewasyncoperation.MSWebViewAsyncOperation) {
	js.Rewrite("$<.capturePreviewToBlobAsync()")
	return m
}

func (*MSHTMLWebViewElement) CaptureSelectedContentToDataPackageAsync() (m *mswebviewasyncoperation.MSWebViewAsyncOperation) {
	js.Rewrite("$<.captureSelectedContentToDataPackageAsync()")
	return m
}

func (*MSHTMLWebViewElement) GetDeferredPermissionRequestByID(id uint) (d *deferredpermissionrequest.DeferredPermissionRequest) {
	js.Rewrite("$<.getDeferredPermissionRequestById($1)", id)
	return d
}

func (*MSHTMLWebViewElement) GetDeferredPermissionRequests() (d []*deferredpermissionrequest.DeferredPermissionRequest) {
	js.Rewrite("$<.getDeferredPermissionRequests()")
	return d
}

func (*MSHTMLWebViewElement) GoBack() {
	js.Rewrite("$<.goBack()")
}

func (*MSHTMLWebViewElement) GoForward() {
	js.Rewrite("$<.goForward()")
}

func (*MSHTMLWebViewElement) InvokeScriptAsync(scriptName string, args interface{}) (m *mswebviewasyncoperation.MSWebViewAsyncOperation) {
	js.Rewrite("$<.invokeScriptAsync($1, $2)", scriptName, args)
	return m
}

func (*MSHTMLWebViewElement) Navigate(uri string) {
	js.Rewrite("$<.navigate($1)", uri)
}

func (*MSHTMLWebViewElement) NavigateFocus(navigationReason *navigationreason.NavigationReason, origin *focusnavigationorigin.FocusNavigationOrigin) {
	js.Rewrite("$<.navigateFocus($1, $2)", navigationReason, origin)
}

func (*MSHTMLWebViewElement) NavigateToLocalStreamURI(source string, streamResolver interface{}) {
	js.Rewrite("$<.navigateToLocalStreamUri($1, $2)", source, streamResolver)
}

func (*MSHTMLWebViewElement) NavigateToString(contents string) {
	js.Rewrite("$<.navigateToString($1)", contents)
}

func (*MSHTMLWebViewElement) NavigateWithHTTPRequestMessage(requestMessage interface{}) {
	js.Rewrite("$<.navigateWithHttpRequestMessage($1)", requestMessage)
}

func (*MSHTMLWebViewElement) Refresh() {
	js.Rewrite("$<.refresh()")
}

func (*MSHTMLWebViewElement) Stop() {
	js.Rewrite("$<.stop()")
}

func (*MSHTMLWebViewElement) GetCanGoBack() (canGoBack bool) {
	js.Rewrite("$<.canGoBack")
	return canGoBack
}

func (*MSHTMLWebViewElement) GetCanGoForward() (canGoForward bool) {
	js.Rewrite("$<.canGoForward")
	return canGoForward
}

func (*MSHTMLWebViewElement) GetContainsFullScreenElement() (containsFullScreenElement bool) {
	js.Rewrite("$<.containsFullScreenElement")
	return containsFullScreenElement
}

func (*MSHTMLWebViewElement) GetDocumentTitle() (documentTitle string) {
	js.Rewrite("$<.documentTitle")
	return documentTitle
}

func (*MSHTMLWebViewElement) GetHeight() (height uint) {
	js.Rewrite("$<.height")
	return height
}

func (*MSHTMLWebViewElement) SetHeight(height uint) {
	js.Rewrite("$<.height = $1", height)
}

func (*MSHTMLWebViewElement) GetSettings() (settings *mswebviewsettings.MSWebViewSettings) {
	js.Rewrite("$<.settings")
	return settings
}

func (*MSHTMLWebViewElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*MSHTMLWebViewElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*MSHTMLWebViewElement) GetWidth() (width uint) {
	js.Rewrite("$<.width")
	return width
}

func (*MSHTMLWebViewElement) SetWidth(width uint) {
	js.Rewrite("$<.width = $1", width)
}
