package coo

import (
	"math"
)

const (
	toRad = math.Pi / 180
	toDeg = 180 / math.Pi
)

func GeographicToGeocentric(lon, lat, h, a, fi float64) (x, y, z float64) {
	lon *= toRad
	lat *= toRad
	N := a / math.Sqrt(1-e2(fi)*math.Pow(math.Sin(lat), 2))
	x = (N + h) * math.Cos(lat) * math.Cos(lon)
	y = (N + h) * math.Cos(lat) * math.Sin(lon)
	z = (N*b2(a, fi)/(a*a) + h) * math.Sin(lat)
	return
}

func GeocentricToGeographic(x, y, z, a, fi float64) (lon, lat, h float64) {
	s := math.Sqrt(x*x + y*y)
	T := math.Atan(z * a / (s * b(a, fi)))
	B := math.Atan((z + e2(fi)*a*a/b(a, fi)*math.Pow(math.Sin(T), 3)) / (s - e2(fi)*a*math.Pow(math.Cos(T), 3)))
	var L float64
	if x >= 0 {
		L = math.Atan(y / x)
	} else if y > 0 {
		L = math.Atan(y/x) + math.Pi
	} else {
		L = math.Atan(y/x) - math.Pi
	}
	N := a / math.Sqrt(1-e2(fi)*math.Pow(math.Sin(B), 2))
	h = s/math.Cos(B) - N
	lon = L * toDeg
	lat = B * toDeg
	return
}

func helmert(x, y, z, Tx, Ty, Tz, Ds, Rx, Ry, Rz float64) (xN, yN, zN float64) {
	Rx, Ry, Rz = Rx*math.Pi/180.0/3600.0, Ry*math.Pi/180.0/3600.0, Rz*math.Pi/180.0/3600.0
	Ds = 1 + Ds/math.Pow(10, 6)
	xN = 1/1000000*Ds*(Ry*z-Rz*y+x) + Ry*z - Rz*y + Tx + x
	yN = -1/1000000*Ds*(Rx*z-Rz*x-y) - Rx*z + Rz*x + Ty + y
	zN = 1/1000000*Ds*(Rx*y-Ry*x+z) + Rx*y - Ry*x + Tz + z
	return
}

func b(a, fi float64) float64 {
	return a * (1 - 1/fi)
}

func b2(a, fi float64) float64 {
	return b(a, fi) * b(a, fi)
}

func f3rd(fi float64) float64 {
	return 0.5 / (fi - 0.5)
}

func e2(fi float64) float64 {
	return 2 * (fi - 0.5) / (fi * fi)
}
