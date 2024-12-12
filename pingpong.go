package pingpong

import "errors"

// PingPong provides a back-and-forth traversal of a slice of strings.
type PingPong struct {
	items     []string
	index     int
	direction int
}

// NewPingPong initializes a PingPong instance.
// It takes a slice of strings as input and returns a pointer to the PingPong instance.
// Returns an error if the input slice is empty.
func NewPingPong(items []string) (*PingPong, error) {
	if len(items) == 0 {
		return nil, errors.New("items slice cannot be empty")
	}
	return &PingPong{
		items:     items,
		index:     0,
		direction: 1, // 1 for forward, -1 for backward
	}, nil
}

// Next retrieves the current item and moves the index in the appropriate direction.
// The direction reverses when the traversal reaches the start or end of the slice.
func (p *PingPong) Next() string {
	current := p.items[p.index]

	// Update the index based on the direction
	p.index += p.direction

	// Reverse direction at boundaries
	if p.index >= len(p.items) {
		p.index -= 2
		p.direction = -1
	} else if p.index < 0 {
		p.index += 2
		p.direction = 1
	}

	return current
}
