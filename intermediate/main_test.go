package main

import "testing"

func TestEx011(t *testing.T) {
	want := "1010,1020"
	got := Ex011("0100,0011,1010,1001,1020")
	if got != want {
		t.Errorf("Ex011() = %v, want %v", got, want)
	}
}
