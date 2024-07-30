package main

import "fmt"

func main() {
	var my_array [5]int
	for i := 0; i < 5; i++ {
		my_array[i] = i
	}
	fmt.Println("my array: ", my_array)
	fmt.Println("my array length:", len(my_array))
}
