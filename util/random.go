package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//RandomInt generates a ramdon integer between min and maximum
func RamdonInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)

	}
	return sb.String()
}

// func RandomName() string{
// 	return RandomString(8)
// }

//Generates random role from roles
func RandomRole() string {
	roles := []string{"ADMIN", "AGT", "SP"}
	n := len(roles)
	return roles[rand.Intn(n)]
}
