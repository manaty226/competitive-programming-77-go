package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const INF = 10000000

type RMQ struct {
	tree []int
	size int
}

func NewRMQ(n int) *RMQ {
	size := 1
	for size < n {
		size <<= 1
	}
	tree := make([]int, size*2)
	for i := 0; i < len(tree); i++ {
		tree[i] = INF
	}
	return &RMQ{
		tree: tree,
		size: size,
	}
}

func (t *RMQ) Update(pos, x int) {
	pos += t.size
	t.tree[pos] = x
	for pos >= 2 {
		pos /= 2
		t.tree[pos] = minInts(t.tree[pos*2], t.tree[pos*2+1])
	}
}

func (t *RMQ) recursiveQuery(l, r, a, b, u int) int {
	if r <= a || b <= l {
		return INF
	}
	if l <= a && b <= r {
		return t.tree[u]
	} else {
		m := (a + b) / 2
		answerL := t.recursiveQuery(l, r, a, m, u*2)
		answerR := t.recursiveQuery(l, r, m, b, u*2+1)
		return minInts(answerL, answerR)
	}
}

func (t *RMQ) Query(l, r int) int {
	return t.recursiveQuery(l, r, 0, t.size, 1)
}

func main() {
	s := NewScanner()
	N := s.NextInt()
	L, R := s.NextInt(), s.NextInt()

	rmq := NewRMQ(N)
	rmq.Update(0, 0)
	X := make([]int, N)
	for i := 0; i < N; i++ {
		X[i] = s.NextInt()
	}

	dp := make([]int, N)
	for i := 1; i < N; i++ {
		x := X[i]
		l := sort.Search(len(X), func(j int) bool { return X[j] >= x-R })
		r := sort.Search(len(X), func(j int) bool { return X[j] > x-L })
		dp[i] = rmq.Query(l, r) + 1
		rmq.Update(i, dp[i])
	}
	fmt.Printf("%d \n", dp[N-1])
}
