package main

import "testing"

func TestSayHello(t *testing.T) {
	want := "Hello, world."
	if got := sayhello(); got != want {
		t.Errorf("got = %s, want = %s\n", got, want)
	}
}
