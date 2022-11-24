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

func sumOfDigits(n int) int {
	s := strconv.Itoa(n)
	sum := 0
	for _, d := range s {
		n, _ := strconv.Atoi(string(d))
		sum += n
	}
	return sum
}

func main() {
	s := NewScanner()
	N, K := s.NextInt(), s.NextInt()

	dp := make([][]int, 30)
	dp[0] = make([]int, N+1)
	for i := 0; i <= N; i++ {
		dp[0][i] = i - sumOfDigits(i)
	}

	for i := 1; i <= 29; i++ {
		dp[i] = make([]int, N+1)
		for j := 0; j <= N; j++ {
			dp[i][j] = dp[i-1][dp[i-1][j]]
		}
	}

	answer := make([]int, N+1)
	for i := 1; i <= N; i++ {
		current := i
		j := 0
		for (1 << j) <= K {
			if (K>>j)&1 == 1 {
				current = dp[j][current]
			}
			j++
		}
		answer[i] = current
	}

	for i := 1; i <= N; i++ {
		fmt.Printf("%d \n", answer[i])
	}

}
