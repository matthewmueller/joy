package mswebviewsettings

import "github.com/matthewmueller/golly/js"

// MSWebViewSettings struct
// js:"MSWebViewSettings,omit"
type MSWebViewSettings struct {
}

// IsIndexedDBEnabled prop
// js:"isIndexedDBEnabled"
func (*MSWebViewSettings) IsIndexedDBEnabled() (isIndexedDBEnabled bool) {
	js.Rewrite("$_.isIndexedDBEnabled")
	return isIndexedDBEnabled
}

// SetIsIndexedDBEnabled prop
// js:"isIndexedDBEnabled"
func (*MSWebViewSettings) SetIsIndexedDBEnabled(isIndexedDBEnabled bool) {
	js.Rewrite("$_.isIndexedDBEnabled = $1", isIndexedDBEnabled)
}

// IsJavaScriptEnabled prop
// js:"isJavaScriptEnabled"
func (*MSWebViewSettings) IsJavaScriptEnabled() (isJavaScriptEnabled bool) {
	js.Rewrite("$_.isJavaScriptEnabled")
	return isJavaScriptEnabled
}

// SetIsJavaScriptEnabled prop
// js:"isJavaScriptEnabled"
func (*MSWebViewSettings) SetIsJavaScriptEnabled(isJavaScriptEnabled bool) {
	js.Rewrite("$_.isJavaScriptEnabled = $1", isJavaScriptEnabled)
}
