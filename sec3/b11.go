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

func binarySearch(A []int, target int) int {

	min, max := 0, len(A)-1
	mid := (max-min)/2 + min
	for {
		if min >= max {
			return mid
		}
		if A[mid] == target {
			return mid
		} else if A[mid] > target {
			max = mid - 1
		} else {
			min = mid + 1
		}
		mid = (max-min)/2 + min
	}
}

func main() {
	s := NewScanner()
	N := s.NextInt()

	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = s.NextInt()
	}

	sort.Ints(A)

	Q := s.NextInt()
	for i := 0; i < Q; i++ {
		X := s.NextInt()
		r := binarySearch(A, X)

		var res int
		switch {
		case X > A[r]:
			res = r + 1
		default:
			res = r
		}
		fmt.Printf("%d \n", res)
	}
}
