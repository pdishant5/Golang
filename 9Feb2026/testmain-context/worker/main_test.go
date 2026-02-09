package worker

import (
	"context"
	"os"
	"testing"
)

var rootCtx context.Context

func TestMain(m *testing.M) {
	rootCtx = context.Background()
	println(">>> Root test context created!")

	exitCode := m.Run()

	// -------- TEARDOWN --------
	println(">>> Test suite finished!")

	os.Exit(exitCode)
}
