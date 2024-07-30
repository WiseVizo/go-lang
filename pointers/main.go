package main

import "fmt"

func main() {
	fmt.Println("pointers 101!")
	num := 30
	ptr := &num
	fmt.Println(ptr)
	fmt.Println(*ptr + 10)
	*ptr += 10
	fmt.Println(num)
}
