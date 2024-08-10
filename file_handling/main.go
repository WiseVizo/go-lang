package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("file handling in go lang")
	// content := "hi can i really write in a file?\n"
	// file, err := os.Create("myfile.txt")
	// if err != nil {
	// 	fmt.Println("bye bye")
	// 	panic(err)
	// }
	// length, err := io.WriteString(file, content)
	// if err != nil {
	// 	fmt.Println("sed")
	// 	panic(err)
	// }
	// fmt.Println("lenght of file is: ", length)

	// defer file.Close()
	read_file("myfile.txt")
}

func read_file(file_path string) {
	databyte, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	fmt.Println("file content: ", databyte)
}
