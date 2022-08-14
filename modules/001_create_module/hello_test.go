package main

import "testing"

func TestGreeting(t *testing.T) {
	want := "hello world"
	if got := greeting(); got != want {
		t.Errorf("want = %s, got = %s\n", got, want)
	}
}
