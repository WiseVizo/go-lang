package main

import "fmt"

func division(numerator, denominator int) (float64, error) {
	if denominator <= 0 {
		return -1, fmt.Errorf("denominator can't be zero or smaller")
	}
	return float64(numerator) / float64(denominator), nil
}

func main() {
	fmt.Println(division(10, 0))
}
