# piper

Yet another point in polygon package. Piper makes use of ray casting and does account for holes in polygons.

## Installation

```
go get -u github.com/iwpnd/piper
```

## Usage

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

## License

MIT

## Maintainer

Benjamin Ramser - [@iwpnd](https://github.com/iwpnd)

Project Link: [https://github.com/iwpnd/piper](https://github.com/iwpnd/piper)

## Acknowledgement

Phillip Lemons - [Ray Casting Algorithm](http://philliplemons.com/posts/ray-casting-algorithm)

Great introduction into the topic with good visualisations.
