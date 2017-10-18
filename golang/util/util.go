package util

import (
	"errors"
	"fmt"
	"go/ast"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/apex/log"
	"github.com/fatih/structtag"
)

var runtimePkg string

// RuntimePath gets the path of our runtime
func RuntimePath() (string, error) {
	if runtimePkg != "" {
		return runtimePkg, nil
	}

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get the filepath")
	}

	root := path.Dir(path.Dir(path.Dir(file)))
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", errors.New("GOPATH must be set in your path")
	}
	gosrc := path.Join(gopath, "src")
	runtimePath := path.Join(root, "runtime")

	// runtime package
	rt, e := filepath.Rel(gosrc, runtimePath)
	if e != nil {
		return "", e
	}
	runtimePkg = rt

	return runtimePkg, nil
}

// JSTag parses a js tag if there is one
func JSTag(n *ast.CommentGroup) (*structtag.Tag, error) {
	if n == nil {
		return nil, nil
	}

	comments := n.List
	for _, comment := range comments {
		if !strings.Contains(comment.Text, "js:") {
			continue
		}

		jstag, err := JSTagFromString(comment.Text[2:])
		if err != nil {
			return nil, err
		}

		return jstag, nil
	}

	return nil, nil
}

// JSTagFromString parse tag
func JSTagFromString(tag string) (*structtag.Tag, error) {
	tags, err := structtag.Parse(tag)
	if err != nil {
		return nil, err
	}

	jstag, err := tags.Get("js")
	if err != nil && err.Error() != "tag does not exist" {
		return nil, err
	}

	return jstag, nil
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
	case *ast.KeyValueExpr:
		log.Infof("%+v", t)
		return nil, fmt.Errorf("GetIdentifier: unhandled %T", n)
	default:
		return nil, fmt.Errorf("GetIdentifier: unhandled %T", n)
	}
}

// ExprToString fn
func ExprToString(n ast.Node) (string, error) {
	switch t := n.(type) {
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
	default:
		return "", fmt.Errorf("ExprToString: unhandled %T", n)
	}
}
