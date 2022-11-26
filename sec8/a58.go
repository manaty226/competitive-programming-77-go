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
	if a > b {
		return a
	}
	return b
}

type RMQ struct {
	tree []int
	size int
}

func NewRMQ(n int) *RMQ {
	size := 1
	for size < n {
		size *= 2
	}
	tree := make([]int, size*2)
	return &RMQ{
		tree: tree,
		size: size,
	}
}

func (t *RMQ) Update(pos, x int) {
	pos = pos + t.size - 1
	t.tree[pos] = x
	for pos >= 2 {
		pos /= 2
		t.tree[pos] = maxInts(t.tree[pos*2], t.tree[pos*2+1])
	}
}

func (t *RMQ) recursiveQuery(l, r, a, b, u int) int {
	if r <= a || b <= l {
		return -1000000000
	}
	if l <= a && b <= r {
		return t.tree[u]
	}
	m := (a + b) / 2
	answerL := t.recursiveQuery(l, r, a, m, u*2)
	answerR := t.recursiveQuery(l, r, m, b, u*2+1)
	return maxInts(answerL, answerR)
}

func (t *RMQ) Query(l, r int) int {
	return t.recursiveQuery(l, r, 1, t.size+1, 1)
}

func main() {
	s := NewScanner()
	N, Q := s.NextInt(), s.NextInt()
	rmq := NewRMQ(N)

	for i := 0; i < Q; i++ {
		q := s.NextInt()
		switch q {
		case 1:
			pos, x := s.NextInt(), s.NextInt()
			rmq.Update(pos, x)
		case 2:
			l, r := s.NextInt(), s.NextInt()
			fmt.Printf("%d \n", rmq.Query(l, r))
		}
	}
}
