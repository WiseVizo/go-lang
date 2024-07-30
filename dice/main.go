package main

import (
	"fmt"
	"math/rand"
)

func main() {

	dice_num := 6
	for dice_num == 6 {
		fmt.Println("Rolling Dice!!!")
		dice_num = rand.Intn(6) + 1
		switch dice_num {
		case 1:
			fmt.Println(dice_num)
		case 2:
			fmt.Println(dice_num)

		case 3:
			fmt.Println(dice_num)

		case 4:
			fmt.Println(dice_num)

		case 5:
			fmt.Println(dice_num)
			fallthrough // fallthrough allows next case to automatically be true

		case 6:
			fmt.Println(dice_num)

		default:
			fmt.Println("wow")
		}
	}

}
