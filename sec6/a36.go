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

func isReachable(N, K int) bool {
	if K > 2*N-2 && K%2 == 0 {
		return true
	}
	return false
}

func main() {
	s := NewScanner()
	N := s.NextInt()
	K := s.NextInt()

	if isReachable(N, K) {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}

}
