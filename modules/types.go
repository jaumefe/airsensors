package modules

/*
Measurement holds the processed sensor data
*/
type Measurement struct {
	Temperature   float64
	Humidity      float64
	Pressure      float64
	GasResistance float64
}

/*
RawData holds the raw data sensor readings into integers
*/
type RawData struct {
	RawTemperature   int
	RawHumidity      int
	RawPressure      int
	RawGasResistance int
}
