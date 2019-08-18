package gosh

import (
	"fmt"
	"runtime"
	"testing"
)

func TestZipUnZip(t *testing.T) {
	var separator string

	if runtime.GOOS == "windows" {
		separator = "\\"
	} else {
		separator = "/"
	}

	RmDir("archive")
	RmDir("repo.zip")

	zipErr := Zip(".", "repo.zip")
	if zipErr != nil {
		t.Errorf("failed to Zip")
	}

	zipExist := Fex("repo.zip")
	if !zipExist {
		t.Errorf("repo.zip not created")
	}

	unZipErr := Unzip("repo.zip", "archive")
	if unZipErr != nil {
		fmt.Println(unZipErr)
		t.Errorf("failed to UnZip")
	}

	unZippedReadme := Fex("archive" + separator + "README.md")
	if !unZippedReadme {
		t.Errorf("README.md not unzipped from repo.zip")
	}

	RmDir("archive")
	RmDir("repo.zip")
}
