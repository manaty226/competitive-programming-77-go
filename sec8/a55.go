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

type OrderedSet []int

func (o *OrderedSet) Push(x int) {
	if pos := sort.Search(len(*o), func(i int) bool { return (*o)[i] == x }); pos == len(*o) {
		*o = append(*o, x)
		sort.Slice(*o, func(i, j int) bool { return (*o)[i] < (*o)[j] })
	}
}

func (o *OrderedSet) Remove(x int) {
	if pos := sort.Search(len(*o), func(i int) bool { return (*o)[i] == x }); pos != len(*o) {
		if pos+1 >= len(*o) {
			*o = (*o)[:pos]
		} else {
			*o = append((*o)[:pos], (*o)[pos+1:]...)
		}
	}
}

func main() {
	s := NewScanner()
	set := OrderedSet{}
	Q := s.NextInt()
	for i := 0; i < Q; i++ {
		q, x := s.NextInt(), s.NextInt()
		switch q {
		case 1:
			set.Push(x)
		case 2:
			set.Remove(x)
		case 3:
			pos := sort.Search(len(set), func(i int) bool { return set[i] >= x })
			if pos >= len(set) {
				fmt.Printf("-1 \n")
			} else {
				fmt.Printf("%d \n", set[pos])
			}
		}
		fmt.Printf("%v \n", set)
	}

}
