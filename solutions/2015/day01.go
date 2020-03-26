package solutions

import (
	"fmt"
	"github.com/pradeepg26/advent/solutions"
	"io/ioutil"

	"github.com/pradeepg26/advent/util"
)

type solution201501 struct{}

func (s solution201501) Solve(user string) {
}

func LoadInputDay01() []rune {
	bytes, err := ioutil.ReadFile("input.txt")
	util.DieOnErr(err)
	return []rune(string(bytes))
}

func Floor(data []rune) int {
	floor := 0
	for _, ch := range data {
		if '(' == ch {
			floor += 1
		} else if ')' == ch {
			floor -= 1
		}
	}
	return floor
}

func IdxOfBasement(data []rune) int {
	floor := 0
	for i, ch := range data {
		if '(' == ch {
			floor += 1
		} else if ')' == ch {
			floor -= 1
		}
		if floor < 0 {
			return i + 1
		}
	}
	return -1
}

func init() {
	solutions.Register(solutions.ProblemID(201501), solution201501{})
}

func main1501() {
	data := LoadInputDay01()
	fmt.Println(Floor(data))
	fmt.Println(IdxOfBasement(data))
}
