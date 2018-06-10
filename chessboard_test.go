package main

import (
	"testing"
)

func TestGetMovesBaseCaseStartOfBoard(t *testing.T) {
	var initial Position = Position{2, 2}
	var capacity int = cap(getMoves(initial))
	if capacity != 0 {
		t.Errorf("Too many moves returned, got: %d, want: %d", capacity, 0)
	}
}

func TestGetMovesReturnsValidMoves(t *testing.T) {
	var initial Position = Position{3, 3}
	var capacity int = cap(getMoves(initial))
	if capacity == 0 {
		t.Errorf("Too many moves returned, got: %d, want: %d", capacity, 0)
	}
}

func TestWhoWins(t *testing.T) {
	var can1Win = canForceWin([]Position{{1, 1}})
	if can1Win == true {
		t.Errorf("Player one should lose")
	}
}

func TestWhoWins2(t *testing.T) {
	var can1Win = canForceWin([]Position{{5, 2}})
	if can1Win == true {
		t.Errorf("Player two should win")
	}
}

func TestWhoWins3(t *testing.T) {
	var can1Win = canForceWin([]Position{{5, 3}})
	if can1Win != true {
		t.Errorf("Player one should win")
	}
}
