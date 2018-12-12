package coo

import (
	"math"
)

var WebMercator = &System{
	toGeographic: func(east, north float64, ell *Ellipsoid) (lon, lat float64) {
		if ell == nil {
			ell = DefaultEllipsoid
		}
		a := ell.A
		lon = east / a * toDeg
		lat = math.Atan(math.Exp(north/a))*toDeg*2 - 90
		return
	},
	fromGeographic: func(lon, lat float64, ell *Ellipsoid) (east, north float64) {
		if ell == nil {
			ell = DefaultEllipsoid
		}
		a := ell.A
		east = lon * a * toRad
		north = math.Log(math.Tan((90+lat)*toRad/2)) * a
		return
	},
}

func TransverseMercator(centralMeridian, scale, falseEasting, falseNorthing float64) *System {
	return &System{
		toGeographic: func(east, north float64, ell *Ellipsoid) (lon, lat float64) {
			if ell == nil {
				ell = DefaultEllipsoid
			}
			a := ell.A
			fi := ell.Fi
			north = (north - falseNorthing) / scale
			n := f3rd(fi)
			n2 := n * n
			n3 := n2 * n
			n4 := n3 * n
			n5 := n4 * n
			alpha := (a + b(a, fi)) / 2 * (1 + n2/4 + n4/64)
			latta := 3*n/2 - 27*n3/32 + 269*n5/512
			gamma := 21*n2/16 - 55*n4/32
			dellta := 151*n3/96 - 417*n5/128
			epsilon := 1097 * n4 / 512
			Y := (east - falseEasting) / scale
			Y2 := Y * Y
			Y3 := Y2 * Y
			Y4 := Y3 * Y
			Y5 := Y4 * Y
			Y6 := Y5 * Y
			Y7 := Y6 * Y
			BO := north / alpha
			BF := BO + latta*math.Sin(2*BO) + gamma*math.Sin(4*BO) + dellta*math.Sin(6*BO) + epsilon*math.Sin(8*BO)
			NF := a / math.Sqrt(1-(e2(fi)*math.Pow(math.Sin(BF), 2)))
			NF2 := NF * NF
			NF3 := NF2 * NF
			NF4 := NF3 * NF
			NF5 := NF4 * NF
			NF6 := NF5 * NF
			NF7 := NF6 * NF
			NPhi := math.Sqrt(a * a / b2(a, fi) * e2(fi) * math.Pow(math.Cos(BF), 2))
			NPhi2 := NPhi * NPhi
			NPhi4 := NPhi2 * NPhi2
			tF := math.Tan(BF)
			tF2 := tF * tF
			tF4 := tF2 * tF2
			tF6 := tF4 * tF2
			B1 := (tF / 2.0) / NF2 * (-1 - NPhi2) * Y2
			B2 := (tF / 24.0) / NF4 * (5 + 3*tF2 + 6*NPhi2 - 6*tF2*NPhi2 - 4*NPhi4 - 9*tF2*NPhi4) * Y4
			B3 := (tF / 720.0) / NF6 * (-61 - 90*tF2 - 45*tF4 - 107*NPhi2 + 162*tF2*NPhi2 + 45*tF4*NPhi2) * Y6
			lat = (BF + B1 + B2 + B3) * toDeg
			L1 := ((1 / NF) / math.Cos(BF)) * Y
			L2 := (1 / 6.0 * 1 / NF3) / math.Cos(BF) * (-1 - 2*tF2 - NPhi2) * Y3
			L3 := (1 / 120.0 * 1 / NF5) / math.Cos(BF) * (5 + 28*tF2 + 24*tF4 + 6*NPhi2 + 8*tF2*NPhi2) * Y5
			L4 := (1 / 15040.0 * 1 / NF7) / math.Cos(BF) * (-61 - 622*tF2 - 1320*tF4 - 720*tF6) * Y7
			lon = centralMeridian + (L1+L2+L3+L4)*toDeg
			return
		},
		fromGeographic: func(lon, lat float64, ell *Ellipsoid) (east, north float64) {
			if ell == nil {
				ell = DefaultEllipsoid
			}
			a := ell.A
			fi := ell.Fi
			n := f3rd(fi)
			n2 := n * n
			n3 := n2 * n
			n4 := n3 * n
			n5 := n4 * n
			alpha := (a + b(a, fi)) / 2 * (1 + n2/4 + n4/64)
			beta := -3*n/2 + 9*n3/16 - 3*n5/32
			gamma := 15*n2/16 - 15*n4/32
			dellta := -35*n3/48 + 105*n5/256
			epsilon := 315 * n4 / 512
			lat *= toRad
			l := (lon - centralMeridian) * toRad
			l2 := l * l
			l3 := l2 * l
			l4 := l3 * l
			Ne := a / math.Sqrt(1-e2(fi)*math.Pow(math.Sin(lat), 2))
			eta := math.Sqrt(a * a / b2(a, fi) * e2(fi) * math.Pow(math.Cos(lat), 2))
			t := math.Tan(lat)
			R := alpha * (lat + beta*math.Sin(2*lat) + gamma*math.Sin(4*lat) + dellta*math.Sin(6*lat) + epsilon*math.Sin(8*lat))
			h1 := t / 2 * Ne * math.Pow(math.Cos(lat), 2) * l2
			h2 := t / 24 * Ne * math.Pow(math.Cos(lat), 4) * (5 - t*t + 9*eta*eta + 4*eta*eta*eta*eta) * l4
			north = (R+h1+h2)*scale + falseNorthing
			r1 := Ne * math.Cos(lat) * l
			r2 := Ne / 6 * math.Pow(math.Cos(lat), 3) * (1 - t*t + eta*eta) * l3
			east = (r1+r2)*scale + falseEasting
			return
		},
	}
}

