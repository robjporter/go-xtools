package xcompression

import (
	"encoding/base32"
	"encoding/base64"
)

// Base32Encode base32 encode
func Encode32(decoded string) string {
	return base32.StdEncoding.EncodeToString([]byte(decoded))
}

// Base32Decode base32 decode
func Decode32(encoded string) (string, error) {
	decoded, err := base32.StdEncoding.DecodeString(encoded)
	return string(decoded), err
}

//Encode base64 encodes the string.
func Encode64(decoded string) string {
	return base64.StdEncoding.EncodeToString([]byte(decoded))
}

//Decode decodes the string.
func Decode64(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	return string(decoded), err
}
