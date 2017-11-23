package mswebviewsettings

import "github.com/matthewmueller/golly/js"

// MSWebViewSettings struct
// js:"MSWebViewSettings,omit"
type MSWebViewSettings struct {
}

// IsIndexedDBEnabled prop
func (*MSWebViewSettings) IsIndexedDBEnabled() (isIndexedDBEnabled bool) {
	js.Rewrite("$<.isIndexedDBEnabled")
	return isIndexedDBEnabled
}

// IsIndexedDBEnabled prop
func (*MSWebViewSettings) SetIsIndexedDBEnabled(isIndexedDBEnabled bool) {
	js.Rewrite("$<.isIndexedDBEnabled = $1", isIndexedDBEnabled)
}

// IsJavaScriptEnabled prop
func (*MSWebViewSettings) IsJavaScriptEnabled() (isJavaScriptEnabled bool) {
	js.Rewrite("$<.isJavaScriptEnabled")
	return isJavaScriptEnabled
}

// IsJavaScriptEnabled prop
func (*MSWebViewSettings) SetIsJavaScriptEnabled(isJavaScriptEnabled bool) {
	js.Rewrite("$<.isJavaScriptEnabled = $1", isJavaScriptEnabled)
}
