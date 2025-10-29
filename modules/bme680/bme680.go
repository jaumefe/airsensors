package bme680

import "airsensors/modules"

type BME680Decoder struct {
	RawTemperature   int32
	RawPressure      int32
	RawHumidity      int16
	RawGasResistance int32
}

func (d *BME680Decoder) Decode() (modules.Measurement, error) {
	return modules.Measurement{}, nil
}

func (d *BME680Decoder) TemperatureFromRaw() (float64, error) {
	return 0, nil
}

func (d *BME680Decoder) HumidityFromRaw() (float64, error) {
	return 0, nil
}

func (d *BME680Decoder) PressureFromRaw() (float64, error) {
	return 0, nil
}

func (d *BME680Decoder) GasResistanceFromRaw() (float64, error) {
	return 0, nil
}
