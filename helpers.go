package proxypaygo

import (
	"math/rand"
)

func generateNineDigitNumber() int64 {
	min := 100000000
	max := 999999999
	randomNumber := rand.Intn(max-min+1) + min
	return int64(randomNumber)
}
