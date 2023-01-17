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

func main() {
	s := NewScanner()
	N := s.NextInt()

	problems := make([]Problem, N)
	for i := 0; i < N; i++ {
		T, D := s.NextInt(), s.NextInt()
		problems[i] = Problem{
			consume: T,
			limit:   D,
		}
	}
	sort.Slice(problems, func(i, j int) bool {
		return problems[i].limit < problems[j].limit
	})

	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, 1441)
		for j := 0; j < 1441; j++ {
			dp[i][j] = -1
		}
	}

	dp[0][0] = 0
	for i := 1; i <= N; i++ {
		for j := 0; j <= 1440; j++ {
			if j > problems[i-1].limit || j < problems[i-1].consume {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = maxInts(dp[i-1][j], dp[i-1][j-problems[i-1].consume]+1)
			}
		}
	}
	answer := 0
	for i := 0; i <= 1440; i++ {
		answer = maxInts(answer, dp[N][i])
	}
	fmt.Println(answer)
}
