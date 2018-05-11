package xhttp

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadFile(url string) (string, error) {
	splittedFileName := strings.Split(url, "/")
	fileName := splittedFileName[len(splittedFileName)-1]
	fmt.Println("Downloading ", fileName, " ... ")
	output, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer output.Close()
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if _, err := io.Copy(output, response.Body); err != nil {
		return "", err
	}
	return fileName, nil
}
