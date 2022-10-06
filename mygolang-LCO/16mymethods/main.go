package main

import "fmt"

func main() {
	raj := User{"Raj", "Raj@dev.go", 21, true}

	fmt.Println(raj)
	raj.getName()
}

type User struct {
	name   string
	email  string
	age    int
	status bool
}

//methods
func (u User) getName() {
	fmt.Println("Name of the user is: ", u.name)
}
