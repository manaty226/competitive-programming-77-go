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

func distanceFrom(start int, G Graph) []int {
	dist := make([]int, len(G.nodes))
	for i := 0; i < len(dist); i++ {
		dist[i] = -1
	}
	dist[start] = 0
	queue := NewQueue()
	queue.Push(start)
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

type State int

func (s State) GetNext(changeList []int, N int) State {
	stateArray := make([]int, 19)
	for i := 0; i < N; i++ {
		wari := (1 << i)
		stateArray[i] = (int(s) / wari) & 1
	}
	for _, c := range changeList {
		stateArray[c] = 1 - stateArray[c]
	}
	ret := 0
	for i, e := range stateArray {
		if e == 1 {
			ret += (1 << i)
		}
	}
	return State(ret)
}

func main() {
	s := NewScanner()
	N, M := s.NextInt(), s.NextInt()

	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = s.NextInt()
	}

	ChangeList := make([][]int, M)
	for i := 0; i < M; i++ {
		ChangeList[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			ChangeList[i][j] = s.NextInt() - 1
		}
	}

	G := NewGraph((1 << N))
	for i := 0; i < (1 << N); i++ {
		for j := 0; j < M; j++ {
			next := State(i).GetNext(ChangeList[j], N)
			G.nodes[i].AddNeighbor(int(next))
		}
	}

	start := 0
	for i, a := range A {
		if a == 1 {
			start += (1 << i)
		}
	}

	dist := distanceFrom(start, G)

	fmt.Println(dist[(1<<N)-1])

}
