package solution

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	got := GetMessage()
	want := string([]rune{72, 101, 108, 108, 111, 32, 58, 119, 111, 108, 114, 100, 95, 109, 97, 112, 58, 33})
	if got != want {
		t.Errorf("Not equal strings, got [%v], expected [%v]", got, want)
	}
}
