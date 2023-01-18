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

func isValid(A []int, K, L, l int) bool {
	cnt := 0
	last := 0
	N := len(A)
	for i := 0; i < N; i++ {
		if A[i]-last >= l && L-A[i] >= l {
			cnt++
			last = A[i]
		}
	}
	return cnt >= K
}

func main() {
	s := NewScanner()

	N, L, K := s.NextInt(), s.NextInt(), s.NextInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = s.NextInt()
	}

	left, right := 1, 1_000_000_000

	for left < right {
		mid := (right + left + 1) / 2
		if isValid(A, K, L, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	fmt.Println(left)
}
