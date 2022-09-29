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
	D := s.NextInt()

	eventDates := make([]int, D+1)

	N := s.NextInt()

	for i := 0; i < N; i++ {
		start := s.NextInt()
		end := s.NextInt()
		eventDates[start-1] += 1
		eventDates[end] -= 1
	}

	fmt.Printf("%d \n", eventDates[0])
	for i := 1; i < D; i++ {
		eventDates[i] += eventDates[i-1]
		fmt.Printf("%d \n", eventDates[i])
	}
}
