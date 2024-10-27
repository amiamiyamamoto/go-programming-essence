package readdata

import (
	"runtime"
	"testing"
)

func TestReadData(t *testing.T) {
	if runtime.GOOS != "windowns" {
		t.Skip("skipping in %v", runtime.GOOS)
	}
	// test code
}
