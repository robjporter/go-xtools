package compression

import "encoding/base64"

//Encode base64 encodes the string.
func Encode(decoded string) string {
	return base64.StdEncoding.EncodeToString([]byte(decoded))
}

//Decode decodes the string.
func Decode(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	return string(decoded), err
}
