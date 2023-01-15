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

type Current struct {
	pos  int
	cost int
}

type IntHeap []Current

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }

func (h *IntHeap) Push(e interface{}) {
	*h = append(*h, e.(Current))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *IntHeap) isEmpty() bool {
	return len(*h) == 0
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Edge struct {
	neighbor int
	cost     int
}

type Node struct {
	neighbors []Edge
}

func (n *Node) AddNeighbor(nb, cost int) {
	n.neighbors = append(n.neighbors, Edge{
		neighbor: nb,
		cost:     cost,
	})
}

type Graph struct {
	nodes []Node
}

func main() {
	s := NewScanner()
	N, M := s.NextInt(), s.NextInt()

	G := Graph{
		nodes: make([]Node, N),
	}
	for i := 0; i < M; i++ {
		A := s.NextInt()
		B := s.NextInt()
		C := s.NextInt()
		D := s.NextInt()

		G.nodes[A-1].AddNeighbor(B-1, 10000*C-D)
		G.nodes[B-1].AddNeighbor(A-1, 10000*C-D)
	}

	visited := make([]bool, N)
	current := make([]int, N)
	for i := 1; i < N; i++ {
		current[i] = 1 << 60
	}

	q := &IntHeap{
		Current{
			pos:  0,
			cost: 1,
		},
	}
	heap.Init(q)

	for !q.isEmpty() {
		cur := heap.Pop(q).(Current)
		if visited[cur.pos] {
			continue
		}
		visited[cur.pos] = true

		for _, n := range G.nodes[cur.pos].neighbors {
			nex := n.neighbor
			cost := n.cost

			if current[nex] > current[cur.pos]+cost {
				current[nex] = current[cur.pos] + cost
				heap.Push(q, Current{
					pos:  nex,
					cost: current[nex],
				})
			}
		}
	}

	dist := (current[N-1] + 9999) / 10000
	tree := dist*10000 - current[N-1]
	fmt.Printf("%d %d \n", dist, tree)

}
