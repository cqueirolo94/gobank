package db

import (
	"math/rand"

	"github.com/cqueirolo94/gobank/util"
)

// Returns a random Owner
func RandomOwner() string {
	return util.RandomString(10)
}

// Returns a random Money sum
func RandomMoney() int64 {
	return util.RandomInt(0, 1_000)
}

// Returns a random currency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "ARG", "YEN"}

	n := len(currencies)
	return currencies[rand.Intn(n)]
}
