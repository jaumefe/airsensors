package modules

type Measurement struct {
	Temperature   float64
	Humidity      float64
	Pressure      float64
	GasResistance float64
}

type RawData struct {
	RawTemperature   uint
	RawHumidity      uint
	RawPressure      uint
	RawGasResistance uint
}
