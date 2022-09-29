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
	T := s.NextInt()

	times := make([]int, T+1)

	N := s.NextInt()
	for i := 0; i < N; i++ {
		start := s.NextInt()
		end := s.NextInt()
		times[start] += 1
		times[end] -= 1
	}

	fmt.Printf("%d \n", times[0])
	for i := 1; i < T; i++ {
		times[i] += times[i-1]
		fmt.Printf("%d \n", times[i])
	}
}
