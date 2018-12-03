package main

import (
	"fmt"
	"io/ioutil"

	"github.com/pradeepg26/advent/util"
)

type Pos struct {
	x, y int
}

func (p Pos) ToString() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p Pos) Move(r rune) Pos {
	switch r {
	case 'v':
		return Pos{p.x, p.y - 1}
	case '^':
		return Pos{p.x, p.y + 1}
	case '<':
		return Pos{p.x - 1, p.y}
	case '>':
		return Pos{p.x + 1, p.y}
	default:
		panic("invalid rune")
	}
}

func LoadInput() []rune {
	bytes, err := ioutil.ReadFile("input.txt")
	util.DieOnErr(err)
	return []rune(string(bytes))
}

func HousesVisited(data []rune) int {
	set := make(map[string]bool)
	pos := Pos{}
	set[pos.ToString()] = true
	for _, r := range data {
		pos = pos.Move(r)
		set[pos.ToString()] = true
	}
	return len(set)
}

func HousesVisitedTwoAtATime(data []rune) int {
	set := make(map[string]bool)
	pos1 := Pos{}
	pos2 := Pos{}
	set[pos1.ToString()] = true
	for i := 0; i < len(data); i = i + 2 {
		pos1 = pos1.Move(data[i])
		pos2 = pos2.Move(data[i+1])
		set[pos1.ToString()] = true
		set[pos2.ToString()] = true
	}
	return len(set)
}

func main() {
	data := LoadInput()
	fmt.Println(HousesVisited(data))
	fmt.Println(HousesVisitedTwoAtATime(data))
}
