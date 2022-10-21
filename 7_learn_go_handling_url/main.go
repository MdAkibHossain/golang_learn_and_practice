package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://akib.dev:9000/learn?coursename=Flutter&paymentid=ghj456hb"

func main() {
	fmt.Println(myUrl)
	result, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	qparams := result.Query()
	fmt.Printf("The types of query params are %T\n", qparams)
	fmt.Println(qparams["coursename"])
	for _, val := range qparams {
		fmt.Println("Parms is :", val)
	}

	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "ak.dev",
		Path:    "about",
		RawPath: "user=akib",
	}
	url := partsofUrl.String()
	fmt.Println(url)

}
