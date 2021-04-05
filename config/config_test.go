package config

import (
	"fmt"
	"testing"
)

func TestNewResult(t *testing.T) {
	got, _ := New("aaa", "bbb")
	fmt.Println(got)

	t.Error("New connection return: ", got)
}
