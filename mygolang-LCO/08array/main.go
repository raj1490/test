package main

import "fmt"

func main() {
	var arrayList [4]string
	arrayList[0] = "Apple"
	arrayList[1] = "mango"
	arrayList[3] = "banana"

	fmt.Println("The array list is: ", arrayList)

	var vegList = [3]string{"tomato", "potato", "ladyfinger"}
	fmt.Println("The vegList is ", vegList)
}
