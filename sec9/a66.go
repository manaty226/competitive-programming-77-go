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

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(N int) UnionFind {
	parent := make([]int, N)
	size := make([]int, N)
	for i := 0; i < N; i++ {
		parent[i] = -1
		size[i] = 1
	}
	return UnionFind{
		parent: parent,
		size:   size,
	}
}

func (uf *UnionFind) root(x int) int {
	for uf.parent[x] != -1 {
		x = uf.parent[x]
	}
	return x
}

func (uf *UnionFind) unite(u, v int) {
	ru := uf.root(u)
	rv := uf.root(v)
	if ru == rv {
		return
	}
	if uf.size[ru] < uf.size[rv] {
		uf.parent[ru] = rv
		uf.size[rv] = uf.size[ru] + uf.size[rv]
	} else {
		uf.parent[rv] = ru
		uf.size[ru] = uf.size[ru] + uf.size[rv]
	}
}

func (uf *UnionFind) isSameRoot(u, v int) bool {
	return uf.root(u) == uf.root(v)
}

func main() {
	s := NewScanner()
	N, Q := s.NextInt(), s.NextInt()

	UF := NewUnionFind(N)

	for i := 0; i < Q; i++ {
		q1, q2, q3 := s.NextInt(), s.NextInt()-1, s.NextInt()-1
		if q1 == 1 {
			UF.unite(q2, q3)
		} else {
			if UF.isSameRoot(q2, q3) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}

}
