package util

import (
	"strconv"
	"strings"
)

func DieOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseInt(s string) int {
	i, e := strconv.ParseInt(s, 10, 32)
	DieOnErr(e)
	return int(i)
}

func StrSliceToIntSlice(arr []string) []int {
	res := make([]int, 0, len(arr))
	for _, s := range arr {
		res = append(res, ParseInt(s))
	}
	return res
}

func ParseIntSliceSep(s string, sep string) []int {
	split := strings.Split(s, sep)
	return StrSliceToIntSlice(split)
}

func Sum(data []int64) int64 {
	var sum int64
	for _, i := range data {
		sum += i
	}
	return sum
}
