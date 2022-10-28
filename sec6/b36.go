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
	buf := make([]byte, 512*1000)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)
	scanner.Split(bufio.ScanWords)
	reader := bufio.NewReader(os.Stdin)
	return &Scanner{scanner, reader}
}

func (s *Scanner) NextInt() int {
	s.Scan()
	word := s.Text()
	n, _ := strconv.Atoi(word)
	return n
}

func (s *Scanner) NextString() string {
	// line := ""

	// for {
	// 	l, p, e := s.ReadLine()
	// 	if e != nil {
	// 		return ""
	// 	}
	// 	line += string(l)
	// 	if !p {
	// 		break
	// 	}
	// }
	// return line
	s.Scan()
	return s.Text()
}

func isReachable(count, K int) bool {
	if (count % 2) == (K % 2) {
		return true
	}
	return false
}

func main() {
	s := NewScanner()
	_ = s.NextInt()
	K := s.NextInt()
	S := s.NextString()

	count := 0
	for _, s := range S {
		if string(s) == "1" {
			count++
		}
	}

	if isReachable(count, K) {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}

}
