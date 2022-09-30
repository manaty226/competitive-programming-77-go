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

	max := 1501
	dim := make([][]int, max)
	for i := 0; i < max; i++ {
		dim[i] = make([]int, max)
	}

	N := s.NextInt()

	for i := 0; i < N; i++ {
		x := s.NextInt()
		y := s.NextInt()
		dim[x][y]++
	}

	for i := 1; i < max; i++ {
		for j := 1; j < max; j++ {
			dim[i][j] += dim[i][j-1]
		}
	}

	for i := 1; i < max; i++ {
		for j := 1; j < max; j++ {
			dim[i][j] += dim[i-1][j]
		}
	}

	Q := s.NextInt()
	for i := 0; i < Q; i++ {
		A := s.NextInt()
		B := s.NextInt()
		C := s.NextInt()
		D := s.NextInt()

		res := dim[A-1][B-1] + dim[C][D] - dim[A-1][D] - dim[C][B-1]
		fmt.Printf("%d \n", res)
	}
}
