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

const inf = 1000_000_000

func main() {
	s := NewScanner()
	N := s.NextInt()
	M := s.NextInt()
	Total := 1 << N

	A := make([][]int, M+1)
	for i := 0; i <= M; i++ {
		A[i] = make([]int, N+1)
	}

	for i := 1; i <= M; i++ {
		for j := 1; j <= N; j++ {
			A[i][j] = s.NextInt()
		}
	}

	dp := make([][]int, M+1)
	for i := 0; i <= M; i++ {
		dp[i] = make([]int, Total)
		for j := 0; j < Total; j++ {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0
	for i := 1; i <= M; i++ {
		for j := 0; j < Total; j++ {
			// すでにクーポン適用済み or 今回クーポン適用
			v := 0
			for k := 1; k <= N; k++ {
				if (j>>(k-1))&1 == 1 || A[i][k] == 1 {
					v += (1 << (k - 1))
				}
			}
			dp[i][j] = minInts(dp[i][j], dp[i-1][j])
			dp[i][v] = minInts(dp[i][v], dp[i][j]+1)
		}
	}

	if dp[M][Total-1] == inf {
		fmt.Printf("%d \n", -1)
	} else {
		fmt.Printf("%d \n", dp[M][Total-1])
	}
}
