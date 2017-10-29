package main

import (
	"encoding/json"

	"github.com/matthewmueller/golly/testdata/45-external-fetch/fetch"
)

func main() {
	test()
}

// Repo struct
type Repo struct {
	Name string `js:"name"`
}

func test() {
	response, err := fetch.Fetch("https://api.github.com/users/matthewmueller/repos", &fetch.Options{
		Method: "GET",
	})
	if err != nil {
		println(err)
		return
	}

	var repos []Repo
	data := response.Text()
	if e := json.Unmarshal(data, &repos); e != nil {
		panic(e)
	}

	for _, repo := range repos {
		if repo.Name == ".dotfiles" {
			println(repo.Name)
		}
	}
}
