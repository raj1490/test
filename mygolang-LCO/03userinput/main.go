package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of the pizza: ")

	//comma ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanx for the rating: ", input)
}
