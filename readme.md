# piper

Yet another point in polygon package

## Installation

```
go get -u github.com/iwpnd/piper
```

## Usage

```
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
