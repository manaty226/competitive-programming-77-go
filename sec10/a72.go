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

type Column struct {
	pos int
	cnt int
}

type IntHeap []Column

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i].cnt > h[j].cnt }

func (h *IntHeap) Push(e interface{}) {
	*h = append(*h, e.(Column))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func paintColumn(remainingSteps int, D [][]string) int {
	H, W := len(D), len(D[0])

	columns := &IntHeap{}
	heap.Init(columns)
	for j := 0; j < W; j++ {
		cnt := 0
		for i := 0; i < H; i++ {
			if D[i][j] == "." {
				cnt++
			}
			heap.Push(columns, Column{
				pos: j,
				cnt: cnt,
			})
		}
	}

	for s := 0; s < remainingSteps; s++ {
		pos := heap.Pop(columns).(Column).pos
		for i := 0; i < H; i++ {
			D[i][pos] = "#"
		}
	}

	ret := 0
	for _, row := range D {
		for _, e := range row {
			if e == "#" {
				ret++
			}
		}
	}
	return ret
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := NewScanner()
	H, W, K := s.NextInt(), s.NextInt(), s.NextInt()

	C := make([][]string, H)
	for i := 0; i < H; i++ {
		C[i] = make([]string, W)
		row := s.NextString()
		for j, e := range row {
			C[i][j] = string(e)
		}
	}

	answer := 0
	for t := 0; t < (1 << H); t++ {
		D := make([][]string, H)
		for i := 0; i < H; i++ {
			D[i] = append([]string{}, C[i]...)
		}

		remainingSteps := K
		for i := 0; i < H; i++ {
			if (t>>i)&1 == 0 {
				continue
			}
			remainingSteps--
			for j := 0; j < W; j++ {
				D[i][j] = "#"
			}
		}

		if remainingSteps >= 0 {
			subAnswer := paintColumn(remainingSteps, D)
			answer = maxInts(answer, subAnswer)
		}
	}

	fmt.Println(answer)

}
