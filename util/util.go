package util

import (
	"math/rand"
)

func Rand(min,max int) int {
	if max < min {
		return 0
	}
	c := max - min
	r := rand.Intn(c+1)
	return r + min
}