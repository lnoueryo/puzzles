package crypto

import "testing"

func TestMakeRandomStr(t *testing.T) {
	_, err := MakeRandomStr(20);if err != nil {
		t.Error("Got an error when we should not have")
	}
}