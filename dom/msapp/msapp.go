package msapp

import (
	"github.com/matthewmueller/golly/dom2/blob"
	"github.com/matthewmueller/golly/dom2/file"
	"github.com/matthewmueller/golly/dom2/msappasyncoperation"
	"github.com/matthewmueller/golly/dom2/msstream"
	"github.com/matthewmueller/golly/js"
)

// MSApp struct
// js:"MSApp,omit"
type MSApp struct {
}

// ClearTemporaryWebDataAsync fn
func (*MSApp) ClearTemporaryWebDataAsync() (m *msappasyncoperation.MSAppAsyncOperation) {
	js.Rewrite("$<.clearTemporaryWebDataAsync()")
	return m
}

// CreateBlobFromRandomAccessStream fn
func (*MSApp) CreateBlobFromRandomAccessStream(kind string, seeker interface{}) (b blob.Blob) {
	js.Rewrite("$<.createBlobFromRandomAccessStream($1, $2)", kind, seeker)
	return b
}

// CreateDataPackage fn
func (*MSApp) CreateDataPackage(object interface{}) (i interface{}) {
	js.Rewrite("$<.createDataPackage($1)", object)
	return i
}

// CreateDataPackageFromSelection fn
func (*MSApp) CreateDataPackageFromSelection() (i interface{}) {
	js.Rewrite("$<.createDataPackageFromSelection()")
	return i
}

// CreateFileFromStorageFile fn
func (*MSApp) CreateFileFromStorageFile(storageFile interface{}) (f *file.File) {
	js.Rewrite("$<.createFileFromStorageFile($1)", storageFile)
	return f
}

// CreateStreamFromInputStream fn
func (*MSApp) CreateStreamFromInputStream(kind string, inputStream interface{}) (m *msstream.MSStream) {
	js.Rewrite("$<.createStreamFromInputStream($1, $2)", kind, inputStream)
	return m
}

// ExecAsyncAtPriority fn
func (*MSApp) ExecAsyncAtPriority(asynchronousCallback func(args interface{}) (MSExecAtPriorityFunctionCallback interface{}), priority string, args interface{}) {
	js.Rewrite("$<.execAsyncAtPriority($1, $2, $3)", asynchronousCallback, priority, args)
}

// ExecAtPriority fn
func (*MSApp) ExecAtPriority(synchronousCallback func(args interface{}) (MSExecAtPriorityFunctionCallback interface{}), priority string, args interface{}) (i interface{}) {
	js.Rewrite("$<.execAtPriority($1, $2, $3)", synchronousCallback, priority, args)
	return i
}

// GetCurrentPriority fn
func (*MSApp) GetCurrentPriority() (s string) {
	js.Rewrite("$<.getCurrentPriority()")
	return s
}

// GetHTMLPrintDocumentSourceAsync fn
func (*MSApp) GetHTMLPrintDocumentSourceAsync(htmlDoc interface{}) (i interface{}) {
	js.Rewrite("await $<.getHtmlPrintDocumentSourceAsync($1)", htmlDoc)
	return i
}

// GetViewID fn
func (*MSApp) GetViewID(view interface{}) (i interface{}) {
	js.Rewrite("$<.getViewId($1)", view)
	return i
}

// IsTaskScheduledAtPriorityOrHigher fn
func (*MSApp) IsTaskScheduledAtPriorityOrHigher(priority string) (b bool) {
	js.Rewrite("$<.isTaskScheduledAtPriorityOrHigher($1)", priority)
	return b
}

// PageHandlesAllApplicationActivations fn
func (*MSApp) PageHandlesAllApplicationActivations(enabled bool) {
	js.Rewrite("$<.pageHandlesAllApplicationActivations($1)", enabled)
}

// SuppressSubdownloadCredentialPrompts fn
func (*MSApp) SuppressSubdownloadCredentialPrompts(suppress bool) {
	js.Rewrite("$<.suppressSubdownloadCredentialPrompts($1)", suppress)
}

// TerminateApp fn
func (*MSApp) TerminateApp(exceptionObject interface{}) {
	js.Rewrite("$<.terminateApp($1)", exceptionObject)
}
