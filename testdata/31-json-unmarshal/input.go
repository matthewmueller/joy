package main

import "encoding/json"

// Person struct
type Person struct {
	Name string `json:"name,omitempty" js:"name"`
	Age  int    `json:"age,omitempty" js:"age"`
}

func main() {
	var p Person
	if e := json.Unmarshal([]byte(`{"name":"matt","age":28}`), &p); e != nil {
		panic(e)
	}
	println(p.Name)
}
