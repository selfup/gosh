package gosh

import (
	"runtime"
	"testing"
)

func TestXosSlash(t *testing.T) {
	separator := Slash()

	if runtime.GOOS == "windows" {
		if separator != "\\" {
			t.Errorf("bad xos (windows) slash")
		}
	} else {
		if separator != "/" {
			t.Errorf("bad xos (non-windows) slash")
		}
	}
}