func UTM(zone float64, hemisphere string) *System {
	for {
		if zone > 0 && zone < 61 {
			break
		}
		if zone < 1 {
			zone += 60
		}
		if zone > 60 {
			zone = math.Mod(zone, 60)
		}
	}
	if hemisphere == "S" {
		return TransverseMercator(zone*6-183, 0.9996, 500000, 10000000)
	}
	return TransverseMercator(zone*6-183, 0.9996, 500000, 0)
}

func GaussKrueger(zone float64) *System {
	for {
		if zone > 0 && zone < 121 {
			break
		}
		if zone < 1 {
			zone += 120
		}
		if zone > 120 {
			zone = math.Mod(zone, 120)
		}
	}
	return TransverseMercator(zone*3, 1, zone*1000000+500000, 0)
}

func ConformalConic(lat1, lat2, falseLat, falseLon, falseEasting, falseNorthing float64) *System {
	return &System{
		toGeographic: func(east, north float64, ell *Ellipsoid) (lon, lat float64) {
			if ell == nil {
				ell = DefaultEllipsoid
			}
			a := ell.A
			fi := ell.Fi
			f := 1 / fi
			ef := math.Sqrt(2*f - f*f)
			m := func(l float64) float64 {
				return math.Cos(l) / math.Sqrt(1-ef*ef*math.Pow(math.Sin(l), 2))
			}
			m1 := m(lat1 * toRad)
			m2 := m(lat2 * toRad)
			t := func(l float64) float64 {
				return math.Tan(math.Pi*0.25-l*0.5) / math.Pow((1-ef*math.Sin(l))/(1+ef*math.Sin(l)), ef/2)
			}
			tO := t(falseLat * toRad)
			t1 := t(lat1 * toRad)
			t2 := t(lat2 * toRad)
			n := (math.Log(m1) - math.Log(m2)) / (math.Log(t1) - math.Log(t2))
			if lat1 == lat2 {
				n = math.Sin(lat1 * toRad)
			}
			F := m1 / (n * math.Pow(t1, n))
			pO := a * F * math.Pow(tO, n)
			Ni := north - falseNorthing
			Ei := east - falseEasting
			pi := math.Sqrt(math.Abs(Ei*Ei + math.Pow(pO-Ni, 2)))
			if n < 0 {
				pi = -pi
			}
			ti := math.Pow(pi/(a*F), 1/n)
			yi := math.Atan(Ei / (pO - Ni))
			lat = math.Pi*0.5 - 2*math.Atan(ti)
			for i := 0; i < 3; i++ {
				lat = math.Pi*0.5 - 2*math.Atan(ti*math.Pow((1-ef*math.Sin(lat))/(1+ef*math.Sin(lat)), ef/2))
			}
			lat *= toDeg
			lon = (yi/n + falseLon*toRad) * toDeg
			return
		},
		fromGeographic: func(lon, lat float64, ell *Ellipsoid) (east, north float64) {
			lon *= toRad
			lat *= toRad
			if ell == nil {
				ell = DefaultEllipsoid
			}
			a := ell.A
			fi := ell.Fi
			f := 1 / fi
			ef := math.Sqrt(2*f - f*f)
			m := func(l float64) float64 {
				return math.Cos(l) / math.Sqrt(1-ef*ef*math.Pow(math.Sin(l), 2))
			}
			m1 := m(lat1 * toRad)
			m2 := m(lat2 * toRad)
			t := func(l float64) float64 {
				return math.Tan(math.Pi*0.25-l*0.5) / math.Pow((1-ef*math.Sin(l))/(1+ef*math.Sin(l)), ef/2)
			}
			t1 := t(lat1 * toRad)
			t2 := t(lat2 * toRad)
			n := (math.Log(m1) - math.Log(m2)) / (math.Log(t1) - math.Log(t2))
			if lat1 == lat2 {
				n = math.Sin(lat1 * toRad)
			}
			F := m1 / (n * math.Pow(t1, n))
			p := func(l float64) float64 {
				return a * F * math.Pow(t(l), n)
			}
			pO := p(falseLat * toRad)
			y := n * (lon - falseLon*toRad)
			north = falseNorthing + pO - p(lat)*math.Cos(y)
			east = falseEasting + p(lat)*math.Sin(y)
			return
		},
	}
}
