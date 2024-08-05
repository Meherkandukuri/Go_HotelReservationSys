package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// as returned type is http.Handler we are fine with the output

	default:
		t.Errorf("Type mismatch: Expected http.Handler, got %T", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// as returned type is http.Handler we are fine with the output

	default:
		t.Errorf("Type mismatch: Expected http.Handler, got %T", v)
	}
}
