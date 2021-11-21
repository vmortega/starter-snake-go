package main

import (
	"log"
	"math/rand"
)

func info() BattlesnakeInfoResponse {
	log.Println("INFO")
	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "vicoar",
		Color:      "#006400",
		Head:       "default",
		Tail:       "default",
	}
}

var games = make(map[string]Game)

func start(state Game) {
	log.Printf("%s START\n", state.ID)
	games[state.ID] = state
}

func end(state Game) {
	log.Printf("%s END\n\n", state.ID)
}

func move(state GameState) BattlesnakeMoveResponse {
	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	// Step 0: Don't let your Battlesnake move back in on it's own neck
	myHead := state.Snakes[0].Coords[0] // Coordinates of your head
	myNeck := state.Snakes[0].Coords[1] // Coordinates of body piece directly behind your head (your "neck")
	if myNeck[0] < myHead[0] {
		possibleMoves["left"] = false
	} else if myNeck[0] > myHead[0] {
		possibleMoves["right"] = false
	} else if myNeck[1] < myHead[1] {
		possibleMoves["down"] = false
	} else if myNeck[1] > myHead[1] {
		possibleMoves["up"] = false
	}

	// TODO: Step 1 - Don't hit walls.
	// Use information in GameState to prevent your Battlesnake from moving beyond the boundaries of the board.
	// boardWidth := state.Board.Width
	// boardHeight := state.Board.Height

	// TODO: Step 2 - Don't hit yourself.
	// Use information in GameState to prevent your Battlesnake from colliding with itself.
	// mybody := state.You.Body

	// TODO: Step 3 - Don't collide with others.
	// Use information in GameState to prevent your Battlesnake from colliding with others.

	// TODO: Step 4 - Find food.
	// Use information in GameState to seek out and find food.

	// Finally, choose a move from the available safe moves.
	// TODO: Step 5 - Select a move to make based on strategy, rather than random.
	var nextMove string

	safeMoves := []string{}
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.GameID, state.Turn, nextMove)
	} else {
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
		log.Printf("%s MOVE %d: %s\n", state.GameID, state.Turn, nextMove)
	}
	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}
