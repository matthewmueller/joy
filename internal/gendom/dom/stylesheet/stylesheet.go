package stylesheet

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/medialist"
	"github.com/matthewmueller/golly/internal/gendom/dom/node"
	"github.com/matthewmueller/golly/js"
)

type StyleSheet struct {
}

func (*StyleSheet) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*StyleSheet) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*StyleSheet) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*StyleSheet) GetMedia() (media *medialist.MediaList) {
	js.Rewrite("$<.media")
	return media
}

func (*StyleSheet) GetOwnerNode() (ownerNode *node.Node) {
	js.Rewrite("$<.ownerNode")
	return ownerNode
}

func (*StyleSheet) GetParentStyleSheet() (parentStyleSheet *StyleSheet) {
	js.Rewrite("$<.parentStyleSheet")
	return parentStyleSheet
}

func (*StyleSheet) GetTitle() (title string) {
	js.Rewrite("$<.title")
	return title
}

func (*StyleSheet) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
