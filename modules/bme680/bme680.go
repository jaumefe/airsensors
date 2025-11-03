package bme680

import "airsensors/modules"

type BME680Decoder struct {
	calibration BME680Cablibration
}

func NewDecoder(calibration BME680Cablibration) *BME680Decoder {
	return &BME680Decoder{
		calibration: calibration,
	}
}

func (d *BME680Decoder) Decode(raw modules.RawData) modules.Measurement {
	meas := modules.Measurement{}
	meas.Temperature = d.TemperatureFromRaw(raw.RawTemperature)
	meas.Humidity = d.HumidityFromRaw(raw.RawHumidity, meas.Temperature)
	meas.Pressure = d.PressureFromRaw(raw.RawPressure, meas.Temperature)
	meas.GasResistance = d.GasResistanceFromRaw(raw.RawGasResistance)
	return meas
}

func (d *BME680Decoder) TemperatureFromRaw(rawTemperature uint) float64 {
	var1 := (((float64)(rawTemperature) / 16384.0) - ((float64)(d.calibration.ParT1) / 1024.0)) * (float64)(d.calibration.ParT2)
	var2 := ((((float64)(rawTemperature) / 131072.0) - ((float64)(d.calibration.ParT1) / 8192.0)) * (((float64)(rawTemperature) / 131072.0) - ((float64)(d.calibration.ParT1) / 8192.0))) * ((float64)(d.calibration.ParT3) * 16.0)
	tFine := var1 + var2
	return tFine / 5120.0
}

func (d *BME680Decoder) HumidityFromRaw(rawHumidity uint, temperature float64) float64 {
	var1 := (float64)(rawHumidity) - (((float64)(d.calibration.ParH1) * 16.0) + (((float64)(d.calibration.ParH3) / 2.0) * temperature))
	var2 := var1 * (((float64)(d.calibration.ParH2) / 262144.0) * (1.0 + (((float64)(d.calibration.ParH4) / 16384.0) * temperature) + (((float64)(d.calibration.ParH5) / 1048576.0) * temperature * temperature)))
	var3 := (float64)(d.calibration.ParH6) / 16384.0
	var4 := (float64)(d.calibration.ParH7) / 2097152.0
	humidity := var2 + ((var3 + (var4 * temperature)) * var2 * var2)
	if humidity > 100.0 {
		humidity = 100.0
	} else if humidity < 0.0 {
		humidity = 0.0
	}
	return humidity
}

func (d *BME680Decoder) PressureFromRaw(rawPressure uint, temperature float64) float64 {
	var1 := (temperature/5120.0)/2.0 - 64000.0
	var2 := var1 * var1 * ((float64)(d.calibration.ParP6)) / 131072.0
	var2 = var2 + (var1 * ((float64)(d.calibration.ParP5)) * 2.0)
	var2 = (var2 / 4.0) + (((float64)(d.calibration.ParP4)) * 65536.0)
	var1 = ((((float64)(d.calibration.ParP3) * var1 * var1) / 16384.0) + ((float64)(d.calibration.ParP2) * var1)) / 524288.0
	var1 = (1.0 + (var1 / 32768.0)) * (float64)(d.calibration.ParP1)
	pressure := 1048576.0 - (float64)(rawPressure)
	pressure = (pressure - (var2 / 4096.0)) * 6250.0 / var1
	var1 = ((float64)(d.calibration.ParP9) * pressure * pressure) / 2147483648.0
	var2 = pressure * ((float64)(d.calibration.ParP8)) / 32768.0
	var3 := (pressure / 256.0) * (pressure / 256.0) * (pressure / 256.0) * ((float64)(d.calibration.ParP10) / 131072.0)
	pressure = pressure + ((var1 + var2 + var3 + ((float64)(d.calibration.ParP7) * 128)) / 16.0)
	return pressure
}

func (d *BME680Decoder) GasResistanceFromRaw(rawGasResistance uint) float64 {
	var1 := (1340.0 + 5.0*(float64)(d.calibration.GasRangeSWErr)) * const_array_1[d.calibration.GasRange]
	gas_res := var1 * const_array_2[d.calibration.GasRange] / ((float64)(rawGasResistance) - 512.0 + var1)
	return gas_res
}

func (d *BME680Decoder) SensorName() string {
	return "BME680"
}
