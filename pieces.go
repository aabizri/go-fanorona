// This file (pieces.go) countains all types & methods for pieces & their moves

package basic

import (
	"errors"
)

// A Piece is either the black or white piece used
//
// It countains references to the current slot & the color
type Piece struct {
	Black bool
	Slot  *Slot
}

// A list of pieces
type PieceList []*Piece

// Creates a new piece
func newPiece(black bool) *Piece {
	piece := &Piece{Black: black}
	return piece
}

// Clears p.Slot
func (p *Piece) uproot() {
	p.Slot = nil
}

// Eliminates the piece from the slot
func (p *Piece) eliminate() {
	p.Slot.clear() // Clear the piece record on the slot
	p.Slot = nil   // Clear the slot record on the piece
}

// Checks if a piece can be moved in a direction
func (p *Piece) CanMove(direction Offset) bool {

	// If it can't move there, return false
	if !p.Slot.Surrounding[direction.dir] {
		return false
	}

	// if there's already a piece there as well
	var (
		virtualNewPosition_h uint = uint(int(p.Slot.Coordinates.Horizontal) + direction.h)
		virtualNewPosition_v uint = uint(int(p.Slot.Coordinates.Vertical) + direction.v)
	)
	destination := p.Slot.Board[virtualNewPosition_h][virtualNewPosition_v]
	if destination.Piece != nil {
		return false
	}

	return true
}

// Moves without evaluation a piece following a direction
func (p *Piece) Move(direction Offset) error {
	if !p.CanMove(direction) {
		return errors.New("p.Move(): Cannot move there\n")
	}

	// Find the destination slot
	destination := p.Slot.Board[uint(int(p.Slot.Coordinates.Horizontal)+direction.h)][uint(int(p.Slot.Coordinates.Vertical)+direction.v)]

	// Clear & repopulate
	p.Slot.clear()          // Clear the slot
	destination.Populate(p) // Populate the destination slot with the piece

	return nil
}

// Evaluates the consequences of a move for other pieces
/* Illustration:

[-][-][-][-][-][-][-][s][-]
[-][-][-][-][-][-][s][-][-]
[-][-][s][-][-][p][-][-][-]
[-][-][-][-][-][-][-][-][-]
[-][-][-][-][-][-][-][-][s]

South-west move would do that:

[-][-][-][-][-][-][-][d][-]
[-][-][-][-][-][-][d][-][-]
[-][-][s][-][-][-][-][-][-]
[-][-][-][-][p][-][-][-][-]
[-][-][-][-][-][-][-][-][s]

As such it returns a list of the pieces affected (d) on illustration

*/
// As such it requires the "destroyer option" a.k.a the sameDirection bool, that is, in the case that there is two ways pieces can be destroyed, to chose the right one
func (p *Piece) Eval(move Offset, sameDirection bool) (PieceList, error) {

	b := p.Slot.Board

	// Hold the pieces destroyed
	pieces := make(PieceList, 0)

	// If it's the same direction start from current place + move , if it's not, start from the previous place + move (invert the original move)
	var ray Offset // THe ray of destruction
	var (
		h uint = p.Slot.Coordinates.Horizontal
		v uint = p.Slot.Coordinates.Vertical
	)
	if sameDirection {
		ray = move
		h = uint(int(h) + ray.h)
		v = uint(int(v) + ray.v)
	} else {
		ray = move.opposite()
		h = uint(int(h) + 2*ray.h)
		v = uint(int(v) + 2*ray.v)
	}

	// For as long as we're within the game board, continue
	for IsInside(h, v) {
		inspected := b[h][v].Piece
		if inspected == nil {
			// Then there's no piece there, so break the loop
			break
		} else {
			// If the two pieces are of the same color, break the loop
			if inspected.Black == p.Black {
				break
			} else {
				pieces = append(pieces, inspected)
			}
		}
		h = uint(int(h) + ray.h)
		v = uint(int(v) + ray.v)
	}
	return pieces, nil
}

// Convenience wrapper around Piece.Move(), Piece.Eval() & PieceList.Eliminate()
//
// Moves & eliminates
func (p *Piece) MovEval(direction Offset, sameDirection bool) error {
	err := p.Move(direction)
	if err != nil {
		return err
	}

	pieces, err := p.Eval(direction, sameDirection)
	if err != nil {
		return err
	}

	pieces.Eliminate()

	return err
}

// Eliminate all pieces in the list
func (p PieceList) Eliminate() {
	for _, k := range p {
		k.eliminate()
	}
}
