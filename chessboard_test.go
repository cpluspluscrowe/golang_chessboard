package main

import (
	"testing"
)

func TestGetMovesBaseCase(t *testing.T) {
	var initial Position = Position{1, 1}
	var capacity int = cap(getMoves(initial))
	if capacity != 0 {
		t.Errorf("Too many moves returned, got: %d, want: %d", capacity, 0)
	}
}

func TestGetMovesReturnsValidMoves(t *testing.T) {
	var initial Position = Position{5, 5}
	var capacity int = cap(getMoves(initial))
	if capacity == 0 {
		t.Errorf("Too many moves returned, got: %d, want: %d", capacity, 0)
	}
}
