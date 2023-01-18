package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Scanner struct {
	*bufio.Scanner
	*bufio.Reader
}

func NewScanner() *Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 10000)
	scanner.Buffer(buf, 256*1000)

	reader := bufio.NewReader(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return &Scanner{scanner, reader}
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

type Problem struct {
	consume int
	limit   int
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const mod = 1000000007

func main() {
	s := NewScanner()
	N, W, L, R := s.NextInt(), s.NextInt(), s.NextInt(), s.NextInt()

	X := make([]int, N+2)
	for i := 1; i <= N; i++ {
		X[i] = s.NextInt()
	}
	X[N+1] = W

	dp := make([]int, N+2)
	sum := make([]int, N+2)

	dp[0], sum[0] = 1, 1
	for i := 1; i <= N+1; i++ {
		x := X[i]
		posL := sort.Search(len(X), func(j int) bool { return X[j] >= x-R })
		posR := sort.Search(len(X), func(j int) bool { return X[j] >= x-L+1 }) - 1

		if posR == -1 {
			dp[i] = 0
		} else {
			dp[i] = sum[posR]
		}

		if posL >= 1 {
			dp[i] -= sum[posL-1]
		}
		dp[i] = (dp[i] + mod) % mod

		sum[i] = sum[i-1] + dp[i]
		sum[i] %= mod
	}

	fmt.Println(dp[N+1])

}
