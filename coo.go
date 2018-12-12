package coo

type Geocentric struct {
	*Geocentric
	Tx, Ty, Tz, Ds, Rx, Ry, Rz float64
}

type Geographic struct {
	*Geocentric
	*Ellipsoid
}

type Ellipsoid struct {
	A, Fi float64
}

var DefaultEllipsoid = WGS84

type Projected struct {
	*Geographic
	*System
}

type System struct {
	toGeographic   func(east, north float64, ell *Ellipsoid) (lon, lat float64)
	fromGeographic func(lon, lat float64, ell *Ellipsoid) (east, north float64)
}

var DefaultSystem = WebMercator

func (geoc *Geocentric) ToSource(x, y, z float64) (xS, yS, zS float64) {
	xS, yS, zS = helmert(x, y, z, geoc.Tx, geoc.Ty, geoc.Tz, geoc.Ds, geoc.Rx, geoc.Ry, geoc.Rz)
	if geoc.Geocentric != nil {
		xS, yS, zS = geoc.Geocentric.ToSource(xS, yS, zS)
	}
	return
}

func (geoc *Geocentric) FromSource(xS, yS, zS float64) (x, y, z float64) {
	if geoc.Geocentric != nil {
		xS, yS, zS = geoc.Geocentric.FromSource(xS, yS, zS)
	}
	x, y, z = helmert(xS, yS, zS, -geoc.Tx, -geoc.Ty, -geoc.Tz, -geoc.Ds, -geoc.Rx, -geoc.Ry, -geoc.Rz)
	return
}

func (geogr *Geographic) ToGeocentric(lon, lat, h float64) (x, y, z float64) {
	geogr.Validate()
	x, y, z = GeographicToGeocentric(lon, lat, h, geogr.A, geogr.Fi)
	return
}

func (geogr *Geographic) FromGeocentric(x, y, z float64) (lon, lat, h float64) {
	geogr.Validate()
	lon, lat, h = GeocentricToGeographic(x, y, z, geogr.A, geogr.Fi)
	return
}

func (geogr *Geographic) ToSource(lon, lat, h float64) (xS, yS, zS float64) {
	geogr.Validate()
	x, y, z := geogr.ToGeocentric(lon, lat, h)
	xS, yS, zS = geogr.Geocentric.ToSource(x, y, z)
	return
}

func (geogr *Geographic) FromSource(xS, yS, zS float64) (lon, lat, h float64) {
	geogr.Validate()
	x, y, z := geogr.Geocentric.FromSource(xS, yS, zS)
	lon, lat, h = geogr.FromGeocentric(x, y, z)
	return
}

func (proj *Projected) ToGeographic(east, north float64) (lon, lat float64) {
	proj.Validate()
	lon, lat = proj.System.ToGeographic(east, north, proj.Geographic.Ellipsoid)
	return
}

func (proj *Projected) FromGeographic(lon, lat float64) (east, north float64) {
	proj.Validate()
	east, north = proj.System.FromGeographic(lon, lat, proj.Geographic.Ellipsoid)
	return
}

func (proj *Projected) ToGeocentric(east, north, h float64) (x, y, z float64) {
	proj.Validate()
	lon, lat := proj.ToGeographic(east, north)
	x, y, z = proj.Geographic.ToGeocentric(lon, lat, h)
	return
}

func (proj *Projected) FromGeocentric(x, y, z float64) (east, north, h float64) {
	proj.Validate()
	lon, lat, h := proj.Geographic.FromGeocentric(x, y, z)
	east, north = proj.FromGeographic(lon, lat)
	return
}

func (proj *Projected) ToSource(east, north, h float64) (xS, yS, zS float64) {
	proj.Validate()
	x, y, z := proj.ToGeocentric(east, north, h)
	xS, yS, zS = proj.Geocentric.ToSource(x, y, z)
	return
}

func (proj *Projected) FromSource(xS, yS, zS float64) (east, north, h float64) {
	proj.Validate()
	x, y, z := proj.Geocentric.FromSource(xS, yS, zS)
	east, north, h = proj.FromGeocentric(x, y, z)
	return
}

func (sys *System) ToGeographic(east, north float64, ell *Ellipsoid) (lon, lat float64) {
	return sys.toGeographic(east, north, ell)
}

func (sys *System) FromGeographic(lon, lat float64, ell *Ellipsoid) (east, north float64) {
	return sys.fromGeographic(lon, lat, ell)
}

func (geogr *Geographic) Validate() {
	if geogr == nil {
		geogr = &Geographic{}
	}
	if geogr.Geocentric == nil {
		geogr.Geocentric = &Geocentric{}
	}
	if geogr.Ellipsoid == nil {
		geogr.Ellipsoid = DefaultEllipsoid
	}
}

func (proj *Projected) Validate() {
	if proj == nil {
		proj = &Projected{}
	}
	if proj.System == nil {
		proj.System = DefaultSystem
	}
	if proj.Geographic == nil {
		proj.Geographic = &Geographic{}
	}
	if proj.Geographic.Ellipsoid == nil {
		proj.Geographic.Ellipsoid = DefaultEllipsoid
	}
	if proj.Geographic.Geocentric == nil {
		proj.Geographic.Geocentric = &Geocentric{}
	}
}
