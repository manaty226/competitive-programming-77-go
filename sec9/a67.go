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

type Edge struct {
	orig int
	dest int
	cost int
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

	if uf.size[ru] > uf.size[rv] {
		ru, rv = rv, ru
	}
	uf.parent[ru] = rv
	uf.size[rv] += uf.size[ru]
}

func (uf *UnionFind) isSameRoot(u, v int) bool {
	return uf.root(u) == uf.root(v)
}

func main() {
	s := NewScanner()
	N, M := s.NextInt(), s.NextInt()

	edges := make([]Edge, M)
	for i := 0; i < M; i++ {
		a, b, c := s.NextInt(), s.NextInt(), s.NextInt()
		edges[i] = Edge{
			orig: a - 1,
			dest: b - 1,
			cost: c,
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

	UF := NewUnionFind(N)

	answer := 0
	for i := 0; i < M; i++ {
		u, v := edges[i].orig, edges[i].dest
		if UF.isSameRoot(u, v) {
			continue
		}
		answer += edges[i].cost
		UF.unite(u, v)
	}
	fmt.Println(answer)
}
