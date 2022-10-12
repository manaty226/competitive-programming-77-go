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

func (s *Scanner) NextLine() string {
	s.Scan()
	return s.Text()
}

func maxInts(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func lower_bound(L []int, target int) int {

	min, max, mid := 0, len(L)-1, 0
	for min <= max {
		mid = (max-min)/2 + min
		if L[mid] >= target {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return mid
}

func main() {
	s := NewScanner()
	H := s.NextInt()
	W := s.NextInt()

	C := make([][]string, H)
	for i := 0; i < H; i++ {
		C[i] = make([]string, W)
		line := s.NextLine()
		for j, c := range line {
			C[i][j] = string(c)
		}
	}

	dp := make([][]int64, H+1)
	for i := 0; i <= H; i++ {
		dp[i] = make([]int64, W+1)
	}

	dp[0][0] = 1
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if C[i][j] == "#" {
				continue
			}
			dp[i+1][j] += dp[i][j]
			dp[i][j+1] += dp[i][j]
		}
	}

	fmt.Printf("%d \n", dp[H-1][W-1])
}
