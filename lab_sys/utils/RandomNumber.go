package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func RandNumber(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	v := " "
	for i := 0; i < n; i++ {
		v += strconv.Itoa(r.Intn(n))

	}
	return v
}
