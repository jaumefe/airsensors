package dps310

type DPS310Calibration struct {
	C0, C1                  int16
	C00, C10                int32
	C01, C11, C20, C21, C30 int16
	KT, KP                  float64
}
