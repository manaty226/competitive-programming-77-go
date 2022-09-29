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

	sumArray := make([]int, N+1)
	for i := 0; i < N; i++ {
		sumArray[i+1] = sumArray[i] + (-1 + 2*s.NextInt())
	}

	Q := s.NextInt()
	for i := 0; i < Q; i++ {
		start := s.NextInt()
		end := s.NextInt()
		res := sumArray[end] - sumArray[start-1]

		switch {
		case res > 0:
			fmt.Printf("win \n")
		case res < 0:
			fmt.Printf("lose \n")
		default:
			fmt.Printf("draw \n")
		}
	}

}
