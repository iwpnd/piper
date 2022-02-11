package piper

func max(a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}

	return b
}

func between(p, min, max float64) bool {
	return p > min && p < max
}

func inExtent(p []float64, ring [][]float64) bool {
	w := ring[0][0]
	s := ring[0][1]
	e := ring[0][0]
	n := ring[0][1]

	for _, p := range ring {
		if w > p[0] {
			w = p[0]
		}

		if s > p[1] {
			s = p[1]
		}

		if e < p[0] {
			e = p[0]
		}

		if n < p[1] {
			n = p[1]
		}
	}
	lon, lat := p[0], p[1]

	return (((w <= lon) && (lon <= w)) ||
		((e <= lon) && (lon <= w)) ||
		((s <= lat) && (lat <= n)) ||
		((n <= lat) && (lat <= s)))
}

func inRing(p []float64, ring [][]float64) bool {
	first, last := ring[0], ring[len(ring)-1]
	if first[0] == last[0] && first[1] == last[1] {
		ring = ring[0 : len(ring)-1]
	}

	lon := p[0]
	lat := p[1]
	counter := 0

	for i, j := 0, len(ring)-1; i < len(ring); i, j = i+1, i {
		iLon := ring[i][0]
		iLat := ring[i][1]
		jLon := ring[j][0]
		jLat := ring[j][1]

		// if p's latitude is not between the edges latitudes it cannot intersect
		min := min(iLat, jLat)
		max := max(iLat, jLat)
		if !between(p[1], min, max) {
			continue
		}

		// if p's longitude is smaller than the longitude of the
		// rays intersection with the current edge then it intersects
		intersects := lon < ((jLon-iLon)*(lat-iLat))/(jLat-iLat)+iLon
		if intersects {
			counter++
		}
	}

	return counter%2 != 0
}

func hasHoles(polygon [][][]float64) bool {
	return len(polygon) > 1
}

// Pip checks if Point p is inside input polygon. Does account for holes.
func Pip(p []float64, polygon [][][]float64) bool {
	outer := polygon[0]
	inPolygon := false

	// speeds up operations on complex polygons, insignifically slows
	// down on simple polygons
	if !inExtent(p, outer) {
		return false
	}

	if inRing(p, outer) {
		inPolygon = true

		// if there inner ring/holes we have to assume
		// that p can be in a hole, and therefor not in polygon
		if hasHoles(polygon) {
			holes := polygon[1:]
			inPolygon = false
			inHole := false

			for i := 0; i < len(holes); i++ {
				if inRing(p, holes[i]) {
					inHole = true
				}
			}

			if !inHole {
				inPolygon = true
			}
		}

		return inPolygon
	}

	return inPolygon
}
