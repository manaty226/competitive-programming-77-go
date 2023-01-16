package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Scanner struct {
	*bufio.Scanner
	*bufio.Reader
}

func NewScanner() *Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 10000)
	scanner.Buffer(buf, 256*1000)

	reader := bufio.NewReader(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return &Scanner{scanner, reader}
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

func main() {
	s := NewScanner()
	N := s.NextInt()

	row := make([]int, N)
	column := make([]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			p := s.NextInt()
			if p != 0 {
				row[i] = p
				column[j] = p
			}
		}
	}

	inversionR := 0
	inversionC := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if row[i] > row[j] {
				inversionR++
			}
			if column[i] > column[j] {
				inversionC++
			}
		}
	}

	fmt.Println(inversionR + inversionC)

}
