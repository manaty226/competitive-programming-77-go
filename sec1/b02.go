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
	numArray := strings.Split(text, " ")

	min, _ := strconv.Atoi(numArray[0])
	max, _ := strconv.Atoi(numArray[1])

	for i := min; i <= max; i++ {
		if 100%i == 0 {
			fmt.Printf("Yes")
			return
		}
	}
	fmt.Printf("No")
}
