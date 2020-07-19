package main

type finder struct {
	//combinations []*Rect
	//result       []*Rect
}

// NewFinder .
func NewFinder() *finder {
	return &finder{}
}

func (f *finder) clear() {
	//	f.combinations = nil
	//	f.result = nil
}

func (f *finder) find(input []Rect) []*Rect {
	// clear if finder instance is already used to find the intersections.
	// 	f.clear()
	_, r := f.internalFind(input, 0)

	return r
}

func (f *finder) internalFind(r []Rect, index int) ([]*Rect, []*Rect) {
	var combinations []*Rect
	var intersected []*Rect
	if (len(r) - 2) == index {
		// left with ast two rectagles
		// create combinations = self + new intersect rectangle if exists
		combinations = append(combinations, &r[index], &r[index+1])
		if overlapped, i := overlap(r[index], r[index+1]); overlapped {
			combinations = append(combinations, i)
			intersected = append(intersected, i)
		}

		return combinations, intersected
	}

	cc, ir := f.internalFind(r, index+1)

	combinations = append(combinations, cc...)
	intersected = append(intersected, ir...)
	// check current element have intersections with combinations
	// if A is not intersect B and A not intersect C, there is no point in checking if A intersect withh B and C combination.
	// TODO - find a way to skip above condition.
	for _, c := range cc {
		if overlapped, i := overlap(r[index], *c); overlapped {
			combinations = append(combinations, i)
			intersected = append(intersected, i)
		}
	}

	return combinations, intersected
}

// overlap - can replace with method receiver to Rect struct
func overlap(a, b Rect) (bool, *Rect) {
	// x not intersecting
	if a.LeftX() < b.LeftX() && a.RightX() < b.LeftX() ||
		a.LeftX() > b.RightX() && a.LeftX() > b.LeftX() {
		return false, nil
	}

	// y not intersecting
	if a.BottomY() < b.BottomY() && a.TopY() < b.BottomY() ||
		a.BottomY() > b.BottomY() && a.BottomY() > b.TopY() {
		return false, nil
	}

	leftX := max(a.LeftX(), b.LeftX())
	rightX := min(a.RightX(), b.RightX())
	bottomY := max(a.BottomY(), b.BottomY())
	topY := min(a.TopY(), b.TopY())

	//log.Printf("left %d, right %d, bottom %d, top %d \n", leftX, rightX, bottomY, topY)

	return true, &Rect{X: leftX, Y: topY, W: (rightX - leftX), H: (topY - bottomY)}
}
