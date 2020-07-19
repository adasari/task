package main

// Rect .
type Rect struct {
	X, Y, W, H int
}

// LeftX .
func (r Rect) LeftX() int {
	return r.X
}

// RightX .
func (r Rect) RightX() int {
	return r.X + r.W
}

// TopY .
func (r Rect) TopY() int {
	return r.Y
}

// BottomY .
func (r Rect) BottomY() int {
	return r.Y - r.H
}

// intersect -.
func (r Rect) intersect(o *Rect) (bool, *Rect) {
	// x not intersecting
	if r.LeftX() < o.LeftX() && r.RightX() < o.LeftX() ||
		r.LeftX() > o.RightX() && r.LeftX() > o.LeftX() {
		return false, nil
	}

	// y not intersecting
	if r.BottomY() < o.BottomY() && r.TopY() < o.BottomY() ||
		r.BottomY() > o.BottomY() && r.BottomY() > o.TopY() {
		return false, nil
	}

	leftX := max(r.LeftX(), o.LeftX())
	rightX := min(r.RightX(), o.RightX())
	bottomY := max(r.BottomY(), o.BottomY())
	topY := min(r.TopY(), o.TopY())

	// log.Printf("left %d, right %d, bottom %d, top %d \n", leftX, rightX, bottomY, topY)

	return true, &Rect{X: leftX, Y: topY, W: (rightX - leftX), H: (topY - bottomY)}
}

// NamedRect .
type NamedRect struct {
	Name string
	*Rect
}
