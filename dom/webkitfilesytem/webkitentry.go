package webkitfilesytem

// WebKitEntry interface
// js:"WebKitEntry"
type WebKitEntry interface {

	// Filesystem prop
	// js:"filesystem"
	// jsrewrite:"$_.filesystem"
	Filesystem() (filesystem *WebKitFileSystem)

	// FullPath prop
	// js:"fullPath"
	// jsrewrite:"$_.fullPath"
	FullPath() (fullPath string)

	// IsDirectory prop
	// js:"isDirectory"
	// jsrewrite:"$_.isDirectory"
	IsDirectory() (isDirectory bool)

	// IsFile prop
	// js:"isFile"
	// jsrewrite:"$_.isFile"
	IsFile() (isFile bool)

	// Name prop
	// js:"name"
	// jsrewrite:"$_.name"
	Name() (name string)
}
