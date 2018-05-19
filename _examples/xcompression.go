package main

import (
	"fmt"

	"../xcompression"
)

func main() {
	test := "This is a random string to test the features of xcompression = test!"
	fmt.Println("== BASE64 =================================================================")
	encoded := xcompression.Encode64(test)
	fmt.Println("ENCODE BASE64: ", encoded)
	decoded, err := xcompression.Decode64(encoded)
	fmt.Println("DECODED BASE64: ", decoded, err)
	encoded = xcompression.Encode32(test)
	fmt.Println("ENCODE BASE32: ", encoded)
	decoded, err = xcompression.Decode32(encoded)
	fmt.Println("DECODED BASE32: ", decoded, err)
	fmt.Println("== COMPRESSION ============================================================")
	compressed, err2 := xcompression.CompressString(test)
	fmt.Println("COMPRESSED: ", compressed, err2)
	decompressed, err3 := xcompression.Decompress(compressed)
	fmt.Println("DECOMPRESSED: ", decompressed, err3)
	fmt.Println("== GZIP ===================================================================")
	zipped := xcompression.GZip(test)
	fmt.Println("GZIP: ", zipped)
	unzipped := xcompression.UnGZip(zipped)
	fmt.Println("UNGZIP: ", unzipped)
}
