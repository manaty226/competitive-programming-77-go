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

func check(A []int, period int, target int) bool {
	var sum int
	for _, s := range A {
		sum += int(period / s)
	}
	return target <= sum
}

func binarySearch(A []int, target int) int {
	min, max := 1, 1000_000_000
	for {
		mid := (max-min)/2 + min
		if min >= max {
			return min
		}
		ok := check(A, mid, target)
		if ok {
			max = mid
		} else {
			min = mid + 1
		}
		// fmt.Printf("min: %d, max: %d, mid: %d \n", min, max, mid)
	}
}

func main() {
	s := NewScanner()
	N := s.NextInt()
	X := s.NextInt()

	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = s.NextInt()
	}

	r := binarySearch(A, X)
	fmt.Printf("%d \n", r)
}
