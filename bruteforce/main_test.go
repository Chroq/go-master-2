package main

import "testing"

func TestAkali(t *testing.T) {
	expected := "@kAl1"
	got := findPasswordNaive("@kAl1")

	if got != expected {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}

func TestShen(t *testing.T) {
	expected := "Sh3n"
	got := findPasswordNaive("Sh3n")

	if got != expected {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}

func TestZed(t *testing.T) {
	expected := "z3D"
	got := findPasswordNaive("z3D")

	if got != expected {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}
