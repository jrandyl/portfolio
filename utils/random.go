package utils

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// Randomly generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Randomly generates a name
func RandomName() string {
	return RandomString(6)
}

// Generates random price for the orders
func RandomNumber() int64 {
	return RandomInt(0, 10000000)
}

// Generates random rating for the orders
func RandomRating() float64 {
	rating := float64(RandomInt(0, 10000000))
	return rating
}

// Generates random status
func RandomStatus() string {
	status := []string{"Offline", "Online"}
	n := len(status)
	return status[rand.Intn(n)]
}

// Generates random order_status
func RandomOrderStatus() string {
	status := []string{"Pending", "Paid", "Cancelled"}
	n := len(status)
	return status[rand.Intn(n)]
}
