package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
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

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func CalcScore(A, B [][]int) int {
	N := len(A)
	sum := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			sum += absInt(A[i][j] - B[i][j])
		}
	}
	return 200_000_000 - sum
}

type Query struct {
	Q int
	H []int
	X []int
	Y []int
}

func NewQuery(B [][]int, Q int) Query {
	X := make([]int, Q)
	Y := make([]int, Q)
	H := make([]int, Q)
	N := len(B)
	for i := 0; i < Q; i++ {
		X[i] = rand.Intn(N)
		Y[i] = rand.Intn(N)
		H[i] = 1
		B[X[i]][Y[i]] += 1
	}

	return Query{
		Q: Q,
		H: H,
		X: X,
		Y: Y,
	}
}

func (q *Query) ChangeTo(m [][]int, t, x, y, h int) {
	N := len(m)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			m[i][j] -= maxInts(0, q.H[t]-absInt(q.X[t]-i)-absInt(q.Y[t]-j))
		}
	}

	q.X[t] = x
	q.Y[t] = y
	q.H[t] = h

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			m[i][j] += maxInts(0, q.H[t]-absInt(q.X[t]-i)-absInt(q.Y[t]-j))
		}
	}

}

type Optimizer struct {
	MaxTimeMilliseconds int64
	Evaluator           func(A, B [][]int) int
	Query               Query
	Target              [][]int
	Given               [][]int
}

func (o Optimizer) Anneeling() {
	startTime := time.Now().UTC()
	Q := len(o.Query.H)
	N := len(o.Target)

	currentScore := o.Evaluator(o.Given, o.Target)

	for time.Since(startTime).Milliseconds() < o.MaxTimeMilliseconds {
		t := rand.Intn(Q)
		ox, nx := o.Query.X[t], o.Query.X[t]+rand.Intn(19)-9
		oy, ny := o.Query.Y[t], o.Query.Y[t]+rand.Intn(19)-9
		oh, nh := o.Query.H[t], o.Query.H[t]+rand.Intn(39)-19
		if nx < 0 || nx >= N || ny < 0 || ny >= N || nh <= 0 || nh > N {
			continue
		}
		o.Query.ChangeTo(o.Target, t, nx, ny, nh)
		newScore := o.Evaluator(o.Given, o.Target)
		temperature := 180.0 - 179.0*(float64(time.Since(startTime).Milliseconds())/float64(o.MaxTimeMilliseconds))
		prob := math.Exp(math.Min(0.0, float64(newScore-currentScore)/temperature))
		if rand.Float64() < prob {
			currentScore = newScore
		} else {
			o.Query.ChangeTo(o.Target, t, ox, oy, oh)
		}
	}

}

func main() {
	s := NewScanner()
	N := 100
	Q := N * 10

	A := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, N)
		for j := 0; j < N; j++ {
			A[i][j] = s.NextInt()
		}
	}

	B := make([][]int, N)
	for i := 0; i < N; i++ {
		B[i] = make([]int, N)
	}

	query := NewQuery(B, Q)

	optimizer := Optimizer{
		MaxTimeMilliseconds: 5.95 * 1000,
		Evaluator:           CalcScore,
		Query:               query,
		Target:              B,
		Given:               A,
	}

	optimizer.Anneeling()

	fmt.Printf("%d \n", Q)
	for i := 0; i < Q; i++ {
		fmt.Printf("%d %d %d \n", query.X[i], query.Y[i], query.H[i])
	}

}
