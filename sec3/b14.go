package main

import (
	"bufio"
	"fmt"
	"math"
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

func nextIntArray(s *Scanner, N int) []int {
	array := make([]int, N)
	for i := 0; i < N; i++ {
		array[i] = s.NextInt()
	}
	return array
}

func bitFullSearch(A []int) []int {
	length := int(math.Pow(2.0, float64(len(A))))
	res := make([]int, length)
	for i := 0; i < length; i++ {
		sum := 0
		for j, a := range A {
			if (i/(1<<j))&1 == 1 {
				sum += a
			}
		}
		res[i] = sum
	}
	return res
}

func array_to_map(A []int) map[int]bool {
	res := map[int]bool{}
	for _, a := range A {
		res[a] = true
	}
	return res
}

func main() {
	s := NewScanner()
	N := s.NextInt()
	K := s.NextInt()

	A := nextIntArray(s, N)

	left := bitFullSearch(A[:N/2])
	right := array_to_map(bitFullSearch(A[N/2:]))

	for i := 0; i < len(left); i++ {
		if _, ok := right[K-left[i]]; ok {
			fmt.Printf("Yes")
			return
		}
	}
	fmt.Printf("No")
}
