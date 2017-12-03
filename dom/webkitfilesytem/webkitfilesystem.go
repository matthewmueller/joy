package webkitfilesytem

import "github.com/matthewmueller/joy/macro"

// WebKitFileSystem struct
// js:"WebKitFileSystem,omit"
type WebKitFileSystem struct {
}

// Name prop
// js:"name"
func (*WebKitFileSystem) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

// Root prop
// js:"root"
func (*WebKitFileSystem) Root() (root *WebKitDirectoryEntry) {
	macro.Rewrite("$_.root")
	return root
}
