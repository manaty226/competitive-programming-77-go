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

	A := make([]int, N+2)
	P := make([]int, N+2)
	for i := 1; i <= N; i++ {
		P[i] = s.NextInt()
		A[i] = s.NextInt()
	}

	dp := make([][]int, N+2)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, N+2)
	}

	for length := N - 2; length >= 0; length-- {
		for l := 1; l <= N-length; l++ {
			r := l + length
			var score_1, score_2 int
			if l <= P[l-1] && P[l-1] <= r {
				score_1 = A[l-1]
			}
			if l <= P[r+1] && P[r+1] <= r {
				score_2 = A[r+1]
			}

			switch {
			case l == 1:
				dp[l][r] = dp[l][r+1] + score_2
			case r == N:
				dp[l][r] = dp[l-1][r] + score_1
			default:
				dp[l][r] = maxInts(dp[l-1][r]+score_1, dp[l][r+1]+score_2)
			}
		}
	}

	var res int
	for i := 1; i <= N; i++ {
		res = maxInts(res, dp[i][i])
	}
	fmt.Printf("%d\n", res)
}
