// TODO:
// - Make the code compilable (done)
// - Cut the code (done)
// - Write a test suite (oops)
// - Upload to GitHub (done)

// Package basic of fanorana implements the board game basic functions in go to allow YOU to implement this game
//
// To start a game, setup a board:
//
//  board := basic.SetupBoard()
//
// Now move a piece:
//
//	piece := board[3][5]
//	piece.Move(basic.North)
//
// Evaluate the piece:
//
//	piece.Eval()
//
// Is there a winner?
//
//  win,black := board.Win()
//  if win && player==black{
//	  fmt.Println("Yay")
//  }
//
// YOU implement turns, permissions, AI, etc. You can take a look at fanorona-cli which implements a very basic CLI turn-by-turn 2-player game.
//
// Or you use a midddleware that I may never do, but it's on my todo-list (which is very, _very_ long, so don't hold your breath).
//
// The source code is fully my work, noone contributed as of now.
//
// Licensed under the UNLICENSE, written by @nodvos <alexandre@bizri.fr>
package basic

const (
	Horizontal uint = 9 // The horizontal measure of the board
	Vertical   uint = 5 // The vertical measure of the board
)

// An Offset() is the technical info behind a direction, that involves the movement
type Offset struct {
	dir string
	h   int
	v   int
}

// Gives the opposite Offset
func (o Offset) opposite() Offset {
	return Offset{"", -o.h, -o.v}
}

// Standard directions, use them!
var (
	North     = Offset{"North", 0, 1}
	NorthEast = Offset{"NorthEast", 1, 1}
	East      = Offset{"East", 1, 0}
	SouthEast = Offset{"SouthEast", 1, -1}
	South     = Offset{"South", 0, -1}
	SouthWest = Offset{"SouthWest", -1, -1}
	West      = Offset{"West", -1, 0}
	NorthWest = Offset{"NorthWest", -1, 1}
)

// The map of directions
//
// Use cases:
//   - Check whether an inputed direction is valid (check the indexes)
//   - From an given inputed direction, get the corresponding Offset
var Directions map[string]Offset = make(map[string]Offset, 8)

func init() {
	Directions = map[string]Offset{
		"North":     North,
		"NorthEast": NorthEast,
		"East":      East,
		"SouthEast": SouthEast,
		"South":     South,
		"SouthWest": SouthWest,
		"West":      West,
		"NorthWest": NorthWest,
	}
}

// Checks if coordinates are insdie the board
func IsInside(h uint, v uint) bool {
	if 0 <= h && h <= Horizontal-1 && 0 <= v && v <= Vertical-1 {
		return true
	}
	return false
}
