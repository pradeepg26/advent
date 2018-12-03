package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pradeepg26/advent/util"
)

func LoadInput() []string {
	return util.LoadInputAsLines()
}

// Process the string and return two booleans as ints
// The first represents the fact that the string contains
// at least on character that occurs in the string exactly
// twice. The second represents the same but with exactly
// three occurances instead
func Process(data string) (int, int) {
	counts := make(map[rune]int)
	for _, c := range data {
		counts[c] += 1
	}
	hasExactlyTwo := 0
	hasExactlyThree := 0
	for _, c := range counts {
		if c == 2 {
			hasExactlyTwo = 1
		} else if c == 3 {
			hasExactlyThree = 1
		}
	}
	return hasExactlyTwo, hasExactlyThree
}

func Checksum(data []string) int64 {
	countTwo := 0
	countThree := 0
	for _, id := range data {
		x, y := Process(id)
		countTwo += x
		countThree += y
	}
	return int64(countTwo) * int64(countThree)
}

func EditDistance(left, right string) int {
	if len(left) != len(right) {
		panic("must be equal length")
	}
	distance := 0
	lhs := []rune(left)
	rhs := []rune(right)
	for i, ch := range lhs {
		if ch != rhs[i] {
			distance++
		}
	}
	return distance
}

func FindCorrectIds(data []string) (string, string) {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if EditDistance(data[i], data[j]) == 1 {
				return data[i], data[j]
			}
		}
	}
	panic("must find a match")
}

func CommonLetters(left, right string) string {
	result := make([]rune, 0, len(left)-1)
	lhs := []rune(left)
	rhs := []rune(right)
	for i, ch := range lhs {
		if ch == rhs[i] {
			result = append(result, ch)
		}
	}
	return string(result)
}

func main() {
	data := LoadInput()
	fmt.Println("checksum =", Checksum(data))
	fmt.Println(CommonLetters(FindCorrectIds(data)))
}
