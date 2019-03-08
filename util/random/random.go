package random

import (
	"math/rand"
	"time"
)

// Constant letters for randomize
const (
	Letters     = "ABCDEFGHJKMNPQRSTUVWXYZ23456789" // ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890
	DigiLetters = "1234567890"                      // 1234567890
)

var Let = [4]string{"HJKMNPQRABCDEFGSTUVWXYZ23456789", "UVWXYZ23456789HJKMNPQRABCDEFGST", "789HJKMNPQUVWXYZ23456RABCDEFGST", "YZ23456789HJKMNPQUVWXRABCDEFGST"}

// GenerateRandomString return random string by n
func GenerateRandomString(n int) string {
	letters := []rune(Letters)
	rand.Seed(time.Now().UTC().UnixNano())
	randomString := make([]rune, n)
	for i := range randomString {
		randomString[i] = letters[rand.Intn(len(letters))]
	}
	return string(randomString)
}

// GenerateRandomDigitString return random digit string by n
func GenerateRandomDigitString(n int) string {
	letters := []rune(DigiLetters)
	rand.Seed(time.Now().UTC().UnixNano())
	randomString := make([]rune, n)
	for i := range randomString {
		randomString[i] = letters[rand.Intn(len(letters))]
	}
	return string(randomString)
}

// GenerateRandomKey return random string by n
func GenerateRandomKey(n int, k int) string {
	letters := []rune(Let[k])
	rand.Seed(time.Now().UTC().UnixNano())
	randomString := make([]rune, n)
	for i := range randomString {
		randomString[i] = letters[rand.Intn(len(Let[k]))]
	}
	return string(randomString)
}
