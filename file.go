package gosh

import (
	"io/ioutil"
	"os"
)

// Wr writes a file with given contents and filemode
func Wr(destination string, contents []byte, mode os.FileMode) error {
	return ioutil.WriteFile(destination, contents, mode)
}

// MkDirAll is like mkdir -p
func MkDirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// Mv moves source to destination
func Mv(source string, destionation string) error {
	return os.Rename(source, destionation)
}

// Cp copies a source to a destination
func Cp(source string, destination string) error {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(destination, input, 0777)
	if err != nil {
		return err
	}

	return nil
}

// Fex checks if a file exists
func Fex(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Rm removes given file/dir from the filesystem
func Rm(filePath string) error {
	return os.Remove(filePath)
}

// RmDir deletes a Directory with all of it's children
func RmDir(dirPath string) error {
	return os.RemoveAll(dirPath)
}

// RmDirChildren only deletes the children of a directory
func RmDirChildren(dirPath string) error {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, fileInDir := range dir {
		RmDir(dirPath + "/" + fileInDir.Name())
	}

	return nil
}

// Chmod takes in a slice of files to change modifications on
func Chmod(files []string, mode os.FileMode) error {
	for _, file := range files {
		err := os.Chmod(file, mode)
		if err != nil {
			return err
		}
	}

	return nil
}
