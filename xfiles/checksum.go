package files

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

//GetChecksum gets the checksum.
func GetChecksum(data string) string {
	checksum := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", checksum)
}

//CheckIntegrity checks the integrity.
func CheckIntegrity(content string, hash string) error {
	if hash != GetChecksum(content) {
		return errors.New("data has been tampered with")
	}
	return nil
}
