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

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func absInt(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	s := NewScanner()
	N := s.NextInt()
	S := s.NextInt()

	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = s.NextInt()
	}

	dp := make([][]bool, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]bool, S+1)
	}

	dp[0][0] = true
	for i := 1; i <= S; i++ {
		dp[0][i] = false
	}

	for i := 1; i <= N; i++ {
		for j := 0; j <= S; j++ {
			if j < A[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-A[i-1]]
			}
		}
	}

	if dp[N][S] == false {
		fmt.Printf("%d", -1)
		return
	}

	i, j := N, S
	cards := []int{}
	for {
		if i == 0 {
			break
		}
		if dp[i-1][j] {
			i--
			continue
		}
		cards = append(cards, i)
		i, j = i-1, j-A[i-1]
	}

	fmt.Printf("%d \n", len(cards))
	for i := len(cards) - 1; i >= 0; i-- {
		fmt.Printf("%d ", cards[i])
	}

}
