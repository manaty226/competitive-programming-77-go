package main

import (
	"bufio"
	"fmt"
	"math"
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

func minInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Edge struct {
	to  int
	cap int
	rev int
}

type Node struct {
	edges []Edge
}

func NewNode() Node {
	return Node{
		edges: []Edge{},
	}
}

func (n *Node) addEdge(to, cap, rev int) {
	n.edges = append(n.edges, Edge{
		to:  to,
		cap: cap,
		rev: rev,
	})
}

type Graph struct {
	nodes []Node
}

func NewGraph(N int) Graph {
	nodes := make([]Node, N)
	for i := 0; i < N; i++ {
		nodes[i] = NewNode()
	}
	return Graph{
		nodes: nodes,
	}
}

func (g *Graph) addEdge(a, b, cap int) {
	Ga, Gb := len(g.nodes[a].edges), len(g.nodes[b].edges)
	g.nodes[a].addEdge(b, cap, Gb)
	g.nodes[b].addEdge(a, 0, Ga)
}

func dfs(pos, goal, F int, G Graph, used []bool) int {
	if pos == goal {
		return F
	}
	used[pos] = true

	for i, _ := range G.nodes[pos].edges {
		edge := &G.nodes[pos].edges[i]
		if edge.cap == 0 {
			continue
		}
		if used[edge.to] {
			continue
		}

		flow := dfs(edge.to, goal, minInts(F, edge.cap), G, used)
		if flow >= 1 {
			edge.cap -= flow
			G.nodes[edge.to].edges[edge.rev].cap += flow
			return flow
		}
	}
	return 0
}

const INF = 10000000

func maxFlow(s, t int, G Graph) int {
	totalFlow := 0
	for {
		used := make([]bool, len(G.nodes))
		F := dfs(s, t, math.MaxInt64, G, used)
		if F == 0 {
			break
		}
		totalFlow += F
	}
	return totalFlow
}

func main() {
	sc := NewScanner()
	N, M := sc.NextInt(), sc.NextInt()
	G := NewGraph(N + 2)
	s, t := N, N+1
	income := 0
	for i := 0; i < N; i++ {
		P := sc.NextInt()
		if P >= 0 {
			income += P
			G.addEdge(s, i, P)
			G.addEdge(i, t, 0)
		} else {
			G.addEdge(i, t, -P)
			G.addEdge(s, i, 0)
		}
	}

	for i := 0; i < M; i++ {
		a, b := sc.NextInt()-1, sc.NextInt()-1
		G.addEdge(a, b, math.MaxInt64)
	}

	fmt.Printf("%d \n", income-maxFlow(s, t, G))
}
