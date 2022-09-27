package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func answer(a, b int) int {
	return a + b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	numArray := strings.Split(text, " ")
	a, err := strconv.Atoi(numArray[0])
	if err != nil {
		return
	}
	b, err := strconv.Atoi(numArray[1])
	if err != nil {
		return
	}
	fmt.Println(answer(a, b))
}
