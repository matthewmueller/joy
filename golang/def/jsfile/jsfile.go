package jsfile

import (
	"go/types"
	"path"

	"github.com/matthewmueller/golly/golang/def"
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
func NewFile(packagePath, filepath string) (def.Definition, error) {

	return &jsfiledef{
		// index:    index,
		// info:     info,
		id:   filepath,
		path: packagePath,
		name: path.Base(filepath),
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

func (d *jsfiledef) FromRuntime() bool {
	return false
}

func (d *jsfiledef) Imports() map[string]string {
	return map[string]string{}
}
