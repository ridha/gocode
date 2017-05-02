package stringutil

import "testing"

func TestReverse(t *testing.T) {
	data := []struct{ in, expected string }{
		{"123 ", " 321"},
		{"Hello, world", "dlrow ,olleH"},
		{"", ""},
	}
	for _, i := range data {
		got := Reverse(i.in)
		if i.expected != got {
			t.Errorf("Reverse(%q) == %q, Expected: %q", i.in, got, i.expected)
		}

	}
}

func TestRotate(t *testing.T) {
	in, expected := "ABCDEFGHI", "DEFGHIABC"
	if got := Rotate(in, 3); got != expected {
		t.Errorf("Rotate(%q) == %q, Expected: %q", in, got, expected)
	}

	in, expected = "", ""
	if got := Rotate(in, 3); got != expected {
		t.Errorf("Rotate(%q) == %q, Expected: %q", in, got, expected)
	}

}
