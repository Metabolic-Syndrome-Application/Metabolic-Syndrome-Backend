package utils

import (
	"math/rand"
)

func GenerateOTP(length int) string {
	chars := "0123456789"
	otp := make([]byte, length)
	for i := range otp {
		otp[i] = chars[rand.Intn(len(chars))]
	}
	return string(otp)
}
