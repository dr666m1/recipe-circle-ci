package main

import "testing"

func Test_hello(t *testing.T) {
	result := hello("world")
	if result != "hello world!" {
		t.Error("incorrect result!")
	}
}
