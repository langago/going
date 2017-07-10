package strings

import (
	"testing"
)

func TestHasSuffix(t *testing.T) {
	s := "foo.bar"

	if !HasSuffix(s, "hello", "golang", ".bar") {
		t.Error("error occurred at HasSuffix")
	}
}