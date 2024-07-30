package main

import "fmt"

func main() {
	// arrays
	var mylist [5]int
	mylist[0] = 1
	mylist[2] = 2
	fmt.Println("array->", mylist)
	fmt.Println("len of array:", len(mylist))

	// slices
	var myslices []int = []int{55, 66, 77}
	fmt.Println("slice->", myslices)
	fmt.Println("len of slice:", len(myslices))
	myslices = append(myslices, 77, 88)
	fmt.Println("slice 2.0->", myslices)
	fmt.Println("len of slice 2.0:", len(myslices))

	slice := make([]int, 5) // Creates a slice with length 5 and capacity 5
	fmt.Println(slice)

	mymarks := []int{20, 30, 40, 50}
	mymarks = append(mymarks[:1], mymarks[2:]...)
	fmt.Println(mymarks)
}
