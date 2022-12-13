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

func searchPath(G Graph, goal int) []int {
	visited := make([]bool, len(G.nodes))
	path := []int{0}
	dfs(G, 0, &visited, &path, goal)
	return path
}

func dfs(G Graph, v int, visited *[]bool, path *[]int, goal int) bool {
	if v == goal {
		return true
	}
	(*visited)[v] = true
	for _, n := range G.nodes[v].neighbors {
		if (*visited)[n] {
			continue
		}
		*path = append(*path, n)
		if dfs(G, n, visited, path, goal) {
			return true
		}
		*path = (*path)[:len(*path)-1]
	}
	return false
}

func main() {
	s := NewScanner()
	N, M := s.NextInt(), s.NextInt()

	G := NewGraph(N)

	for i := 0; i < M; i++ {
		a, b := s.NextInt(), s.NextInt()
		G.nodes[a-1].AddNeighbor(b - 1)
		G.nodes[b-1].AddNeighbor(a - 1)
	}

	path := searchPath(G, len(G.nodes)-1)
	for _, p := range path {
		fmt.Printf("%d \n", p+1)
	}

}
