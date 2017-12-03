package util

import (
	"fmt"
	"go/ast"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/apex/log"
	"github.com/fatih/structtag"
	"github.com/matthewmueller/joy/internal/compiler/def"
	"github.com/pkg/errors"
)

var joyPath string
var goSourcePath string
var runtimePkg string
var jsSourcePkg string
var vdomSourcePkg string

// JoyPath absolute path
func JoyPath() (string, error) {
	if joyPath != "" {
		return joyPath, nil
	}

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get the joy source root")
	}
	root := path.Dir(path.Dir(path.Dir(path.Dir(file))))
	joyPath = root
	return root, nil
}

// RuntimePath gets the path of our runtime
func RuntimePath() (string, error) {
	if runtimePkg != "" {
		return runtimePkg, nil
	}

	root, e := JoyPath()
	if e != nil {
		return "", e
	}
	runtimePath := path.Join(root, "internal", "runtime")

	gosrc, e := GoSourcePath()
	if e != nil {
		return "", e
	}

	// runtime package
	rt, e := filepath.Rel(gosrc, runtimePath)
	if e != nil {
		return "", e
	}
	runtimePkg = rt

	return runtimePkg, nil
}

// JSTag struct
type JSTag struct {
	Rename  string
	Omit    bool
	Async   bool
	Rewrite string
}

// JSTagFromComment parses a js tag if there is one
func JSTagFromComment(n *ast.CommentGroup) (tag JSTag, err error) {
	if n == nil {
		return tag, nil
	}

	comments := n.List
	for _, comment := range comments {
		text := strings.TrimLeft(comment.Text, "/ ")

		if strings.HasPrefix(text, "js:") {
			t, err := JSTagFromString(text)
			if err != nil {
				return tag, errors.Wrapf(err, "error getting tag")
			}
			tag.Rename = t.Rename
			tag.Async = t.Async
			tag.Omit = t.Omit
		}

		if strings.HasPrefix(text, "jsrewrite:") {
			t, err := JSTagFromString(text)
			if err != nil {
				return tag, errors.Wrapf(err, "error getting tag")
			}
			tag.Rewrite = t.Rewrite
		}
	}

	return tag, nil
}

// JSTagFromString parse tag
func JSTagFromString(tagstring string) (tag JSTag, err error) {
	tags, err := structtag.Parse(tagstring)
	if err != nil {
		return tag, err
	} else if tags == nil {
		return tag, nil
	}

	js, err := tags.Get("js")
	if err == nil {
		tag.Rename = js.Name
		tag.Omit = js.HasOption("omit")
		tag.Async = js.HasOption("async")
	}

	// TODO: fix this hacky code
	// structtag treats "," as options (rightfully so)
	// but that'll break functions with parameters
	rewrite, err := tags.Get("jsrewrite")
	if err == nil {
		fulltag := rewrite.String()
		fulltag = fulltag[len("jsrewrite:"):]

		unquoted, err := strconv.Unquote(fulltag)
		if err != nil {
			return tag, err
		}

		tag.Rewrite = unquoted
	}

	return tag, nil
}

// GetIdentifier gets rightmost identifier if there is one
func GetIdentifier(n ast.Node) (*ast.Ident, error) {
	switch t := n.(type) {
	case *ast.Ident:
		return t, nil
	case *ast.StarExpr:
		return GetIdentifier(t.X)
	case *ast.UnaryExpr:
		return GetIdentifier(t.X)
	case *ast.SelectorExpr:
		return GetIdentifier(t.Sel)
	case *ast.IndexExpr:
		return GetIdentifier(t.X)
	case *ast.CallExpr:
		return GetIdentifier(t.Fun)
	case *ast.CompositeLit:
		return GetIdentifier(t.Type)
	case *ast.FuncDecl:
		return t.Name, nil
	case *ast.ParenExpr:
		return GetIdentifier(t.X)
	case *ast.ArrayType, *ast.MapType, *ast.StructType,
		*ast.ChanType, *ast.FuncType, *ast.InterfaceType,
		*ast.FuncLit, *ast.BinaryExpr:
		return nil, nil
	case *ast.SliceExpr:
		return GetIdentifier(t.X)
	default:
		_, file, line, _ := runtime.Caller(2)
		log.Warnf("GetIdentifier: file=%s line=%s", file, line)
		return nil, fmt.Errorf("GetIdentifier: unhandled %T", n)
	}
}

