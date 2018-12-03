package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/pradeepg26/advent/util"
)

func LoadInput() []int64 {
	file, e := os.Open("input.txt")
	util.DieOnErr(e)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := make([]int64, 0)
	for scanner.Scan() {
		i, e := strconv.ParseInt(scanner.Text(), 10, 32)
		util.DieOnErr(e)
		res = append(res, i)
	}
	util.DieOnErr(scanner.Err())
	return res
}

func FindFirstRepeatedFrequency(data []int64) int64 {
	set := make(map[int64]bool)
	var freq int64
	for {
		for _, i := range data {
			freq += i
			if set[freq] {
				return freq
			}
			set[freq] = true
		}
	}
}

func main() {
	data := LoadInput()
	fmt.Println("final freq =", util.Sum(data))
	fmt.Println("first repeated freq =", FindFirstRepeatedFrequency(data))
}
