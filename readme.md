# piper

Yet another point in polygon package. Piper makes use of ray casting and does account for holes in polygons.

## Installation

```
go get -u github.com/iwpnd/piper
```

## Usage

Vanilla

```go
package main

import (
  "fmt"

  "github.com/iwpnd/piper"
  )

func main() {
  p := []float64{0.5,0.5}
  polygon := [][][]float64{{{0, 0}, {0, 1}, {1, 1}, {1, 0}, {0, 0}}}

  pip := piper.Pip(p, polygon)
  fmt.Printf("Point in Polygon: %+v\n", pip)
}
```

Or using [github/paulmach/go.geosjon](https://github.com/paulmach/go.geojson)

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/iwpnd/piper"
	"github.com/paulmach/go.geojson"
)

func main() {
	p := []float64{0.5, 0.5}

	raw, err := ioutil.ReadFile("my_feature.geojson")
	if err != nil {
		panic("something went wrong")
	}

	var feature geojson.Feature
	err = json.Unmarshal(raw, &feature)
	if err != nil {
		panic("something went wrong")
	}

	pip := piper.Pip(p, feature.Geometry.Polygon)
	fmt.Printf("Point in Polygon: %+v\n", pip)
}
```

## Benchmark

```
goos: darwin
goarch: amd64
pkg: github.com/iwpnd/piper
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkPipSimpleInside-12              	43108360	        27.73 ns/op
BenchmarkPipSimpleOutside-12             	44025870	        27.53 ns/op
BenchmarkPipSimpleInsideWithHoles-12     	27355524	        42.84 ns/op
BenchmarkPipSimpleOutsideWithHoles-12    	42239286	        28.22 ns/op
BenchmarkPipComplexInside-12             	  474601	        2323 ns/op
```

## License

MIT

## Maintainer

Benjamin Ramser - [@iwpnd](https://github.com/iwpnd)

Project Link: [https://github.com/iwpnd/piper](https://github.com/iwpnd/piper)

## Acknowledgement

Phillip Lemons - [Ray Casting Algorithm](http://philliplemons.com/posts/ray-casting-algorithm)

Great introduction into the topic with good visualisations.
