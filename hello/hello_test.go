package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Serj")
	want := "Hello, Serj"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
