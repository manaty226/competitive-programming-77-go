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

func maxInts(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	s := NewScanner()
	N := s.NextInt()

	A := make([]int, N+1)
	for i := 1; i < N; i++ {
		A[i] = s.NextInt()
	}
	B := make([]int, N+1)
	for i := 1; i < N; i++ {
		B[i] = s.NextInt()
	}
	dp := make([]int, N+1)
	for i := 2; i < N+1; i++ {
		dp[i] = -100000000
	}

	for i := 1; i < N; i++ {
		dp[A[i]] = maxInts(dp[A[i]], dp[i]+100)
		dp[B[i]] = maxInts(dp[B[i]], dp[i]+150)
	}

	fmt.Printf("%d \n", dp[N])
}
