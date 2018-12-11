package coo

type FromSource interface {
	FromSource(xS, yS, zS float64) (x, y, z float64)
}

type ToSource interface {
	ToSource(x, y, z float64) (xS, yS, zS float64)
}

func Transform(x, y, z float64, from ToSource, to FromSource) (xN, yN, zN float64) {
	x, y, z = from.ToSource(x, y, z)
	xN, yN, zN = to.FromSource(x, y, z)
	return
}
