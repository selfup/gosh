package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	currentTime := time.Now()

	scnnr := FindGoFiles{
		Directory: ".",
	}

	poll(scnnr, currentTime)
}

func poll(scnnr FindGoFiles, currentTime time.Time) {
	err := scnnr.Scan()
	if err != nil {
		panic(err)
	}

	for _, file := range scnnr.Found {
		info, err := os.Stat(file)
		if err != nil {
			fmt.Println(err)
		}

		if info.ModTime().Unix() > currentTime.Unix() {
			cmd := exec.Command("go", "test")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			cmdErr := cmd.Run()
			if cmdErr != nil {
				log.Fatal(cmdErr)
			}

			currentTime = time.Now()
		}
	}

	time.Sleep(1 * time.Second)

	poll(scnnr, currentTime)
}

// FindGoFiles holds cli args
type FindGoFiles struct {
	Directory string
	Found     []string
}

// Scan walks the given directory tree
func (f *FindGoFiles) Scan() error {
	err := filepath.Walk(f.Directory, f.scan)
	if err != nil {
		return err
	}

	return nil
}

func (f *FindGoFiles) scan(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !info.IsDir() {
		if filepath.Ext(path) == ".go" {
			f.Found = append(f.Found, path)
		}
	}

	return nil
}
