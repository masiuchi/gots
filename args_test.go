package gots

import "testing"

func TestParseArgsWithEmpty(t *testing.T) {
	argsSlice := []string{"gots"}
	args := NewArgs()
	args.ParseArgs(argsSlice)
	if args.Rel != false {
		t.Error("rel is not false")
	}
	if args.UseFormat != false {
		t.Error("useFormat is not false")
	}
	if args.Format != DefaultFormat {
		t.Error("format is not default")
	}
}

func TestParseArgsWithRel(t *testing.T) {
	argsSlice := []string{"gots", "-r"}
	args := NewArgs()
	args.ParseArgs(argsSlice)
	if args.Rel != true {
		t.Error("rel is not true")
	}
	if args.UseFormat != false {
		t.Error("useFormat is not false")
	}
	if args.Format != DefaultFormat {
		t.Error("format is not default")
	}
}

func TestParseArgsWithFormat(t *testing.T) {
	argsSlice := []string{"gots", "%Y"}
	args := NewArgs()
	args.ParseArgs(argsSlice)
	if args.Rel != false {
		t.Error("rel is not false")
	}
	if args.UseFormat != true {
		t.Error("useFormat is not true")
	}
	if args.Format != "%Y" {
		t.Error("format is not \"%Y\"")
	}
}

func TestParseArgsWithRelAndFormat(t *testing.T) {
	argsSlice := []string{"gots", "-r", "%Y"}
	args := NewArgs()
	args.ParseArgs(argsSlice)
	if args.Rel != true {
		t.Error("rel is not true")
	}
	if args.UseFormat != true {
		t.Error("useFormat is not true")
	}
	if args.Format != "%Y" {
		t.Error("format is not \"%Y\"")
	}
}

func TestIsHiresFalse(t *testing.T) {
	argsSlice := []string{"gots"}
	args := NewArgs()
	args.ParseArgs(argsSlice)
	if args.IsHires() != false {
		t.Error("IsHires() is not false")
	}
}

func TestIsHiresTrue(t *testing.T) {
	argsSlice := []string{"gots", "%.S"}
	args := NewArgs()
	args.ParseArgs(argsSlice)
	if args.IsHires() != true {
		t.Error("IsHires() is not true")
	}
}
