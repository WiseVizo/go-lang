package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://example.com/"

func main() {
	fmt.Println("web requests in go!")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The type of response is %T\n", res)
	defer res.Body.Close()
	databytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
