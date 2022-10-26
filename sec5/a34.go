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

func (s *Scanner) NextString() string {
	s.Scan()
	return s.Text()
}

func isFirstVictory(grundy, A []int) bool {
	xorSum := grundy[A[0]]
	for i := 1; i < len(A); i++ {
		xorSum ^= grundy[A[i]]
	}
	return xorSum != 0
}

func main() {
	s := NewScanner()
	N := s.NextInt()
	X := s.NextInt()
	Y := s.NextInt()

	A := make([]int, N)
	maxNum := 0
	for i := 0; i < N; i++ {
		A[i] = s.NextInt()
		if maxNum < A[i] {
			maxNum = A[i]
		}
	}

	grundy := make([]int, maxNum+1)
	for i := 0; i < maxNum+1; i++ {
		transit := [3]bool{false, false, false}
		if i >= X {
			transit[grundy[i-X]] = true
		}
		if i >= Y {
			transit[grundy[i-Y]] = true
		}

		if transit[0] == false {
			grundy[i] = 0
		} else if transit[1] == false {
			grundy[i] = 1
		} else {
			grundy[i] = 2
		}
	}

	if isFirstVictory(grundy, A) {
		fmt.Printf("First")
	} else {
		fmt.Printf("Second")
	}

}
