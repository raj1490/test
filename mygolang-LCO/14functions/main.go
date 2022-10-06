package main

import "fmt"

func main() {
	fmt.Println("Functions")
	fmt.Println("The sum is: ", adder(5, 3))
	fmt.Println("The sum is: ", proAdder(5, 5, 5, 5, 6, 4))
	greeting()
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	result := 0

	for _, val := range values {
		result += val
	}

	return result
}

func greeting() {
	fmt.Println("Hello from golang")
}
