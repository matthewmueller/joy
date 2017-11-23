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

// ClearTemporaryWebDataAsync
func (*MSApp) ClearTemporaryWebDataAsync() (m *msappasyncoperation.MSAppAsyncOperation) {
	js.Rewrite("$<.ClearTemporaryWebDataAsync()")
	return m
}

// CreateBlobFromRandomAccessStream
func (*MSApp) CreateBlobFromRandomAccessStream(kind string, seeker interface{}) (b blob.Blob) {
	js.Rewrite("$<.CreateBlobFromRandomAccessStream($1, $2)", kind, seeker)
	return b
}

// CreateDataPackage
func (*MSApp) CreateDataPackage(object interface{}) (i interface{}) {
	js.Rewrite("$<.CreateDataPackage($1)", object)
	return i
}

// CreateDataPackageFromSelection
func (*MSApp) CreateDataPackageFromSelection() (i interface{}) {
	js.Rewrite("$<.CreateDataPackageFromSelection()")
	return i
}

// CreateFileFromStorageFile
func (*MSApp) CreateFileFromStorageFile(storageFile interface{}) (f *file.File) {
	js.Rewrite("$<.CreateFileFromStorageFile($1)", storageFile)
	return f
}

// CreateStreamFromInputStream
func (*MSApp) CreateStreamFromInputStream(kind string, inputStream interface{}) (m *msstream.MSStream) {
	js.Rewrite("$<.CreateStreamFromInputStream($1, $2)", kind, inputStream)
	return m
}

// ExecAsyncAtPriority
func (*MSApp) ExecAsyncAtPriority(asynchronousCallback func(args interface{}) (MSExecAtPriorityFunctionCallback interface{}), priority string, args interface{}) {
	js.Rewrite("$<.ExecAsyncAtPriority($1, $2, $3)", asynchronousCallback, priority, args)
}

// ExecAtPriority
func (*MSApp) ExecAtPriority(synchronousCallback func(args interface{}) (MSExecAtPriorityFunctionCallback interface{}), priority string, args interface{}) (i interface{}) {
	js.Rewrite("$<.ExecAtPriority($1, $2, $3)", synchronousCallback, priority, args)
	return i
}

// GetCurrentPriority
func (*MSApp) GetCurrentPriority() (s string) {
	js.Rewrite("$<.GetCurrentPriority()")
	return s
}

// GetHTMLPrintDocumentSourceAsync
func (*MSApp) GetHTMLPrintDocumentSourceAsync(htmlDoc interface{}) (i interface{}) {
	js.Rewrite("await $<.GetHTMLPrintDocumentSourceAsync($1)", htmlDoc)
	return i
}

// GetViewID
func (*MSApp) GetViewID(view interface{}) (i interface{}) {
	js.Rewrite("$<.GetViewID($1)", view)
	return i
}

// IsTaskScheduledAtPriorityOrHigher
func (*MSApp) IsTaskScheduledAtPriorityOrHigher(priority string) (b bool) {
	js.Rewrite("$<.IsTaskScheduledAtPriorityOrHigher($1)", priority)
	return b
}

// PageHandlesAllApplicationActivations
func (*MSApp) PageHandlesAllApplicationActivations(enabled bool) {
	js.Rewrite("$<.PageHandlesAllApplicationActivations($1)", enabled)
}

// SuppressSubdownloadCredentialPrompts
func (*MSApp) SuppressSubdownloadCredentialPrompts(suppress bool) {
	js.Rewrite("$<.SuppressSubdownloadCredentialPrompts($1)", suppress)
}

// TerminateApp
func (*MSApp) TerminateApp(exceptionObject interface{}) {
	js.Rewrite("$<.TerminateApp($1)", exceptionObject)
}
