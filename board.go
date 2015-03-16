// This file (board.go) countains all types & methods for boards

package basic

// A Board is obviously what you play on, silly (wo)man
type Board [Horizontal][Vertical]*Slot

// Returns a clean board
func NewBoard() *Board {
	var board *Board = new(Board)
	for h := uint(0); h < Horizontal; h++ {
		for v := uint(0); v < Vertical; v++ {

			// Create the slots
			slot := SetupSlot(h, v, board)

			// Place the slots
			board[h][v] = slot
		}
	}
	return board
}

// Populates the board
func (board *Board) Setup() {
	for h := uint(0); h < Horizontal; h++ {
		for v := uint(0); v < Vertical; v++ {

			// Pick the slot
			slot := board[h][v]

			// Place the pieces, regular first, irregular after that
			if v != 2 { // If it's not middle
				if v <= 1 {
					slot.Populate(newPiece(false))
				}
				if v >= 3 {
					slot.Populate(newPiece(true))
				}
			}
			if v == 2 && h != 4 { // If it's on the middle horizontal line
				if h == 0 || h == 2 || h == 5 || h == 7 {
					slot.Populate(newPiece(true))
				} else {
					slot.Populate(newPiece(false))
				}
			}
		}
	}
}

// Convenience wrapper for NewBoard() & Board.Setup()
func SetupBoard() *Board {
	b := NewBoard()
	b.Setup()
	return b
}

// Reset() resets the board
func (b *Board) Reset() {
	b = NewBoard()
	b.Setup()
}

// Count() counts the number of white & black pieces on the board
func (b *Board) Count() (total, n_white, n_black uint) {
	for h := uint(0); h < Horizontal; h++ {
		for v := uint(0); v < Vertical; v++ {
			s := b[h][v]
			if s.Piece != nil {
				total += 1
				if s.Piece.Black {
					n_black += 1
				}
			}
		}
	}
	return total, total - n_black, n_white
}

// Returns if there's a win and then who won
func (b *Board) Win() (finished, black bool) {
	total, _, n_black := b.Count()
	if total == 1 {
		finished = true
		if n_black == 1 {
			black = true
		} else {
			black = false
		}
	}
	return finished, black
}
