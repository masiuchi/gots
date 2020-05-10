package gots

import "testing"

func Test1(t *testing.T) {
	args := []string{}
	a := NewArgs()
	a.ParseArgs(args)
	if a.Rel != false {
		t.Error("rel is not false")
	}
	if a.UseFormat != false {
		t.Error("useFormat is not false")
	}
	if a.Format != DefaultFormat {
		t.Error("format is not empty")
	}
}
