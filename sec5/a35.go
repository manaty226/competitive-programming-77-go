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

func isFirstVictory(grundy, A []int) bool {
	xorSum := grundy[A[0]]
	for i := 1; i < len(A); i++ {
		xorSum ^= grundy[A[i]]
	}
	return xorSum != 0
}

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInts(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	s := NewScanner()
	N := s.NextInt()

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		dp[N-1][i] = s.NextInt()
	}

	for i := N - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if (i+1)%2 == 1 {
				dp[i][j] = maxInts(dp[i+1][j], dp[i+1][j+1])
			} else {
				dp[i][j] = minInts(dp[i+1][j], dp[i+1][j+1])
			}
		}
	}

	fmt.Printf("%d \n", dp[0][0])

}
