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

type Stack []interface{}

func (s *Stack) Push(x interface{}) {
	*s = append((*s), x)
}

func (s *Stack) Pop() interface{} {
	old := *s
	size := len(old)
	x := old[size-1]
	*s = old[:size-1]
	return x
}

type Bracket struct {
	index int
	value string
}

func main() {
	s := NewScanner()
	brackets := s.NextString()

	stack := Stack{}
	for i, b := range brackets {
		bracket := Bracket{
			index: i + 1,
			value: string(b),
		}
		if bracket.value == ")" {
			ob := stack.Pop().(Bracket)
			fmt.Printf("%d %d \n", ob.index, bracket.index)
		} else {
			stack.Push(bracket)
		}
	}
}
