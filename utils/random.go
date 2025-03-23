package utils

import (
	"math/rand"
	"time"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomURL(size int) string {
	rand.Seed(time.Now().UnixNano())
	str := make([]rune, size)
	for i := range str {
		str[i] = runes[rand.Intn(len(runes))]
	}
	return string(str)
}
