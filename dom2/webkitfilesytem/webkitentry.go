package webkitfilesytem

// js:"WebKitEntry,omit"
type WebKitEntry interface {

	// Filesystem
	Filesystem() (filesystem *WebKitFileSystem)

	// FullPath
	FullPath() (fullPath string)

	// IsDirectory
	IsDirectory() (isDirectory bool)

	// IsFile
	IsFile() (isFile bool)

	// Name
	Name() (name string)
}
