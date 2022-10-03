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

func check(x float64, target int) (float64, bool) {
	r := math.Pow(float64(x), 3.0) + float64(x)
	res := r - float64(target)
	return res, math.Abs(res) < 0.001
}

func binarySearch(target int) float64 {
	min, max := 0.0, 100.0
	epsilon := 0.00000001
	for {
		mid := (max-min)/2 + min
		if min >= max {
			return min
		}
		res, ok := check(mid, target)
		if ok {
			return mid
		}

		if res > 0 {
			max = mid
		} else {
			min = mid + epsilon
		}
	}
}

func main() {
	s := NewScanner()
	X := s.NextInt()

	r := binarySearch(X)
	if _, ok := check(r, X); !ok {
		fmt.Printf("%f \n", -1.0)
		return
	}
	fmt.Printf("%f \n", r)
}
