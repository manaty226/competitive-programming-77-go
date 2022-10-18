package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return &Scanner{scanner}
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

const diviser int64 = 1000000007

func CalcPower(a, b int64) int64 {
	p, res := a, int64(1)
	for i := 0; i < 30; i++ {
		if (b>>i)&1 == 1 {
			res = (res * p) % diviser
		}
		p = (p * p) % diviser
	}
	return res
}

func main() {
	s := NewScanner()
	a, b := int64(s.NextInt()), int64(s.NextInt())

	answer := CalcPower(a, b)
	fmt.Printf("%d \n", answer)
}
