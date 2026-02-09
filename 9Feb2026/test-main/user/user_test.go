package user

import "testing"

func TestGetAppMode(t *testing.T) {
	mode := GetAppMode()
	if mode != "test" {
		t.Fatalf("Expected app mode: 'test'; Got: '%s'", mode)
	}
}

func TestIsProduction(t *testing.T) {
	if IsProduction() {
		t.Fatal("Expected IsProduction to return false in test mode!")
	}
}
