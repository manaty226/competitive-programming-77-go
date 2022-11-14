package main

import (
	"bufio"
	"container/heap"
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
	if a > b {
		return b
	}
	return a
}

type State struct {
	Score    int
	X        []int
	LastMove string
	LastPos  int
}

func (s *State) Copy() State {
	newX := make([]int, len(s.X))
	copy(newX, s.X)
	newState := State{
		Score:    s.Score,
		X:        newX,
		LastMove: s.LastMove,
		LastPos:  s.LastPos,
	}
	return newState
}

func (s *State) Operation(p, q, r int, command string, lastPos int) {
	var adding int
	if command == "A" {
		adding = 1
	} else {
		adding = -1
	}
	newX := s.X
	newX[p-1] += adding
	newX[q-1] += adding
	newX[r-1] += adding
	s.X = newX
	s.LastPos = lastPos
	s.LastMove = command
	score := 0
	for _, x := range s.X {
		score += int(math.Abs(float64(x)))
	}
	s.Score = score
}

type StateHeap []State

func (h StateHeap) Len() int           { return len(h) }
func (h StateHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h StateHeap) Less(i, j int) bool { return h[i].Score < h[j].Score }

func (h *StateHeap) Push(e interface{}) {
	*h = append(*h, e.(State))
}

func (h *StateHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

const MAXWIDTH = 1000

func main() {
	s := NewScanner()
	T := s.NextInt()

	beam := make([][]State, T+1)
	for i := 0; i <= T; i++ {
		beam[i] = []State{}
	}

	beam[0] = append(beam[0], State{
		Score: 0,
		X:     make([]int, 20),
	})

	for i := 1; i <= T; i++ {
		candidate := &StateHeap{}
		heap.Init(candidate)
		P, Q, R := s.NextInt(), s.NextInt(), s.NextInt()
		for j := 0; j < len(beam[i-1]); j++ {
			newStateWithA := beam[i-1][j].Copy()
			newStateWithA.Operation(P, Q, R, "A", j)

			newStateWithB := beam[i-1][j].Copy()
			newStateWithB.Operation(P, Q, R, "B", j)

			heap.Push(candidate, newStateWithA)
			heap.Push(candidate, newStateWithB)
		}
		for j := 0; j < minInts(len(*candidate), MAXWIDTH); j++ {
			state := heap.Pop(candidate)
			beam[i] = append(beam[i], state.(State))
		}
	}

	currentPlace := 0
	answer := make([]string, T+1)
	for i := T; i >= 1; i-- {
		answer[i] = beam[i][currentPlace].LastMove
		currentPlace = beam[i][currentPlace].LastPos
	}
	for i := 1; i <= T; i++ {
		fmt.Printf("%s \n", answer[i])
	}
}
