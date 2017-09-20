package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if add(1, 2) == 3 {
		t.Log("1+2=3")
	}

	if add(1, 1) == 3 {
		t.Error("1+1=3")
	}
}
