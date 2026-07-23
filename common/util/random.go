package util

import (
	"math"
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

func RandomNumberString(length int) string {
	var sb strings.Builder
	k := len(number)

	for i := 0; i < length; i++ {
		c := number[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomNumberInt(length int) int64 {
	if length <= 0 {
		return 0
	}
	if length > 18 {
		length = 18
	}
	min := int64(math.Pow10(length - 1))
	max := int64(math.Pow10(length))

	return min + rand.Int63n(max-min)
}

func RandomName() string {
	return RandomAlphabet(4)
}

func RandomToken() string {
	return RandomAlphabet(32)
}

func RandomRefreshToken() string {
	return RandomAlphabet(3) + RandomNumberString(5) + RandomAlphabet(5) + RandomNumberString(5)
}

func RandomEmail() string {
	return RandomAlphabet(4) + RandomNumberString(3) + "@gmail.com"
}

func RandomPassword() string {
	return RandomAlphabet(5) + RandomNumberString(5)
}

func RandomFullName() string {
	return RandomAlphabet(4) + " " + RandomAlphabet(4) + " " + RandomAlphabet(4)
}

func RandomNameCategory() string {
	return RandomAlphabet(5)
}

func RandomSlug() string {
	return RandomAlphabet(6) + "-" + RandomAlphabet(8) + "-" + RandomAlphabet(6)
}

func RandomSku() string {
	return RandomAlphabet(3) + "-" + RandomAlphabet(3) + "-" + RandomAlphabet(3) + RandomNumberString(2)
}

func RandomPrice() int64 {
	return RandomNumberInt(8)
}

func RandomUrl() string {
	return "https://" + RandomAlphabet(10) + ".com/image.jpg"
}

func RandomUserID() int64 {
	return RandomNumberInt(2)
}

func RandomProductID() int64 {
	return RandomNumberInt(2)
}

func RandomStatus() string {
	status := []string{"PENDING", "SHIPPING", "DELIVERED"}
	n := len(status)
	return status[rand.Intn(n)]
}

func RandomAddress() string {
	return RandomAlphabet(20)
}

func RandomPhone() string {
	return "0" + RandomNumberString(9)
}

func RandomRoles() string {
	roles := []string{"ADMIN", "STAFF", "CUSTOMER"}
	n := len(roles)
	return roles[rand.Intn(n)]
}

func RandomDescription() string {
	description := []string{"System Administrator with full access rights", "Store Staff with limited operational access", "Default Customer account for general users"}
	n := len(description)
	return description[rand.Intn(n)]
}