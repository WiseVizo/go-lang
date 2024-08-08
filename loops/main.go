package main

import "fmt"

func main() {
	fmt.Println("loops in go lang")
	my_list := []int{6, 7, 8, 9}
	for i := range my_list {
		fmt.Println("item is:", my_list[i])
	}
	// for each
	for i, item := range my_list {
		fmt.Printf("index is %d and item is %d\n", i, item)
	}

	// while
	index := 0
	for index < 10 {
		if index == 5 {
			index++
			continue
		}
		if index == 8 {
			goto checkpoint
		}
		fmt.Println(index)
		index++
	}

checkpoint:
	fmt.Println("at checkpoint")

}
