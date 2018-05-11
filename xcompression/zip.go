package compression

import (
	"bytes"
	"compress/gzip"
)

//GZip gzips data.
func GZip(data string) string {
	var buffer bytes.Buffer
	gzip := gzip.NewWriter(&buffer)
	gzip.Write([]byte(data))
	gzip.Close()
	return buffer.String()
}

//UnGZip ungzips data.
func UnGZip(data string) string {
	var buffer bytes.Buffer
	buffer.Write([]byte(data))
	gzip, _ := gzip.NewReader(&buffer)
	var outputBytes bytes.Buffer
	outputBytes.ReadFrom(gzip)
	gzip.Close()
	return outputBytes.String()
}
