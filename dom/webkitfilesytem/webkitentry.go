package webkitfilesytem

import "github.com/matthewmueller/golly/js"

// js:"WebKitEntry,omit"
type WebKitEntry interface {

	// Filesystem prop
	// js:"filesystem,rewrite=filesystem"
	Filesystem() (filesystem *WebKitFileSystem)

	// FullPath prop
	// js:"fullPath,rewrite=fullpath"
	FullPath() (fullPath string)

	// IsDirectory prop
	// js:"isDirectory,rewrite=isdirectory"
	IsDirectory() (isDirectory bool)

	// IsFile prop
	// js:"isFile,rewrite=isfile"
	IsFile() (isFile bool)

	// Name prop
	// js:"name,rewrite=name"
	Name() (name string)
}

// filesystem prop
func filesystem() (filesystem *WebKitFileSystem) {
	js.Rewrite("$<.filesystem")
	return filesystem
}

// fullpath prop
func fullpath() (fullPath string) {
	js.Rewrite("$<.fullPath")
	return fullPath
}

// isdirectory prop
func isdirectory() (isDirectory bool) {
	js.Rewrite("$<.isDirectory")
	return isDirectory
}

// isfile prop
func isfile() (isFile bool) {
	js.Rewrite("$<.isFile")
	return isFile
}

// name prop
func name() (name string) {
	js.Rewrite("$<.name")
	return name
}
