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

type GradeStore map[string]int

func main() {
	s := NewScanner()
	Q := s.NextInt()

	store := GradeStore{}

	for i := 0; i < Q; i++ {
		command := s.NextInt()

		switch command {
		case 1:
			name, grade := s.NextString(), s.NextInt()
			store[name] = grade
		case 2:
			name := s.NextString()
			fmt.Printf("%d \n", store[name])
		}
	}

}
