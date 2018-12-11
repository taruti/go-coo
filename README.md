# go-coo

Coordinate conversion and transformation package in pure Go(lang).

## Usage

```go
go get github.com/wroge/go-coo
```

## Examples

```go
east, north, _ := coo.UTM(32, "N").FromGeographic(9, 52, coo.GRS80)

east, north, _ := coo.ConformalConic(43, 62, 30, 10, 0, 0).FromGeographic(9, 52, coo.Hayford)

proj := &coo.Projected{}
geogr := &coo.Geographic{}
east, north, _ := coo.Transform(9, 52, 100, geogr, proj)

webmercator := &coo.Projected{
    System: coo.WebMercator,
}
dhdnGK3 := &coo.Projected{
    Geographic: &coo.Geographic{
        Geocentric: &coo.Geocentric{
            Geocentric: nil,
            Tx:         598.1,
            Ty:         73.7,
            Tz:         418.2,
            Rx:         0.202,
            Ry:         0.045,
            Rz:         -2.455,
            Ds:         6.7,
        },
        Ellipsoid: coo.Bessel,
    },
    System: coo.GaussKrueger(3)
}
east, north, _ := coo.Transform(1.0018754171394621e+06, 6.800125454397305e+06, 100, webmercator, dhdnGK3)
```

## EPSG

```go
wgs84geogr, err := epsg.Code(4326)
```