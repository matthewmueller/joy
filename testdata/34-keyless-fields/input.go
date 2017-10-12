package main

// Phone struct
type Phone struct {
	Type   string
	Number string `js:"number"`
}

// User struct
type User struct {
	Name  string `js:"name"`
	Age   int
	Phone *Phone `js:"phone"`
}

func main() {
	phone := Phone{"Android", "8674205"}
	age := 28
	u := User{"matt", age, &phone}
	println(u.Name, u.Age, u.Phone.Number)
}
