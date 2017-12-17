package run_test

import (
	"bytes"
	"context"
	gobuild "go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/matthewmueller/joy/api/run"
	"github.com/matthewmueller/joy/internal/paths"
	"github.com/matthewmueller/joy/internal/testutil"
	"github.com/stretchr/testify/assert"

	"github.com/matthewmueller/joy/internal/mains"
)

func Test(t *testing.T) {
	ctx := context.Background()

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	gopath := gobuild.Default.GOPATH
	gosrc := path.Join(gopath, "src")

	dirs, err := ioutil.ReadDir(path.Join(cwd, "testdata"))
	if err != nil {
		t.Fatal(err)
	}

	tmp, err := ioutil.TempDir("", "run_test")
	if err != nil {
		t.Fatal(err)
	}

	// this is where joy's state lives
	// when it's not installed from source
	prefs, err := paths.Preferences()
	if err != nil {
		t.Fatal(err)
	}
	// remove all existing state
	if err := os.RemoveAll(prefs); err != nil {
		t.Fatal(err)
	}

	// find the go binary
	gobin, err := exec.LookPath("go")
	if err != nil {
		t.Fatal(err)
	}
	joyPath, err := paths.Joy()
	if err != nil {
		t.Fatal(err)
	}

	// copy all tests outside gopath for certain tests
	outsidePath := path.Join(tmp, "testdata")
	if err := testutil.Copy(path.Join(cwd, "testdata"), path.Join(tmp, "testdata")); err != nil {
		t.Fatal(err)
	}

	// rewrite our deps
	srcpaths, err := filepath.Glob(path.Join(outsidePath, "*", "*", "*.go"))
	if err != nil {
		t.Fatal(err)
	}
	for _, srcpath := range srcpaths {
		src, err := ioutil.ReadFile(srcpath)
		if err != nil {
			t.Fatalf("no '%s' file found", src)
		}
		rel, err := filepath.Rel(gosrc, path.Join(cwd, "testdata"))
		if err != nil {
			t.Fatal(err)
		}

		relgopath, err := filepath.Rel(gosrc, outsidePath)
		if err != nil {
			t.Fatal(err)
		}

		str := strings.Replace(string(src), rel, relgopath, -1)
		if err := ioutil.WriteFile(srcpath, []byte(str), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// create a binary of the latest joy
	joybin := path.Join(tmp, "joy")
	cmd := exec.Command(gobin, "build", "-o", joybin, path.Join(joyPath, "cmd", "joy", "main.go"))
	if err != nil {
		t.Fatal(err)
	}
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		// find all our package mains inside gopath
		mainpaths, err := mains.Find(path.Join(cwd, "testdata", dir.Name(), "..."))
		if err != nil {
			t.Fatal(err)
		}

		for _, main := range mainpaths {
			expected, err := ioutil.ReadFile(path.Join(main, "expected.txt"))
			if err != nil {
				t.Fatal(err)
			}

			// 1.
			t.Run(dir.Name()+" joy inside source inside", func(t *testing.T) {
				res, err := run.Run(&run.Config{
					FilePath: path.Join(main, "input.go"),
					JoyPath:  joyPath,
					Context:  ctx,
				})
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, string(expected), res, path.Join(main, "input.go"))
			})

			// 2.
			t.Run(dir.Name()+" joy outside source inside", func(t *testing.T) {
				cmd := exec.Command(joybin, "run", "--joy", prefs, main)
				buf := bytes.NewBufferString("")

				cmd.Stderr = os.Stderr
				cmd.Stdout = buf

				err := cmd.Run()
				if err != nil {
					t.Fatal(err)
				}

				assert.Equal(t, string(expected), strings.TrimSpace(buf.String()), path.Join(main, "input.go"))
			})
		}

		// find all our package mains inside gopath
		outsiders, err := mains.Find(path.Join(outsidePath, dir.Name(), "..."))
		if err != nil {
			t.Fatal(err)
		}

		for _, main := range outsiders {
			expected, err := ioutil.ReadFile(path.Join(main, "expected.txt"))
			if err != nil {
				t.Fatal(err)
			}

			// 3.
			t.Run(dir.Name()+" joy inside source outside", func(t *testing.T) {
				res, err := run.Run(&run.Config{
					FilePath: path.Join(main, "input.go"),
					JoyPath:  joyPath,
					Context:  ctx,
				})
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, string(expected), res, path.Join(main, "input.go"))
			})

			// 4.
			t.Run(dir.Name()+" joy outside source outside", func(t *testing.T) {
				cmd := exec.Command(joybin, "run", "--joy", prefs, main)
				buf := bytes.NewBufferString("")

				cmd.Stderr = os.Stderr
				cmd.Stdout = buf

				err := cmd.Run()
				if err != nil {
					t.Fatal(err)
				}

				assert.Equal(t, string(expected), strings.TrimSpace(buf.String()), path.Join(main, "input.go"))
			})
		}

	}
}
