package coo

var WGS84 = &Ellipsoid{
	A:  6378137,
	Fi: 298.257223563,
}

var GRS80 = &Ellipsoid{
	A:  6378137,
	Fi: 298.257222101,
}

var Bessel = &Ellipsoid{
	A:  6377397.155,
	Fi: 299.1528153513233,
}

var Krassowski = &Ellipsoid{
	A:  6378245,
	Fi: 298.3,
}

var Airy1830 = &Ellipsoid{
	A:  6377340.189,
	Fi: 299.3249514,
}

var Clarke1866 = &Ellipsoid{
	A:  6378206.400,
	Fi: 294.9786982,
}

var Clarke1880 = &Ellipsoid{
	A:  6378249.17,
	Fi: 293.4663,
}

var Hayford = &Ellipsoid{
	A:  6378388.000,
	Fi: 297.0,
}
