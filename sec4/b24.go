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

func maxInts(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type Box struct {
	X int
	Y int
}

func (b Box) isBiggerThan(c Box) bool {
	return ((b.X > c.X) && (b.Y > c.Y))
}

func lower_bound(L []Box, target int) int {

	min, max, mid := 0, len(L)-1, 0
	for min <= max {
		mid = (max-min)/2 + min
		if L[mid].Y >= target {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return mid
}

func main() {
	s := NewScanner()
	N := s.NextInt()

	A := make([]Box, N)
	for i := 0; i < N; i++ {
		A[i] = Box{
			X: s.NextInt(),
			Y: s.NextInt(),
		}
	}

	sort.Slice(A, func(i, j int) bool {
		if A[i].X == A[j].X {
			return A[i].Y > A[j].Y
		}
		return A[i].X < A[j].X
	})

	dp := make([]int, N+1)
	L := make([]Box, N+1)
	len := 1
	L[0] = A[0]
	for i := 1; i < N; i++ {
		pos := lower_bound(L[:len], A[i].Y)
		if A[i].isBiggerThan(L[pos]) {
			pos++
		}
		dp[i] = pos
		L[dp[i]] = A[i]
		if dp[i] >= len {
			len++
		}
	}
	fmt.Printf("%d \n", len)
}
