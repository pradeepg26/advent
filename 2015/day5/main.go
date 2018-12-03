package main

import (
	"fmt"

	"github.com/pradeepg26/advent/util"
)

func LoadInput() []string {
	return util.LoadInputAsLines()
}

func NiceString(s string) bool {
	vowelCount := 0
	hasDoubleChar := false
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		cur := str[i]
		if cur == 'a' || cur == 'e' || cur == 'i' || cur == 'o' || cur == 'u' {
			vowelCount++
		}
		if i > 0 {
			prev := str[i-1]
			hasDoubleChar = hasDoubleChar || (prev == cur)
			// ab, cd, pq, or xy
			if (prev == 'a' && cur == 'b') ||
				(prev == 'c' && cur == 'd') ||
				(prev == 'p' && cur == 'q') ||
				(prev == 'x' && cur == 'y') {
				return false
			}
		}
	}
	return vowelCount > 2 && hasDoubleChar
}

func CountNiceStrings(data []string) int {
	count := 0
	for _, s := range data {
		if NiceString(s) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(NiceString("ugknbfddgicrmopn"))
	fmt.Println(NiceString("aaa"))
	fmt.Println(NiceString("jchzalrnumimnmhp"))
	fmt.Println(NiceString("haegwjzuvuyypxyu"))
	fmt.Println(NiceString("dvszwmarrgswjxmb"))
	data := LoadInput()
	fmt.Println(CountNiceStrings(data))
}
