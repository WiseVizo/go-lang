package main

import "fmt"

type Food struct {
	Name      string
	Price     float32
	Quantity  int64
	Avaibalbe bool
}

// to create methods for structures we need to pass the copy of struct of given type in a fn

func (food Food) Is_Available() bool { // any updation of value within struct method won't reflect in original struct
	return food.Avaibalbe
}

func main() {
	fmt.Println("struct in golang!")
	var food1 Food = Food{"apple", 3.99, 5, true}
	fmt.Println("food1:", food1)
	fmt.Printf("food1 details are: %+v\n", food1)
	fmt.Printf("food1 name is %v, price is %v\n", food1.Name, food1.Price)
	fmt.Println(food1.Is_Available())
}
