package modules

type TemperatureDecoder interface {
	TemperatureFromRaw(rawTemperature []byte) float64
}

type HumidityDecoder interface {
	HumidityFromRaw(rawHumidity []byte) float64
}

type PressureDecoder interface {
	PressureFromRaw(rawPressure []byte) float64
}

type GasResistanceDecoder interface {
	GasResistanceFromRaw(rawGasResistance []byte) float64
}

type SensorDecoder interface {
	TemperatureDecoder
	HumidityDecoder
	PressureDecoder
	GasResistanceDecoder

	Decode(raw RawData) Measurement
	SensorName() string
}
