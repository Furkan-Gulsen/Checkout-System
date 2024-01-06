package utils

import (
	"math/rand"
	"time"
)

func GenerateID() int {
	seed := time.Now().UnixNano()
	randomGenerator := rand.New(rand.NewSource(seed))
	return randomGenerator.Intn(9000) + 1000
}
