package test

import "testing"

// go test                   run all tests in this package
// go test -v                run all tests with verbose output
// go test -v -run TestName  run a specific test
// go test ./...             run all tests in all packages

// TestSquare demonstrates t.Errorf — logs a formatted message and
// continues execution. The test will fail but keep running.
func TestSquare(t *testing.T) {
	result := Square(5)
	expected := 25

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// TestNegativeSquare demonstrates t.Error — same as Errorf but
// without format-string parsing. Uses plain arguments like fmt.Print.
func TestNegativeSquare(t *testing.T) {
	result := Square(-5)
	expected := 25

	if result != expected {
		t.Error("Expected", expected, "got", result)
	}
}

// TestZeroSquare demonstrates t.Fatalf — logs a formatted message
// then stops the test immediately with FailNow().
// Lines AFTER t.Fatalf / t.Fatal / t.FailNow are NEVER executed.
func TestZeroSquare(t *testing.T) {
	result := Square(0)
	expected := 4 // deliberately wrong — this WILL fail

	if result != expected {
		t.Fatalf("Expected %d, got %d — stopping test now", expected, result)
	}

	// Unreachable because Fatalf triggers FailNow which calls Goexit.
	t.Log("This line will never print")
}

// TestFatal demonstrates t.Fatal — same as Fatalf but without
// format-string. Stops immediately when called.
func TestFatal(t *testing.T) {
	if 1 != 2 {
		t.Fatal("1 should equal 2 — cannot proceed, aborting")
	}

	t.Log("This line will never print")
}

func TestFailBehavior(t *testing.T) {
	// t.Fail() — marks the test as failed but does NOT stop.
	if 1 != 2 {
		t.Fail()
	}
	t.Log("After t.Fail() — execution continues")

	// t.FailNow() — marks the test as failed AND stops.
	if 2 != 3 {
		t.FailNow()
	}

	// Unreachable because FailNow stops the goroutine.
	t.Log("After t.FailNow() — this will never run")
}
