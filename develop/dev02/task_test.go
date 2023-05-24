package string_unpacker

import (
	"testing"
)

func TestStringUnpack(t *testing.T) {
	t.Run("#1", func(t *testing.T) {
		result, err := Unpack("a4bc2d5e")
		if err != nil {
			t.Error(err)
		}
		expected := "aaaabccddddde"

		if result != expected {
			t.Errorf("expected string is not equal to result:\nexpected: %v\result: %v\n", expected, result)
		}
	})

	t.Run("#2", func(t *testing.T) {
		result, err := Unpack("abcd")
		if err != nil {
			t.Error(err)
		}
		expected := "abcd"

		if result != expected {
			t.Errorf("expected string is not equal to result:\nexpected: %v\result: %v\n", expected, result)
		}
	})

	t.Run("#3", func(t *testing.T) {
		_, err := Unpack("45")
		if err == nil {
			t.Errorf("string not detected as invalid")
		}
	})

	t.Run("#4", func(t *testing.T) {
		result, err := Unpack("")
		if err != nil {
			t.Error(err)
		}
		expected := ""

		if result != expected {
			t.Errorf("expected string is not equal to result:\nexpected: %v\result: %v\n", expected, result)
		}
	})
}