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

	A := make([]int, N)
	for i := 1; i < N; i++ {
		A[i] = s.NextInt()
	}
	B := make([]int, N)
	for i := 2; i < N; i++ {
		B[i] = s.NextInt()
	}

	T := make([]int, N)
	T[1] = A[1]
	for i := 2; i < N; i++ {
		a := A[i] + T[i-1]
		b := B[i] + T[i-2]
		T[i] = int(math.Min(float64(a), float64(b)))
	}

	place := N - 1
	path := []int{}
	for place > 0 {
		path = append(path, place+1)
		if T[place]-A[place] == T[place-1] {
			place = place - 1
		} else {
			place = place - 2
		}
	}
	path = append(path, 1)

	fmt.Printf("%d \n", len(path))
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Printf("%d ", path[i])
	}
}
