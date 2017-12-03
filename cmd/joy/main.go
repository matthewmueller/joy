package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/matthewmueller/joy/internal/cli"

	"github.com/apex/log"
	logcli "github.com/apex/log/handlers/cli"
)

var version = "master"

func main() {
	ctx := trap(syscall.SIGINT, syscall.SIGTERM)
	log.SetHandler(logcli.Default)

	err := cli.Run(ctx, version)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}

func trap(sig ...os.Signal) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
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
