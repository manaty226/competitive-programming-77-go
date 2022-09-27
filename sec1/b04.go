package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	var n int
	for i := 0; i < len(text); i++ {
		n += int(text[i]-48) * (int(1 << (len(text) - i - 1)))
	}
	fmt.Printf("%d", n)
}
