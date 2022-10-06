package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case")
	//to genrate random numbers
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("The number at dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Yeah you can either open or move by 1")
	case 2:
		fmt.Println("you can move by 2")
	case 3:
		fmt.Println("you can move by 3")
	case 4:
		fmt.Println("you can move by 4")
	case 5:
		fmt.Println("you can move by 5")
	case 6:
		fmt.Println("you can move by 6 and roll the dice again")
	}
}
