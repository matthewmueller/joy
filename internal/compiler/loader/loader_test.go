package loader_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/matthewmueller/joy/internal/compiler/loader"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name     string
	packages []string
	outpaths []string
}{
	// {"empty", []string{}, []string{}},
	{"00-basic", []string{"$PWD/testdata/00-basic"}, []string{
		"$PWD/testdata/00-basic",
		"$PWD/testdata/internal/runtime",
		"$PWD/testdata/macro",
		"$PWD/testdata/stdlib/fmt",
	}},
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
	os.Unsetenv("GOPATH")

	tmpdir, err := ioutil.TempDir("", "loader_test")
	if err != nil {
		t.Fatal(err)
	}

	tmp := path.Join(tmpdir, "testdata")
	if err := copyDir(path.Join(cwd, "testdata"), tmp); err != nil {
		t.Fatal(err)
	}

	for _, test := range tests {

		// 1) has a $GOPATH
		os.Setenv("GOPATH", gopath)
		t.Run(test.name, func(t *testing.T) {
			var packages []string
			for _, pkg := range test.packages {
				packages = append(packages, strings.Replace(pkg, "$PWD", cwd, -1))
			}
			var outpaths []string
			for _, pkg := range test.outpaths {
				outpaths = append(outpaths, strings.Replace(pkg, "$PWD", cwd, -1))
			}

			p, err := loader.Load(&loader.Config{
				JoyPath:  path.Join(cwd, "testdata"),
				Packages: packages,
			})
			if err != nil {
				t.Fatal(err)
			}

			var paths []string
			for _, pkg := range p.AllPackages {
				paths = append(paths, path.Join(gopath, "src", pkg.Pkg.Path()))
			}
			sort.Strings(paths)

			for i, p := range paths {
				assert.Equal(t, outpaths[i], p)
			}
		})

		// No $GOPATH
		os.Unsetenv("GOPATH")
		t.Run(test.name+" no gopath", func(t *testing.T) {
			var packages []string
			for _, pkg := range test.packages {
				packages = append(packages, strings.Replace(pkg, "$PWD", cwd, -1))
			}
			var outpaths []string
			for _, pkg := range test.outpaths {
				outpaths = append(outpaths, strings.Replace(pkg, "$PWD", cwd, -1))
			}

			p, err := loader.Load(&loader.Config{
				JoyPath:  path.Join(cwd, "testdata"),
				Packages: packages,
			})
			if err != nil {
				t.Fatal(err)
			}

			var paths []string
			for _, pkg := range p.AllPackages {
				paths = append(paths, path.Join(gopath, "src", pkg.Pkg.Path()))
			}
			sort.Strings(paths)

			for i, p := range paths {
				assert.Equal(t, outpaths[i], p)
			}
		})

		// run against a directory outside of the gopath
		if err := os.Chdir(tmpdir); err != nil {
			t.Fatal(err)
		}

		// outside the $GOPATH
		t.Run(test.name+" outside gopath", func(t *testing.T) {
			var packages []string
			for _, pkg := range test.packages {
				packages = append(packages, strings.Replace(pkg, "$PWD", tmpdir, -1))
			}
			var outpaths []string
			for _, pkg := range test.outpaths {
				outpaths = append(outpaths, strings.Replace(pkg, "$PWD", tmpdir, -1))
			}

			p, err := loader.Load(&loader.Config{
				JoyPath:  path.Join(tmpdir, "testdata"),
				Packages: packages,
			})
			if err != nil {
				t.Fatal(err)
			}

			var paths []string
			for _, pkg := range p.AllPackages {
				paths = append(paths, path.Join(gopath, "src", pkg.Pkg.Path()))
			}
			sort.Strings(paths)

			for i, p := range paths {
				assert.Equal(t, outpaths[i], p)
			}
		})

		if err := os.Chdir(cwd); err != nil {
			t.Fatal(err)
		}
	}
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
