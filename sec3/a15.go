package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	for i := 0; i < N; i++ {
		A[i] = s.NextInt()
	}

	T := make([]int, N)
	copy(T, A)
	sort.Ints(T)
	T = unique(T)

	table := map[int]int{}
	for i, t := range T {
		table[t] = i + 1
	}

	for i := 0; i < len(A); i++ {
		n, _ := table[A[i]]
		fmt.Printf("%d ", n)
	}
}
