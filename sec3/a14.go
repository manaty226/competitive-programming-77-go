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

func nextIntArray(s *Scanner, N int) []int {
	array := make([]int, N)
	for i := 0; i < N; i++ {
		array[i] = s.NextInt()
	}
	return array
}

func main() {
	s := NewScanner()
	N := s.NextInt()
	K := s.NextInt()

	A := nextIntArray(s, N)
	B := nextIntArray(s, N)
	C := nextIntArray(s, N)
	D := nextIntArray(s, N)

	Q := map[int]bool{}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			Q[C[i]+D[j]] = true
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if _, ok := Q[K-A[i]-B[j]]; ok {
				fmt.Printf("Yes")
				return
			}
		}
	}
	fmt.Printf("No")
}
