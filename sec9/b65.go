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

type Node struct {
	neighbors []int
}

func NewNode() Node {
	return Node{
		neighbors: []int{},
	}
}

func (n *Node) AddNeighbor(v int) {
	n.neighbors = append(n.neighbors, v)
}

type Graph struct {
	nodes []Node
}

func NewGraph(n int) Graph {
	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = NewNode()
	}
	return Graph{
		nodes: nodes,
	}
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(G Graph, rank *[]int, n, p int) int {
	r := *rank
	if r[n] != -1 {
		return r[n]
	}
	r[n] = 0

	for _, v := range G.nodes[n].neighbors {
		if v == p {
			continue
		}
		r[n] = maxInts(dfs(G, rank, v, n)+1, r[n])
	}
	return r[n]
}

func main() {
	s := NewScanner()
	N, T := s.NextInt(), s.NextInt()

	G := NewGraph(N)

	for i := 1; i < N; i++ {
		a, b := s.NextInt(), s.NextInt()
		G.nodes[a-1].AddNeighbor(b - 1)
		G.nodes[b-1].AddNeighbor(a - 1)
	}

	rank := make([]int, N)
	for i := 0; i < N; i++ {
		rank[i] = -1
	}

	dfs(G, &rank, T-1, -1)
	for i, r := range rank {
		if i >= 1 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", r)
	}
	fmt.Printf("\n")
}
