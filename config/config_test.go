package config

import "testing"

func TestLoad(t *testing.T) {
	want := true
	got, err := Load()

	if got != want {
		t.Error("Load test return false but want true")
	} else if err != nil {
		t.Error("Load test return a error.")
	}
}
