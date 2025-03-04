package random

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	math_rand "math/rand"
	"time"
)

var defLetters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var defLettersRune = []rune(defLetters)

// Bytes generates n random bytes.
func Bytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

// Base62Str generates a random base62 string of length n.
func Base62Str(s int) string {
	return String(s, defLetters)
}

// Base64Str generates a random base64 string of length n.
func Base64Str(n int) string {
	return String(n, defLetters)
}

// HexStr generates a random hex string of length n.
func HexStr(n int) string {
	return hex.EncodeToString(Bytes(n))
}

func String(n int, letters ...string) string {
	// Default letters if none are passed.
	var letterRunes []rune
	if len(letters) > 0 {
		letterRunes = []rune(letters[0])
	} else {
		letterRunes = defLettersRune
	}

	// Create a random source.
	randSource := math_rand.NewSource(time.Now().UnixNano())
	r := math_rand.New(randSource)

	var bb bytes.Buffer
	bb.Grow(n)
	l := len(letterRunes)

	// Generate random index and append.
	for i := 0; i < n; i++ {
		randomIndex := r.Intn(l)
		bb.WriteRune(letterRunes[randomIndex])
	}

	return bb.String()
}

// Digits returns a random numeric string of length n.
func Digits(n int) string {
	const digits = "0123456789"
	return String(n, digits)
}
