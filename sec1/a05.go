package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	text := scanner.Text()
	array := strings.Split(text, " ")
	N, _ := strconv.Atoi(array[0])
	K, _ := strconv.Atoi(array[1])

	var count int
	for i := 1; i <= N && i <= K-2; i++ {
		for j := 1; j <= N && i+j <= K-1; j++ {
			if (K - i - j) <= N {
				count++
			}
		}
	}
	fmt.Printf("%d", count)
}
