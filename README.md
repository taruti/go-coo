# go-coo

Coordinate conversion and transformation package in pure Go(lang).

## Usage

```go
go get github.com/wroge/go-coo

import (
	"github.com/wroge/go-coo"
)
```

## Features

- Coordinate conversion between projected coordinate systems and geographical coordinates.

```go
texasLambert := coo.ConformalConic(27.5, 35, 18, -100, 1500000, 5000000)
lon, lat := texasLambert.ToGeographic(1492401.186, 6484663.429, coo.GRS80)
east, north := texasLambert.FromGeographic(-100.080, 31.316, coo.GRS80)
```

- Coordinate conversion between different projected coordinate systems.

```go
east, north := coo.Convert(1001875.417, 6800125.454, coo.WebMercator, 
    coo.UTM(32, "N"), coo.WGS84)
east, north = coo.Convert(500000.000, 5761038.212, coo.UTM(32, "N"), 
    coo.WebMercator, coo.WGS84)
```

- Coordinate transformation between different coordinate systems.

```go
etrs89utm32N := &coo.Projected{
    Geographic: &coo.Geographic{
        Ellipsoid: coo.GRS80,
    },
    System: coo.UTM(32, "N"),
}
dhdn2001gk3 := &coo.Projected{
    Geographic: &coo.Geographic{
        Geocentric: de.DHDN2001,
        Ellipsoid:  coo.Bessel,
    },
    System: coo.GaussKrueger(3),
}
east, north, _ := coo.Transform(500000.000, 5761038.212, 0, etrs89utm32N, dhdn2001gk3)
```

- Standard coordinate systems

```go
// &coo.Geocentric{} -> WGS84, &coo.Geographic{} -> WGS84 Geographic, &coo.Projected{} -> Web Mercator WGS84
x, y, z := coo.Transform(1001875.417, 6800125.454, 0, &coo.Projected{}, &coo.Geocentric{})
lon, lat, h := coo.Transform(x, y, z, &coo.Geocentric{}, &coo.Geographic{})
fmt.Println(lon, lat, h)
```

## EPSG

EPSG is a sub-package to which EPSG codes are added on a regular basis.

```go
import (
	"github.com/wroge/go-coo/epsg"
)
```