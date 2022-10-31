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

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *IntHeap) Push(e interface{}) {
	*h = append(*h, e.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxInts(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func sumOfHeights(H []int) int {
	sum := 0
	for _, h := range H {
		sum += h
	}
	return sum
}

func main() {
	s := NewScanner()
	N, D := s.NextInt(), s.NextInt()

	works := make([][]int, D)
	for i := 0; i < D; i++ {
		works[i] = []int{}
	}
	for i := 0; i < N; i++ {
		X, Y := s.NextInt(), s.NextInt()
		works[X-1] = append(works[X-1], Y)
	}

	que := &IntHeap{}
	heap.Init(que)
	answer := 0
	for i := 0; i < D; i++ {
		for _, w := range works[i] {
			heap.Push(que, -w)
		}
		if que.Len() == 0 {
			continue
		}
		answer += -heap.Pop(que).(int)
	}

	fmt.Printf("%d \n", answer)
}
