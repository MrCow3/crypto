package crypto

import "testing"

func testDecodeHex(t *testing.T) {
	got := DecodeHex("48656c6c6f20576f726c64")
	want := "Hello World"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
