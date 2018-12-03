package main

import (
	"fmt"
	"io/ioutil"

	"github.com/pradeepg26/advent/util"
)

func LoadInput() []rune {
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

func main() {
	data := LoadInput()
	fmt.Println(Floor(data))
	fmt.Println(IdxOfBasement(data))
}
