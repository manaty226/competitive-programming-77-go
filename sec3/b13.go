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

func main() {
	s := NewScanner()
	N := s.NextInt()
	K := s.NextInt()

	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		A[i] = s.NextInt() + A[i-1]
	}

	R := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		R[i] = R[i-1]
		for R[i] < N && A[R[i]+1]-A[i-1] <= K {
			R[i]++
		}
	}

	var res uint64
	for i, r := range R[1:] {
		res += uint64(r - i)
	}
	fmt.Printf("%d \n", res)
}
