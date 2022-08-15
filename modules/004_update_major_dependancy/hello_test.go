package major

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := hello(); got != want {
		t.Errorf("got = %#v want = %#v", got, want)
	}
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := proverb(); got != want {
		t.Errorf("got = %#v want = %#v", got, want)
	}
}
