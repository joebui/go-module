package gomodule

import "testing"

func TestHello(t *testing.T) {
	Hello()
}

func TestNewHello(t *testing.T) {
	NewHello("hello world")
}
