package loader_test

import (
	"path"
	"runtime"
	"testing"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/golang/loader"
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
	// file, err := ioutil.ReadFile(path.Join(dirname, name))
	// if err != nil {
	// 	t.Fatal(err)
	// }

	return path.Join(dirname, "testdata", name)
}

func TestDeepInterfaces(t *testing.T) {
	program, err := loader.Load(fixture("dom"))
	if err != nil {
		t.Fatal(err)
	}
	log.Infof("got program! %+v", program)
}
