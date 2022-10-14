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

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	s := NewScanner()
	N := s.NextInt()

	primes := []int{2}
	for i := 3; i <= N; i++ {
		IsPrime := true
		for _, p := range primes {
			if i%p == 0 {
				IsPrime = false
				break
			}
			if p*p > i {
				break
			}
		}
		if IsPrime {
			primes = append(primes, i)
		}
	}
	for _, p := range primes {
		fmt.Printf("%d \n", p)
	}
}
