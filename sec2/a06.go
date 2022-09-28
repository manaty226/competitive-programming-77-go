package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	return &Scanner{
		bufio.NewScanner(os.Stdin),
	}
}

func (s *Scanner) nextInt() int {
	s.Scan()
	text := s.Text()
	n, _ := strconv.Atoi(text)
	return n
}

func main() {
	scanner := NewScanner()
	scanner.Split(bufio.ScanWords)

	N, Q := scanner.nextInt(), scanner.nextInt()

	sumArray := make([]int, N+1)
	for i := 0; i < N; i++ {
		sumArray[i+1] = sumArray[i] + scanner.nextInt()
	}

	for i := 0; i < Q; i++ {
		start, end := scanner.nextInt(), scanner.nextInt()
		fmt.Printf("%d\n", sumArray[end]-sumArray[start-1])
	}
}
