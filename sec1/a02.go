package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func linearSearch(array []int, target int) bool {
	for _, e := range array {
		if target == e {
			return true
		}
	}
	return false
}

func scan(scanner *bufio.Scanner) []int {
	scanner.Scan()
	text := scanner.Text()
	array := strings.Split(text, " ")

	var numArray []int
	for _, e := range array {
		n, err := strconv.Atoi(e)
		if err != nil {
			return []int{}
		}
		numArray = append(numArray, n)
	}
	return numArray
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	targetArray := scan(scanner)
	numArray := scan(scanner)

	if ok := linearSearch(numArray, targetArray[1]); ok {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}

}
