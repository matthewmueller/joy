package mains_test

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/matthewmueller/joy/internal/mains"
	"github.com/matthewmueller/joy/internal/testutil"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	in   []string
	out  []string
}{
	// basics
	{"no prefix file", []string{"testdata/one/main.go"}, []string{"$PWD/testdata/one"}},
	{"no prefix dir", []string{"testdata/one"}, []string{"$PWD/testdata/one"}},
	{"no prefix dir slash", []string{"testdata/one/"}, []string{"$PWD/testdata/one"}},
	{"relative file", []string{"./testdata/one/main.go"}, []string{"$PWD/testdata/one"}},
	{"relative dir", []string{"./testdata/one"}, []string{"$PWD/testdata/one"}},
	{"relative dir slash", []string{"./testdata/one/"}, []string{"$PWD/testdata/one"}},
	{"absolute file", []string{"$PWD/testdata/one/main.go"}, []string{"$PWD/testdata/one"}},
	{"absolute dir", []string{"$PWD/testdata/one"}, []string{"$PWD/testdata/one"}},
	{"absolute dir slash", []string{"$PWD/testdata/one/"}, []string{"$PWD/testdata/one"}},

	// ellipsis
	{"no prefix file ellipsis", []string{"testdata/one/..."}, []string{"$PWD/testdata/one"}},
	{"relative dir ellipsis", []string{"./testdata/one/..."}, []string{"$PWD/testdata/one"}},
	{"absolute file ellipsis", []string{"$PWD/testdata/one/..."}, []string{"$PWD/testdata/one"}},
	{"no prefix file ellipsis", []string{"testdata/two/..."}, []string{"$PWD/testdata/two/a", "$PWD/testdata/two/b"}},
	{"relative dir ellipsis", []string{"./testdata/two/..."}, []string{"$PWD/testdata/two/a", "$PWD/testdata/two/b"}},
	{"absolute file ellipsis", []string{"$PWD/testdata/two/..."}, []string{"$PWD/testdata/two/a", "$PWD/testdata/two/b"}},
}

func Test(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		t.Fatal("mains test expects a gopath")
	}

	tmpdir, err := ioutil.TempDir("", "mains_test")
	if err != nil {
		t.Fatal(err)
	}

	tmp := path.Join(tmpdir, "testdata")
	if err := testutil.Copy(path.Join(cwd, "testdata"), tmp); err != nil {
		t.Fatal(err)
	}

	os.Unsetenv("GOPATH")
	for _, test := range tests {

		os.Setenv("GOPATH", gopath)
		t.Run(test.name, func(t *testing.T) {
			ins, outs := inouts(t, test)
			mns, err := mains.Find(append([]string{}, ins...)...)
			if err != nil {
				t.Fatal(err)
			}

			assert.Len(t, mns, len(outs), "wrong length")
			for i, main := range mns {
				assert.Equal(t, outs[i], main)
			}
		})

		os.Unsetenv("GOPATH")
		t.Run(test.name+" no gopath", func(t *testing.T) {
			ins, outs := inouts(t, test)
			mns, err := mains.Find(append([]string{}, ins...)...)
			if err != nil {
				t.Fatal(err)
			}

			assert.Len(t, mns, len(outs), "wrong length")
			for i, main := range mns {
				assert.Equal(t, outs[i], main)
			}
		})

		// run against a directory outside of the gopath
		if err := os.Chdir(tmpdir); err != nil {
			t.Fatal(err)
		}

		t.Run(test.name+" outside gopath", func(t *testing.T) {
			ins, outs := inouts(t, test)
			mns, err := mains.Find(append([]string{}, ins...)...)
			if err != nil {
				t.Fatal(err)
			}

			assert.Len(t, mns, len(outs), "wrong length")
			for i, main := range mns {
				assert.Equal(t, outs[i], main)
			}
		})

		if err := os.Chdir(cwd); err != nil {
			t.Fatal(err)
		}
	}
}

func inouts(t *testing.T, test struct {
	name string
	in   []string
	out  []string
}) (ins []string, outs []string) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	for _, in := range test.in {
		ins = append(ins, strings.Replace(in, "$PWD", cwd, -1))
	}
	for _, out := range test.out {
		outs = append(outs, strings.Replace(out, "$PWD", cwd, -1))
	}
	return
}
