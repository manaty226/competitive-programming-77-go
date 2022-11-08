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

func maxInts(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type Matrix struct {
	matrix [][]int
	rows   []int
}

func NewMatrix(A [][]int) *Matrix {
	N := len(A[0])
	rows := make([]int, N)
	for i := 0; i < N; i++ {
		rows[i] = i + 1
	}
	return &Matrix{
		matrix: A,
		rows:   rows,
	}
}

func (m *Matrix) GetValue(i, j int) int {
	ii := m.rows[i-1]
	return m.matrix[ii-1][j-1]
}

func (m *Matrix) SwitchRows(x, y int) {
	m.rows[y-1], m.rows[x-1] = m.rows[x-1], m.rows[y-1]
}

type Query struct {
	Command int
	X       int
	Y       int
}

func NewQuery(c, x, y int) Query {
	return Query{
		Command: c,
		X:       x,
		Y:       y,
	}
}

func (q Query) Execute(matrix *Matrix) *int {
	var res *int = nil
	switch q.Command {
	case 1:
		matrix.SwitchRows(q.X, q.Y)
	case 2:
		v := matrix.GetValue(q.X, q.Y)
		res = &v
	}
	return res
}

func main() {
	s := NewScanner()
	N := s.NextInt()

	A := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, N)
		for j := 0; j < N; j++ {
			A[i][j] = s.NextInt()
		}
	}

	matrix := NewMatrix(A)
	Q := s.NextInt()

	for i := 0; i < Q; i++ {
		query := NewQuery(s.NextInt(), s.NextInt(), s.NextInt())
		if res := query.Execute(matrix); res != nil {
			fmt.Printf("%d \n", *res)
		}
	}
}
