package util

import (
	"math/rand"
	"strings"

	"github.com/shopspring/decimal"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Generate a random int between the min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomAccountOwner() string {
	return RandomString(6)
}

func RandomAmount() decimal.Decimal {
	amount := decimal.NewFromInt(RandomInt(0, 1000))
	fraction := decimal.NewFromInt(RandomInt(0, 99))
	return amount.Add(fraction.Div(decimal.NewFromInt(100)))
}

func RandomCurrency() string {
	currencies := []string{"gbp", "eur", "usd"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
