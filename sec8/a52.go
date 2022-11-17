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

type Queue []interface{}

func (q *Queue) Push(x interface{}) {
	*q = append((*q), x)
}

func (q *Queue) Pop() interface{} {
	old := *q
	x := old[0]
	*q = old[1:]
	return x
}

func (q *Queue) Top() interface{} {
	return (*q)[0]
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

func (q Query) Execute(queue *Queue) {
	switch q.Command {
	case 1:
		queue.Push(q.Title)
	case 2:
		title := queue.Top().(string)
		fmt.Printf("%s \n", title)
	case 3:
		queue.Pop()
	}
}

func main() {
	s := NewScanner()
	Q := s.NextInt()

	queue := Queue{}
	for i := 0; i < Q; i++ {
		query := ReadQuery(s)
		query.Execute(&queue)
	}
}
