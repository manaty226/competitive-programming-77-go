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
	s.Scan()
	return s.Text()
}

func maxInts(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	s := NewScanner()
	N := s.NextInt()

	M := map[int]int{}
	for i := 0; i < N; i++ {
		A := s.NextInt()
		if _, ok := M[A]; !ok {
			M[A] = 1
		} else {
			M[A]++
		}
	}

	answer := 0
	for _, v := range M {
		answer += v * (v - 1) * (v - 2) / 6
	}

	fmt.Printf("%d \n", answer)
}
