package webdev

import "testing"

func TestExecuteTpl(t *testing.T) {
	err:= ExecuteTpl()
	if err != nil {
		t.Error("Error")
	}
}
