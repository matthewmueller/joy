package main

import "encoding/json"

// Person struct
type Person struct {
	Name string `json:"name,omitempty" js:"name"`
	Age  int    `json:"age,omitempty" js:"age"`
}

func main() {
	bytes, err := json.Marshal(&Person{
		Name: "Matt",
		Age:  28,
	})
	if err != nil {
		panic(err)
	}
	println(bytes)
}
