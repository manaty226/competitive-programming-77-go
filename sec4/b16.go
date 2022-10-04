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

	H := make([]int, N)
	for i := 0; i < N; i++ {
		H[i] = s.NextInt()
	}

	T := make([]int, N)
	T[0] = 0
	T[1] = absInt(H[1] - H[0])
	for i := 2; i < N; i++ {
		a := T[i-1] + absInt(H[i]-H[i-1])
		b := T[i-2] + absInt(H[i]-H[i-2])
		T[i] = minInts(a, b)
	}
	fmt.Printf("%d", T[len(T)-1])
}
