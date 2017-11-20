package msapp

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/blob"
	"github.com/matthewmueller/golly/internal/gendom/dom/file"
	"github.com/matthewmueller/golly/internal/gendom/dom/msappasyncoperation"
	"github.com/matthewmueller/golly/internal/gendom/dom/msstream"
	"github.com/matthewmueller/golly/js"
)

type MSApp struct {
}

func (*MSApp) ClearTemporaryWebDataAsync() (m *msappasyncoperation.MSAppAsyncOperation) {
	js.Rewrite("$<.clearTemporaryWebDataAsync()")
	return m
}

func (*MSApp) CreateBlobFromRandomAccessStream(kind string, seeker interface{}) (b *blob.Blob) {
	js.Rewrite("$<.createBlobFromRandomAccessStream($1, $2)", kind, seeker)
	return b
}

func (*MSApp) CreateDataPackage(object interface{}) (i interface{}) {
	js.Rewrite("$<.createDataPackage($1)", object)
	return i
}

func (*MSApp) CreateDataPackageFromSelection() (i interface{}) {
	js.Rewrite("$<.createDataPackageFromSelection()")
	return i
}

func (*MSApp) CreateFileFromStorageFile(storageFile interface{}) (f *file.File) {
	js.Rewrite("$<.createFileFromStorageFile($1)", storageFile)
	return f
}

func (*MSApp) CreateStreamFromInputStream(kind string, inputStream interface{}) (m *msstream.MSStream) {
	js.Rewrite("$<.createStreamFromInputStream($1, $2)", kind, inputStream)
	return m
}

func (*MSApp) ExecAsyncAtPriority(asynchronousCallback func(args interface{}) interface{}, priority string, args interface{}) {
	js.Rewrite("$<.execAsyncAtPriority($1, $2, $3)", asynchronousCallback, priority, args)
}

func (*MSApp) ExecAtPriority(synchronousCallback func(args interface{}) interface{}, priority string, args interface{}) (i interface{}) {
	js.Rewrite("$<.execAtPriority($1, $2, $3)", synchronousCallback, priority, args)
	return i
}

func (*MSApp) GetCurrentPriority() (s string) {
	js.Rewrite("$<.getCurrentPriority()")
	return s
}

func (*MSApp) GetHTMLPrintDocumentSourceAsync(htmlDoc interface{}) (i interface{}) {
	js.Rewrite("await $<.getHtmlPrintDocumentSourceAsync($1)", htmlDoc)
	return i
}

func (*MSApp) GetViewID(view interface{}) (i interface{}) {
	js.Rewrite("$<.getViewId($1)", view)
	return i
}

func (*MSApp) IsTaskScheduledAtPriorityOrHigher(priority string) (b bool) {
	js.Rewrite("$<.isTaskScheduledAtPriorityOrHigher($1)", priority)
	return b
}

func (*MSApp) PageHandlesAllApplicationActivations(enabled bool) {
	js.Rewrite("$<.pageHandlesAllApplicationActivations($1)", enabled)
}

func (*MSApp) SuppressSubdownloadCredentialPrompts(suppress bool) {
	js.Rewrite("$<.suppressSubdownloadCredentialPrompts($1)", suppress)
}

func (*MSApp) TerminateApp(exceptionObject interface{}) {
	js.Rewrite("$<.terminateApp($1)", exceptionObject)
}
