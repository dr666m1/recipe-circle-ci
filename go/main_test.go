package main

import "testing"

func Test_hello(t *testing.T) {
	result := hello("dr666m1")
	if result != "hello dr666m1!" {
		t.Error("incorrect result!")
	}
}
