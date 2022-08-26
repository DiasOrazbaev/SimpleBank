package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "qwertyuiopasdfghjklzxcvbnm"
)

var (
	currencies = []string{"EUR", "USD", "KZT", "RUB"}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random interger between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string with len length which interger
func RandomString(length int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random string with length 6
func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
