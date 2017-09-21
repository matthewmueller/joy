package main

// User struct
type User struct {
	FirstName string
	LastName  string
	Phones    []Phone
}

// Phone struct
type Phone struct {
	Type   string
	Number string
	Main   bool
}

func main() {
	user := User{
		FirstName: "Matt",
		Phones: []Phone{
			Phone{
				Type:   "HOME",
				Number: "1234125",
				Main:   false,
			},
			Phone{
				Type:   "CELL",
				Number: "1231411",
				Main:   true,
			},
		},
	}
	println(user.Phones[1].Number)
}
