package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

var transformations [][]int = [][]int{
	{-2, 1},
	{-2, -1},
	{1, -2},
	{-1, -2},
}

func applyTransformation(initial Position, index int) (Position, error) {
	initial.Y += transformations[index][0]
	initial.X += transformations[index][1]
	if initial.X > 15 || initial.Y > 15 || initial.X < 1 || initial.Y < 1 {
		return initial, errors.New("out of bounds")
	} else {
		return initial, nil
	}
}

func getMoves(initial Position) []Position {
	var moves []Position
	for index := 0; index < 4; index++ {
		var newPosition, error = applyTransformation(initial, index)
		if error == nil {
			moves = append(moves, newPosition)
		}
	}
	return moves
}

func canForceWinGame(positions Position) bool {
	var nextMoves = getMoves(positions)
	for _, move := range nextMoves { // check for winning condition
		if !canForceWinGame(move) {
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	var input, _, _ = reader.ReadLine()
	var numberOfLines, _ = strconv.Atoi(string(input))
	for i := 0; i < numberOfLines; i++ {
		var startingPosition, _, _ = reader.ReadLine()
		var xy []string = strings.Split(string(startingPosition), " ")
		var x, _ = strconv.Atoi(xy[0])
		var y, _ = strconv.Atoi(xy[1])
		var position = Position{x, y}
		if canForceWinGame(position) == true {
			fmt.Println("First")
		} else {
			fmt.Println("Second")
		}
	}
}
