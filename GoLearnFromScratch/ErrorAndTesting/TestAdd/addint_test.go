package main

import "testing"

func TestAddInt(t *testing.T) {
	got, _ := addInt(2, -1)

	want := 5

	if got != want {
		t.Errorf("Expected '%d', but got '%d'", want, got)
	}
}
func TestAddInt_InvalidType(t *testing.T) {
	_, err := addInt(2, "three")
	if err == nil {
		t.Error("Expected error for invalid input types, but got none")
	}
}
