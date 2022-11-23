package main

import (
	"bufio"
	"fmt"
	"os"
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

const MOD = 214783647

func Hash(l, r int64, H, B []int64) int64 {
	res := H[r] - (H[l-1] * B[r-l+1] % MOD)
	if res < 0 {
		res += MOD
	}
	return res
}

func main() {
	s := NewScanner()
	N, Q := s.NextInt(), s.NextInt()
	A := make([]int, N+1)
	for i := 0; i < N; i++ {
		A[i] = s.NextInt()
	}

	dp := make([][]int, 30)
	dp[0] = make([]int, N)
	for i := 0; i < N; i++ {
		dp[0][i] = A[i]
	}
	for i := 1; i < 30; i++ {
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = dp[i-1][dp[i-1][j]-1]
		}
	}

	for i := 0; i < Q; i++ {
		current, days := s.NextInt(), s.NextInt()
		for d := 29; d >= 0; d-- {
			if (days>>d)&1 == 1 {
				current = dp[d][current-1]
			}
		}
		fmt.Printf("%d \n", current)
	}

}
