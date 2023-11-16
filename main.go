package main

// Welcome to
// __________         __    __  .__                               __
// \______   \_____ _/  |__/  |_|  |   ____   ______ ____ _____  |  | __ ____
//  |    |  _/\__  \\   __\   __\  | _/ __ \ /  ___//    \\__  \ |  |/ // __ \
//  |    |   \ / __ \|  |  |  | |  |_\  ___/ \___ \|   |  \/ __ \|    <\  ___/
//  |________/(______/__|  |__| |____/\_____>______>___|__(______/__|__\\_____>
//
// This file can be a nice home for your Battlesnake logic and helper functions.
//
// To get you started we've included code to prevent your Battlesnake from moving backwards.
// For more info see docs.battlesnake.com

import (
	"fmt"
	"log"
	"math/rand"
)

// info is called when you create your Battlesnake on play.battlesnake.com
// and controls your Battlesnake's appearance
func info() BattlesnakeInfoResponse {
	log.Println("INFO")

	colour := fmt.Sprintf("#%02X%02X%02X", rand.Intn(255), rand.Intn(255), rand.Intn(255))

	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "chris-hoefgen",
		Color:      colour,
		Head:       snakeHeads[rand.Intn(len(snakeHeads))],
		Tail:       snakeTails[rand.Intn(len(snakeTails))],
	}
}

// start is called when your Battlesnake begins a game
func start(state GameState) {
	log.Println("GAME START")
}

// end is called when your Battlesnake finishes a game
func end(state GameState) {
	log.Printf("GAME OVER\n\n")
}

func applyMove(move Move, snake *Battlesnake) Coord {
	return Coord{X: snake.Head.X + move.Direction.X, Y: snake.Head.Y + move.Direction.Y}
}


// move is called on every turn and returns your next move
// Valid moves are "up", "down", "left", or "right"
// See https://docs.battlesnake.com/api/example-move for available data
func move(state GameState) BattlesnakeMoveResponse {
	gameBoard, err := NewGameBoard(&state.Board, &state.You)

	if err != nil {
		log.Fatal(err)
		return BattlesnakeMoveResponse{Move: "up"}
	}

	left := Move{Name: "left", Direction: Coord{ X: -1, Y: 0}}
	right := Move{Name: "right", Direction: Coord{ X: 1, Y: 0}}
	up := Move{Name: "up", Direction: Coord{ X: 0, Y: 1}}
	down := Move{Name: "down", Direction: Coord{ X: 0, Y: -1}}

	allMoves := []Move{left, right, up, down}
	moveOptions := []Move{}
	copy(moveOptions, allMoves)

	for i := 0; i < len(allMoves); i++ {
		newPos := applyMove(allMoves[i], &state.You)

		if gameBoard.IsOutOfBounds(newPos) {
			continue
		} else if gameBoard.IsASnake(newPos) {
			continue
		}
		
		// A legal move, add it to the candidate list
		moveOptions = append(moveOptions, allMoves[i])
	}

	if len(moveOptions) == 0 {
		log.Printf("MOVE %d: No safe moves detected! Moving down\n", state.Turn)
		return BattlesnakeMoveResponse{Move: "down"}
	}

	// Choose a random move from the safe ones
	nextMove := moveOptions[rand.Intn(len(moveOptions))]

	// TODO: Step 4 - Move towards food instead of random, to regain health and survive longer
	// food := state.Board.Food

	log.Printf("MOVE %d: %s\n", state.Turn, nextMove.Name)
	return BattlesnakeMoveResponse{Move: nextMove.Name}
}

func main() {
	RunServer()
}
