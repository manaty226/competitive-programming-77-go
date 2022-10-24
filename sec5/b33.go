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

func isFirstVictory(A, B []int) bool {
	xorSum := A[0] ^ B[0]
	for i := 1; i < len(A); i++ {
		xorSum ^= A[i] ^ B[i]
	}
	return xorSum != 0
}

func main() {
	s := NewScanner()
	N := s.NextInt()
	_ = s.NextInt()
	_ = s.NextInt()

	A := make([]int, N)
	B := make([]int, N)

	for i := 0; i < N; i++ {
		A[i] = s.NextInt() - 1
		B[i] = s.NextInt() - 1
	}

	if isFirstVictory(A, B) {
		fmt.Printf("First")
	} else {
		fmt.Printf("Second")
	}

}
