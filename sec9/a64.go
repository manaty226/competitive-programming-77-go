package main

import (
	"bufio"
	"container/heap"
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

type Edge struct {
	destination *Node
	distance    int
}

type Node struct {
	id        int
	neighbors []Edge
}

func NewNode(id int) Node {
	return Node{
		id:        id,
		neighbors: []Edge{},
	}
}

func (n *Node) AddEdge(e Edge) {
	if n.neighbors == nil {
		n.neighbors = []Edge{e}
	}
	n.neighbors = append(n.neighbors, e)
}

type Graph struct {
	nodes []*Node
}

func NewGraph(n int) Graph {
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		node := NewNode(i)
		nodes[i] = &node
	}
	return Graph{
		nodes: nodes,
	}
}

func minInts(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type Cost struct {
	id   int
	cost int
}

type CostHeap []Cost

func (h CostHeap) Len() int           { return len(h) }
func (h CostHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h CostHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }

func (h *CostHeap) Push(e interface{}) {
	*h = append(*h, e.(Cost))
}

func (h *CostHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

const INF = 2000000000

func main() {
	s := NewScanner()
	N, M := s.NextInt(), s.NextInt()

	G := NewGraph(N)

	for i := 0; i < M; i++ {
		a, b, c := s.NextInt(), s.NextInt(), s.NextInt()
		G.nodes[a-1].AddEdge(Edge{
			destination: G.nodes[b-1],
			distance:    c,
		})
		G.nodes[b-1].AddEdge(Edge{
			destination: G.nodes[a-1],
			distance:    c,
		})
	}

	determined := make([]bool, N)
	current := make([]int, N)
	for i := 1; i < N; i++ {
		current[i] = INF
	}

	q := &CostHeap{Cost{
		id:   0,
		cost: current[0],
	}}
	heap.Init(q)

	for q.Len() != 0 {
		cost := heap.Pop(q).(Cost)

		if determined[cost.id] {
			continue
		}

		determined[cost.id] = true
		for _, e := range G.nodes[cost.id].neighbors {
			nex := (*e.destination).id
			distance := e.distance
			if current[nex] > current[cost.id]+distance {
				current[nex] = current[cost.id] + distance
				heap.Push(q, Cost{
					id:   nex,
					cost: current[nex],
				})
			}
		}
	}

	for _, c := range current {
		if c == INF {
			fmt.Println("-1")
		}
		fmt.Println(c)
	}

}
