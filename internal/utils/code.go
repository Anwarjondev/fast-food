package utils

import (
	"math/rand"
	"time"
)


func GenerateCode() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(900000) + 100000
}