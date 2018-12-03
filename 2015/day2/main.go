package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pradeepg26/advent/util"
)

type Dim struct {
	L, W, H int
}

func (d Dim) SurfaceArea() int {
	return 2*d.L*d.W + 2*d.L*d.H + 2*d.W*d.H
}

func (d Dim) MinFaceArea() int {
	a1 := d.L * d.W
	a2 := d.L * d.H
	a3 := d.W * d.H
	return min(a1, min(a2, a3))
}

func (d Dim) MinFacePerimeter() int {
	p1 := 2*d.L + 2*d.W
	p2 := 2*d.L + 2*d.H
	p3 := 2*d.W + 2*d.H
	return min(p1, min(p2, p3))
}

func (d Dim) Volume() int {
	return d.L * d.W * d.H
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func LoadInput() []Dim {
	file, e := os.Open("input.txt")
	util.DieOnErr(e)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := make([]Dim, 0)
	for scanner.Scan() {
		dims := util.ParseIntSliceSep(scanner.Text(), "x")
		d := Dim{dims[0], dims[1], dims[2]}
		res = append(res, d)
	}
	util.DieOnErr(scanner.Err())
	return res
}

func MinAreaNeeded(data []Dim) int {
	total := 0
	for _, d := range data {
		total += d.SurfaceArea() + d.MinFaceArea()
	}
	return total
}

func MinRibbonNeeded(data []Dim) int {
	total := 0
	for _, d := range data {
		total += d.MinFacePerimeter() + d.Volume()
	}
	return total
}

func main() {
	data := LoadInput()
	fmt.Println(MinAreaNeeded(data))
	fmt.Println(MinRibbonNeeded(data))
}
