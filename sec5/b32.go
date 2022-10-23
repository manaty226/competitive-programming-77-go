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
	K := s.NextInt()

	A := make([]int, K)
	for i := 0; i < K; i++ {
		A[i] = s.NextInt()
	}

	dp := make([]bool, N+1)
	for i := 0; i <= N; i++ {
		for j := 0; j < K; j++ {
			if i >= A[j] && dp[i-A[j]] == false {
				dp[i] = true
			}
		}
	}

	if dp[N] {
		fmt.Printf("First")
	} else {
		fmt.Printf("Second")
	}
}
