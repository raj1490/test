package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This file needs to go in the server."

	file, err := os.Create("./mylcogofile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text inside the data file is: ", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
