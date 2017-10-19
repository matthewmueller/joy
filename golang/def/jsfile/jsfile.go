package jsfile

import (
	"go/types"
	"path"

	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/util"
)

// JSFile interface
type JSFile interface {
	def.Definition
}

type jsfiledef struct {
	id   string
	path string
	name string
}

var _ JSFile = (*jsfiledef)(nil)

// NewFile fn
func NewFile(packagePath, relative string) (def.Definition, error) {
	src, e := util.GoSourcePath()
	if e != nil {
		return nil, e
	}

	pkgpath := path.Join(packagePath, relative)
	fullpath := path.Join(src, packagePath, relative)

	return &jsfiledef{
		// index:    index,
		// info:     info,
		id:   fullpath,
		path: pkgpath,
		name: path.Base(fullpath),
		// exported: exported,
		// node:     n,
		// params:   params,
		// kind:     info.TypeOf(n.Name),
		// runtime:  fromRuntime,
		// tag:      tag,
		// imports:  map[string]string{},
	}, nil
}

func (d *jsfiledef) ID() string {
	return d.id
}

func (d *jsfiledef) Name() string {
	return d.name
}

func (d *jsfiledef) Path() string {
	return d.path
}

func (d *jsfiledef) Dependencies() (defs []def.Definition, err error) {
	return defs, nil
}

func (d *jsfiledef) Exported() bool {
	return false
}

func (d *jsfiledef) Type() types.Type {
	return nil
}

func (d *jsfiledef) Omitted() bool {
	return false
}

func (d *jsfiledef) Kind() string {
	return "JSFILE"
}

func (d *jsfiledef) FromRuntime() bool {
	return false
}

func (d *jsfiledef) Imports() map[string]string {
	return map[string]string{}
}
