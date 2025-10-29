package modules

type SensorDecoder interface {
	Decode() (Measurement, error)
	TemperatureFromRaw() (float64, error)
	HumidityFromRaw() (float64, error)
	PressureFromRaw() (float64, error)
	GasResistanceFromRaw() (float64, error)
	SensorName() string
}
