package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	if ReturnHelloWorld() != "Hello World" {
		t.Error("Expected Hello World")
	}
}
