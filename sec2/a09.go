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

	H := s.NextInt()
	W := s.NextInt()
	N := s.NextInt()

	dim := make([][]int, H+1)
	for i := 0; i < H+1; i++ {
		dim[i] = make([]int, W+1)
	}

	for i := 0; i < N; i++ {
		A := s.NextInt()
		B := s.NextInt()
		C := s.NextInt()
		D := s.NextInt()

		dim[A-1][B-1]++
		dim[A-1][D]--
		dim[C][B-1]--
		dim[C][D]++
	}

	for i := 0; i < H+1; i++ {
		for j := 1; j < W+1; j++ {
			dim[i][j] += dim[i][j-1]
		}
		if i > 0 {
			for j := 0; j < W+1; j++ {
				dim[i][j] += dim[i-1][j]
			}
		}
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Printf("%d ", dim[i][j])
		}
		fmt.Printf("\n")
	}
}
