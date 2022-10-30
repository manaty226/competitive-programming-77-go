package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Movie struct {
	start int
	end   int
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
	N := s.NextInt()

	M := make([]Movie, N)
	for i := 0; i < N; i++ {
		M[i] = Movie{
			start: s.NextInt(),
			end:   s.NextInt(),
		}
	}

	sort.Slice(M, func(i, j int) bool {
		return M[i].end < M[j].end
	})

	seenMovies := 1
	var lastMovie Movie = M[0]
	for i := 1; i < N; i++ {
		if lastMovie.end <= M[i].start {
			seenMovies++
			lastMovie = M[i]
		}
	}
	fmt.Printf("%d \n", seenMovies)
}
