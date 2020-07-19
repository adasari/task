# task

Flow
1. Prepare the points with left edge and right edge of the rectagle
2. Sort the points (used priority queue)
3. iterate over the points from step#2 and prepare the active rectangles 
4. if right edge of the rectangles is found, determine the active intersected rectangles by checking if it is intersecting with active rectangles (other than self) and other intersected rectangles
5. Add every intersected rectangle to result

