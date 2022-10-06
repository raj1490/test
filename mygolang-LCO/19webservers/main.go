package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://lco.dev"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)

	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println(content)
}
