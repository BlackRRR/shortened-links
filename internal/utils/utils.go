package utils

import (
	"math/rand"
	"time"
)

const (
	symbols       = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789_"
	symbolsLength = 10
)

func GetShortURL() string {
	var (
		key string
	)

	rand.Seed(time.Now().UnixNano())

	rs := []rune(symbols)
	lenOfArray := len(rs)

	for i := 0; i < symbolsLength; i++ {
		key += string(rs[rand.Intn(lenOfArray)])
	}

	return key

}
