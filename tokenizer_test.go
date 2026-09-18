package main

import "testing"

func TestTokenizerRoundTrip(t *testing.T) {
	tests := []string{
		"cat",
		"café",
		"hello\nworld",
		"你好",
		"🤘",
		"",
	}

	for _, input := range tests {
		encoded := encode(input)
		decoded := decode(encoded)

		if decoded != input {
			t.Errorf("round trip failed: input=%q decoded=%q", input, decoded)
		}
	}
}

func TestEncodeASCII(t *testing.T) {
	input := "cat"
	expected := []byte{99, 97, 116}

	got := encode(input)

	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("encode(%q) = %v, want %v", input, got, expected)
		}
	}
}

func TestAnalyzeBytesTotal(t *testing.T) {
	data := []byte("ttab")

	stats := analyzeBytes(data)

	if stats.TotalBytes != 4 {
		t.Fatalf("expected 4 total bytes, got %d", stats.TotalBytes)
	}

	if stats.Counts['t'] != 2 {
		t.Fatalf("expected a count of 2, got %d", stats.Counts['t'])
	}

	if stats.Probability('t') != 0.5 {
		t.Fatalf("expected probability of 0.5, got %f", stats.Probability('t'))
	}

	if stats.Counts['a'] != 1 {
		t.Fatalf("expected a count of 1, got %d", stats.Counts['a'])
	}

	if stats.Probability('a') != 0.25 {
		t.Fatalf("expected probability of 0.25, got %f", stats.Probability('a'))
	}

}

func TestAnalyzeBytesEmptyCorpus(t *testing.T) {
	data := []byte("")

	stats := analyzeBytes(data)

	if stats.TotalBytes != 0 {
		t.Fatalf("expected 0 total bytes, got %d", stats.TotalBytes)
	}

	if stats.Probability('t') != 0 {
		t.Fatalf("expected probability of 0, got %f", stats.Probability('t'))
	}
}