// ExprToString fn
func ExprToString(n ast.Node) (string, error) {
	switch t := n.(type) {
	case *ast.BasicLit:
		return t.Value, nil
	case *ast.Ident:
		return t.Name, nil
	case *ast.StarExpr:
		return ExprToString(t.X)
	case *ast.UnaryExpr:
		return ExprToString(t.X)
	case *ast.SelectorExpr:
		s, e := ExprToString(t.X)
		if e != nil {
			return "", e
		}
		return s + "." + t.Sel.Name, nil
	case *ast.IndexExpr:
		s, e := ExprToString(t.X)
		if e != nil {
			return "", e
		}
		i, e := ExprToString(t.Index)
		if e != nil {
			return "", e
		}
		return s + "[" + i + "]", nil
	case *ast.CallExpr:
		c, e := ExprToString(t.Fun)
		if e != nil {
			return "", e
		}
		var args []string
		for _, arg := range t.Args {
			a, e := ExprToString(arg)
			if e != nil {
				return "", e
			}
			args = append(args, a)
		}
		return c + "(" + strings.Join(args, ", ") + ")", nil
	case *ast.CompositeLit:
		c, e := ExprToString(t.Type)
		if e != nil {
			return "", e
		}
		var args []string
		for _, arg := range t.Elts {
			a, e := ExprToString(arg)
			if e != nil {
				return "", e
			}
			args = append(args, a)
		}
		return c + "{" + strings.Join(args, ", ") + "}", nil
	case *ast.ArrayType:
		c, e := ExprToString(t.Elt)
		if e != nil {
			return "", e
		}
		return `[]` + c, nil
	case *ast.FuncLit:
		return "func(){}()", nil
	case *ast.ParenExpr:
		x, e := ExprToString(t.X)
		if e != nil {
			return "", e
		}
		return "(" + x + ")", nil
	case *ast.KeyValueExpr:
		k, e := ExprToString(t.Key)
		if e != nil {
			return "", e
		}
		v, e := ExprToString(t.Value)
		if e != nil {
			return "", e
		}
		return "{" + k + ":" + v + "}", nil
	case *ast.MapType:
		k, e := ExprToString(t.Key)
		if e != nil {
			return "", e
		}
		v, e := ExprToString(t.Value)
		if e != nil {
			return "", e
		}
		return "{" + k + ":" + v + "}", nil
	case *ast.BinaryExpr:
		x, e := ExprToString(t.X)
		if e != nil {
			return "", e
		}
		y, e := ExprToString(t.Y)
		if e != nil {
			return "", e
		}
		return x + " " + t.Op.String() + " " + y, nil
	case *ast.InterfaceType:
		return "interface{}", nil
	default:
		return "", fmt.Errorf("util/ExprToString: unhandled %T", n)
	}
}

// GoSourcePath gets the fullpath to Go's source files
func GoSourcePath() (string, error) {
	if goSourcePath != "" {
		return goSourcePath, nil
	}

	p := os.Getenv("GOPATH")
	if p == "" {
		return "", errors.New("GOPATH not set")
	}

	goSourcePath = path.Join(p, "src")
	return goSourcePath, nil
}

// JSSourcePath gets the fullpath to Go's source files
func JSSourcePath() (string, error) {
	if jsSourcePkg != "" {
		return jsSourcePkg, nil
	}

	root, e := JoyPath()
	if e != nil {
		return "", e
	}
	runtimePath := path.Join(root, "js")

	gosrc, e := GoSourcePath()
	if e != nil {
		return "", e
	}

	// runtime package
	rel, e := filepath.Rel(gosrc, runtimePath)
	if e != nil {
		return "", e
	}
	jsSourcePkg = rel

	return jsSourcePkg, nil
}

// VDOMSourcePath gets the fullpath to the VDOM source files
func VDOMSourcePath() (string, error) {
	if vdomSourcePkg != "" {
		return vdomSourcePkg, nil
	}

	root, e := JoyPath()
	if e != nil {
		return "", e
	}
	runtimePath := path.Join(root, "vdom")

	gosrc, e := GoSourcePath()
	if e != nil {
		return "", e
	}

	// runtime package
	rel, e := filepath.Rel(gosrc, runtimePath)
	if e != nil {
		return "", e
	}
	vdomSourcePkg = rel

	return vdomSourcePkg, nil
}

// MethodsFromInterface fn
func MethodsFromInterface(n *ast.InterfaceType, path, name string) (methods []string) {
	// link the methods to the interface declaration
	for _, m := range n.Methods.List {
		for _, n := range m.Names {
			mid := path + " " + n.Name + " " + name
			methods = append(methods, mid)
		}
	}
	return methods
}

// Unique returns definitions unique by their id
func Unique(defs []def.Definition) []def.Definition {
	u := make([]def.Definition, 0, len(defs))
	seen := make(map[string]bool)

	for _, def := range defs {
		if _, ok := seen[def.ID()]; !ok {
			seen[def.ID()] = true
			u = append(u, def)
		}
	}

	return u
}

// GetExprCaller gets the expression caller
// e.g. x.Cool() => x
func GetExprCaller(n ast.Node) (ast.Expr, error) {
	switch t := n.(type) {
	case *ast.SelectorExpr:
		return t.X, nil
	case *ast.Ident:
		return nil, nil
	default:
		return nil, fmt.Errorf("util/GetExprCaller: unhandled %T", t)
	}
}
