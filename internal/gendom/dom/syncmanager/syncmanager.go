package syncmanager

import "github.com/matthewmueller/golly/js"

type SyncManager struct {
}

func (*SyncManager) GetTags() (s []string) {
	js.Rewrite("await $<.getTags()")
	return s
}

func (*SyncManager) Register(tag string) {
	js.Rewrite("await $<.register($1)", tag)
}
