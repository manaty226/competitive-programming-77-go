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
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return &Scanner{scanner}
}

func (s *Scanner) NextInt() int {
	s.Scan()
	word := s.Text()
	n, _ := strconv.Atoi(word)
	return n
}

func (s *Scanner) NextString() string {
	s.Scan()
	return s.Text()
}

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sumOfWorkHours(W []int) int {
	sum := 0
	for _, w := range W {
		sum += w
	}
	return sum
}

func main() {
	s := NewScanner()
	D := s.NextInt()
	N := s.NextInt()

	W := make([]int, D)
	for i := 0; i < D; i++ {
		W[i] = 24
	}

	for i := 0; i < N; i++ {
		L := s.NextInt()
		R := s.NextInt()
		H := s.NextInt()
		for j := L - 1; j < R; j++ {
			W[j] = minInts(W[j], H)
		}
	}

	fmt.Printf("%d \n", sumOfWorkHours(W))
}
