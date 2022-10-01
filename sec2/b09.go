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
	max := 1600
	s := NewScanner()

	H := max
	W := max
	N := s.NextInt()

	dim := make([][]int, H+1)
	for i := 0; i < H+1; i++ {
		dim[i] = make([]int, W+1)
	}

	lp := []int{max, max}
	rp := []int{0, 0}

	for i := 0; i < N; i++ {
		// 左下座標(A,B)，右上座標(C,D)なので
		// 左上座標(C,B)，右下座標(A,D)となる
		A := s.NextInt()
		B := s.NextInt()
		C := s.NextInt()
		D := s.NextInt()

		if lp[0] > A {
			lp[0] = A
		}
		if lp[1] > B {
			lp[1] = B
		}
		if rp[0] < C {
			rp[0] = C
		}
		if rp[1] < D {
			rp[1] = D
		}

		dim[A][B]++
		dim[A][D]--
		dim[C][B]--
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

	// for i := lp[0]; i <= rp[0]; i++ {
	// 	for j := lp[1]; j < rp[1]; j++ {
	// 		fmt.Printf("%d ", dim[i][j])
	// 	}
	// 	fmt.Printf("\n")
	// }

	var count int
	for i := lp[0]; i <= rp[0]; i++ {
		for j := lp[1]; j < rp[1]; j++ {
			if dim[i][j] > 0 {
				count++
			}
		}
	}
	fmt.Printf("%d \n", count)
}
