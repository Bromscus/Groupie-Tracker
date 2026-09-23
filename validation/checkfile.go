package validation

import (
	"crypto/sha256"
	"os"
)

func CheckFile(fileName string, expectedHash [32]byte) bool {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return false
	}

	currentHash := sha256.Sum256(data)

	return currentHash == expectedHash
}
