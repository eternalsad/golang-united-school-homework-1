package solution

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	got := GetMessage()
	want := "Hello 🗺️ !"
	if got != want {
		t.Errorf("Not equal strings, got [%v], expected [%v]", got, want)
	}
}
