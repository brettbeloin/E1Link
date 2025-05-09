package main

import (
	"testing"
)

func TestAddInt(t *testing.T) {
	w := singleLinkList[int]()
	w.Add(1)
	expect := "1"

	if w.ToString() != expect {
		t.Errorf("Expected: %v, found: %v", expect, w.ToString())
	}
}

func TestAddEmpty(t *testing.T) {
	w := singleLinkList[string]()
	w.Add("")
	expect := ""

	if w.ToString() != expect {
		t.Errorf("Expected: %v, found: %v", expect, w.ToString())
	}
}
