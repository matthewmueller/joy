package mswebviewsettings

import "github.com/matthewmueller/golly/js"

// MSWebViewSettings struct
// js:"MSWebViewSettings,omit"
type MSWebViewSettings struct {
}

// IsIndexedDBEnabled
func (*MSWebViewSettings) IsIndexedDBEnabled() (isIndexedDBEnabled bool) {
	js.Rewrite("$<.IsIndexedDBEnabled")
	return isIndexedDBEnabled
}

// IsIndexedDBEnabled
func (*MSWebViewSettings) SetIsIndexedDBEnabled(isIndexedDBEnabled bool) {
	js.Rewrite("$<.IsIndexedDBEnabled = $1", isIndexedDBEnabled)
}

// IsJavaScriptEnabled
func (*MSWebViewSettings) IsJavaScriptEnabled() (isJavaScriptEnabled bool) {
	js.Rewrite("$<.IsJavaScriptEnabled")
	return isJavaScriptEnabled
}

// IsJavaScriptEnabled
func (*MSWebViewSettings) SetIsJavaScriptEnabled(isJavaScriptEnabled bool) {
	js.Rewrite("$<.IsJavaScriptEnabled = $1", isJavaScriptEnabled)
}
