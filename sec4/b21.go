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

func maxInts(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := NewScanner()
	N := s.NextInt()
	S := s.NextString()

	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, N+1)
		dp[i][i] = 1
	}

	for l := 1; l <= N-1; l++ {
		r := l + 1
		if S[l-1] == S[r-1] {
			dp[l][r] = 2
		} else {
			dp[l][r] = 1
		}
	}

	for length := 2; length <= N; length++ {
		for l := 1; l <= N-length; l++ {
			r := l + length
			var score int
			if S[l-1] == S[r-1] {
				score = 2
			}
			dp[l][r] = maxInts(maxInts(dp[l+1][r], dp[l][r-1]), dp[l+1][r-1]+score)
		}
	}

	fmt.Printf("%d\n", dp[1][N])
}
