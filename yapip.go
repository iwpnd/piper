package yapip

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
	if !inExtent(p, ring) {
		return false
	}

	first, last := ring[0], ring[len(ring)-1]
	if first[0] == last[0] && first[1] == last[1] {
		ring = ring[0 : len(ring)-1]
	}

	lon := p[0]
	lat := p[1]
	counter := 0

	for i, j := 0, len(ring)-1; i < len(ring); i, j = i+1, i {
		startLon := ring[i][0]
		startLat := ring[i][1]
		endLon := ring[j][0]
		endLat := ring[j][1]

		intersects := ((startLat > lat) != (endLat > lat)) &&
			(lon < ((endLon-startLon)*(lat-startLat))/(endLat-startLat)+startLon)

		if intersects {
			counter++
		}
	}

	return counter%2 != 0
}
