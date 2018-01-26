package ebics

import (
	"math/rand"
)

// const lettersAndNumbers = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const numbers = "0123456789"
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ssings = "+-?:().,'"

func RestrictedIdentificationSEPA1() string {
	b := make([]byte, 35)
	selected := letters + numbers + ssings
	for i := range b {
		b[i] = selected[rand.Int63()%int64(len(selected))]
	}
	return string(b)
}
