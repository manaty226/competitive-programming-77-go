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

type Walker struct {
	position  int
	direction string
}

func main() {
	s := NewScanner()
	N, L := s.NextInt(), s.NextInt()

	answer := 0
	for i := 0; i < N; i++ {
		walker := Walker{
			position:  s.NextInt(),
			direction: s.NextString(),
		}
		if walker.direction == "E" {
			answer = maxInts(L-walker.position, answer)
		} else {
			answer = maxInts(walker.position, answer)
		}
	}
	fmt.Printf("%d \n", answer)
}
