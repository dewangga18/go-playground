package test

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting Test")
	m.Run()
	fmt.Println("Test Finished")
}