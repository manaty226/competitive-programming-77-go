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

func main() {
	s := NewScanner()
	N, M := s.NextInt(), s.NextInt()

	G := make([][]int, N)
	for i := 0; i < N; i++ {
		G[i] = []int{}
	}

	for i := 0; i < M; i++ {
		a, b := s.NextInt(), s.NextInt()
		G[a-1] = append(G[a-1], b-1)
		G[b-1] = append(G[b-1], a-1)
	}

	for i := 0; i < N; i++ {
		fmt.Printf("%d: {", i+1)
		for i, v := range G[i] {
			if i > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%d", v+1)
		}
		fmt.Printf("} \n")
	}
}
