package gots

import (
	"testing"
	"time"
)

func TestGetTimeString(t *testing.T) {
	timeT, _ := time.Parse("2006-01-02", "2014-12-31")
	f := DefaultFormat
	timeString := GetTimeString(&timeT, f)
	if timeString != "Dec 31 00:00:00" {
		t.Error("GetTimeString() is wrong")
	}
}

func TestGetMicroTimeString(t *testing.T) {
	timeT, _ := time.Parse("2006-01-02", "2014-12-31")
	f := "%b %d %H:%M:%.S"
	timeString := GetMicroTimeString(&timeT, f)
	if timeString != "Dec 31 00:00:00.000000" {
		t.Error("GetMicroTimeString() is wrong")
	}
}
