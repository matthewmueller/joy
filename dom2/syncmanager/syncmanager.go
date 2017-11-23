package syncmanager

import "github.com/matthewmueller/golly/js"

// js:"SyncManager,omit"
type SyncManager struct {
}

// GetTags
func (*SyncManager) GetTags() (s []string) {
	js.Rewrite("await $<.GetTags()")
	return s
}

// Register
func (*SyncManager) Register(tag string) {
	js.Rewrite("await $<.Register($1)", tag)
}
