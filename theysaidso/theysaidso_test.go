package theysaidso

import (
	"testing"
)

func TestPrintQuoteOfTheDay(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	PrintQuoteOfTheDay()
}

func BenchmarkPrintQuoteOfTheDay(b *testing.B) {
	for i := 0; i <b.N; i++ {
		PrintQuoteOfTheDay()
	}
}