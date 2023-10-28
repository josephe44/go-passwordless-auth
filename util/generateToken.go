package util

import (
	"crypto/rand"
	"math/big"
)

func GenerateOTPCode() (string, error) {
	seed := "012345679"
	byteSlice := make([]byte, 6)

	for i := 0; i < 6; i++ {
		max := big.NewInt(int64(len(seed)))
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		byteSlice[i] = seed[num.Int64()]
	}

	return string(byteSlice), nil
}
