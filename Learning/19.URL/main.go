package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3002/learn?coursename=reactjs&paymentid=ghbj456"

func main() {
	fmt.Println("URLs in Golang")
	fmt.Println(myUrl)

	// parsing the url
	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println(qparams["coursename"])

	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsofUrl.String()
	fmt.Println(anotherUrl)
}
