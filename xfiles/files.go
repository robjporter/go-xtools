package files

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var topath string
var frompath string

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//CheckIfExists returns true if file/dir exists and false otherwise
func CheckIfExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

// DirIsEmpty checks if the given directory is empty.
func DirIsEmpty(dir string) bool {
	var err error
	var f *os.File

	if f, err = os.Open(dir); err == nil {
		var names []string
		if names, err = f.Readdirnames(0); err != nil {
			panic(err)
		}

		if len(names) > 0 {
			return false
		}

		return true
	}

	panic(err)
}

//LookupPath looks the file up in the specified path.
func LookupPath(pathVar string, file string) (string, error) {
	pathSplit := strings.Split(pathVar, ";")
	for _, pathdir := range pathSplit {
		newPath := path.Join(pathdir, file)
		_, err := os.Stat(newPath)
		if err == nil {
			return newPath, nil
		}
	}
	return "", errors.New("file not found")
}

//CopyDir copies a directory to another.
func CopyDir(from string, to string) error {
	var err error
	frompath, err = filepath.Abs(from)
	if err != nil {
		return err
	}
	topath, err = filepath.Abs(to)
	if err != nil {
		return err
	}
	return filepath.Walk(from, visit)
}

func visit(pathName string, fileInfo os.FileInfo, err error) error {
	pathName, err = filepath.Abs(pathName)
	if err != nil {
		return err
	}
	relativePath := strings.Replace(pathName, frompath, "", 1)
	newPath := path.Join(topath, relativePath)
	if fileInfo.IsDir() {
		return os.MkdirAll(newPath, 0700)
	}
	content, err := ioutil.ReadFile(pathName)
	if err != nil {
		return err
	}
	file, err := os.Create(newPath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	return err
}
