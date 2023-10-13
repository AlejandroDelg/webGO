package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH MyHandler
	h := NoSurf(&myH)
	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintln("type is not http.Handler ", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH MyHandler
	h := SessionLoad(&myH)
	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintln("type is not http.Handler ", v))
	}
}
