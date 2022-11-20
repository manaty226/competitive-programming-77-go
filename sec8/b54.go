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
	m := map[int]int{}

	sum := 0
	for i := 0; i < N; i++ {
		a := s.NextInt()
		v, ok := m[a]
		if !ok {
			m[a] = 1
		} else {
			sum += v
			m[a] = v + 1
		}
	}
	fmt.Printf("%d \n", sum)
}
