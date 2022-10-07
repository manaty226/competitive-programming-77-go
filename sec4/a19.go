package main

import (
	"bufio"
	"fmt"
	"math"
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

func unique(X []int) []int {
	unique := []int{X[0]}

	for i := 1; i < len(X); i++ {
		if X[i-1] != X[i] {
			unique = append(unique, X[i])
		}
	}

	return unique
}

func maxInts(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func absInt(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	s := NewScanner()
	N := s.NextInt()
	W := s.NextInt()

	w := make([]int, N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		w[i] = s.NextInt()
		v[i] = s.NextInt()
	}

	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
	}

	dp[0][0] = 0
	for i := 1; i <= W; i++ {
		dp[0][i] = -1000_000_000_000_000_000
	}

	for i := 1; i <= N; i++ {
		for j := 0; j <= W; j++ {
			if j < w[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = maxInts(dp[i-1][j], dp[i-1][j-w[i-1]]+v[i-1])
			}
		}
	}

	maxVal := -1
	for _, v := range dp[N] {
		maxVal = maxInts(maxVal, v)
	}

	fmt.Printf("%d \n", maxVal)
}
