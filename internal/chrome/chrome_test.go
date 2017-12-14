package chrome_test

import (
	"context"
	"io/ioutil"
	"path"
	"testing"

	"github.com/matthewmueller/joy/internal/paths"

	"github.com/matthewmueller/joy/internal/chrome"
	"github.com/matthewmueller/log"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	ctx := context.Background()

	root, err := paths.Joy()
	if err != nil {
		t.Fatal(err)
	}

exists:
	chromePath, err := chrome.Exists(path.Join(root, "chrome"))
	if err != nil {
		t.Fatal(err)
	} else if chromePath == "" {
		log.Infof("downloading headless chrome (this only needs to be done once)")
		if err := chrome.Download(path.Join(root, "chrome")); err != nil {
			t.Fatal(err)
		}
		goto exists
	}

	ch, err := chrome.New(ctx, &chrome.Settings{
		ExecutablePath: chromePath,
		Stderr:         ioutil.Discard,
		Stdout:         ioutil.Discard,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer ch.Close()

	target, err := ch.Target()
	if err != nil {
		t.Fatal(err)
	}

	result, err := target.Run(`
		console.log("hiya", 5)
	`)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "hiya 5", result)

	if e := target.Close(); e != nil {
		t.Fatal(e)
	}
	if e := ch.Close(); e != nil {
		t.Fatal(e)
	}
}
