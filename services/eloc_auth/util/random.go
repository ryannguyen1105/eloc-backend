package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const number = "0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInit(min, max int64) int64 {
	if min > max {
		return min
	}
	return min + rand.Int63n(max-min+1)
}

func RandomAlphabet(length int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < length; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomNumber(length int) string {
	var sb strings.Builder
	k := len(number)

	for i := 0; i < length; i++ {
		c := number[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomName() string {
	return RandomAlphabet(4)
}

func RandomRefreshToken() string {
	return RandomAlphabet(3) + RandomNumber(5) + RandomAlphabet(5) + RandomNumber(5)
}

func RandomEmail() string {
	return RandomAlphabet(4) + RandomNumber(3) +"@gmail.com"
}

func RandomPassword() string {
	return RandomAlphabet(5) + RandomNumber(5)
}

func RandomFullName() string {
	return RandomAlphabet(4) + " " + RandomAlphabet(4) + " " + RandomAlphabet(4)
}
