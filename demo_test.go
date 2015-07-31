package main

import (
	"testing"
)

func TestFormatString(t *testing.T) {
	format := getFormatString()

	if format != "Hi there, I love %s" {
		t.Error("The format string is broken.")
	}
}
