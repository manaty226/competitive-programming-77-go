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

const diviser int64 = 1000000007

func Calculator() func(n int) int64 {
	var a, a1, a2 int64 = 1, 1, 1
	return func(n int) int64 {
		for i := 2; i < n; i++ {
			a = (a1 + a2) % diviser
			a2 = a1
			a1 = a
		}
		return a
	}
}

func main() {
	s := NewScanner()
	N := s.NextInt()

	c := Calculator()
	answer := c(N)
	fmt.Printf("%d \n", answer)
}
