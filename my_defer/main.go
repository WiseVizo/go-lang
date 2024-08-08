package main

import "fmt"

// defer follows the first in last out stack for storing defered lines and call them at the end of fn

func main() {
	defer fmt.Println("i got defered 1st")
	fmt.Print("defer in go lang\n")
	defer fmt.Print("i got defered 2nd\n")
	my_defer()
}

func my_defer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}
