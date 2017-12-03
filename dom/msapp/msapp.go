package msapp

import (
	"github.com/matthewmueller/joy/dom/blob"
	"github.com/matthewmueller/joy/dom/file"
	"github.com/matthewmueller/joy/dom/msappasyncoperation"
	"github.com/matthewmueller/joy/dom/msstream"
	"github.com/matthewmueller/joy/macro"
)

// MSApp struct
// js:"MSApp,omit"
type MSApp struct {
}

// ClearTemporaryWebDataAsync fn
// js:"clearTemporaryWebDataAsync"
func (*MSApp) ClearTemporaryWebDataAsync() (m *msappasyncoperation.MSAppAsyncOperation) {
	macro.Rewrite("$_.clearTemporaryWebDataAsync()")
	return m
}

// CreateBlobFromRandomAccessStream fn
// js:"createBlobFromRandomAccessStream"
func (*MSApp) CreateBlobFromRandomAccessStream(kind string, seeker interface{}) (b blob.Blob) {
	macro.Rewrite("$_.createBlobFromRandomAccessStream($1, $2)", kind, seeker)
	return b
}

// CreateDataPackage fn
// js:"createDataPackage"
func (*MSApp) CreateDataPackage(object interface{}) (i interface{}) {
	macro.Rewrite("$_.createDataPackage($1)", object)
	return i
}

// CreateDataPackageFromSelection fn
// js:"createDataPackageFromSelection"
func (*MSApp) CreateDataPackageFromSelection() (i interface{}) {
	macro.Rewrite("$_.createDataPackageFromSelection()")
	return i
}

// CreateFileFromStorageFile fn
// js:"createFileFromStorageFile"
func (*MSApp) CreateFileFromStorageFile(storageFile interface{}) (f *file.File) {
	macro.Rewrite("$_.createFileFromStorageFile($1)", storageFile)
	return f
}

// CreateStreamFromInputStream fn
// js:"createStreamFromInputStream"
func (*MSApp) CreateStreamFromInputStream(kind string, inputStream interface{}) (m *msstream.MSStream) {
	macro.Rewrite("$_.createStreamFromInputStream($1, $2)", kind, inputStream)
	return m
}

// ExecAsyncAtPriority fn
// js:"execAsyncAtPriority"
func (*MSApp) ExecAsyncAtPriority(asynchronousCallback func(args interface{}) (MSExecAtPriorityFunctionCallback interface{}), priority string, args interface{}) {
	macro.Rewrite("$_.execAsyncAtPriority($1, $2, $3)", asynchronousCallback, priority, args)
}

// ExecAtPriority fn
// js:"execAtPriority"
func (*MSApp) ExecAtPriority(synchronousCallback func(args interface{}) (MSExecAtPriorityFunctionCallback interface{}), priority string, args interface{}) (i interface{}) {
	macro.Rewrite("$_.execAtPriority($1, $2, $3)", synchronousCallback, priority, args)
	return i
}

// GetCurrentPriority fn
// js:"getCurrentPriority"
func (*MSApp) GetCurrentPriority() (s string) {
	macro.Rewrite("$_.getCurrentPriority()")
	return s
}

// GetHTMLPrintDocumentSourceAsync fn
// js:"getHtmlPrintDocumentSourceAsync"
func (*MSApp) GetHTMLPrintDocumentSourceAsync(htmlDoc interface{}) (i interface{}) {
	macro.Rewrite("await $_.getHtmlPrintDocumentSourceAsync($1)", htmlDoc)
	return i
}

// GetViewID fn
// js:"getViewId"
func (*MSApp) GetViewID(view interface{}) (i interface{}) {
	macro.Rewrite("$_.getViewId($1)", view)
	return i
}

// IsTaskScheduledAtPriorityOrHigher fn
// js:"isTaskScheduledAtPriorityOrHigher"
func (*MSApp) IsTaskScheduledAtPriorityOrHigher(priority string) (b bool) {
	macro.Rewrite("$_.isTaskScheduledAtPriorityOrHigher($1)", priority)
	return b
}

// PageHandlesAllApplicationActivations fn
// js:"pageHandlesAllApplicationActivations"
func (*MSApp) PageHandlesAllApplicationActivations(enabled bool) {
	macro.Rewrite("$_.pageHandlesAllApplicationActivations($1)", enabled)
}

// SuppressSubdownloadCredentialPrompts fn
// js:"suppressSubdownloadCredentialPrompts"
func (*MSApp) SuppressSubdownloadCredentialPrompts(suppress bool) {
	macro.Rewrite("$_.suppressSubdownloadCredentialPrompts($1)", suppress)
}

// TerminateApp fn
// js:"terminateApp"
func (*MSApp) TerminateApp(exceptionObject interface{}) {
	macro.Rewrite("$_.terminateApp($1)", exceptionObject)
}
