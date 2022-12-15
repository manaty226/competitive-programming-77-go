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

type Queue []interface{}

func NewQueue() Queue {
	return Queue{}
}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

func distanceFromStart(G Graph) []int {
	dist := make([]int, len(G.nodes))
	for i := 1; i < len(dist); i++ {
		dist[i] = -1
	}
	queue := NewQueue()
	queue.Push(0)
	for !queue.Empty() {
		pos := queue.Pop().(int)
		for _, to := range G.nodes[pos].neighbors {
			if dist[to] == -1 {
				dist[to] = dist[pos] + 1
				queue.Push(to)
			}
		}
	}
	return dist
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

	dist := distanceFromStart(G)

	for _, d := range dist {
		fmt.Printf("%d \n", d)
	}

}
