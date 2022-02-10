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
