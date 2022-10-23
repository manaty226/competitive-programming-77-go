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
	A := s.NextInt()
	B := s.NextInt()

	dp := make([]bool, N+1)

	for i := 0; i <= N; i++ {
		if i >= A && dp[i-A] == false {
			dp[i] = true
		} else if i >= B && dp[i-B] == false {
			dp[i] = true
		} else {
			dp[i] = false
		}
	}

	if dp[N] {
		fmt.Printf("First")
	} else {
		fmt.Printf("Second")
	}
}
