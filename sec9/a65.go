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

func main() {
	s := NewScanner()
	N := s.NextInt()

	G := NewGraph(N)

	for i := 1; i < N; i++ {
		a := s.NextInt()
		G.nodes[a-1].AddNeighbor(i)
	}

	dp := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		for _, n := range G.nodes[i].neighbors {
			dp[i] += dp[n] + 1
		}
	}

	for i, n := range dp {
		if i >= 1 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", n)
	}
	fmt.Printf("\n")
}
