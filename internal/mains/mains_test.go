package mains_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/matthewmueller/joy/internal/mains"
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
	if err := copyDir(path.Join(cwd, "testdata"), tmp); err != nil {
		t.Fatal(err)
	}

	os.Unsetenv("GOPATH")
	for _, test := range tests {

		os.Setenv("GOPATH", gopath)
		t.Run(test.name, func(t *testing.T) {
			ins, outs := inouts(t, test)
			mns, err := mains.Find(append([]string{}, ins...))
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
			mns, err := mains.Find(append([]string{}, ins...))
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
			mns, err := mains.Find(append([]string{}, ins...))
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

// CopyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
// Symlinks are ignored and skipped.
func copyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("destination already exists")
	}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = copyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = copyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}

// CopyFile copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file. The file mode will be copied from the source and
// the copied data is synced/flushed to stable storage.
func copyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}
