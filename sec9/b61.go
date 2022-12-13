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

type Graph struct {
	neighbors []int
}

func NewGraph() Graph {
	return Graph{
		neighbors: []int{},
	}
}

func (g *Graph) AddNeighbor(v int) {
	g.neighbors = append(g.neighbors, v)
}

func (g *Graph) CountNeighbors() int {
	return len(g.neighbors)
}

func main() {
	s := NewScanner()
	N, M := s.NextInt(), s.NextInt()

	G := make([]Graph, N)
	for i := 0; i < N; i++ {
		G[i] = NewGraph()
	}

	for i := 0; i < M; i++ {
		a, b := s.NextInt(), s.NextInt()
		G[a-1].AddNeighbor(b - 1)
		G[b-1].AddNeighbor(a - 1)
	}

	currentNode, maxNeighbors := -1, 0
	for i, v := range G {
		count := v.CountNeighbors()
		if count > maxNeighbors {
			currentNode = i + 1
			maxNeighbors = count
		}
	}
	fmt.Printf("%d \n", currentNode)

}
