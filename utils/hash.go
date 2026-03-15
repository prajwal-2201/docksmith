package utils

import (
	"crypto/sha256"
	"fmt"
)

func ComputeDigest(data []byte) string {

	hash := sha256.Sum256(data)

	return fmt.Sprintf("sha256:%x", hash)
}
