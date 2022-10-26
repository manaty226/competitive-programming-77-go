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

func modulo(a, m int) int {
	return a % m
}

func isFirstVictory(grundy, A []int) bool {
	xorSum := grundy[modulo(A[0], 5)]
	for i := 1; i < len(A); i++ {
		xorSum ^= grundy[modulo(A[i], 5)]
	}
	return xorSum != 0
}

func main() {
	s := NewScanner()
	N := s.NextInt()
	_ = s.NextInt()
	_ = s.NextInt()

	A := make([]int, N)
	maxNum := 0
	for i := 0; i < N; i++ {
		A[i] = s.NextInt()
		if maxNum < A[i] {
			maxNum = A[i]
		}
	}

	grundy := []int{0, 0, 1, 1, 2}

	if isFirstVictory(grundy, A) {
		fmt.Printf("First")
	} else {
		fmt.Printf("Second")
	}

}
