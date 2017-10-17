package golly_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/matthewmueller/golly/api"
	"github.com/matthewmueller/golly/chrome"
	"github.com/pkg/errors"
)

// little invisibles helper
func invisibles(str string) string {
	str = strings.Replace(str, " ", "·", -1)
	str = strings.Replace(str, "\r", "¬", -1)
	str = strings.Replace(str, "\n", "¬", -1)
	str = strings.Replace(str, "\t", "‣", -1)
	return str
}

func Test(t *testing.T) {
	ctx := context.Background()
	log.SetHandler(text.New(os.Stderr))

	dirs, err := ioutil.ReadDir("./testdata")
	if err != nil {
		t.Fatal(err)
	}

	cwd, e := os.Getwd()
	if e != nil {
		t.Fatal(e)
	}

	gosrc := path.Join(os.Getenv("GOPATH"), "src")

	ch, err := chrome.New(ctx, &chrome.Settings{
		ExecutablePath: os.Getenv("GOLLY_CHROME_PATH"),
		Stderr:         ioutil.Discard,
		Stdout:         ioutil.Discard,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer ch.Close()

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		// subtests
		t.Run(dir.Name(), func(t *testing.T) {
			if strings.HasPrefix(dir.Name(), "_") {
				log.Warnf("Skipping '%s'", dir.Name())
				return
			}

			testdir := path.Join(cwd, "testdata", dir.Name())
			var inputs []string

			// support single and multipage tests
			inps, e := filepath.Glob(path.Join(testdir, "input.go"))
			if e != nil {
				t.Fatal(e)
			}

			inputs = append(inputs, inps...)
			inps, e = filepath.Glob(path.Join(testdir, "*", "input.go"))
			if e != nil {
				t.Fatal(e)
			}
			inputs = append(inputs, inps...)

			var pages []string
			for _, input := range inputs {
				pages = append(pages, path.Dir(input))
			}

			// compile the file
			scripts, e := api.Build(ctx, &api.BuildSettings{
				Packages: pages,
			})
			if e != nil {
				t.Fatal(errors.Wrap(e, "compile error"))
			}
			if len(scripts) == 0 {
				t.Fatal("expected atleast 1 script to be built")
			}

			for _, script := range scripts {
				jspath := path.Join(gosrc, script.Name(), "expected.js.txt")
				resultpath := path.Join(gosrc, script.Name(), "expected.txt")

				// read the expected js source
				js, err := ioutil.ReadFile(jspath)
				if err != nil {
					t.Fatalf("no '%s' file found", jspath)
				}

				// compile the code
				if script.Source() != string(js) {
					// if err := ioutil.WriteFile(jspath, []byte(file.Source), 0755); err != nil {
					// 	t.Fatal(err)
					// }
					t.Fatal(fmt.Sprintf("\n## Expected ##\n\n%s\n\n## Actual ##\n\n%s", string(js), script.Source()))
				}

				// try reading the result path
				if _, err := os.Stat(resultpath); os.IsNotExist(err) {
					log.Warnf("'%s' doesn't exist, skipping node execution", resultpath)
					return
				}

				// create a target
				tar, err := ch.Target()
				if err != nil {
					t.Fatal(err)
				}

				// run the code in a headless chrome target
				actual, err := tar.Run(script.Source())
				if err != nil {
					t.Fatal(err)
				}

				expected, e := ioutil.ReadFile(resultpath)
				if e != nil {
					t.Fatalf("no '%s' file found", resultpath)
				}

				// compare the result path
				if strings.TrimSpace(string(actual)) != strings.TrimSpace(string(expected)) {
					t.Fatal(fmt.Sprintf("\n## Expected ##\n\n%s\n\n## Actual ##\n\n%s", string(expected), string(actual)))
				}

				if e := tar.Close(); e != nil {
					t.Fatal(e)
				}
			}
		})
	}
}
