package main

import (
	"github.com/matthewmueller/joy/testdata/57-rewrite-file-dep/fetch"
)

// TODO: move tests off external dependencies
// var api = "https://hacker-news.firebaseio.com"

// Repo struct
type Repo struct {
	Name string `js:"name"`
}

func main() {
	response, err := fetch.Get("https://api.github.com/users/matthewmueller/repos")
	if err != nil {
		println(err)
		return
	}

	var repos []Repo
	if e := response.JSON(&repos); e != nil {
		panic(e)
	}

	for _, repo := range repos {
		if repo.Name == ".dotfiles" {
			println(repo.Name)
		}
	}
}
