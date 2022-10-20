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

func CalcComb(n, r int64) int64 {

	// 分子の計算
	a := int64(1)
	for i := int64(1); i <= n; i++ {
		a = (a * i) % diviser
	}

	// 分母の計算
	b := int64(1)
	for i := int64(1); i <= r; i++ {
		b = (b * i) % diviser
	}
	for i := int64(1); i <= n-r; i++ {
		b = (b * i) % diviser
	}

	return (a * CalcPower(b, diviser-2)) % diviser
}

func main() {
	s := NewScanner()
	H, W := int64(s.NextInt()), int64(s.NextInt())
	answer := CalcComb(H+W-2, W-1)
	fmt.Printf("%d \n", answer)
}
