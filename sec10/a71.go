package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Scanner struct {
	*bufio.Scanner
	*bufio.Reader
}

func NewScanner() *Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 10000)
	scanner.Buffer(buf, 256*1000)

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
	s.Scan()
	return s.Text()
}

func main() {
	s := NewScanner()
	N := s.NextInt()

	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = s.NextInt()
	}
	for i := 0; i < N; i++ {
		B[i] = s.NextInt()
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	sort.Slice(B, func(i, j int) bool {
		return B[i] > B[j]
	})

	answer := 0
	for i := 0; i < N; i++ {
		answer += A[i] * B[i]
	}
	fmt.Println(answer)
}
