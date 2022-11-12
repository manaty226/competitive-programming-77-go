package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Scanner struct {
	*bufio.Scanner
	*bufio.Reader
}

func NewScanner() *Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 10000)
	scanner.Buffer(buf, 256*1000)

	reader := bufio.NewReader(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return &Scanner{scanner, reader}
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

type City struct {
	X int
	Y int
}

func (c City) DistanceFrom(d City) float64 {
	return math.Sqrt(math.Pow(float64(c.X-d.X), 2) + math.Pow(float64(c.Y-d.Y), 2))
}

func CalcScore(paths []int, cities []City) int {
	N := len(paths) - 1
	sum := 0
	for i := 0; i < N; i++ {
		sum += int(cities[paths[i]-1].DistanceFrom(cities[paths[i+1]-1]))
	}
	return sum
}

func Reverse(A []int) {
	for i := 0; i < len(A)/2; i++ {
		A[i], A[len(A)-i-1] = A[len(A)-i-1], A[i]
	}
}

func GreedyCalcPath(cities []City) []int {
	N := len(cities)
	paths := make([]int, N+1)
	paths[0] = 1
	visited := make([]bool, N)
	visited[0] = true
	currentCity := cities[0]
	for i := 1; i < N; i++ {
		minDist := math.MaxFloat64
		var nextCityIndex int = -1
		for j := 0; j < N; j++ {
			if visited[j] {
				continue
			}
			dist := currentCity.DistanceFrom(cities[j])
			if minDist > dist {
				minDist = dist
				nextCityIndex = j
			}
		}
		visited[nextCityIndex] = true
		currentCity = cities[nextCityIndex]
		paths[i] = nextCityIndex + 1
	}
	paths[N] = 1
	return paths
}

func main() {
	s := NewScanner()
	N := s.NextInt()

	cities := make([]City, N)
	for i := 0; i < N; i++ {
		cities[i] = City{
			X: s.NextInt(),
			Y: s.NextInt(),
		}
	}

	answer := GreedyCalcPath(cities)
	currentScore := CalcScore(answer, cities)
	rand.Seed(time.Now().UnixNano())
	maxIter := 50000
	for i := 0; i < maxIter; i++ {
		l := rand.Intn(N-1) + 1
		r := rand.Intn(N-1) + 1
		if l > r {
			l, r = r, l
		}
		Reverse(answer[l : r+1])
		newScore := CalcScore(answer, cities)

		T := 30.0 - 28.0*float64(i)/float64(maxIter)
		P := math.Exp(math.Min(0.0, float64(currentScore-newScore)/T))
		if rand.Float64() < P {
			currentScore = newScore
		} else {
			Reverse(answer[l : r+1])
		}
	}

	for i := 0; i <= N; i++ {
		fmt.Printf("%d \n", answer[i])
	}
}
