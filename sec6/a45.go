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

func Color2Int(color string) int {
	var res int
	switch color {
	case "W":
		res = 0
	case "B":
		res = 1
	case "R":
		res = 2
	}
	return res
}

func main() {
	s := NewScanner()
	_, C, A := s.NextInt(), s.NextString(), s.NextString()

	var score int = 0
	for _, a := range A {
		score += Color2Int(string(a))
	}

	if score%3 == Color2Int(C) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
