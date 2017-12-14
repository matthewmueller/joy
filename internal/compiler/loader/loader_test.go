package loader_test

import (
	"path"
	"runtime"
	"testing"

	"github.com/matthewmueller/joy/internal/compiler/loader"
)

var dirname string

func init() {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		panic("file not found")
	}
	dirname = path.Dir(f)
}

func fixture(name string) string {
	return path.Join(dirname, "testdata", name)
}

func TestDeepInterfaceFix(t *testing.T) {
	_, err := loader.Load(fixture("dom"))
	if err != nil {
		t.Fatal(err)
	}
}
