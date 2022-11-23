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

const MOD = 214783647

func Hash(l, r int64, H, B []int64) int64 {
	res := H[r] - (H[l-1] * B[r-l+1] % MOD)
	if res < 0 {
		res += MOD
	}
	return res
}

func main() {
	s := NewScanner()
	N, Q := s.NextInt(), s.NextInt()
	S := s.NextString()

	T := make([]int64, N+1)
	for i, s := range S {
		T[i+1] = int64((s - 'a') + 1)
	}
	B := make([]int64, N+1)
	B[0] = 1
	for i := 1; i <= N; i++ {
		B[i] = 100 * B[i-1] % MOD
	}

	H := make([]int64, N+1)
	H[0] = 0
	for i := 1; i <= N; i++ {
		H[i] = (int64(100)*H[i-1] + T[i]) % MOD
	}

	for i := 0; i < Q; i++ {
		a, b, c, d := s.NextInt(), s.NextInt(), s.NextInt(), s.NextInt()
		h1 := Hash(int64(a), int64(b), H, B)
		h2 := Hash(int64(c), int64(d), H, B)
		if h1 == h2 {
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
	}

}
