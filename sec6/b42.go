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

type Card struct {
	Face int
	Back int
}

func main() {
	s := NewScanner()
	N := s.NextInt()

	cards := make([]Card, N)
	pp, pm, mp, mm := 0, 0, 0, 0
	for i := 0; i < N; i++ {
		cards[i] = Card{
			Face: s.NextInt(),
			Back: s.NextInt(),
		}
		if cards[i].Face+cards[i].Back > 0 {
			pp += cards[i].Face + cards[i].Back
		}
		if cards[i].Face-cards[i].Back > 0 {
			pm += cards[i].Face - cards[i].Back
		}
		if -cards[i].Face+cards[i].Back > 0 {
			mp += -cards[i].Face + cards[i].Back
		}
		if -cards[i].Face-cards[i].Back > 0 {
			mm += -cards[i].Face - cards[i].Back
		}
	}

	fmt.Printf("%d \n", maxInts(maxInts(pp, pm), maxInts(mp, mm)))
}
