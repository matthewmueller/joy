package golly_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/matthewmueller/golly"
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
	log.SetHandler(text.New(os.Stderr))

	dirs, err := ioutil.ReadDir("./testdata")
	if err != nil {
		t.Fatal(err)
	}

	cwd, e := os.Getwd()
	if e != nil {
		t.Fatal(e)
	}

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

			// compile the file
			src, e := golly.Compile(path.Join(cwd, "testdata", dir.Name()))
			if e != nil {
				t.Fatal(e)
			}

			jspath := path.Join(cwd, "testdata", dir.Name(), "expected.js.txt")
			js, e := ioutil.ReadFile(jspath)
			if e != nil {
				t.Fatalf("no '%s' file found", jspath)
			}

			// compile the code
			if src != string(js) {
				t.Fatal(fmt.Sprintf("\n## Expected ##\n\n%s\n\n## Actual ##\n\n%s", string(js), src))
			}

			resultpath := path.Join(cwd, "testdata", dir.Name(), "expected.txt")
			if _, err := os.Stat(resultpath); os.IsNotExist(err) {
				log.Warnf("'%s' doesn't exist, skipping node execution", resultpath)
				return
			}

			// run the code
			cmd := exec.Command("./nodejs/node", jspath)
			buf, e := cmd.CombinedOutput()
			if e != nil {
				t.Fatalf("javascript error: %s", buf)
			}

			result, e := ioutil.ReadFile(resultpath)
			if e != nil {
				t.Fatalf("no '%s' file found", resultpath)
			}

			// compare the result path
			if string(buf) != string(result) {
				t.Fatal(fmt.Sprintf("\n## Expected ##\n\n%s\n\n## Actual ##\n\n%s", string(result), string(buf)))
			}
		})
	}
}
