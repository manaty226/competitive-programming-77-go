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

func main() {
	s := NewScanner()
	N := s.NextInt()

	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = s.NextInt()
	}

	l2r := make([]int, N+1)
	r2l := make([]int, N+1)

	for i := 0; i < N; i++ {
		l2r[i+1] = int(math.Max(float64(l2r[i]), float64(A[i])))
		r2l[N-1-i] = int(math.Max(float64(r2l[N-i]), float64(A[N-1-i])))
	}

	D := s.NextInt()
	for i := 0; i < D; i++ {
		L := s.NextInt()
		R := s.NextInt()
		fmt.Printf("%d \n", int(math.Max(float64(l2r[L-1]), float64(r2l[R]))))
	}
}
