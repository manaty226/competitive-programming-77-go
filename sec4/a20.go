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
	S := s.NextString()
	T := s.NextString()

	lenS := len(S)
	lenT := len(T)

	dp := make([][]int, lenS+1)
	for i := 0; i <= lenS; i++ {
		dp[i] = make([]int, lenT+1)
	}

	for i := 1; i <= lenS; i++ {
		for j := 1; j <= lenT; j++ {
			if S[i-1] == T[j-1] {
				dp[i][j] = maxInts(maxInts(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]+1)
			} else {
				dp[i][j] = maxInts(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Printf("%d\n", dp[lenS][lenT])
}
