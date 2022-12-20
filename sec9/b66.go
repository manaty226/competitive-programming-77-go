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
	edges := make([][2]int, M)
	for i := 0; i < M; i++ {
		a, b := s.NextInt()-1, s.NextInt()-1
		edges[i] = [2]int{a, b}
	}

	UF := NewUnionFind(N)

	Q := s.NextInt()
	queries := make([][]int, Q)
	isBreak := make([]bool, M)
	for i := 0; i < Q; i++ {
		q1 := s.NextInt()
		if q1 == 1 {
			en := s.NextInt() - 1
			queries[i] = []int{q1, edges[en][0], edges[en][1]}
			isBreak[en] = true
		} else {
			a, b := s.NextInt()-1, s.NextInt()-1
			queries[i] = []int{q1, a, b}
		}
	}

	for i, b := range isBreak {
		if !b {
			UF.unite(edges[i][0], edges[i][1])
		}
	}

	answer := []string{}
	for i := Q - 1; i >= 0; i-- {
		if queries[i][0] == 1 {
			UF.unite(queries[i][1], queries[i][2])
		} else {
			if UF.isSameRoot(queries[i][1], queries[i][2]) {
				answer = append(answer, "Yes")
			} else {
				answer = append(answer, "No")
			}
		}
	}

	for i := len(answer) - 1; i >= 0; i-- {
		fmt.Println(answer[i])
	}

}
