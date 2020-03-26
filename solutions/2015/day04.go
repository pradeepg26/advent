package solutions

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Hash(key string, i int) string {
	data := []byte(fmt.Sprintf("%s%d", key, i))
	return fmt.Sprintf("%x", md5.Sum(data))
}

func Mine(key string, numLeading int) int {
	prefix := strings.Repeat("0", numLeading)
	for i := 1; i < 2000000000; i++ {
		hash := Hash(key, i)
		if strings.HasPrefix(hash, prefix) {
			return i
		}
	}
	panic("must find num")
}

func main1504() {
	fmt.Println(Mine("abcdef", 5))
	fmt.Println(Mine("pqrstuv", 5))
	fmt.Println(Mine("bgvyzdsv", 5))
	fmt.Println(Mine("bgvyzdsv", 6))
}
