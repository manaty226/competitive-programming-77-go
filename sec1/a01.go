package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func answer(n int) int {
	return n * n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}
	fmt.Println(answer(n))
}
