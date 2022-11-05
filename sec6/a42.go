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

type Player struct {
	HP int
	MP int
}

type Team struct {
	count      int
	baselineHP int
	baselineMP int
	difference int
}

func (t *Team) AddMember(p Player) {
	if t.baselineHP <= p.HP && p.HP <= t.baselineHP+t.difference &&
		t.baselineMP <= p.MP && p.MP <= t.baselineMP+t.difference {
		t.count++
	}
}

func main() {
	s := NewScanner()
	N, K := s.NextInt(), s.NextInt()

	players := make([]Player, N)
	for i := 0; i < N; i++ {
		players[i] = Player{
			HP: s.NextInt(),
			MP: s.NextInt(),
		}
	}

	answer := 0
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			team := Team{
				count:      0,
				baselineHP: a,
				baselineMP: b,
				difference: K,
			}
			for _, p := range players {
				team.AddMember(p)
			}
			if answer < team.count {
				answer = team.count
			}
		}
	}
	fmt.Printf("%d \n", answer)
}
