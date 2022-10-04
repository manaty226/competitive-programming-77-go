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

func main() {
	s := NewScanner()
	N := s.NextInt()

	A := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		A[i] = s.NextInt()
	}
	B := make([]int, N-2)
	for i := 0; i < N-2; i++ {
		B[i] = s.NextInt()
	}

	T := make([]int, N)
	T[1] = A[0]
	for i := 2; i < N; i++ {
		a := A[i-1] + T[i-1]
		b := B[i-2] + T[i-2]
		T[i] = int(math.Min(float64(a), float64(b)))
	}
	fmt.Printf("%d", T[len(T)-1])
}
