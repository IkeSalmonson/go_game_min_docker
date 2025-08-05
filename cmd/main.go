package main

import (
	"fmt"
	// Import other packages as needed, e.g., board
	"go_game_min_docker/cmd/board"
	"go_game_min_docker/cmd/piece"
)

func main() {
	fmt.Println("Welcome to the Go Game!")
	piece := piece.PrintNewPiece("1", "Black")
	fmt.Println(piece)
	board := board.PrintNewBoard(11, 11)
	fmt.Println(board)

	// Main game initialization and loop will go here.
}
