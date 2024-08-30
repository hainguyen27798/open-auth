package utils

import (
	"golang.org/x/exp/rand"
	"time"
)

func GenerateOTO() int {
	random := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	return 100000 + random.Intn(900000)
}
