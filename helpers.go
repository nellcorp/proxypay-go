package proxypaygo

import (
	"math/rand"
	"time"
)

func generateNineDigitNumber() int64 {
	rand.Seed(time.Now().UnixNano())
	min := 100000000
	max := 999999999
	randomNumber := rand.Intn(max-min+1) + min
	return int64(randomNumber)
}
