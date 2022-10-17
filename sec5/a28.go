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

const diviser int64 = 10000

func Calculator() func(operator string, x int64) int64 {
	var answer int64 = 0
	return func(operator string, x int64) int64 {
		switch operator {
		case "+":
			answer += x
		case "-":
			answer -= x
		case "*":
			answer *= x
		}
		if answer < 0 {
			answer += diviser
		}
		answer %= diviser
		return answer
	}
}

func main() {
	s := NewScanner()
	N := s.NextInt()

	c := Calculator()
	var answer int64 = 0
	for i := 0; i < N; i++ {
		operator, x := s.NextString(), s.NextInt()
		answer = c(operator, int64(x))
		fmt.Printf("%d \n", answer)
	}
}
