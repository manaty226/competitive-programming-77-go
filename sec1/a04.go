package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	n, _ := strconv.Atoi(text)

	var answer string
	for n != 0 {
		nn := n % 2
		answer = fmt.Sprintf("%d%s", nn, answer)
		n /= 2
	}
	n, _ = strconv.Atoi(answer)
	fmt.Printf("%010d", n)
}
