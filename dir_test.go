package gosh

import "testing"

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
