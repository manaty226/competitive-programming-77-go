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

func NumDivisers(N int) int {
	return (N / 3) + (N / 5) + (N / 7) - (N / 15) - (N / 21) - (N / 35) + (N / (3 * 5 * 7))
}

func main() {
	s := NewScanner()
	N := s.NextInt()
	answer := NumDivisers(N)
	fmt.Printf("%d \n", answer)
}
