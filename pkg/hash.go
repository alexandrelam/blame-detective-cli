package pkg

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

func GenerateRandomHash() string {
	// Generate a random number to use as input for hash
	randNum, _ := rand.Int(rand.Reader, big.NewInt(1000000))

	// Convert the random number to bytes
	randBytes := []byte(randNum.String())

	// Compute the SHA-256 hash of the random bytes
	hash := sha256.New()
	hash.Write(randBytes)
	hashBytes := hash.Sum(nil)

	// Convert the hash bytes to a hexadecimal string
	hashString := hex.EncodeToString(hashBytes)

	// keep only the first 8 characters
	hashString = hashString[:8]

	return hashString
}
