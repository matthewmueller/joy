package defs

import (
	"go/types"
	"path"

	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/util"
)

// Filer interface
type Filer interface {
	def.Definition
	Implicit() bool
}

// files struct
type files struct {
	id   string
	path string
	name string
	implicit bool
}

var _ Filer = (*files)(nil)

var filecache = map[string]def.Definition{}

// File constructor
// 
// NOTE: implicit includes are files that have been added 
// during the process, that have no representation
// in the source code (e.g. preact := js.File("./preact.js"))
// you shouldn't need this option unless you're working on
// the compiler
func File(packagePath, relative string, implicit bool) (def.Definition, error) {
	src, e := util.GoSourcePath()
	if e != nil {
		return nil, e
	}

	pkgpath := path.Join(packagePath, relative)
	fullpath := path.Join(src, packagePath, relative)

	if filecache[fullpath] != nil {
		return filecache[fullpath], nil
	}
	
	def := &files{
		id:   fullpath,
		path: pkgpath,
		implicit: implicit,
		name: path.Base(fullpath),
	}

	filecache[fullpath] = def
	return def, nil
}

func (d *files) ID() string {
	return d.id
}

func (d *files) Name() string {
	return d.name
}

func (d *files) Path() string {
	return d.path
}

func (d *files) Dependencies() (deps []def.Definition, err error) {
	return deps, nil
}

func (d *files) Exported() bool {
	return false
}

func (d *files) Type() types.Type {
	return nil
}

func (d *files) Omitted() bool {
	return false
}

func (d *files) Kind() string {
	return "FILE"
}

func (d *files) FromRuntime() bool {
	return false
}

func (d *files) Imports() map[string]string {
	return map[string]string{}
}

func (d *files) Implicit() bool {
	return d.implicit
}
