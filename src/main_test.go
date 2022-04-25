package main

import "testing"

func TestHello(t *testing.T) {
    expected := "Listening on port 5050..."
    if got := print(); got != expected {
        t.Errorf("print() = %q, want %q", got, expected)
    }
}
