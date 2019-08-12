package gosh

import (
	"testing"
)

func TestMvCpFex(t *testing.T) {
	mkDirAllErr := MkDirAll("fixtures")
	if mkDirAllErr != nil {
		t.Errorf("failed to MkDirAll")
	}

	wrErr := Wr("fixtures/file.txt", []byte("contents"), 0777)
	if wrErr != nil {
		t.Errorf("failed to Wr")
	}

	fileTxt := Fex("fixtures/file.txt")
	if !fileTxt {
		t.Errorf("needed fixture/file.txt does not exist")
	}

	cpErr := Cp("fixtures/file.txt", "fixtures/cp_file.txt")
	if cpErr != nil {
		t.Errorf("failed to Cp file")
	}

	cpFileTxt := Fex("fixtures/cp_file.txt")
	if !cpFileTxt {
		t.Errorf("needed fixture/cp_file.txt does not exist")
	}

	mvErr := Mv("fixtures/cp_file.txt", "fixtures/mv_file.txt")
	if mvErr != nil {
		t.Errorf("failed to Mv file")
	}

	mvFileTxt := Fex("fixtures/mv_file.txt")
	if !mvFileTxt {
		t.Errorf("needed fixture/mv_file.txt does not exist")
	}

	rmErr := Rm("fixtures/mv_file.txt")
	if rmErr != nil {
		t.Errorf("failed to Rm file")
	}

	rmFileTxt := Fex("fixtures/mv_file.txt")
	if rmFileTxt {
		t.Errorf("fixture cleanup was not successful")
	}
}

func TestRmDirs(t *testing.T) {
	Wr("fixtures/rmDirChildOne.txt", []byte("contents"), 0777)
	Wr("fixtures/rmDirChildTwo.txt", []byte("contents"), 0777)
	Wr("fixtures/rmDirChildThree.txt", []byte("contents"), 0777)

	rdcErr := RmDirChildren("fixtures")
	if rdcErr != nil {
		t.Errorf("RmDirChildren failed")
	}

	fexOne := Fex("fixtures/rmDirChildOne.txt")
	if fexOne {
		t.Errorf("rdc failed to remove first child")
	}

	fexTwo := Fex("fixtures/rmDirChildTwo.txt")
	if fexTwo {
		t.Errorf("rdc failed to remove second child")
	}

	fexThree := Fex("fixtures/rmDirChildThree.txt")
	if fexThree {
		t.Errorf("rdc failed to remove third child")
	}
}
