package main

import (
	"fmt"
	"regexp"

	"github.com/pradeepg26/advent/util"
)

type Claim struct {
	id         string
	x, y, w, h int
}

func LoadInput() []Claim {
	result := make([]Claim, 0)
	for _, str := range util.LoadInputAsLines() {
		re := regexp.MustCompile(`(?P<id>#\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<w>\d+)x(?P<h>\d+)`)
		extr := re.FindStringSubmatch(str)
		id := extr[1]
		x := util.ParseInt(extr[2])
		y := util.ParseInt(extr[3])
		w := util.ParseInt(extr[4])
		h := util.ParseInt(extr[5])
		result = append(result, Claim{id, x, y, w, h})
	}
	return result
}

func Process(data []Claim) int {
	counts := make(map[string]int)
	for _, c := range data {
		// iterate each square and mark it
		for i := c.x; i < c.x+c.w; i++ {
			for j := c.y; j < c.y+c.h; j++ {
				coord := str(i, j)
				counts[coord] = counts[coord] + 1
			}
		}
	}
	total := 0
	for _, count := range counts {
		if count > 1 {
			total++
		}
	}
	return total
}

func FindNonOverlappingClaim(data []Claim) string {
	claims := make(map[string][]string)
	for _, c := range data {
		// iterate each square and mark it
		for i := c.x; i < c.x+c.w; i++ {
			for j := c.y; j < c.y+c.h; j++ {
				coord := str(i, j)
				claims[coord] = append(claims[coord], c.id)
			}
		}
	}
	maxOverlap := make(map[string]int)
	for _, claimsInSquare := range claims {
		sz := len(claimsInSquare)
		for _, cl := range claimsInSquare {
			if c, ok := maxOverlap[cl]; ok {
				maxOverlap[cl] = max(c, sz)
			} else {
				maxOverlap[cl] = sz
			}
		}
	}
	for cl, lap := range maxOverlap {
		if lap == 1 {
			return cl
		}
	}
	return ""
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func str(x, y int) string {
	return fmt.Sprintf("(%d, %d)", x, y)
}

func main() {
	data := LoadInput()
	fmt.Println(Process(data))
	fmt.Println(FindNonOverlappingClaim(data))
}
