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

func isConnected(G Graph) bool {
	visited := make([]bool, len(G.nodes))
	dfs(G, 0, &visited)
	for _, b := range visited {
		if b == false {
			return false
		}
	}
	return true
}

func dfs(G Graph, v int, visited *[]bool) {
	(*visited)[v] = true
	for _, n := range G.nodes[v].neighbors {
		if (*visited)[n] == false {
			dfs(G, n, visited)
		}
	}
	return
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

	if isConnected(G) {
		fmt.Printf("The graph is connected.")
	} else {
		fmt.Printf("The graph is not connected.")
	}

}
