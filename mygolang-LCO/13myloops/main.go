package main

import "fmt"

func main() {
	fmt.Println("For loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Println("Index and days are: ", index, day)
	}
}
