package random

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const number = "0123456789"

func init() {
	rand.Seed(time.Now().Unix())
}

// RandomInt generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomStringWithLetter generate a random string of length n by alphabet
func RandomStringWithLetter(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomStringWithNumber generage a random string of length n by number
func RandomStringWithNumber(n int) string {
	var sb strings.Builder
	k := len(number)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomStudentID generage a random StudentID
func RandomStudentID() string {
	return RandomStringWithNumber(8)
}

// RandomUserName generage a random Username
func RandomUserName() string {
	return RandomStringWithLetter(8)
}

// RandomPassword generate a random password
func RandomPassword() string {
	length := RandomInt(6, 12)
	return RandomStringWithLetter(int(length))
}
