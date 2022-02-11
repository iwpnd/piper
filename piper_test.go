package piper

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestInExtent(t *testing.T) {
	test := []struct {
		ring     [][]float64
		p        []float64
		expected bool
	}{
		// inside
		{
			ring:     [][]float64{{0, 0}, {0, 1}, {1, 1}, {1, 0}, {0, 0}},
			p:        []float64{0.5, 0.5},
			expected: true,
		},
		// outside
		{
			ring:     [][]float64{{0, 0}, {0, 1}, {1, 1}, {1, 0}, {0, 0}},
			p:        []float64{2, 2},
			expected: false,
		},
		// touches
		{
			ring:     [][]float64{{0, 0}, {0, 1}, {1, 1}, {1, 0}, {0, 0}},
			p:        []float64{0, 0},
			expected: true,
		},
	}

	for _, test := range test {
		got := inExtent(test.p, test.ring)

		if got != test.expected {
			t.Errorf("expected %+v, got: %+v", test.expected, got)
		}
	}
}

func TestInRing(t *testing.T) {
	test := []struct {
		ring     [][]float64
		p        []float64
		expected bool
	}{
		// inside
		{
			ring:     [][]float64{{0, 0}, {0, 1}, {1, 1}, {1, 0}, {0, 0}},
			p:        []float64{0.5, 0.5},
			expected: true,
		},
		// outside
		{
			ring:     [][]float64{{0, 0}, {0, 1}, {1, 1}, {1, 0}, {0, 0}},
			p:        []float64{2, 2},
			expected: false,
		},
		// complex inside
		{
			ring: [][]float64{
				{13.324699401855469, 52.495741489296144},
				{13.3209228515625, 52.491769893960836},
				{13.321266174316406, 52.48717075649203},
				{13.325729370117188, 52.485080080472514},
				{13.332252502441404, 52.484243782241144},
				{13.34014892578125, 52.485080080472514},
				{13.343925476074217, 52.48717075649203},
				{13.344268798828125, 52.49093372292005},
				{13.342552185058594, 52.494905393770985},
				{13.337745666503906, 52.49741363265356},
				{13.336715698242188, 52.49344218834781},
				{13.33740234375, 52.490306584206124},
				{13.335342407226562, 52.48863417058138},
				{13.331222534179686, 52.48884322576231},
				{13.328132629394531, 52.49114276717071},
				{13.329505920410156, 52.49511441914269},
				{13.330535888671875, 52.4976226461021},
				{13.324699401855469, 52.495741489296144},
			},
			p:        []float64{13.32517147064209, 52.491743763856476},
			expected: true,
		},
		// complex outside
		{
			ring: [][]float64{
				{13.324699401855469, 52.495741489296144},
				{13.3209228515625, 52.491769893960836},
				{13.321266174316406, 52.48717075649203},
				{13.325729370117188, 52.485080080472514},
				{13.332252502441404, 52.484243782241144},
				{13.34014892578125, 52.485080080472514},
				{13.343925476074217, 52.48717075649203},
				{13.344268798828125, 52.49093372292005},
				{13.342552185058594, 52.494905393770985},
				{13.337745666503906, 52.49741363265356},
				{13.336715698242188, 52.49344218834781},
				{13.33740234375, 52.490306584206124},
				{13.335342407226562, 52.48863417058138},
				{13.331222534179686, 52.48884322576231},
				{13.328132629394531, 52.49114276717071},
				{13.329505920410156, 52.49511441914269},
				{13.330535888671875, 52.4976226461021},
				{13.324699401855469, 52.495741489296144},
			},
			p:        []float64{13.332552909851074, 52.491848284180826},
			expected: false,
		},
	}

	for _, test := range test {
		got := inRing(test.p, test.ring)

		if got != test.expected {
			t.Errorf("expected %+v, got: %+v", test.expected, got)
		}
	}
}

func TestInPolygonWithoutHoles(t *testing.T) {
	polygon := [][][]float64{
		{
			{-3.0988311767578125, 40.837710162420045},
			{-3.121490478515625, 40.820045086716505},
			{-3.0978012084960938, 40.80237530523985},
			{-3.0754852294921875, 40.8210843390845},
			{-3.0988311767578125, 40.837710162420045},
		},
	}

	test := []struct {
		tcase    string
		p        []float64
		expected bool
	}{
		{
			tcase:    "outside polygon",
			p:        []float64{-3.0816650390625, 40.80809251416925},
			expected: false,
		},
		{
			tcase:    "inside polygon",
			p:        []float64{-3.0878448486328125, 40.81497849824719},
			expected: true,
		},
	}

	for _, test := range test {
		got := Pip(test.p, polygon)

		if got != test.expected {
			t.Errorf("expected %+v, got: %+v", test.expected, got)
		}
	}
}

