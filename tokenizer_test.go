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
