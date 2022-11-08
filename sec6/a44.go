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

type State int

const (
	FORWARD State = 0
	REVERSE State = 1
)

type ArrayWithState struct {
	data  []int
	state State
}

func NewArrayWithState(size int) *ArrayWithState {
	d := make([]int, size)
	for i := 0; i < size; i++ {
		d[i] = i + 1
	}
	return &ArrayWithState{
		data:  d,
		state: FORWARD,
	}
}

func (a ArrayWithState) GetPosition(p int) int {
	if a.state == FORWARD {
		return p - 1
	} else {
		return (len(a.data) - p)
	}
}

func (a *ArrayWithState) ChangeOrder() {
	a.state = State((int(a.state) + 1) % 2)
}

func (a *ArrayWithState) Change(position int, value int) {
	p := a.GetPosition(position)
	a.data[p] = value
}

func (a *ArrayWithState) GetValue(position int) int {
	p := a.GetPosition(position)
	return a.data[p]
}

type Query int

const (
	Change  Query = 1
	Reverse Query = 2
	Get     Query = 3
)

type QueryBuilder struct {
	Array        *ArrayWithState
	Query        Query
	Position     int
	ChangedValue int
}

func NewQuery(query Query, array *ArrayWithState) *QueryBuilder {
	return &QueryBuilder{
		Array:        array,
		Query:        Query(query),
		Position:     -1,
		ChangedValue: -1,
	}
}

func (q *QueryBuilder) WithPosition(p int) *QueryBuilder {
	q.Position = p
	return q
}

func (q *QueryBuilder) WithValue(v int) *QueryBuilder {
	q.ChangedValue = v
	return q
}

func (q *QueryBuilder) Execute() *int {
	var res *int = nil
	switch q.Query {
	case Change:
		q.Array.Change(q.Position, q.ChangedValue)
	case Reverse:
		q.Array.ChangeOrder()
	case Get:
		v := q.Array.GetValue(q.Position)
		res = &v
	}
	return res
}

func main() {
	s := NewScanner()
	N, Q := s.NextInt(), s.NextInt()

	array := NewArrayWithState(N)

	for i := 0; i < Q; i++ {
		command := Query(s.NextInt())
		query := NewQuery(command, array)
		if command == Change {
			query.WithPosition(s.NextInt()).WithValue(s.NextInt())
		} else if command == Get {
			query.WithPosition(s.NextInt())
		}
		if res := query.Execute(); res != nil {
			fmt.Printf("%d \n", *res)
		}
	}
}
