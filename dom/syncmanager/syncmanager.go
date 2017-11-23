package syncmanager

import "github.com/matthewmueller/golly/js"

// SyncManager struct
// js:"SyncManager,omit"
type SyncManager struct {
}

// GetTags fn
func (*SyncManager) GetTags() (s []string) {
	js.Rewrite("await $<.getTags()")
	return s
}

// Register fn
func (*SyncManager) Register(tag string) {
	js.Rewrite("await $<.register($1)", tag)
}
