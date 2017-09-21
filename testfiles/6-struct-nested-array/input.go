package main

// User struct
type User struct {
	FirstName string
	LastName  string
	Phones    []Phone
}

// Phone struct
type Phone struct {
	User
	Type   string
	Number string
	Main   bool
}

func main() {
	user := User{
		FirstName: "Matt",
		LastName:  "Mueller",
		Phones: []Phone{
			Phone{
				Type:   "HOME",
				Number: "1234123",
			},
			Phone{
				Type:   "CELL",
				Number: "0923123",
				Main:   true,
			},
		},
	}

	println(user.FirstName + ": " + user.Phones[0].Number)
}
