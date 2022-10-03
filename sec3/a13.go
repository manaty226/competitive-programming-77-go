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

	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = s.NextInt()
	}

	R := make([]int, N)
	for i := 0; i < N; i++ {
		if i > 0 {
			R[i] = R[i-1]
		}
		for R[i] < N-1 && A[R[i]+1]-A[i] <= K {
			R[i]++
		}
	}

	var res uint64
	for i, r := range R {
		res += uint64(r - i)
	}
	fmt.Printf("%d \n", res)
}
