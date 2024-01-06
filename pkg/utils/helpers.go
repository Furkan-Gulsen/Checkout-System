package utils

import (
	"math/rand"
	"time"
)

func GenerateID() int {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(9000) + 1000
	return randomInt
}
