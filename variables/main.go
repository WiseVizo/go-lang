package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome!")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	numinput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err == nil {
		fmt.Println("ans is", numinput+1)
	} else {
		fmt.Println("err: ", err)
	}
	fmt.Println("thx for raiting of", input)

}
