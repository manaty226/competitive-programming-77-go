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

func maxInts(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type Student struct {
	WrongAnswers int
}

func main() {
	s := NewScanner()
	N, M := s.NextInt(), s.NextInt()

	students := make([]Student, N)

	for i := 0; i < M; i++ {
		A := s.NextInt()
		students[A-1].WrongAnswers++
	}

	for i := 0; i < N; i++ {
		fmt.Printf("%d \n", M-students[i].WrongAnswers)
	}
}
