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

type Map [][]string

func NewMap(R, C int, rows []string) Map {
	m := make([][]string, R)
	for i := 0; i < R; i++ {
		m[i] = make([]string, C)
		for j, c := range rows[i] {
			m[i][j] = string(c)
		}
	}
	return Map(m)
}

type Position struct {
	x int
	y int
}

func (p Position) Move(dx, dy int) Position {
	return Position{
		x: p.x + dx,
		y: p.y + dy,
	}
}

func searchPath(m Map, sx, sy, gx, gy int) int {
	dist := make([][]int, len(m))
	for i := 0; i < len(dist); i++ {
		dist[i] = make([]int, len(m[0]))
		for j := 0; j < len(m[0]); j++ {
			dist[i][j] = -1
		}
	}
	dist[sy][sx] = 0

	queue := NewQueue()
	queue.Push(Position{x: sx, y: sy})

	move := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for !queue.Empty() {
		pos := queue.Pop().(Position)
		for _, mv := range move {
			next := pos.Move(mv[0], mv[1])
			if next.x < 0 || next.y < 0 {
				continue
			}
			if m[next.y][next.x] == "#" || dist[next.y][next.x] != -1 {
				continue
			}
			dist[next.y][next.x] = dist[pos.y][pos.x] + 1
			if gx == next.x && gy == next.y {
				return dist[next.y][next.x]
			}
			queue.Push(next)
		}
	}
	return -1
}

func main() {
	s := NewScanner()
	R, C := s.NextInt(), s.NextInt()
	sy, sx := s.NextInt(), s.NextInt()
	gy, gx := s.NextInt(), s.NextInt()

	rows := make([]string, R)
	for i := 0; i < R; i++ {
		rows[i] = s.NextString()
	}

	m := NewMap(R, C, rows)
	dist := searchPath(m, sx-1, sy-1, gx-1, gy-1)

	fmt.Printf("%d\n", dist)

}
