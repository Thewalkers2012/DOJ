package random

import (
	"math/rand"
	"strings"
	"time"
)

// Accept                  = 0
// WrongAnswer             = -1
// TimeLimitExceeded       = 1
// RunTimeError            = 4
// CompileError            = 6

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const number = "0123456789"

const (
	easy       = "简单"
	middle     = "中等"
	difficulty = "困难"
)

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

// RandomDiffcultyLevel generage a random level
func RandomDiffcultyLevel() string {
	k := RandomInt(0, 2)
	if k == 0 {
		return easy
	} else if k == 1 {
		return middle
	} else {
		return difficulty
	}
}

// RandomAnswer generate a random answer
func RandomAnswer() int {
	state := []int{-1, 0, 1, 4, 6}

	k := len(state)
	code := state[RandomInt(0, int64(k-1))]

	return code
}

// RandomLanguage generate a random language
func RandomLanguage() string {
	k := RandomInt(0, 2)
	if k == 0 {
		return "cpp"
	} else if k == 1 {
		return "java"
	} else {
		return "python"
	}
}
