package main

import "fmt"

func main() {
	fmt.Println("maps in go lang!")
	var my_map = make(map[string]string)
	my_map["en"] = "english"
	my_map["jp"] = "japanese"
	fmt.Println("my map:", my_map)
	fmt.Println("en:", my_map["en"])
	delete(my_map, "jp") // works for slice too
	fmt.Println("my map:", my_map)
}
