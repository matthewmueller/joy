package syncmanager

import "github.com/matthewmueller/joy/js"

// SyncManager struct
// js:"SyncManager,omit"
type SyncManager struct {
}

// GetTags fn
// js:"getTags"
func (*SyncManager) GetTags() (s []string) {
	js.Rewrite("await $_.getTags()")
	return s
}

// Register fn
// js:"register"
func (*SyncManager) Register(tag string) {
	js.Rewrite("await $_.register($1)", tag)
}
