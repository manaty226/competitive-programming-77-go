package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Scanner struct {
	*bufio.Scanner
	*bufio.Reader
}

func NewScanner() *Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return &Scanner{scanner, reader}
}

func (s *Scanner) NextInt() int {
	s.Scan()
	word := s.Text()
	n, _ := strconv.Atoi(word)
	return n
}

func (s *Scanner) NextString() string {
	// line := ""
	// for {
	// 	l, p, _ := s.ReadLine()
	// 	line += string(l)
	// 	if !p {
	// 		break
	// 	}
	// }
	// return line
	s.Scan()
	return s.Text()
}

func maxInts(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func sumOfHeights(H []int) int {
	sum := 0
	for _, h := range H {
		sum += h
	}
	return sum
}

func main() {
	s := NewScanner()
	N := s.NextInt()
	S := s.NextString()

	H := make([]int, N)
	for i := 0; i < N; i++ {
		H[i] = 1
	}

	for i, s := range S {
		if string(s) == "A" {
			H[i+1] = H[i] + 1
		}
	}

	for i := N - 2; i >= 0; i-- {
		if string(S[i]) == "B" {
			H[i] = maxInts(H[i+1]+1, H[i])
		}
	}

	fmt.Printf("%d \n", sumOfHeights(H))
}
