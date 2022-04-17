package main

import "fmt"

type User struct {
	Name    string
	Email   string
	Country string
	age     int
}

func main() {
	// create user using the above struct
	user1 := User{
		Name:    "Milu",
		Email:   "milu@af.com",
		Country: "Narnia",
		age:     123,
	}
	fmt.Printf("%+v\n", user1)

	// create user using new(User)
	var user2 = new(User)
	user2.Name = "James"
	user2.Email = "james@you.com"
	user2.Country = "USA"
	user2.age = 54
	fmt.Printf("%+v\n", user2)
}
