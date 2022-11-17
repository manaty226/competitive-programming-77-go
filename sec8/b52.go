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

func (q *Queue) Empty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

type Ball struct {
	Color rune
	Pos   int
}

func main() {
	s := NewScanner()
	N, X := s.NextInt(), s.NextInt()
	A := []rune(s.NextString())

	A[X-1] = '@'
	queue := Queue{}

	queue.Push(
		Ball{
			Color: A[X-1],
			Pos:   X,
		},
	)

	for !queue.Empty() {
		b := queue.Pop().(Ball)
		pb, pa := b.Pos-2, b.Pos
		if pb >= 0 && A[pb] == '.' {
			A[pb] = '@'
			queue.Push(
				Ball{
					Color: A[pa],
					Pos:   pb + 1,
				},
			)
		}
		if pa < N && A[pa] == '.' {
			A[pa] = '@'
			queue.Push(
				Ball{
					Color: A[pa],
					Pos:   pa + 1,
				},
			)
		}
	}

	fmt.Printf("%s \n", string(A))

}
