package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func nextInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	text := scanner.Text()
	n, _ := strconv.Atoi(text)
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	N, Q := nextInt(scanner), nextInt(scanner)

	sumArray := make([]int, N+1)
	sumArray = append(sumArray, 0)
	for i := 0; i < N; i++ {
		sumArray[i+1] = sumArray[i] + nextInt(scanner)
	}

	for i := 0; i < Q; i++ {
		start, end := nextInt(scanner), nextInt(scanner)
		fmt.Printf("%d\n", sumArray[end]-sumArray[start-1])
	}
}
