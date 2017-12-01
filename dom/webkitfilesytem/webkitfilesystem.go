package webkitfilesytem

import "github.com/matthewmueller/golly/js"

// WebKitFileSystem struct
// js:"WebKitFileSystem,omit"
type WebKitFileSystem struct {
}

// Name prop
// js:"name"
func (*WebKitFileSystem) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// Root prop
// js:"root"
func (*WebKitFileSystem) Root() (root *WebKitDirectoryEntry) {
	js.Rewrite("$_.root")
	return root
}
