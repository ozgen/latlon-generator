package geo

func CalculatePolygonCentroid(ring LinearRing) Point {
	var cx, cy, area float64
	n := len(ring)

	if n < 3 {
		return ring[0] // fallback to first point
	}

	for i := 0; i < n-1; i++ {
		x0, y0 := ring[i][0], ring[i][1]
		x1, y1 := ring[i+1][0], ring[i+1][1]

		a := x0*y1 - x1*y0
		area += a
		cx += (x0 + x1) * a
		cy += (y0 + y1) * a
	}

	area *= 0.5
	if area == 0 {
		return ring[0]
	}
	cx /= 6 * area
	cy /= 6 * area

	return Point{cx, cy}
}
