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

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

type StockPrice struct {
	date  int
	price int
}

func main() {
	s := NewScanner()
	N := s.NextInt()

	stack := Stack{
		StockPrice{
			date:  1,
			price: s.NextInt(),
		},
	}
	answer := make([]int, N)
	answer[0] = -1
	for i := 1; i < N; i++ {
		a := s.NextInt()
		if stack.Empty() {
			answer[i] = -1
			break
		}
		x := stack.Pop().(StockPrice)
		for {
			if x.price > a {
				answer[i] = x.date
				stack.Push(x)
				break
			}
			if stack.Empty() {
				answer[i] = -1
				break
			}
			x = stack.Pop().(StockPrice)
		}
		stack.Push(
			StockPrice{
				date:  i + 1,
				price: a,
			},
		)
	}
	for _, v := range answer {
		fmt.Printf("%d ", v)
	}
}
