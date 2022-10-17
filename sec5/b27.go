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

func Gcd(x, y int) int {
	a, b := x, y
	if x > y {
		a, b = y, x
	}
	residual := b % a
	if residual == 0 {
		return a
	}
	return Gcd(a, residual)
}

func Lcm(x, y int) int {
	gcd := Gcd(x, y)
	return (x * y) / gcd
}

func main() {
	s := NewScanner()
	A, B := s.NextInt(), s.NextInt()

	lcm := Lcm(A, B)
	fmt.Printf("%d \n", lcm)
}
