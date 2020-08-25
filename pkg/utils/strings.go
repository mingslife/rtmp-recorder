package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"

	"github.com/lithammer/shortuuid"
)

const (
	base       = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	baseLength = len(base)
)

func NewId() string {
	return shortuuid.New()
}

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func RandString(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = base[rand.Intn(baseLength)]
	}
	randString := string(b)
	return randString
}
