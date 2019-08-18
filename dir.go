package gosh

import (
	"io/ioutil"
	"os"
)

// MkDir is like mkdir -p
func MkDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
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
		RmDir(dirPath + Slash() + fileInDir.Name())
	}

	return nil
}
