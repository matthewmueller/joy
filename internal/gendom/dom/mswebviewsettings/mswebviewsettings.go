package mswebviewsettings

import "github.com/matthewmueller/golly/js"

type MSWebViewSettings struct {
}

func (*MSWebViewSettings) GetIsIndexedDBEnabled() (isIndexedDBEnabled bool) {
	js.Rewrite("$<.isIndexedDBEnabled")
	return isIndexedDBEnabled
}

func (*MSWebViewSettings) SetIsIndexedDBEnabled(isIndexedDBEnabled bool) {
	js.Rewrite("$<.isIndexedDBEnabled = $1", isIndexedDBEnabled)
}

func (*MSWebViewSettings) GetIsJavaScriptEnabled() (isJavaScriptEnabled bool) {
	js.Rewrite("$<.isJavaScriptEnabled")
	return isJavaScriptEnabled
}

func (*MSWebViewSettings) SetIsJavaScriptEnabled(isJavaScriptEnabled bool) {
	js.Rewrite("$<.isJavaScriptEnabled = $1", isJavaScriptEnabled)
}
