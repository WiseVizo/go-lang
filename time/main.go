package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome!")
	var current_time time.Time = time.Now()
	fmt.Println("time: ", current_time)
	fmt.Printf("type: %T", current_time)
	fmt.Println(current_time.Format("01-02-2006 15-04-05 Monday"))
	new_date := time.Date(2022, time.February, 27, 0, 0, 0, 0, time.Local)
	fmt.Println(new_date.Format("Monday"))
}
