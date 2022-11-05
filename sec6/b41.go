package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func maxInts(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	s := NewScanner()
	X, Y := s.NextInt(), s.NextInt()
	answer := [][2]int{{X, Y}}

	if X == 1 && Y == 1 {
		fmt.Printf("%d \n", 0)
		return
	}

	for X > 0 && Y > 0 {
		if X > Y {
			X -= Y
		} else {
			Y -= X
		}
		if X == 1 && Y == 1 {
			break
		}
		answer = append(answer, [2]int{X, Y})
	}
	fmt.Printf("%d \n", len(answer))

	for i := len(answer) - 1; i >= 0; i-- {
		fmt.Printf("%d %d \n", answer[i][0], answer[i][1])
	}

}
