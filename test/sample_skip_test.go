package test

import (
	"runtime"
	"testing"
)

func TestSkipFunction(t *testing.T) {
	goos := runtime.GOOS
	if goos != "linux" {
		t.Skip("Skipping this test because it's not implemented for ", goos)
	}

	t.Log("This will not be printed")
}