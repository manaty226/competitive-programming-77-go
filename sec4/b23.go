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

func minFloat(a, b float64) float64 {
	return math.Min(a, b)
}

const inf = 1000_000_000

type City struct {
	X int
	Y int
}

func (c City) DistanceFrom(d City) float64 {
	return math.Sqrt(math.Pow(float64(c.X)-float64(d.X), 2) + math.Pow(float64(c.Y)-float64(d.Y), 2))
}

func main() {
	s := NewScanner()
	N := s.NextInt()
	Total := 1 << N

	A := make([]City, N+1)
	for i := 1; i <= N; i++ {
		A[i] = City{
			X: s.NextInt(),
			Y: s.NextInt(),
		}
	}

	dp := make([][]float64, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]float64, Total+1)
		for j := 0; j <= Total; j++ {
			dp[i][j] = inf
		}
	}

	dp[1][0] = 0
	for j := 0; j < Total; j++ {
		for i := 1; i <= N; i++ {
			if dp[i][j] >= inf {
				continue
			}
			// すでに通った都市の場合スキップ
			for k := 1; k <= N; k++ {
				if j == k || (j&(1<<(k-1))) != 0 {
					continue
				}
				passed := j + (1 << (k - 1))
				dp[k][passed] = math.Min(dp[k][passed], dp[i][j]+A[k].DistanceFrom(A[i]))
			}
		}
	}

	fmt.Printf("%f \n", dp[1][Total-1])
}
