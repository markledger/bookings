package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var mh myHandler
	h := NoSurf(&mh)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Error(fmt.Printf("Type is not http.Handler it is:  %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var mh myHandler
	h := SessionLoad(&mh)
	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Printf("Type is not http.Handler it is:  %T", v))
	}

}
