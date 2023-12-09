package helper

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomCode() (string, error) {
	const charset = "0123456789"
	codeLength := 4

	randomBytes := make([]byte, codeLength)

	for i := range randomBytes {
		randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		randomBytes[i] = charset[randIndex.Int64()]
	}

	return string(randomBytes), nil
}
