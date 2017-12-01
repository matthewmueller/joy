package msapp

import (
	"github.com/matthewmueller/golly/dom/blob"
	"github.com/matthewmueller/golly/dom/file"
	"github.com/matthewmueller/golly/dom/msappasyncoperation"
	"github.com/matthewmueller/golly/dom/msstream"
	"github.com/matthewmueller/golly/js"
)

// MSApp struct
// js:"MSApp,omit"
type MSApp struct {
}

// ClearTemporaryWebDataAsync fn
// js:"clearTemporaryWebDataAsync"
func (*MSApp) ClearTemporaryWebDataAsync() (m *msappasyncoperation.MSAppAsyncOperation) {
	js.Rewrite("$_.clearTemporaryWebDataAsync()")
	return m
}

// CreateBlobFromRandomAccessStream fn
// js:"createBlobFromRandomAccessStream"
func (*MSApp) CreateBlobFromRandomAccessStream(kind string, seeker interface{}) (b blob.Blob) {
	js.Rewrite("$_.createBlobFromRandomAccessStream($1, $2)", kind, seeker)
	return b
}

// CreateDataPackage fn
// js:"createDataPackage"
func (*MSApp) CreateDataPackage(object interface{}) (i interface{}) {
	js.Rewrite("$_.createDataPackage($1)", object)
	return i
}

// CreateDataPackageFromSelection fn
// js:"createDataPackageFromSelection"
func (*MSApp) CreateDataPackageFromSelection() (i interface{}) {
	js.Rewrite("$_.createDataPackageFromSelection()")
	return i
}

// CreateFileFromStorageFile fn
// js:"createFileFromStorageFile"
func (*MSApp) CreateFileFromStorageFile(storageFile interface{}) (f *file.File) {
	js.Rewrite("$_.createFileFromStorageFile($1)", storageFile)
	return f
}

// CreateStreamFromInputStream fn
// js:"createStreamFromInputStream"
func (*MSApp) CreateStreamFromInputStream(kind string, inputStream interface{}) (m *msstream.MSStream) {
	js.Rewrite("$_.createStreamFromInputStream($1, $2)", kind, inputStream)
	return m
}

// ExecAsyncAtPriority fn
// js:"execAsyncAtPriority"
func (*MSApp) ExecAsyncAtPriority(asynchronousCallback func(args interface{}) (MSExecAtPriorityFunctionCallback interface{}), priority string, args interface{}) {
	js.Rewrite("$_.execAsyncAtPriority($1, $2, $3)", asynchronousCallback, priority, args)
}

// ExecAtPriority fn
// js:"execAtPriority"
func (*MSApp) ExecAtPriority(synchronousCallback func(args interface{}) (MSExecAtPriorityFunctionCallback interface{}), priority string, args interface{}) (i interface{}) {
	js.Rewrite("$_.execAtPriority($1, $2, $3)", synchronousCallback, priority, args)
	return i
}

// GetCurrentPriority fn
// js:"getCurrentPriority"
func (*MSApp) GetCurrentPriority() (s string) {
	js.Rewrite("$_.getCurrentPriority()")
	return s
}

// GetHTMLPrintDocumentSourceAsync fn
// js:"getHtmlPrintDocumentSourceAsync"
func (*MSApp) GetHTMLPrintDocumentSourceAsync(htmlDoc interface{}) (i interface{}) {
	js.Rewrite("await $_.getHtmlPrintDocumentSourceAsync($1)", htmlDoc)
	return i
}

// GetViewID fn
// js:"getViewId"
func (*MSApp) GetViewID(view interface{}) (i interface{}) {
	js.Rewrite("$_.getViewId($1)", view)
	return i
}

// IsTaskScheduledAtPriorityOrHigher fn
// js:"isTaskScheduledAtPriorityOrHigher"
func (*MSApp) IsTaskScheduledAtPriorityOrHigher(priority string) (b bool) {
	js.Rewrite("$_.isTaskScheduledAtPriorityOrHigher($1)", priority)
	return b
}

// PageHandlesAllApplicationActivations fn
// js:"pageHandlesAllApplicationActivations"
func (*MSApp) PageHandlesAllApplicationActivations(enabled bool) {
	js.Rewrite("$_.pageHandlesAllApplicationActivations($1)", enabled)
}

// SuppressSubdownloadCredentialPrompts fn
// js:"suppressSubdownloadCredentialPrompts"
func (*MSApp) SuppressSubdownloadCredentialPrompts(suppress bool) {
	js.Rewrite("$_.suppressSubdownloadCredentialPrompts($1)", suppress)
}

// TerminateApp fn
// js:"terminateApp"
func (*MSApp) TerminateApp(exceptionObject interface{}) {
	js.Rewrite("$_.terminateApp($1)", exceptionObject)
}
