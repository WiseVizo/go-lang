package main

import "fmt"

type Food struct {
	Name      string
	Price     float32
	Quantity  int64
	Avaibalbe bool
}

func main() {
	fmt.Println("struct in golang!")
	var food1 Food = Food{"apple", 3.99, 5, true}
	fmt.Println("food1:", food1)
	fmt.Printf("food1 details are: %+v\n", food1)
	fmt.Printf("food1 name is %v, price is %v", food1.Name, food1.Price)
}
