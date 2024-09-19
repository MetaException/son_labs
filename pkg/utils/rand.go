package utils

import (
	"math/rand"
)

func GenerateRandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func GenerateRandom(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
