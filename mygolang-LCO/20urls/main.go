package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghd675bij"

func main() {
	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)

	// fmt.Println(result.RawQuery) //we can grab all the queries or we can say params
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	fmt.Println(result.Scheme)

	//to extract more specifically information
	params := result.Query() //data stored in key value pairs

	fmt.Println(params["paymentid"])

	for index, val := range params {
		fmt.Println(index, "->", val)
	}

	//to make url from parts of url
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}
