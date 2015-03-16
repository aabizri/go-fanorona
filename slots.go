// This file (slots.go) countains all types & methods for slots

package basic

// A Slot is where a piece CAN go.
type Slot struct {
	Coordinates struct {
		Horizontal uint
		Vertical   uint
	}
	Surrounding map[string]bool
	Piece       *Piece
	Board       *Board
}

func NewSlot(h, v uint) *Slot {
	slot := new(Slot)
	slot.Coordinates.Horizontal = h
	slot.Coordinates.Vertical = v
	return slot
}

// setupSlot(h,v) setups a slot (creates one)
// It is a convenience wrapper for NewSlot() + Slot.setup() + setup of board field
// WARNING: Doesn't populates it with pieces
func SetupSlot(h uint, v uint, b *Board) *Slot {
	slot := NewSlot(h, v)
	slot.setup()
	slot.Board = b
	return slot
}

// Populate the slot with the piece p
func (s *Slot) Populate(p *Piece) {
	// Change the piece record on the slot
	s.Piece = p
	// Change the slot record on the board
	p.Slot = s
}

// Clears the slot, doesn't do anything about the piece on it
func (s *Slot) clear() error {
	s.Piece = nil
	return nil
}

// Checks if coordinates valid
func (s *Slot) IsInside() bool {
	return IsInside(s.Coordinates.Horizontal, s.Coordinates.Vertical)
}

// discoverPattern() discovers the pattern (the allowed movements around a slot), and integrates it to the slot
func (s *Slot) setup() {
	//Shorthands
	var (
		h uint = s.Coordinates.Horizontal
		v uint = s.Coordinates.Vertical
	)

	// If true, there are diagonals
	var diagonable bool = (((v-1)%2 == 0 && (h-1)%2 == 0) || ((v-1)%2 != 0 && (h-1)%2 != 0))

	// First fill all
	s.Surrounding = make(map[string]bool, 8)
	s.Surrounding = map[string]bool{
		"North":     true,
		"NorthEast": true,
		"East":      true,
		"SouthEast": true,
		"South":     true,
		"SouthWest": true,
		"West":      true,
		"NorthWest": true,
	}

	// Eliminate based on position
	// If its on the right edge
	if s.Coordinates.Horizontal == Horizontal-1 {
		s.Surrounding["East"] = false
		if diagonable {
			s.Surrounding["SouthEast"] = false
			s.Surrounding["NorthEast"] = false
		}
	}
	// If its on the left edge
	if s.Coordinates.Horizontal == 0 {
		s.Surrounding["West"] = false
		if diagonable {
			s.Surrounding["SouthWest"] = false
			s.Surrounding["NorthWest"] = false
		}
	}
	// If its on the higher edge
	if s.Coordinates.Vertical == Vertical-1 {
		s.Surrounding["North"] = false
		if diagonable {
			s.Surrounding["NorthEast"] = false
			s.Surrounding["NorthWest"] = false
		}
	}
	// If its on the lower edge
	if s.Coordinates.Vertical == 0 {
		s.Surrounding["South"] = false
		if diagonable {
			s.Surrounding["SouthWest"] = false
			s.Surrounding["SouthEast"] = false
		}
	}
}
