package chrome_test

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/matthewmueller/golly/chrome"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	ctx := context.Background()
	ch, err := chrome.New(ctx, &chrome.Settings{
		ExecutablePath: os.Getenv("GOLLY_CHROME_PATH"),
		Stdout:         ioutil.Discard,
		Stderr:         ioutil.Discard,
	})
	if err != nil {
		t.Fatal(err)
	}

	target, err := ch.Target()
	if err != nil {
		t.Fatal(err)
	}

	result, err := target.Run(`
		new Promise(function (resolve, reject) {
			console.log("hiya", 5)
			setTimeout(() => {
				console.log("hello")
				setTimeout(() => {
					resolve("hi world!")
				}, 300)
			}, 300)
		})
	`)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "hiya 5\nhello", result)

	if e := target.Close(); e != nil {
		t.Fatal(e)
	}
	if e := ch.Close(); e != nil {
		t.Fatal(e)
	}
}
