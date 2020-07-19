package sl

import (
	"container/heap"
)

// Finder .
type Finder struct{}

// NewFinder .
func NewFinder() *Finder {
	return &Finder{}
}

// FindIntersections .
func (f Finder) FindIntersections(input []NamedRect) []*IntersectedRect {
	// initialize the piority queue.
	pq := &PriorityQueue{}
	heap.Init(pq)

	for _, r := range input {
		//	log.Printf("input : %s, %+v", r.Name, r.Rect)
		heap.Push(pq, Point{flag: true, priority: r.LeftX(), NamedRect: r})
		heap.Push(pq, Point{flag: false, priority: r.RightX(), NamedRect: r})
	}

	// best to use linked list ? // TODO: revisit.
	var activeRectangles []*NamedRect
	var result []*IntersectedRect
	var activeIntersectedRecangles []*IntersectedRect

	for pq.Len() > 0 {
		p := heap.Pop(pq).(Point)
		if p.flag {
			// start point
			// if there are no active rectangles ?
			// case 1 : all rectangles are ended.
			// case 2 : no rectangles are before.
			// in both cases active rectangles and intersected rectangles shoud be nil to avoid checking additional intersection combinations.
			if len(activeRectangles) == 0 {
				activeRectangles = nil
				activeIntersectedRecangles = nil
			}

			activeRectangles = append(activeRectangles, &p.NamedRect)
			continue
		}

		//fmt.Printf("checking end point %d\n", len(activeRectangles))
		// end point -> if temp is not empty, means it is intersecting with existing points in temp.
		// remove the start point of self. determine the intersections.
		toBeRemoved := 0
		for i, tp := range activeRectangles {
			if tp.Name == p.Name {
				//fmt.Printf("same %s %s\n", tp.Name, p.Name)
				toBeRemoved = i
				continue
			}

			if intersected, intersectedRect := p.Rect.Intersect(tp.Rect); intersected {
				activeIntersectedRecangles = append(activeIntersectedRecangles, &IntersectedRect{
					Names: []string{p.Name, tp.Name},
					Rect:  intersectedRect,
				})

				result = append(result, &IntersectedRect{
					Names: []string{p.Name, tp.Name},
					Rect:  intersectedRect,
				})
			}
		}

		// remove toBeRemoved entry from temp
		// log.Printf("finished rectangle %d, size : %d", toBeRemoved, len(tempPoints))
		if len(activeRectangles) > 0 {
			copy(activeRectangles[toBeRemoved:], activeRectangles[toBeRemoved+1:])
			activeRectangles[len(activeRectangles)-1] = nil
			activeRectangles = activeRectangles[:len(activeRectangles)-1]
		}

		// check if the current point has intersection with active rectangles as well
		newIntersections := findIntersections(p.NamedRect, activeIntersectedRecangles)
		// add them to current active intersected rectangles and result
		activeIntersectedRecangles = append(activeIntersectedRecangles, newIntersections...)
		result = append(result, newIntersections...)
	}

	return result
}

// findIntersections returns the intersections of a rectangle with intersected rectangles.
func findIntersections(current NamedRect, activeIntersectedRecangles []*IntersectedRect) []*IntersectedRect {
	var intersections []*IntersectedRect
	for _, r := range activeIntersectedRecangles {
		if r.Exists(current.Name) {
			// this combination is already calculated - ignore.
			continue
		}

		if intersected, intersectedRect := current.Rect.Intersect(r.Rect); intersected {
			intersections = append(intersections, &IntersectedRect{
				Names: append(r.Names, current.Name),
				Rect:  intersectedRect,
			})
		}

	}

	return intersections
}

/* // TODO: pass the input.
func find() {
	input := []Rect{
		Rect{x: 100, y: 100, w: 250, h: 80},
		Rect{x: 120, y: 200, w: 250, h: 150},
		Rect{x: 140, y: 160, w: 250, h: 100},
		Rect{x: 160, y: 140, w: 350, h: 190},
	}

	// initialize the piority queue.
	pq := &PriorityQueue{}
	heap.Init(pq)

	// create points
	for i, r := range input {
		heap.Push(pq, &Point{name: i + 1, flag: true, Rect: r})
		heap.Push(pq, &Point{name: i + 1, flag: false, Rect: r})
	}

	// best to use linked list ? // TODO: revisit.
	var tempPoints []*Point

	var result []NamedRect

	for pq.Len() > 0 {
		p := heap.Pop(pq).(*Point)
		if p.flag {
			// start point
			tempPoints = append(tempPoints, p)
			continue
		}

		// end point -> if temp is not empty, means it is intersecting with existing points in temp.
		// remove the start point of self. determine the intersections.

		toBeRemoved := 0
		for i, tp := range tempPoints {
			if tp.name == p.name {
				toBeRemoved = i
				continue
			}

			intersected, intersectedRect := p.intersect(tp.Rect)
			if !intersected {
				// impossible. panic ?
			}

			result = append(result, NamedRect{
				name: fmt.Sprintf("%d,%d", p.name, tp.name),
				Rect: intersectedRect,
			})

			/// 1,2,
			// TODO : what if current point and temp has common intersect point.
		}

		// remove toBeRemoved entry from temp
		log.Printf("finished rectangle %d", toBeRemoved)

		if toBeRemoved == 0 {
			tempPoints = tempPoints[1:]
		} else {
			t := tempPoints[:toBeRemoved]
			t = append(t, tempPoints[toBeRemoved+1:]...)
			tempPoints = t
		}
	}

} */
