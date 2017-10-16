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
	default:
		return nil, fmt.Errorf("GetIdentifier: unhandled %T", n)
	}
}
