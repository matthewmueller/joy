package main

// User struct
type User struct {
	FirstName string
	LastName  string
	Phone     Phone
}

// Phone struct
type Phone struct {
	Type   string
	Number string
	Main   bool
}

func main() {
	user := User{}
	println(user.Phone.Number)
}
