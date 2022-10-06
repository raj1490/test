package main

import "fmt"

func main() {
	fmt.Println("mystructs")

	raju := User{"Raj", "Raj@zocket.com", true, 21}
	fmt.Println(raju)
	fmt.Printf("raju goes like: %+v \n", raju)
}

type User struct {
	Name   string
	Email  string
	status bool
	age    int
}