func TestInPolygonWithHoles(t *testing.T) {

	polygon := [][][]float64{
		{
			{-3.0988311767578125, 40.837710162420045},
			{-3.121490478515625, 40.820045086716505},
			{-3.0978012084960938, 40.80237530523985},
			{-3.0754852294921875, 40.8210843390845},
			{-3.0988311767578125, 40.837710162420045},
		},
		{
			{-3.0988311767578125, 40.82783908257347},
			{-3.1098175048828125, 40.820045086716505},
			{-3.0988311767578125, 40.81147063339219},
			{-3.086471557617187, 40.820304901335035},
			{-3.0988311767578125, 40.82783908257347},
		},
	}

	test := []struct {
		tcase    string
		p        []float64
		expected bool
	}{
		{
			tcase:    "south-east outside polygon, in bbox",
			p:        []float64{-3.0816650390625, 40.80809251416925},
			expected: false,
		},
		{
			tcase:    "south-east inside polygon, inside bbox",
			p:        []float64{-3.0878448486328125, 40.81497849824719},
			expected: true,
		},
		{
			tcase:    "south-east outside polygon, outside bbox",
			p:        []float64{-3.0713653564453125, 40.800945926051526},
			expected: false,
		},
		{
			tcase:    "south outside polygon, outside bbox",
			p:        []float64{-3.0978012084960938, 40.79769722250925},
			expected: false,
		},
		{
			tcase:    "south inside polygon, inside bbox",
			p:        []float64{-3.098316192626953, 40.8067931917519},
			expected: true,
		},
		{
			tcase:    "south-west outside polygon, inside bbox",
			p:        []float64{-3.116168975830078, 40.807702720115294},
			expected: false,
		},
		{
			tcase:    "south-west outside polygon, outside bbox",
			p:        []float64{-3.1250953674316406, 40.80068603561921},
			expected: false,
		},
		{
			tcase:    "south-west inside polygon, inside bbox",
			p:        []float64{-3.10810089111328, 40.814198988751876},
			expected: true,
		},
		{
			tcase:    "west outside polygon, outside bbox",
			p:        []float64{-3.1266403198242188, 40.8197852710803},
			expected: false,
		},
		{
			tcase:    "west inside polygon, inside bbox",
			p:        []float64{-3.1141090393066406, 40.82017499415298},
			expected: true,
		},
		{
			tcase:    "north-west inside polygon, inside bbox",
			p:        []float64{-3.1070709228515625, 40.82667004158603},
			expected: true,
		},
		{
			tcase:    "north-west outside polygon, inside bbox",
			p:        []float64{-3.1141090393066406, 40.83199550584334},
			expected: false,
		},
		{
			tcase:    "north-west outside polygon, inside bbox",
			p:        []float64{-3.1266403198242188, 40.84082704020004},
			expected: false,
		},
		{
			tcase:    "north inside polygon, inside bbox",
			p:        []float64{-3.0988311767578125, 40.83264492344398},
			expected: true,
		},
		{
			tcase:    "north outside polygon, outside bbox",
			p:        []float64{-3.0988311767578125, 40.8425152878029},
			expected: false,
		},
		{
			tcase:    "north-east inside polygon, inside bbox",
			p:        []float64{-3.0895614624023438, 40.826799936046804},
			expected: true,
		},
		{
			tcase:    "north-east outside polygon, inside bbox",
			p:        []float64{-3.0816650390625, 40.83160585222969},
			expected: false,
		},
		{
			tcase:    "north-east outside polygon, outside bbox",
			p:        []float64{-3.07016372680664, 40.84147637129013},
			expected: false,
		},
		{
			tcase:    "east inside polygon, inside bbox",
			p:        []float64{-3.080635070800781, 40.82056471493589},
			expected: true,
		},
		{
			tcase:    "east outside polygon, outside bbox",
			p:        []float64{-3.069477081298828, 40.8210843390845},
			expected: false,
		},
		{
			tcase:    "east outside polygon but in hole, inside bbox",
			p:        []float64{-3.098487854003906, 40.81874599835864},
			expected: false,
		},
	}

	for _, test := range test {
		got := Pip(test.p, polygon)

		if got != test.expected {
			t.Errorf("case: %s - expected %+v, got: %+v", test.tcase, test.expected, got)
		}
	}
}

var simplePolygon = [][][]float64{
	{
		{-3.0988311767578125, 40.837710162420045},
		{-3.121490478515625, 40.820045086716505},
		{-3.0978012084960938, 40.80237530523985},
		{-3.0754852294921875, 40.8210843390845},
		{-3.0988311767578125, 40.837710162420045},
	},
}

var simplePolygonWithHoles = [][][]float64{
	{
		{-3.0988311767578125, 40.837710162420045},
		{-3.121490478515625, 40.820045086716505},
		{-3.0978012084960938, 40.80237530523985},
		{-3.0754852294921875, 40.8210843390845},
		{-3.0988311767578125, 40.837710162420045},
	},
	{
		{-3.0988311767578125, 40.82783908257347},
		{-3.1098175048828125, 40.820045086716505},
		{-3.0988311767578125, 40.81147063339219},
		{-3.086471557617187, 40.820304901335035},
		{-3.0988311767578125, 40.82783908257347},
	},
}

func BenchmarkPipSimpleInside(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Pip([]float64{-3.0878448486328125, 40.81497849824719}, simplePolygon)
	}
}

func BenchmarkPipSimpleOutside(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Pip([]float64{-3.069477081298828, 40.8210843390845}, simplePolygon)
	}
}

func BenchmarkPipSimpleInsideWithHoles(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Pip([]float64{-3.080635070800781, 40.82056471493589}, simplePolygonWithHoles)
	}
}

func BenchmarkPipSimpleOutsideWithHoles(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Pip([]float64{-3.0816650390625, 40.80809251416925}, simplePolygonWithHoles)
	}
}

type Geometry struct {
	Type        string
	Coordinates [][][]float64
}

type Feature struct {
	Type     string
	Geometry Geometry
}

func BenchmarkPipComplexInside(b *testing.B) {
	raw, err := ioutil.ReadFile("./testdata/berlin.geojson")
	if err != nil {
		b.Fatal("could not load test data", err)
	}

	var complexPolygon Feature
	err = json.Unmarshal(raw, &complexPolygon)

	if err != nil {
		b.Fatal("could not load test feature", err)
	}

	for n := 0; n < b.N; n++ {
		Pip([]float64{-3.080635070800781, 40.82056471493589}, complexPolygon.Geometry.Coordinates)
	}
}
