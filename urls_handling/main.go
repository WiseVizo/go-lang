package main

import (
	"fmt"
	"net/url"
)

const my_url string = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("url handling in go :)")
	fmt.Println(my_url)
	result, _ := url.Parse(my_url)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	query_params := result.Query()
	fmt.Println(query_params["v"])
	fmt.Println(query_params["list"])
	fmt.Println(query_params["index"])

}
