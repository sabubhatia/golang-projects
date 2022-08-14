package major

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if get := hello(); get != want {
		t.Errorf("Got = %s Want = %s\n", get, want)
	}
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := proverb(); got != want {
		t.Errorf("Got = %s Want = %s\n", got, want)
	}
}
