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

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const inf = 10000000000

func main() {
	s := NewScanner()
	N := s.NextInt()
	W := s.NextInt()
	V := 100000

	w := make([]int, N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		w[i] = s.NextInt()
		v[i] = s.NextInt()
	}

	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, V+1)
		for j := 0; j <= V; j++ {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0
	for i := 1; i <= N; i++ {
		for j := 0; j <= V; j++ {
			if dp[i-1][j] < W {
				dp[i][j] = dp[i-1][j]
			}
			if j-v[i-1] >= 0 && dp[i][j-v[i-1]] < W {
				dp[i][j] = minInts(dp[i-1][j], dp[i-1][j-v[i-1]]+w[i-1])
			}
			if dp[i][j] > W {
				dp[i][j] = inf
			}
		}
	}

	for i := V; i >= 0; i-- {
		if dp[N][i] != inf {
			fmt.Printf("%d \n", i)
			return
		}
	}

}
