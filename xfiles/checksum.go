package xfiles

import (
	"crypto/sha256"
	"fmt"
)

//GetChecksum gets the checksum.
func GetChecksum(data string) string {
	checksum := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", checksum)
}

//CheckIntegrity checks the integrity.
func CheckIntegrity(content string, hash string) bool {
	if hash != GetChecksum(content) {
		return false
	}
	return true
}
