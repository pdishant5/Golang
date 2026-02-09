package user

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// -------- SETUP (runs once) --------
	AppMode = "test"
	println(">>> Global test setup complete!")

	// Run all tests
	exitCode := m.Run()

	// -------- TEARDOWN (runs once) --------
	AppMode = ""
	println(">>> Global test teardown complete!")

	// Important: exit with test result
	os.Exit(exitCode)
}
