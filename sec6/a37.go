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

func main() {
	s := NewScanner()
	N := s.NextInt()
	M := s.NextInt()
	B := s.NextInt()

	sumOfA, sumObC := 0, 0
	for i := 0; i < N; i++ {
		sumOfA += s.NextInt() * M
	}
	for i := 0; i < M; i++ {
		sumObC += s.NextInt() * N
	}
	sumOfB := B * M * N

	fmt.Printf("%d \n", sumOfA+sumOfB+sumObC)

}
