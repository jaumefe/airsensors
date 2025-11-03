package bme680

var const_array_1 = []float64{
	1, 1, 1, 1, 1, 0.99, 1, 0.992, 1, 1, 0.998, 0.995, 1, 0.99, 1, 1,
}

var const_array_2 = []float64{
	8000000, 4000000, 2000000, 1000000, 499500.4995, 248262.1648, 125000, 63004.03226, 31281.28128, 15625, 7812.5, 3906.25, 1953.125, 976.5625, 488.28125, 244.140625,
}

type BME680Cablibration struct {
	ParT1 uint16
	ParT2 int16
	ParT3 int8

	ParP1  uint16
	ParP2  int16
	ParP3  int8
	ParP4  int16
	ParP5  int16
	ParP6  int8
	ParP7  int8
	ParP8  int16
	ParP9  int16
	ParP10 uint8

	ParH1 uint16
	ParH2 uint16
	ParH3 int8
	ParH4 int8
	ParH5 int8
	ParH6 uint8
	ParH7 int8

	GasRange      uint8
	GasRangeSWErr int8
}
