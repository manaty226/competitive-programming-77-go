package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	scanner.Scan()
	array := scan(scanner)

	for i, x := range array {
		for _, y := range array[i+1:] {
			for _, z := range array[i+2:] {
				if x+y+z == 1000 {
					fmt.Printf("Yes")
					return
				}
			}
		}
	}
	fmt.Printf("No")
}
