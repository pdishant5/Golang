package parser

import "testing"

func FuzzParseInt(f *testing.F) {
	// Seed corpus (important)
	f.Add("123")
	f.Add("0")
	f.Add("-99")
	f.Add("999999999999")

	f.Fuzz(func(t *testing.T, input string) {
		_ = ParseInt(input)
	})
}
