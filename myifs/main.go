package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome! enter a number")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	numinput, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err == nil {
		if numinput%2 == 0 {
			fmt.Println(numinput, "is even number")
		} else {
			fmt.Println(numinput, "is odd number")
		}
	} else {
		fmt.Println("error sed")
	}
}
