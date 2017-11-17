package main

import "github.com/apex/log"

type Track interface {
	Track()
}

type track struct{}

func (*track) Track() {}

type Video struct {
	*track
}

type Audio struct {
	*track
}

type Text struct {
	*track
}

func Something() Track {
	return &Video{}
}

func main() {
	s := Something()
	log.Infof("%T", s)
}
