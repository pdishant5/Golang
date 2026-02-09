package stringutils

import "testing"

func BenchmarkReverse(b *testing.B) {
	input := "GoLangBenchmarkTesting"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reverse(input)
	}
}

func BenchmarkReverseUsingBuilder(b *testing.B) {
	input := "GoLangBenchmarkTesting"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseUsingBuilder(input)
	}
}

func BenchmarkReverse_TableDriven(b *testing.B) {
	tests := []struct {
		name  string
		input string
	}{
		{"short string", "go"},
		{"medium string", "golangbenchmark"},
		{"long string", "thisisaverylongstringusedforbenchmarktesting"},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Reverse(tt.input)
			}
		})
	}
}
