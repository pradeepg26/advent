package util

func DieOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Sum(data []int64) int64 {
	var sum int64
	for _, i := range data {
		sum += i
	}
	return sum
}
