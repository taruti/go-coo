package epsg

import (
	"errors"

	"github.com/wroge/go-coo/de"

	"github.com/wroge/go-coo"
)

type FromToSource interface {
	coo.FromSource
	coo.ToSource
}

var Code4326 = &coo.Geographic{
	Ellipsoid: coo.WGS84,
}

var Code4258 = &coo.Geographic{
	Ellipsoid: coo.GRS80,
}

var Code3857 = &coo.Projected{
	Geographic: Code4326,
	System:     coo.WebMercator,
}

var Code900913 = Code3857

var Code25832 = &coo.Projected{
	Geographic: Code4258,
	System:     coo.UTM(32, "N"),
}

var Code25833 = &coo.Projected{
	Geographic: Code4258,
	System:     coo.UTM(33, "N"),
}

var Code5668 = &coo.Projected{
	Geographic: &coo.Geographic{
		Geocentric: de.RD83,
		Ellipsoid:  coo.Bessel,
	},
	System: coo.GaussKrueger(4),
}

var Code5669 = &coo.Projected{
	Geographic: &coo.Geographic{
		Geocentric: de.RD83,
		Ellipsoid:  coo.Bessel,
	},
	System: coo.GaussKrueger(5),
}

var Code31466 = &coo.Projected{
	Geographic: &coo.Geographic{
		Geocentric: de.DHDN2001,
		Ellipsoid:  coo.Bessel,
	},
	System: coo.GaussKrueger(2),
}

var Code31467 = &coo.Projected{
	Geographic: &coo.Geographic{
		Geocentric: de.DHDN2001,
		Ellipsoid:  coo.Bessel,
	},
	System: coo.GaussKrueger(3),
}

var Code31468 = &coo.Projected{
	Geographic: &coo.Geographic{
		Geocentric: de.DHDN2001,
		Ellipsoid:  coo.Bessel,
	},
	System: coo.GaussKrueger(4),
}

var Code31469 = &coo.Projected{
	Geographic: &coo.Geographic{
		Geocentric: de.DHDN2001,
		Ellipsoid:  coo.Bessel,
	},
	System: coo.GaussKrueger(5),
}

var Code32632 = &coo.Projected{
	Geographic: Code4326,
	System:     coo.UTM(32, "N"),
}

var Code32633 = &coo.Projected{
	Geographic: Code4326,
	System:     coo.UTM(33, "N"),
}

var Code4647 = &coo.Projected{
	Geographic: Code4258,
	System:     coo.TransverseMercator(9, 0.9996, 32500000, 0),
}

var Code5650 = &coo.Projected{
	Geographic: Code4258,
	System:     coo.TransverseMercator(15, 0.9996, 33500000, 0),
}

var Code3067 = &coo.Projected{
	Geographic: Code4258,
	System:     coo.TransverseMercator(27, 0.9996, 500000, 0),
}

var codes = map[int]FromToSource{
	4326:   Code4326,
	3857:   Code3857,
	900913: Code900913,
	25832:  Code25832,
	25833:  Code25833,
	5668:   Code5668,
	5669:   Code5669,
	31466:  Code31466,
	31467:  Code31467,
	31468:  Code31468,
	31469:  Code31469,
	32632:  Code32632,
	32633:  Code32633,
	4647:   Code4647,
	5650:   Code5650,
	3067:   Code3067,
}

func Code(code int) (FromToSource, error) {
	for _, c := range List() {
		if c == code {
			return codes[code], nil
		}
	}
	return nil, errors.New("Code-Error: EPSG code is not provided")
}

func List() []int {
	var result = make([]int, len(codes))
	i := 0
	for b := range codes {
		result[i] = b
		i++
	}
	return result
}
