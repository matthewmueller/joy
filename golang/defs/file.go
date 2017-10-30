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
}

// files struct
type files struct {
	id   string
	path string
	name string
}

var _ Filer = (*files)(nil)

// File constructor
func File(packagePath, relative string) (def.Definition, error) {
	src, e := util.GoSourcePath()
	if e != nil {
		return nil, e
	}

	pkgpath := path.Join(packagePath, relative)
	fullpath := path.Join(src, packagePath, relative)

	return &files{
		id:   fullpath,
		path: pkgpath,
		name: path.Base(fullpath),
	}, nil
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

func (d *files) Dependencies() (edges []def.Edge, err error) {
	return edges, nil
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
