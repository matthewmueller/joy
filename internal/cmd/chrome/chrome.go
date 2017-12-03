package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/matthewmueller/joy/internal/chrome"
)

func main() {
	log.SetHandler(text.New(os.Stderr))

	cwd, err := os.Getwd()
	if err != nil {
		log.WithError(err).Fatalf("couldn't get cwd")
	}

	ctx := signalContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	chromePath := os.Getenv("GOLLY_CHROME_PATH")
	filepath := path.Join(cwd, os.Args[1])
	source, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.WithError(err).Fatalf("couldn't read file: %s", filepath)
	}

	ch, err := chrome.New(ctx, &chrome.Settings{
		ExecutablePath: chromePath,
	})
	if err != nil {
		log.WithError(err).Fatalf("unable start chrome")
	}
	defer ch.Close()

	tar, err := ch.Target()
	if err != nil {
		log.WithError(err).Fatalf("unable create target")
	}
	defer tar.Close()

	// run the source code
	if res, err := tar.Run(string(source)); err != nil {
		log.Fatalf(err.Error())
	} else {
		fmt.Println(res)
	}
}

func signalContext(ctx context.Context, sig ...os.Signal) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, sig...)
		defer signal.Stop(c)

		select {
		case <-ctx.Done():
		case <-c:
			cancel()
		}
	}()

	return ctx
}
