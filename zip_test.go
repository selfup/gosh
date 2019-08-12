package gosh

import (
	"fmt"
	"testing"
)

func TestZipUnZip(t *testing.T) {
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

	unZippedReadme := Fex("archive/README.md")
	if !unZippedReadme {
		t.Errorf("README.md not unzipped from repo.zip")
	}

	unZippedTestPoll := Fex("archive/scripts/cmd/test_poll.go")
	if !unZippedTestPoll {
		t.Errorf("test_poll.go not unzipped from repo.zip")
	}

	RmDir("archive")
	RmDir("repo.zip")
}
