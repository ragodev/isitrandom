package isitrandom

type normalLookupTable struct {
	P []float64
	X []float64
}

var (
	// NormalTable stores the frequently needed values for looking up N(0,1) tests
	NormalTable = normalLookupTable{
		[]float64{0.000100, 0.001000, 0.002000, 0.005000, 0.010000, 0.020000, 0.025000, 0.030000, 0.040000, 0.050000, 0.060000, 0.070000, 0.080000, 0.090000, 0.100000, 0.100000, 0.110000, 0.120000, 0.130000, 0.140000, 0.150000, 0.160000, 0.170000, 0.180000, 0.190000, 0.200000, 0.210000, 0.220000, 0.230000, 0.240000, 0.250000, 0.260000, 0.270000, 0.280000, 0.290000, 0.300000, 0.310000, 0.320000, 0.330000, 0.340000, 0.350000, 0.360000, 0.370000, 0.380000, 0.390000, 0.400000, 0.410000, 0.420000, 0.430000, 0.440000, 0.450000, 0.460000, 0.470000, 0.480000, 0.490000, 0.500000, 0.510000, 0.520000, 0.530000, 0.540000, 0.550000, 0.560000, 0.570000, 0.580000, 0.590000, 0.600000, 0.610000, 0.620000, 0.630000, 0.640000, 0.650000, 0.660000, 0.670000, 0.680000, 0.690000, 0.700000, 0.710000, 0.720000, 0.730000, 0.740000, 0.750000, 0.760000, 0.770000, 0.780000, 0.790000, 0.800000, 0.810000, 0.820000, 0.830000, 0.840000, 0.850000, 0.860000, 0.870000, 0.880000, 0.890000, 0.900000, 0.910000, 0.920000, 0.930000, 0.940000, 0.950000, 0.960000, 0.970000, 0.975000, 0.980000, 0.990000, 0.995000, 0.999000, 0.999900},
		[]float64{-3.719016, -3.090232, -2.878162, -2.575829, -2.326348, -2.053749, -1.959964, -1.880794, -1.750686, -1.644854, -1.554774, -1.475791, -1.405072, -1.340755, -1.281552, -1.281552, -1.226528, -1.174987, -1.126391, -1.080319, -1.036433, -0.994458, -0.954165, -0.915365, -0.877896, -0.841621, -0.806421, -0.772193, -0.738847, -0.706303, -0.674490, -0.643345, -0.612813, -0.582842, -0.553385, -0.524401, -0.495850, -0.467699, -0.439913, -0.412463, -0.385320, -0.358459, -0.331853, -0.305481, -0.279319, -0.253347, -0.227545, -0.201893, -0.176374, -0.150969, -0.125661, -0.100434, -0.075270, -0.050154, -0.025069, 0.000000, 0.025069, 0.050154, 0.075270, 0.100434, 0.125661, 0.150969, 0.176374, 0.201893, 0.227545, 0.253347, 0.279319, 0.305481, 0.331853, 0.358459, 0.385320, 0.412463, 0.439913, 0.467699, 0.495850, 0.524401, 0.553385, 0.582842, 0.612813, 0.643345, 0.674490, 0.706303, 0.738847, 0.772193, 0.806421, 0.841621, 0.877896, 0.915365, 0.954165, 0.994458, 1.036433, 1.080319, 1.126391, 1.174987, 1.226528, 1.281552, 1.340755, 1.405072, 1.475791, 1.554774, 1.644854, 1.750686, 1.880794, 1.959964, 2.053749, 2.326348, 2.575829, 3.090232, 3.719016},
	}
)
