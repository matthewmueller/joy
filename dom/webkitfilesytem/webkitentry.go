package webkitfilesytem

// js:"WebKitEntry,omit"
type WebKitEntry interface {

	// Filesystem prop
	// js:"filesystem"
	Filesystem() (filesystem *WebKitFileSystem)

	// FullPath prop
	// js:"fullPath"
	FullPath() (fullPath string)

	// IsDirectory prop
	// js:"isDirectory"
	IsDirectory() (isDirectory bool)

	// IsFile prop
	// js:"isFile"
	IsFile() (isFile bool)

	// Name prop
	// js:"name"
	Name() (name string)
}
