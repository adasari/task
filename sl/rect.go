package sl

import "fmt"

// Rect .
type Rect struct {
	X int `json:"X"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
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

// String .
func (r Rect) String() string {
	return fmt.Sprintf("(%d, %d), w=%d, h=%d", r.X, r.Y, r.W, r.Y)
}

// Intersect -.
func (r Rect) Intersect(o *Rect) (bool, *Rect) {
	// x not intersecting. TODO : refactor
	if r.LeftX() < o.LeftX() && r.RightX() < o.LeftX() ||
		r.LeftX() > o.RightX() && r.LeftX() > o.LeftX() {
		return false, nil
	}

	/* if r.RightX() < o.LeftX() ||
		r.LeftX() > o.RightX() {
		return false, nil
	} */

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

// IntersectedRect .
type IntersectedRect struct {
	Names []string
	*Rect
}

// Exists .
func (r IntersectedRect) Exists(name string) bool {
	for _, n := range r.Names {
		if name == n {
			return true
		}
	}

	return false
}
