package chrome_test

import (
	"context"
	"testing"

	"github.com/matthewmueller/golly/chrome"
	"github.com/stretchr/testify/assert"
)

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	ch, err := chrome.New(ctx, &chrome.Settings{
		Image: "yukinying/chrome-headless-browser:latest",
	})
	if err != nil {
		t.Fatal(err)
	}

	cancel()
	e := ch.Wait()
	assert.EqualError(t, e, "container exited with error code: 137")
}

func TestClose(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch, err := chrome.New(ctx, &chrome.Settings{
		Image: "yukinying/chrome-headless-browser:latest",
	})
	if err != nil {
		t.Fatal(err)
	}

	if e := ch.Close(); e != nil {
		t.Fatal(e)
	}
}

func TestLeaseMultiple(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch, err := chrome.New(ctx, &chrome.Settings{
		Image: "yukinying/chrome-headless-browser:latest",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer ch.Close()

	// lease once
	tar, err := ch.Lease(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if e := tar.Close(); e != nil {
		t.Fatal(e)
	}

	// lease twice
	tar, err = ch.Lease(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if e := tar.Close(); e != nil {
		t.Fatal(e)
	}
}
