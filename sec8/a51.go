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

type Query struct {
	Command int
	Title   string
}

func ReadQuery(scanner *Scanner) Query {
	cmd := scanner.NextInt()
	title := ""
	switch cmd {
	case 1:
		title = scanner.NextString()
	default:
	}
	return Query{
		Command: cmd,
		Title:   title,
	}
}

func (q Query) Execute(stack *Stack) {
	switch q.Command {
	case 1:
		stack.Push(q.Title)
	case 2:
		title := stack.Pop().(string)
		fmt.Printf("%s \n", title)
		stack.Push(title)
	case 3:
		stack.Pop()
	}
}

func main() {
	s := NewScanner()
	Q := s.NextInt()

	stack := Stack{}
	for i := 0; i < Q; i++ {
		query := ReadQuery(s)
		query.Execute(&stack)
	}
}
